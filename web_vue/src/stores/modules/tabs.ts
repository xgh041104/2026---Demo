import router from "@/routers";
import { defineStore } from "pinia";
import { getUrlWithParams } from "@/utils";
import { useKeepAliveStore } from "./keepAlive";
import { useAuthStore } from "@/stores/modules/auth";
import { TabsState, TabsMenuProps } from "@/stores/interface";
import piniaPersistConfig from "@/stores/helper/persist";

const getKeepAliveStore = () => useKeepAliveStore();

/** 与菜单 path 对齐（兼容 hash 模式下 #/path） */
const tabPathKey = (path: string) => {
  const noQuery = path.split("?")[0];
  if (noQuery.startsWith("#")) return noQuery.slice(1) || "/";
  return noQuery;
};

export const useTabsStore = defineStore({
  id: "geeker-tabs",
  state: (): TabsState => ({
    tabsMenuList: []
  }),
  actions: {
    // Add Tabs
    async addTabs(tabItem: TabsMenuProps) {
      if (this.tabsMenuList.every(item => item.path !== tabItem.path)) {
        this.tabsMenuList.push(tabItem);
      }
      // add keepalive
      if (!getKeepAliveStore().keepAliveName.includes(tabItem.name) && tabItem.isKeepAlive) {
        getKeepAliveStore().addKeepAliveName(tabItem.path);
      }
    },
    // Remove Tabs
    async removeTabs(tabPath: string, isCurrent: boolean = true) {
      if (isCurrent) {
        this.tabsMenuList.forEach((item, index) => {
          if (item.path !== tabPath) return;
          const nextTab = this.tabsMenuList[index + 1] || this.tabsMenuList[index - 1];
          if (!nextTab) return;
          router.push(nextTab.path);
        });
      }
      // remove keepalive
      const tabItem = this.tabsMenuList.find(item => item.path === tabPath);
      if (tabItem?.isKeepAlive) getKeepAliveStore().removeKeepAliveName(tabItem.path);
      // set tabs
      this.tabsMenuList = this.tabsMenuList.filter(item => item.path !== tabPath);
    },
    // Close Tabs On Side
    async closeTabsOnSide(path: string, type: "left" | "right") {
      const currentIndex = this.tabsMenuList.findIndex(item => item.path === path);
      if (currentIndex !== -1) {
        const range = type === "left" ? [0, currentIndex] : [currentIndex + 1, this.tabsMenuList.length];
        this.tabsMenuList = this.tabsMenuList.filter((item, index) => {
          return index < range[0] || index >= range[1] || !item.close;
        });
      }
      // set keepalive
      const KeepAliveList = this.tabsMenuList.filter(item => item.isKeepAlive);
      getKeepAliveStore().setKeepAliveName(KeepAliveList.map(item => item.path));
    },
    // Close MultipleTab
    async closeMultipleTab(tabsMenuValue?: string) {
      this.tabsMenuList = this.tabsMenuList.filter(item => {
        return item.path === tabsMenuValue || !item.close;
      });
      // set keepalive
      const KeepAliveList = this.tabsMenuList.filter(item => item.isKeepAlive);
      getKeepAliveStore().setKeepAliveName(KeepAliveList.map(item => item.path));
    },
    // Set Tabs
    async setTabs(tabsMenuList: TabsMenuProps[]) {
      this.tabsMenuList = tabsMenuList;
    },
    // Set Tabs Title
    async setTabsTitle(title: string) {
      this.tabsMenuList.forEach(item => {
        if (item.path == getUrlWithParams()) item.title = title;
      });
    },
    /**
     * 移除当前权限菜单中不存在的标签（例如菜单改版后残留的「登录界面」等）
     */
    pruneTabsToAuthMenu() {
      const authStore = useAuthStore();
      const allowed = new Set(authStore.flatMenuListGet.map(m => m.path));
      const next = this.tabsMenuList.filter(tab => allowed.has(tabPathKey(tab.path)));
      this.tabsMenuList = next;
      const ka = next.filter(t => t.isKeepAlive).map(t => t.path);
      getKeepAliveStore().setKeepAliveName(ka);
    }
  },
  persist: piniaPersistConfig("geeker-tabs")
});
