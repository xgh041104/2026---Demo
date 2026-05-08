<template>
  <div class="page-wrap">
    <div class="toolbar">
      <el-button type="primary" @click="openCreateRootDialog">+ 新增分类</el-button>
      <el-button @click="toggleExpandAll">{{ expandAll ? "全部折叠" : "全部展开" }}</el-button>
    </div>

    <el-table
      ref="tableRef"
      :data="categoryList"
      row-key="id"
      :tree-props="{ children: 'children' }"
      border
      class="category-table"
    >
      <el-table-column prop="id" label="分类ID" width="120" />
      <el-table-column label="分类名称" min-width="420">
        <template #default="{ row }">
          <div class="name-cell">
            <el-icon class="name-icon"><Folder /></el-icon>
            <span>{{ row.name }}</span>
            <el-tag v-if="row.isTop" size="small" type="success">置顶</el-tag>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="320" fixed="right">
        <template #default="{ row }">
          <el-button type="success" link @click="setTop(row)">{{ row.isTop ? "已置顶" : "置顶" }}</el-button>
          <el-button type="primary" link @click="openCreateChildDialog(row)">+新增子项</el-button>
          <el-button type="warning" link @click="openEditDialog(row)">编辑</el-button>
          <el-button type="danger" link @click="removeCategory(row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="420px">
      <el-form ref="formRef" :model="formModel" :rules="formRules" label-width="80px">
        <el-form-item label="分类名称" prop="name">
          <el-input v-model="formModel.name" maxlength="20" placeholder="请输入分类名称" clearable />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="submitDialog">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts" name="CategoryList">
import { nextTick, reactive, ref } from "vue";
import { ElMessage, ElMessageBox, FormInstance, FormRules } from "element-plus";
import { Folder } from "@element-plus/icons-vue";

interface CategoryItem {
  id: number;
  name: string;
  isTop?: boolean;
  children?: CategoryItem[];
}

type DialogMode = "createRoot" | "createChild" | "edit";

const STORAGE_KEY = "medi-link-category-list";
const tableRef = ref<any>();
const formRef = ref<FormInstance>();
const expandAll = ref(false);
const dialogVisible = ref(false);
const dialogTitle = ref("新增分类");
const dialogMode = ref<DialogMode>("createRoot");
const currentTarget = ref<CategoryItem | null>(null);
const idSeed = ref(100);

const defaultData: CategoryItem[] = [
  {
    id: 1,
    name: "影像素材",
    isTop: true,
    children: [
      { id: 2, name: "高清视频" },
      { id: 5, name: "摄影图片" }
    ]
  },
  { id: 12, name: "123" },
  { id: 6, name: "严谨设计", children: [{ id: 8, name: "办公文档" }] },
  { id: 11, name: "相机素材" }
];

const formModel = reactive({ name: "" });
const formRules = reactive<FormRules>({
  name: [{ required: true, message: "请输入分类名称", trigger: "blur" }]
});

const loadLocalData = (): CategoryItem[] => {
  const raw = localStorage.getItem(STORAGE_KEY);
  if (!raw) return defaultData;
  try {
    const parsed = JSON.parse(raw);
    return Array.isArray(parsed) ? parsed : defaultData;
  } catch {
    return defaultData;
  }
};

const categoryList = ref<CategoryItem[]>(loadLocalData());

const flattenIds = (list: CategoryItem[]): number[] =>
  list.flatMap(item => [item.id, ...(item.children ? flattenIds(item.children) : [])]);

const refreshIdSeed = () => {
  const ids = flattenIds(categoryList.value);
  idSeed.value = ids.length ? Math.max(...ids) + 1 : 1;
};

const persist = () => {
  localStorage.setItem(STORAGE_KEY, JSON.stringify(categoryList.value));
  refreshIdSeed();
};

refreshIdSeed();

const createItem = (name: string): CategoryItem => ({ id: idSeed.value++, name });

