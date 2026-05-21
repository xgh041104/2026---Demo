<template>
  <div class="fodder-container">
    <!-- 左侧树形菜单 -->
    <div class="sidebar">
      <el-tree
        :data="treeData"
        :props="treeProps"
        @node-click="handleNodeClick"
        default-expand-all
        highlight-current
        :current-node-key="currentNodeKey"
        node-key="id"
      >
        <template #default="{ node, data }">
          <span class="tree-node-label">
            <el-icon v-if="data.icon" class="node-icon">
              <component :is="data.icon" />
            </el-icon>
            {{ node.label }}
          </span>
        </template>
      </el-tree>
    </div>

    <!-- 右侧内容区域 -->
    <div class="content-area">
      <!-- 面包屑导航 -->
      <el-breadcrumb separator="/" class="breadcrumb">
        <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
        <el-breadcrumb-item>素材</el-breadcrumb-item>
      </el-breadcrumb>

      <!-- 搜索和操作区域 -->
      <div class="toolbar">
        <el-input
          v-model="searchKeyword"
          placeholder="请输入素材名称搜索"
          class="search-input"
          clearable
          @keyup.enter="handleSearch"
        >
          <template #prefix>
            <el-icon>
              <Search />
            </el-icon>
          </template>
        </el-input>

        <el-select v-model="selectedTag" placeholder="请选择标签筛选" class="tag-select" clearable>
          <el-option v-for="item in tagOptions" :key="item.value" :label="item.label" :value="item.value" />
        </el-select>

        <el-button type="primary" class="action-btn" @click="handleAdd">
          <el-icon>
            <Plus />
          </el-icon>
          新增素材
        </el-button>

        <el-button type="primary" class="action-btn" @click="handleBatchImport">
          <el-icon>
            <Upload />
          </el-icon>
          批量导入
        </el-button>
      </div>

      <!-- 标签页切换 -->
      <el-tabs v-model="activeTab" class="content-tabs" @tab-click="handleTabClick">
        <el-tab-pane label="图片" name="image">
          <!-- 图片网格 -->
          <div class="image-grid">
            <div v-for="item in filteredImages" :key="item.id" class="image-card" @click="handleCardClick(item)">
              <div class="image-wrapper">
                <img :src="item.thumbnail" :alt="item.name" class="card-image" />
              </div>
              <div class="card-info">
                <div class="card-name">{{ item.name }}</div>
                <div class="card-meta">
                  <el-tag size="small" type="info">{{ item.type }}</el-tag>
                  <span class="card-time">{{ item.time }}</span>
                </div>
              </div>
            </div>
          </div>
        </el-tab-pane>

        <el-tab-pane label="视频" name="video">
          <div class="image-grid">
            <div v-for="item in filteredVideos" :key="item.id" class="image-card" @click="handleCardClick(item)">
              <div class="image-wrapper">
                <img :src="item.thumbnail" :alt="item.name" class="card-image" />
                <div class="video-overlay">
                  <el-icon class="play-icon">
                    <VideoPlay />
                  </el-icon>
                </div>
              </div>
              <div class="card-info">
                <div class="card-name">{{ item.name }}</div>
                <div class="card-meta">
                  <el-tag size="small" type="info">{{ item.type }}</el-tag>
                  <span class="card-time">{{ item.time }}</span>
                </div>
              </div>
            </div>
          </div>
        </el-tab-pane>
      </el-tabs>
    </div>

    <!-- 新增/编辑素材弹窗 -->
    <formDetail v-model:visible="dialogVisible" :type="dialogType" :info="currentItem || undefined" @refresh="handleRefresh" />
  </div>
</template>

<script lang="ts" setup>
import { ref, computed } from "vue";
import { Search, Plus, Upload, VideoPlay, Picture, Folder, Document, Camera } from "@element-plus/icons-vue";
import formDetail from "./leftPage/formDetail.vue";

interface TreeNode {
  id: string;
  label: string;
  icon?: any;
  children?: TreeNode[];
}

interface FodderItem {
  id: number;
  name: string;
  thumbnail: string;
  type: string;
  time: string;
  category: string;
}

const treeProps = { children: "children", label: "label" };
const currentNodeKey = ref<string>("");

const treeData: TreeNode[] = [
  {
    id: "1",
    label: "影像素材",
    icon: Picture,
    children: [
      { id: "1-1", label: "高清视频", icon: VideoPlay },
      { id: "1-2", label: "摄影图片", icon: Picture }
    ]
  },
  {
    id: "2",
    label: "平面设计",
    icon: Folder,
    children: [
      { id: "2-1", label: "图标", icon: Picture },
      { id: "2-2", label: "宣传海报", icon: Picture }
    ]
  },
  {
    id: "3",
    label: "办公文档",
    icon: Document,
    children: [{ id: "3-1", label: "项目报告", icon: Document }]
  },
  {
    id: "4",
    label: "相机素材",
    icon: Camera
  }
];

const searchKeyword = ref("");
const selectedTag = ref("");
const activeTab = ref("image");

const tagOptions = [
  { label: "风景", value: "风景" },
  { label: "建筑", value: "建筑" },
  { label: "人物", value: "人物" },
  { label: "花卉", value: "花卉" },
  { label: "动物", value: "动物" }
];

const dialogVisible = ref(false);
const dialogType = ref(1);
const currentItem = ref<FodderItem | null>(null);

