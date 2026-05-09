<template>
  <!--搜索框-->
  <div style="display: flex; justify-content: flex-start; align-items: center; margin-bottom: 20px; gap: 20px">
    <el-input v-model="searchKeyword" placeholder="输入关键字搜索" style="width: 300px" clearable @keyup.enter="handleSearch">
      <template #append>
        <el-button :icon="Search" style="width: 100px" @click="handleSearch">搜索</el-button>
      </template>
    </el-input>

    <!--标签选择器-->
    <el-select v-model="value" placeholder="Select" style="width: 240px">
      <el-option v-for="item in tagOptions" :key="item.value" :label="item.label" :value="item.value" />
    </el-select>

    <!-- 表格 header 按钮 -->
    <div style="margin-top: 20px; margin-bottom: 20px">
      <el-button type="primary" :icon="CirclePlus" @click="dialogType = 1">新增素材</el-button>
    </div>

    <div style="margin-top: 20px; margin-bottom: 20px">
      <el-button type="primary" :icon="CirclePlus" @click="dialogType = 0">批量导入</el-button>
    </div>
  </div>

  <!--标签页-->
  <el-tabs v-model="activeName" class="demo-tabs" @tab-click="handleClick">
    <el-tab-pane label="User" name="first">User</el-tab-pane>
    <el-tab-pane label="Config" name="second">Config</el-tab-pane>
    <el-tab-pane label="Role" name="third">Role</el-tab-pane>
    <el-tab-pane label="Task" name="fourth">Task</el-tab-pane>
  </el-tabs>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import type { TabsPaneContext } from "element-plus";
import { Search } from "@element-plus/icons-vue";
import { Fodder } from "@/api/interface";
import { CirclePlus } from "@element-plus/icons-vue"; //Delete, EditPen, View

const value = ref("");

const searchKeyword = ref("");

const handleSearch = () => {
  console.log("搜索关键词：", searchKeyword.value);
  // 在这里编写你的搜索逻辑
};

const tagOptions = [
  {
    value: "Option1",
    label: "Option1"
  },
  {
    value: "Option2",
    label: "Option2"
  },
  {
    value: "Option3",
    label: "Option3"
  },
  {
    value: "Option4",
    label: "Option4"
  },
  {
    value: "Option5",
    label: "Option5"
  }
];

const activeName = ref("first");
const handleClick = (tab: TabsPaneContext, event: Event) => {
  console.log(tab, event);
};

//增删改
// 打开 drawer(新增、查看、编辑)
// ProTable 实例
const bookInfo = ref<Fodder.fodderList>();
const dialogType = ref(0); //0-关闭  1-新增  2-查看  3-编辑
const getDialogType = (type: number, info: Fodder.fodderList) => {
  dialogType.value = type;
  bookInfo.value = info;
};
</script>

<style>
.demo-tabs > .el-tabs__content {
  padding: 32px;
  color: #6b778c;
  font-size: 32px;
  font-weight: 600;
}

.search-container {
  padding: 16px;
  border-bottom: 1px solid #e4e7ed;
}
</style>
