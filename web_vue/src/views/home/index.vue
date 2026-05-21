<template>
  <div class="medilink-home">
    <el-row :gutter="16" class="medilink-home__stats">
      <el-col v-for="item in statCards" :key="item.title" :xs="24" :sm="12" :lg="6">
        <el-card shadow="hover" class="medilink-home__stat-card">
          <div class="medilink-home__stat-body">
            <div class="medilink-home__stat-text">
              <span class="medilink-home__stat-label">{{ item.title }}</span>
              <el-input
                v-model.number="item.value"
                type="number"
                size="small"
                class="medilink-home__stat-input"
                placeholder="请输入数字"
              />
            </div>
            <div class="medilink-home__stat-icon" :style="{ background: item.iconBg, color: item.iconColor }">
              <el-icon :size="26">
                <component :is="item.icon" />
              </el-icon>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <el-card shadow="hover" class="medilink-home__chart-card">
      <template #header>
        <span class="medilink-home__chart-title">近七日素材上传趋势</span>
      </template>
      <ECharts :option="chartOption" :height="320" />
    </el-card>
  </div>
</template>

<script setup lang="ts" name="home">
import { ref, watchEffect } from "vue";
import { FolderOpened, Document, Calendar, PriceTag } from "@element-plus/icons-vue";
import ECharts from "@/components/ECharts/index.vue";
import type { ECOption } from "@/components/ECharts/config";

const statCards = ref([
  {
    title: "素材总数",
    value: 100,
    icon: FolderOpened,
    iconBg: "rgba(0, 150, 136, 0.12)",
    iconColor: "#009688",
    valueClass: ""
  },
  {
    title: "待审核任务",
    value: 6,
    icon: Document,
    iconBg: "rgba(64, 158, 255, 0.12)",
    iconColor: "#409eff",
    valueClass: "is-warn"
  },
  {
    title: "今日新增",
    value: 0,
    icon: Calendar,
    iconBg: "rgba(245, 108, 108, 0.12)",
    iconColor: "#f56c6c",
    valueClass: ""
  },
  {
    title: "已发布素材",
    value: 78,
    icon: PriceTag,
    iconBg: "rgba(103, 194, 58, 0.12)",
    iconColor: "#67c23a",
    valueClass: ""
  }
]);

const chartOption = ref<ECOption>({
  color: ["#009688", "#409eff"],
  tooltip: { trigger: "axis" },
  legend: {
    data: ["图片", "视频"],
    bottom: 0
  },
  grid: { left: "3%", right: "4%", top: "12%", bottom: "14%", containLabel: true },
  xAxis: {
    type: "category",
    boundaryGap: false,
    data: ["04-11", "04-12", "04-13", "04-14", "04-15", "04-16", "04-17"]
  },
  yAxis: {
    type: "value",
    min: 0,
    splitLine: { lineStyle: { type: "dashed", opacity: 0.35 } }
  },
  series: [
    {
      name: "图片",
      type: "line",
      smooth: true,
      symbol: "circle",
      symbolSize: 6,
      data: [2, 5, 3, 8, 4, 6, 7]
    },
    {
      name: "视频",
      type: "line",
      smooth: true,
      symbol: "circle",
      symbolSize: 6,
      data: [0, 1, 0, 2, 1, 1, 2]
    }
  ]
});

watchEffect(() => {
  const values = statCards.value.map(item => Number(item.value));
  const [total, review, today, published] = values;
  if (!chartOption.value.series) return;

  chartOption.value.series[0].data = [total, review, today, published, total, review, today];
  chartOption.value.series[1].data = [published, today, review, total, published, today, review];
});
</script>

<style scoped lang="scss">
.medilink-home {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.medilink-home__stats {
  margin: 0 !important;
}

.medilink-home__stat-card {
  border-radius: 10px;
  :deep(.el-card__body) {
    padding: 18px 20px;
  }
}

.medilink-home__stat-body {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
}

.medilink-home__stat-label {
  display: block;
  margin-bottom: 6px;
  font-size: 14px;
  color: var(--el-text-color-secondary);
}

.medilink-home__stat-value {
  font-size: 26px;
  font-weight: 700;
  color: var(--el-text-color-primary);
  &.is-warn {
    color: var(--el-color-warning);
  }
}

.medilink-home__stat-input {
  width: 100%;
  max-width: 140px;
}

.medilink-home__stat-icon {
  display: flex;
  flex-shrink: 0;
  align-items: center;
  justify-content: center;
  width: 48px;
  height: 60px;
  border-radius: 10px;
}

.medilink-home__chart-card {
  flex: 1;
  border-radius: 10px;
  :deep(.el-card__header) {
    padding: 14px 20px;
    font-weight: 600;
    border-bottom: 1px solid var(--el-border-color-lighter);
  }
  :deep(.el-card__body) {
    padding: 8px 16px 16px;
  }
}

.medilink-home__chart-title {
  font-size: 15px;
}
</style>