const mockImages: FodderItem[] = [
  {
    id: 1,
    name: "龟山电视塔.JPG",
    thumbnail: "https://picsum.photos/300/200?random=1",
    type: "图片",
    time: "2026-03-29 08:12:07",
    category: "摄影图片"
  },
  {
    id: 2,
    name: "风景图片4.JPG",
    thumbnail: "https://picsum.photos/300/200?random=2",
    type: "图片",
    time: "2026-03-27 07:53:48",
    category: "摄影图片"
  },
  {
    id: 3,
    name: "风景照片3.JPG",
    thumbnail: "https://picsum.photos/300/200?random=3",
    type: "图片",
    time: "2026-03-27 10:25:38",
    category: "摄影图片"
  },
  {
    id: 4,
    name: "风景图片2.JPG",
    thumbnail: "https://picsum.photos/300/200?random=4",
    type: "图片",
    time: "2026-03-25 05:27:46",
    category: "摄影图片"
  },
  {
    id: 5,
    name: "革命烈士纪念馆",
    thumbnail: "https://picsum.photos/300/200?random=5",
    type: "图片",
    time: "2026-04-14 03:24:54",
    category: "摄影图片"
  },
  {
    id: 6,
    name: "荷花特写6.JPG",
    thumbnail: "https://picsum.photos/300/200?random=6",
    type: "图片",
    time: "2026-04-14 03:24:54",
    category: "摄影图片"
  },
  {
    id: 7,
    name: "荷花特写5.JPG",
    thumbnail: "https://picsum.photos/300/200?random=7",
    type: "图片",
    time: "2026-04-14 03:24:54",
    category: "摄影图片"
  },
  {
    id: 8,
    name: "荷花特写4.JPG",
    thumbnail: "https://picsum.photos/300/200?random=8",
    type: "图片",
    time: "2026-04-14 03:24:54",
    category: "摄影图片"
  }
];

const mockVideos: FodderItem[] = [
  {
    id: 101,
    name: "高清视频1.mp4",
    thumbnail: "https://picsum.photos/300/200?random=11",
    type: "视频",
    time: "2026-03-29 08:12:07",
    category: "高清视频"
  },
  {
    id: 102,
    name: "高清视频2.mp4",
    thumbnail: "https://picsum.photos/300/200?random=12",
    type: "视频",
    time: "2026-03-27 07:53:48",
    category: "高清视频"
  }
];

const filteredImages = computed(() => {
  let result = mockImages;
  if (searchKeyword.value) {
    result = result.filter(item => item.name.includes(searchKeyword.value));
  }
  return result;
});

const filteredVideos = computed(() => {
  let result = mockVideos;
  if (searchKeyword.value) {
    result = result.filter(item => item.name.includes(searchKeyword.value));
  }
  return result;
});

const handleNodeClick = (data: TreeNode) => {
  currentNodeKey.value = data.id;
  console.log("点击节点:", data.label);
};

const handleSearch = () => {
  console.log("搜索:", searchKeyword.value);
};

const handleTabClick = () => {
  console.log("切换标签页:", activeTab.value);
};

const handleAdd = () => {
  dialogType.value = 1;
  currentItem.value = null;
  dialogVisible.value = true;
};

const handleBatchImport = () => {
  console.log("批量导入");
};

const handleCardClick = (item: FodderItem) => {
  console.log("点击卡片:", item.name);
};

const handleRefresh = () => {
  console.log("刷新列表");
};
</script>

<style scoped>
.fodder-container {
  display: flex;
  height: 100vh;
  overflow: hidden;
  background-color: #f5f7fa;
}
.sidebar {
  width: 220px;
  padding: 16px 0;
  overflow-y: auto;
  background-color: #ffffff;
  border-right: 1px solid #e4e7ed;
}
.tree-node-label {
  display: flex;
  gap: 6px;
  align-items: center;
}
.node-icon {
  font-size: 16px;
}
.content-area {
  flex: 1;
  padding: 20px 24px;
  overflow-y: auto;
  background-color: #ffffff;
}
.breadcrumb {
  margin-bottom: 20px;
}
.toolbar {
  display: flex;
  gap: 16px;
  align-items: center;
  padding-bottom: 20px;
  margin-bottom: 20px;
  border-bottom: 1px solid #e4e7ed;
}
.search-input {
  width: 300px;
}
.tag-select {
  width: 200px;
}
.action-btn {
  background-color: #009688;
  border-color: #009688;
}
.action-btn:hover {
  background-color: #00796b;
  border-color: #00796b;
}
.content-tabs {
  margin-top: 16px;
}
.image-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(240px, 1fr));
  gap: 20px;
  padding: 16px 0;
}
.image-card {
  overflow: hidden;
  cursor: pointer;
  background-color: #ffffff;
  border: 1px solid #e4e7ed;
  border-radius: 8px;
  transition: all 0.3s;
}
.image-card:hover {
  box-shadow: 0 4px 12px rgb(0 0 0 / 10%);
  transform: translateY(-2px);
}
.image-wrapper {
  position: relative;
  width: 100%;
  height: 160px;
  overflow: hidden;
  background-color: #f5f7fa;
}
.card-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}
.video-overlay {
  position: absolute;
  inset: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: rgb(0 0 0 / 30%);
}
.play-icon {
  font-size: 48px;
  color: #ffffff;
}
.card-info {
  padding: 12px;
}
.card-name {
  margin-bottom: 8px;
  overflow: hidden;
  font-size: 14px;
  color: #303133;
  text-overflow: ellipsis;
  white-space: nowrap;
}
.card-meta {
  display: flex;
  gap: 8px;
  align-items: center;
  justify-content: space-between;
}
.card-time {
  font-size: 12px;
  color: #909399;
}
:deep(.el-tabs__header) {
  margin-bottom: 0;
}
:deep(.el-tabs__nav-wrap::after) {
  height: 1px;
}
</style>
