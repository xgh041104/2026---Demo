<template>
  <div>
    <ProTable ref="proTable" :columns="columns" :request-api="getTableList" :data-callback="dataCallback">
      <template #tableHeader="">
        <div style="font-size: 20px">待审核列表</div>
      </template>
      <!-- 表格操作 -->
      <template #operation="scope">
        <el-button type="success" size="small" @click="openDialog(1, scope.row)">通过</el-button>
        <el-button type="danger" size="small" @click="openDialogBH(1, scope.row)">驳回</el-button>
      </template>
    </ProTable>
    <UserDrawer ref="drawerRef" />
    <ImportExcel ref="dialogRef" />
    <passThrough v-model:visible="DialogVisible" :user-info="CurrentRow" />
    <turnDown v-model:visible="DialogVisibleBH" :user-info="CurrentRowBH" />
  </div>
</template>

<script setup lang="tsx">
import { getMaterialList } from "@/api/modules/material";
import ProTable from "@/components/ProTable/index.vue";
import { reactive, ref } from "vue";
import UserDrawer from "@/views/proTable/components/UserDrawer.vue";
import turnDown from "./components/turnDown.vue";
import passThrough from "./components/passThrough.vue";
import { Material } from "@/api/interface/material";

const DialogVisible = ref(0);
const DialogVisibleBH = ref(0);

const CurrentRow = ref<Material.ReqMaterial>({} as Material.ReqMaterial);
const CurrentRowBH = ref<Material.ReqMaterial>({} as Material.ReqMaterial);

const getTableList = (params: any) => {
  return getMaterialList(params);
};

const dataCallback = (data: any) => {
  return {
    list: data.list,
    total: data.total
  };
};

const columns = reactive([
  {
    prop: "name",
    label: "素材名称",
    width: "auto",
    render: scope => (
      <el-button
        type="primary"
        link
        onClick={() => {
          openDialog(2, scope.row);
        }}
      >
        {scope.row.name}
      </el-button>
    )
  },
  { prop: "creator_name", label: "上传者", width: "auto" },
  { prop: "created_at", label: "上传时间", width: "auto" },
  // { prop: "image_url", label: "图片路径", width: "auto" },
  {
    label: "操作",
    prop: "operation",
    width: "auto"
  }
]);

const openDialog = (type: number, row: Material.ReqMaterial = {} as Material.ReqMaterial) => {
  DialogVisible.value = type;
  CurrentRow.value = { ...row };
};

const openDialogBH = (type: number, row: Material.ReqMaterial = {} as Material.ReqMaterial) => {
  DialogVisibleBH.value = type;
  CurrentRowBH.value = { ...row };
};
</script>
