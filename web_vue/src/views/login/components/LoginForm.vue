<template>
  <div class="login-form-inner">
    <el-form ref="loginFormRef" :model="loginForm" :rules="loginRules" size="large">
      <el-form-item prop="username">
        <el-input v-model="loginForm.username" placeholder="用户名:">
          <template #prefix>
            <el-icon class="el-input__icon">
              <User />
            </el-icon>
          </template>
        </el-input>
      </el-form-item>
      <el-form-item prop="password">
        <el-input v-model="loginForm.password" type="password" placeholder="密码:" show-password autocomplete="new-password">
          <template #prefix>
            <el-icon class="el-input__icon">
              <Lock />
            </el-icon>
          </template>
        </el-input>
      </el-form-item>
    </el-form>
    <div class="login-btn">
      <el-button :icon="RefreshRight" round size="large" @click="resetForm(loginFormRef)">重置</el-button>
      <el-button :icon="UserFilled" round size="large" type="primary" :loading="loading" @click="login(loginFormRef)">
        登录
      </el-button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, onBeforeUnmount } from "vue";
import { useRouter } from "vue-router";
import { HOME_URL } from "@/config";
import { Login } from "@/api/interface";
import { ElNotification } from "element-plus";
import { loginApi } from "@/api/modules/login";
import { useUserStore } from "@/stores/modules/user";
import { useTabsStore } from "@/stores/modules/tabs";
import { useKeepAliveStore } from "@/stores/modules/keepAlive";
import { initDynamicRouter } from "@/routers/modules/dynamicRouter";
import { UserFilled, RefreshRight, User, Lock } from "@element-plus/icons-vue";
import type { ElForm } from "element-plus";
import md5 from "md5";

const router = useRouter();
const userStore = useUserStore();
const tabsStore = useTabsStore();
const keepAliveStore = useKeepAliveStore();

type FormInstance = InstanceType<typeof ElForm>;
const loginFormRef = ref<FormInstance>();
const loginRules = reactive({
  username: [{ required: true, message: "请输入用户名", trigger: "blur" }],
  password: [{ required: true, message: "请输入密码", trigger: "blur" }]
});

const loading = ref(false);
const loginForm = reactive<Login.ReqLoginForm>({
  username: "",
  password: ""
});

const login = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.validate(async valid => {
    if (!valid) return;
    loading.value = true;
    try {
      const { data } = await loginApi({ ...loginForm, password: md5(loginForm.password) });
      userStore.setToken(data.access_token);
      await initDynamicRouter();
      tabsStore.setTabs([]);
      keepAliveStore.setKeepAliveName([]);
      router.push(HOME_URL);
      ElNotification({
        title: "登录成功",
        message: "欢迎使用 MediLink",
        type: "success",
        duration: 2800
      });
    } finally {
      loading.value = false;
    }
  });
};

const resetForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.resetFields();
};

onMounted(() => {
  document.onkeydown = (e: KeyboardEvent) => {
    if (e.code === "Enter" || e.code === "enter" || e.code === "NumpadEnter") {
      if (loading.value) return;
      login(loginFormRef.value);
    }
  };
});

onBeforeUnmount(() => {
  document.onkeydown = null;
});
</script>

<style scoped lang="scss">
.login-form-inner {
  width: 100%;
}
</style>