const openCreateRootDialog = () => {
  dialogTitle.value = "新增根分类";
  dialogMode.value = "createRoot";
  currentTarget.value = null;
  formModel.name = "";
  dialogVisible.value = true;
};

const openCreateChildDialog = (row: CategoryItem) => {
  dialogTitle.value = `新增子分类 - ${row.name}`;
  dialogMode.value = "createChild";
  currentTarget.value = row;
  formModel.name = "";
  dialogVisible.value = true;
};

const openEditDialog = (row: CategoryItem) => {
  dialogTitle.value = `编辑分类 - ${row.name}`;
  dialogMode.value = "edit";
  currentTarget.value = row;
  formModel.name = row.name;
  dialogVisible.value = true;
};

const submitDialog = async () => {
  if (!formRef.value) return;
  await formRef.value.validate(async valid => {
    if (!valid) return;
    if (dialogMode.value === "createRoot") categoryList.value.unshift(createItem(formModel.name.trim()));
    if (dialogMode.value === "createChild" && currentTarget.value) {
      currentTarget.value.children ??= [];
      currentTarget.value.children.unshift(createItem(formModel.name.trim()));
      expandAll.value = false;
      await nextTick();
      expandAll.value = true;
    }
    if (dialogMode.value === "edit" && currentTarget.value) currentTarget.value.name = formModel.name.trim();
    dialogVisible.value = false;
    persist();
    ElMessage.success("操作成功");
  });
};

const sortRecursiveTop = (list: CategoryItem[]) => {
  list.sort((a, b) => Number(Boolean(b.isTop)) - Number(Boolean(a.isTop)));
  list.forEach(item => item.children && sortRecursiveTop(item.children));
};

const clearTopRecursive = (list: CategoryItem[]) => {
  list.forEach(item => {
    item.isTop = false;
    if (item.children?.length) clearTopRecursive(item.children);
  });
};

const setTop = (row: CategoryItem) => {
  if (row.isTop) {
    ElMessage.info("当前分类已置顶");
    return;
  }
  clearTopRecursive(categoryList.value);
  row.isTop = true;
  sortRecursiveTop(categoryList.value);
  persist();
  ElMessage.success("已置顶");
};

const removeNodeById = (list: CategoryItem[], id: number): boolean => {
  const index = list.findIndex(item => item.id === id);
  if (index !== -1) {
    list.splice(index, 1);
    return true;
  }
  return list.some(item => (item.children ? removeNodeById(item.children, id) : false));
};

const removeCategory = async (row: CategoryItem) => {
  await ElMessageBox.confirm(`确认删除分类「${row.name}」吗？`, "删除确认", { type: "warning" });
  removeNodeById(categoryList.value, row.id);
  persist();
  ElMessage.success("删除成功");
};

const toggleExpandAll = async () => {
  expandAll.value = !expandAll.value;
  if (!tableRef.value) return;
  if (!expandAll.value) {
    const flatten = (list: CategoryItem[]) => list.flatMap(item => [item, ...(item.children ? flatten(item.children) : [])]);
    flatten(categoryList.value).forEach(item => tableRef.value.toggleRowExpansion(item, false));
    return;
  }
  await nextTick();
  const expand = (list: CategoryItem[]) => {
    list.forEach(item => {
      tableRef.value.toggleRowExpansion(item, true);
      if (item.children?.length) expand(item.children);
    });
  };
  expand(categoryList.value);
};
</script>

<style scoped lang="scss">
.page-wrap {
  padding: 16px;
  background: #ffffff;
}
.toolbar {
  display: flex;
  gap: 12px;
  margin-bottom: 14px;
}
.toolbar :deep(.el-button + .el-button) {
  margin-left: 0;
}
.category-table {
  width: 100%;
}
.name-cell {
  display: flex;
  gap: 8px;
  align-items: center;
}
.name-icon {
  color: #909399;
}
.category-table :deep(.el-button.is-link) {
  font-size: 13px;
}
</style>
