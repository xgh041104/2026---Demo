<template>
  <el-dialog v-model="dialogVisible" :title="dialogTitle" @close="handleClose" width="600px">
    <el-form
      :model="formData"
      :rules="formRules"
      label-width="100px"
      ref="formRef"
    >
      <el-form-item label="素材标题" prop="title">
        <el-input v-model="formData.title" placeholder="请输入素材标题" />
      </el-form-item>
      <el-form-item label="素材分类" prop="category_id">
        <el-select v-model="formData.category_id" placeholder="请选择分类" style="width: 100%">
          <el-option label="摄影图片" :value="1" />
          <el-option label="高清视频" :value="2" />
          <el-option label="图标" :value="3" />
          <el-option label="宣传海报" :value="4" />
          <el-option label="项目报告" :value="5" />
        </el-select>
      </el-form-item>
      <el-form-item label="素材标签" prop="label_id">
        <el-select v-model="formData.label_id" placeholder="请选择标签" style="width: 100%" multiple>
          <el-option label="风景" :value="1" />
          <el-option label="建筑" :value="2" />
          <el-option label="人物" :value="3" />
          <el-option label="花卉" :value="4" />
          <el-option label="动物" :value="5" />
        </el-select>
      </el-form-item>
      <el-form-item label="上传文件" prop="file">
        <el-upload
          class="upload-demo"
          drag
          :auto-upload="false"
          :on-change="handleFileChange"
          :limit="1"
          :accept="acceptTypes"
        >
          <el-icon class="el-icon--upload"><upload-filled /></el-icon>
          <div class="el-upload__text">
            将文件拖到此处，或<em>点击上传</em>
          </div>
          <template #tip>
            <div class="el-upload__tip">
              支持 jpg/png/mp4 格式，文件大小不超过 10MB
            </div>
          </template>
        </el-upload>
      </el-form-item>
      <el-form-item label="备注" prop="remark">
        <el-input v-model="formData.remark" type="textarea" :rows="3" placeholder="请输入备注信息" />
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="handleCancel">取消</el-button>
        <el-button type="primary" @click="handleSubmit" :loading="submitLoading">确定</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup>
import { ref, computed, watch } from "vue";
import { FormRules, ElMessage } from "element-plus";
import { UploadFilled } from "@element-plus/icons-vue";
import { saveMaterialApi } from "@/api/modules/material";
import { useUserStore } from "@/stores/modules/user";

interface FodderFormData {
  title: string;
  category_id: number | null;
  label_id: number[];
  file: File | null;
  remark: string;
}

const props = defineProps({
  visible: { type: Boolean, default: false },
  type: { type: Number, default: 0 },
  info: { type: Object, default: () => ({}) }
});

const emits = defineEmits(["update:visible", "refresh"]);

const userStore = useUserStore();

const dialogVisible = computed({
  get: () => props.visible,
  set: (val) => emits("update:visible", val)
});

const dialogTitle = computed(() => {
  switch (props.type) {
    case 1:
      return "新增素材";
    case 2:
      return "查看素材";
    case 3:
      return "编辑素材";
    default:
      return "素材管理";
  }
});

const formRef = ref();
const submitLoading = ref(false);
const formData = ref<FodderFormData>({
  title: "",
  category_id: null,
  label_id: [],
  file: null,
  remark: ""
});

const acceptTypes = ".jpg,.jpeg,.png,.mp4";

const formRules: FormRules = {
  title: [
    { required: true, message: "请输入素材标题", trigger: "blur" },
    { min: 1, max: 100, message: "长度在 1 到 100 个字符", trigger: "blur" }
  ],
  category_id: [{ required: true, message: "请选择素材分类", trigger: "change" }],
  label_id: [{ required: true, message: "请选择素材标签", trigger: "change" }]
};

watch(
  () => props.visible,
  (val) => {
    if (val && props.info && Object.keys(props.info).length > 0) {
      formData.value = {
        title: props.info.title || "",
        category_id: props.info.category_id || null,
        label_id: props.info.label_id || [],
        file: null,
        remark: props.info.remark || ""
      };
    } else if (val) {
      resetForm();
    }
  }
);

const resetForm = () => {
  formData.value = {
    title: "",
    category_id: null,
    label_id: [],
    file: null,
    remark: ""
  };
  formRef.value?.resetFields();
  formRef.value?.clearValidate();
};

const handleFileChange = (file: any) => {
  const isLt10M = file.size / 1024 / 1024 < 10;
  if (!isLt10M) {
    ElMessage.error("文件大小不能超过 10MB");
    return;
  }
  formData.value.file = file.raw;
};

const handleSubmit = () => {
  formRef.value?.validate(async (valid: boolean) => {
    if (valid) {
      if (!formData.value.file) {
        ElMessage.error("请上传素材文件");
        return;
      }

      submitLoading.value = true;
      try {
        const submitData = new FormData();
        submitData.append("file", formData.value.file);
        submitData.append("creator_id", String(userStore.userInfo.id || 1));
        submitData.append("status", "0");
        submitData.append("label_id", formData.value.label_id.join(","));
        if (formData.value.category_id) {
          submitData.append("category_id", String(formData.value.category_id));
        }
        if (formData.value.title) {
          submitData.append("title", formData.value.title);
        }
        if (formData.value.remark) {
          submitData.append("remark", formData.value.remark);
        }

        await saveMaterialApi(submitData);
        ElMessage.success("新增素材成功");
        emits("refresh");
        dialogVisible.value = false;
      } catch (error) {
        console.error("新增素材失败:", error);
        ElMessage.error("新增素材失败");
      } finally {
        submitLoading.value = false;
      }
    } else {
      ElMessage.error("请填写完整信息");
    }
  });
};

const handleCancel = () => {
  dialogVisible.value = false;
};

const handleClose = () => {
  resetForm();
};
</script>

<style scoped>
.upload-demo {
  width: 100%;
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}
</style>
