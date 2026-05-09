<template>
  <el-dialog v-model="dialogVisible" :title="dialogTitle" @close="handleClose" width="600px">
    <el-form
      :model="formData"
      :rules="formRules"
      label-width="100px"
      ref="formRef"
    >
      <el-form-item label="素材名称" prop="name">
        <el-input v-model="formData.name" placeholder="请输入素材名称" />
      </el-form-item>
      <el-form-item label="素材分类" prop="category">
        <el-select v-model="formData.category" placeholder="请选择分类" style="width: 100%">
          <el-option label="摄影图片" value="摄影图片" />
          <el-option label="高清视频" value="高清视频" />
          <el-option label="图标" value="图标" />
          <el-option label="宣传海报" value="宣传海报" />
          <el-option label="项目报告" value="项目报告" />
        </el-select>
      </el-form-item>
      <el-form-item label="素材标签" prop="tags">
        <el-select v-model="formData.tags" placeholder="请选择标签" style="width: 100%" multiple>
          <el-option label="风景" value="风景" />
          <el-option label="建筑" value="建筑" />
          <el-option label="人物" value="人物" />
          <el-option label="花卉" value="花卉" />
          <el-option label="动物" value="动物" />
        </el-select>
      </el-form-item>
      <el-form-item label="上传文件" prop="file">
        <el-upload
          class="upload-demo"
          drag
          :auto-upload="false"
          :on-change="handleFileChange"
          :limit="1"
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
        <el-button type="primary" @click="handleSubmit">确定</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup>
import { ref, computed, watch } from "vue";
import { FormRules, ElMessage } from "element-plus";
import { UploadFilled } from "@element-plus/icons-vue";

interface FodderFormData {
  name: string;
  category: string;
  tags: string[];
  file: File | null;
  remark: string;
}

const props = defineProps({
  visible: { type: Boolean, default: false },
  type: { type: Number, default: 0 },
  info: { type: Object, default: () => ({}) }
});

const emits = defineEmits(["update:visible", "refresh"]);

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
const formData = ref<FodderFormData>({
  name: "",
  category: "",
  tags: [],
  file: null,
  remark: ""
});

const formRules: FormRules = {
  name: [
    { required: true, message: "请输入素材名称", trigger: "blur" },
    { min: 1, max: 100, message: "长度在 1 到 100 个字符", trigger: "blur" }
  ],
  category: [{ required: true, message: "请选择素材分类", trigger: "change" }],
  tags: [{ required: true, message: "请选择素材标签", trigger: "change" }]
};

watch(
  () => props.visible,
  (val) => {
    if (val && props.info && Object.keys(props.info).length > 0) {
      formData.value = {
        name: props.info.name || "",
        category: props.info.category || "",
        tags: props.info.tags || [],
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
    name: "",
    category: "",
    tags: [],
    file: null,
    remark: ""
  };
  formRef.value?.resetFields();
  formRef.value?.clearValidate();
};

const handleFileChange = (file: any) => {
  formData.value.file = file.raw;
};

const handleSubmit = () => {
  formRef.value?.validate((valid: boolean) => {
    if (valid) {
      console.log("提交表单数据:", formData.value);
      ElMessage.success("操作成功");
      emits("refresh");
      dialogVisible.value = false;
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
