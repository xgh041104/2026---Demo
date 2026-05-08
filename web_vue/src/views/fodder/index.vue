<template>
  <div class="main-layout">
    <!-- 左侧：固定宽度的树 -->
    <div class="sidebar">
      <el-tree :data="data" :props="props" @node-click="handleNodeClick" default-expand-all />
    </div>

    <!-- 右侧：占据剩余宽度，显示内容 -->
    <div class="content-area">
      <!-- 方式A：如果是内部组件切换 -->
      <component :is="currentComponent" v-if="currentComponent" />

      <!-- 方式B：如果是外部链接或独立页面 -->
      <iframe v-if="currentUrl" :src="currentUrl" class="page-iframe" frameborder="0"></iframe>

      <!-- 默认提示 -->
      <div v-if="!currentComponent && !currentUrl" class="placeholder">请点击左侧菜单</div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref } from "vue";

// 引入你想要展示的页面组件 (示例)
// import ImageView from './ImageView.vue';
// import DesignView from './DesignView.vue';

interface Tree {
  id: string;
  label: string;
  component?: string; // 标记对应哪个组件
  url?: string; // 或者标记对应哪个链接
  children?: Tree[];
}

const props = { children: "children", label: "label" };

// 模拟数据：给节点加上 component 或 url 标记
const data: Tree[] = [
  {
    id: "1",
    label: "影像素材",
    children: [
      { id: "1-1", label: "高清视频", url: "https://www.example.com/video" },
      { id: "1-2", label: "摄影图片", component: "ImageView" }
    ]
  },
  {
    id: "2",
    label: "平面设计",
    children: [
      { id: "2-1", label: "图标", url: "/posters.html" },
      { id: "2-2", label: "宣传海报", url: "/posters.html" }
    ]
  },
  {
    id: "3",
    label: "办公文档",
    children: [{ id: "3-1", label: "项目报告", url: "/manual.html" }]
  },
  {
    id: "3",
    label: "相机素材",
    url: ""
  }
];

const currentComponent = ref<string | null>(null);
const currentUrl = ref<string | null>(null);

const handleNodeClick = (node: Tree) => {
  // 点击节点时，根据节点数据决定右边显示什么
  console.log("点击了节点：", node.label);

  if (node.component) {
    currentComponent.value = node.component;
    currentUrl.value = null;
  } else if (node.url) {
    currentUrl.value = node.url;
    currentComponent.value = null;
  } else {
    // 如果点击的是父级目录，可以清空内容或保持原样
    currentComponent.value = null;
    currentUrl.value = null;
  }
};
</script>

<style scoped>
.main-layout {
  display: flex;
  height: 100vh; /* 全屏高度 */
  overflow: hidden;
}

.sidebar {
  width: 250px; /* 左侧固定宽度 */
  border-right: 1px solid #ddd;
  overflow-y: auto;
  background-color: #f5f7fa;
}

.content-area {
  flex: 1; /* 右侧自动填满 */
  padding: 20px;
  overflow-y: auto;
  position: relative;
}

.page-iframe {
  width: 100%;
  height: 100%;
  min-height: 600px;
}
</style>
