<template>
  <div class="page-wrap">
    <div class="toolbar">
      <el-button type="primary" @click="openCreateDialog">+ 新增标签</el-button>
    </div>

    <el-table :data="tagList" border>
      <el-table-column prop="id" label="标签ID" width="120" />
      <el-table-column label="标签名称" min-width="300">
        <template #default="{ row }">
          <el-tag :type="row.colorType" effect="dark">{{ row.name }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="140" fixed="right">
        <template #default="{ row }">
          <el-button type="danger" link @click="removeTag(row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="420px">
      <el-form ref="formRef" :model="formModel" :rules="formRules" label-width="80px">
        <el-form-item label="标签名称" prop="name">
          <el-input v-model="formModel.name" maxlength="20" placeholder="请输入标签名称" clearable />
        </el-form-item>
        <el-form-item label="标签颜色" prop="colorType">
          <el-select v-model="formModel.colorType" style="width: 100%">
            <el-option label="黄色" value="warning" />
            <el-option label="红色" value="danger" />
            <el-option label="绿色" value="success" />
            <el-option label="蓝色" value="primary" />
            <el-option label="灰色" value="info" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="submitDialog">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts" name="TagList">
import { reactive, ref } from "vue";
import { ElMessage, ElMessageBox, FormInstance, FormRules } from "element-plus";

interface TagItem {
  id: number;
  name: string;
  colorType: "success" | "warning" | "danger" | "primary" | "info";
}

const STORAGE_KEY = "medi-link-tag-list";
const formRef = ref<FormInstance>();
const dialogVisible = ref(false);
const dialogTitle = ref("新增标签");
const idSeed = ref(100);

const defaultList: TagItem[] = [
  { id: 8, name: "科技感", colorType: "warning" },
  { id: 9, name: "扁平化设计", colorType: "danger" },
  { id: 10, name: "动态特效", colorType: "primary" },
  { id: 11, name: "红色主题", colorType: "success" },
  { id: 12, name: "素材模版", colorType: "info" },
  { id: 13, name: "低碳环保", colorType: "warning" },
  { id: 14, name: "人工智能", colorType: "danger" },
  { id: 15, name: "立体3D", colorType: "primary" },
  { id: 16, name: "极简主义", colorType: "success" },
  { id: 17, name: "商用授权", colorType: "info" },
  { id: 18, name: "独家资源", colorType: "warning" },
  { id: 19, name: "待优化", colorType: "danger" },
  { id: 20, name: "优质精选", colorType: "primary" }
];

const loadLocalData = () => {
  // 开发环境直接使用 defaultList，方便调试
  if (import.meta.env.DEV) {
    return defaultList;
  }

  // 生产环境才使用 localStorage 缓存
  const raw = localStorage.getItem(STORAGE_KEY);
  if (!raw) return defaultList;
  try {
    const parsed = JSON.parse(raw);
    return Array.isArray(parsed) ? parsed : defaultList;
  } catch {
    return defaultList;
  }
};

const tagList = ref<TagItem[]>(loadLocalData());
idSeed.value = tagList.value.length ? Math.max(...tagList.value.map(item => item.id)) + 1 : 1;

const persist = () => {
  localStorage.setItem(STORAGE_KEY, JSON.stringify(tagList.value));
  idSeed.value = tagList.value.length ? Math.max(...tagList.value.map(item => item.id)) + 1 : 1;
};

const formModel = reactive<TagItem>({ id: 0, name: "", colorType: "warning" });
const formRules = reactive<FormRules>({
  name: [{ required: true, message: "请输入标签名称", trigger: "blur" }],
  colorType: [{ required: true, message: "请选择标签颜色", trigger: "change" }]
});

const openCreateDialog = () => {
  dialogTitle.value = "新增标签";
  formModel.name = "";
  formModel.colorType = "warning";
  dialogVisible.value = true;
};

const submitDialog = async () => {
  if (!formRef.value) return;
  await formRef.value.validate(valid => {
    if (!valid) return;
    tagList.value.unshift({ id: idSeed.value++, name: formModel.name.trim(), colorType: formModel.colorType });
    dialogVisible.value = false;
    persist();
    ElMessage.success("操作成功");
  });
};

const removeTag = async (row: TagItem) => {
  await ElMessageBox.confirm(`确认删除标签「${row.name}」吗？`, "删除确认", { type: "warning" });
  tagList.value = tagList.value.filter(item => item.id !== row.id);
  persist();
  ElMessage.success("删除成功");
};
</script>

<style scoped lang="scss">
.page-wrap {
  padding: 16px;
  background: #ffffff;
}
</style>
