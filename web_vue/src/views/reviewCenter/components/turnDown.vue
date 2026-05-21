<template>
  <el-dialog v-model="Visible" :title="formTitle" width="900px" :before-close="handleClose" center :destroy-on-close="true">
    <el-form ref="Ref" :model="tableForm" label-width="auto" style="max-width: 750px" :disabled="props.visible === 2">
      <el-form-item label="图片名称：" prop="name">
        <el-text>{{ tableForm.name }}</el-text>
      </el-form-item>
      <el-form-item label="上传者：" prop="creator_name">
        <el-text>{{ tableForm.creator_name }}</el-text>
      </el-form-item>
      <el-form-item label="素材：">
        <el-image v-if="IsImage" :src="url" fit="contain" style="width: 700px; height: auto" />
        <video v-else controls style="width: 700px; height: auto">
          <source :src="url" type="video/mp4" />
          你的浏览器不支持视频播放
        </video>
      </el-form-item>
      <div class="dialog-footer">
        <el-button type="primary" @click="handleSubmit"> 确认 </el-button>
        <el-button @click="handleClose"> 取消 </el-button>
      </div>
    </el-form>
  </el-dialog>
</template>

<script lang="ts" setup>
import { Material } from "@/api/interface/material";
import { ElMessage } from "element-plus";
import { computed, ref, watch } from "vue";
import { upAuditMaterialApi } from "@/api/modules/material";

const Ref = ref();

let url = "";

const tableForm = ref({
  id: 0,
  name: "",
  image_url: "",
  creator_name: ""
});

const emit = defineEmits(["update:visible", "refresh"]);
const props = defineProps({
  visible: {
    type: Number,
    default: 0
  },
  userInfo: {
    type: Object as () => Material.ReqMaterial,
    default: () => ({})
  }
});
const Visible = ref(false);
const formTitle = ref("");

const handleSubmit = async () => {
  try {
    const params = {
      id: tableForm.value.id,
      status: 1,
      remark: ""
    };
    await upAuditMaterialApi(params);
    ElMessage.success("操作成功");
    handleClose();
    window.location.reload();
  } catch {
    ElMessage.error("操作失败");
  }
};

const handleClose = () => {
  Visible.value = false;
  emit("update:visible", 0);
};

// 判断是否是图片
const IsImage = computed(() => {
  if (!tableForm.value?.image_url) return true;
  const suffix = tableForm.value.image_url.split(".").pop()?.toLowerCase();
  const imageTypes = ["jpg", "jpeg", "png", "gif", "webp", "bmp", "svg"];
  return imageTypes.includes(suffix || "");
});

watch(
  () => props.visible,
  newVal => {
    if (newVal !== 0) {
      Visible.value = true;
      switch (newVal) {
        case 1:
          formTitle.value = "确认通过审核";
          tableForm.value = { ...props.userInfo };
          url = `http://localhost:8000` + tableForm.value.image_url;
          console.log(tableForm.value.image_url);
          break;
        case 2:
          formTitle.value = "查看审核图片";
          tableForm.value = { ...props.userInfo };
          url = `http://localhost:8000` + tableForm.value.image_url;
          break;
        case 3:
          formTitle.value = "编辑用户";
          tableForm.value = { ...props.userInfo };
          break;
      }
    }
  }
);
</script>
