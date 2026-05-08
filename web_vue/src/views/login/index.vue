<template>
  <div class="medilink-login-standalone">
    <SwitchDark class="medilink-login-standalone__dark" />
    <div class="medilink-login-standalone__panel">
      <div class="medilink-login-standalone__header">
        <el-icon class="medilink-login-standalone__brand" :size="28">
          <Camera />
        </el-icon>
        <h1 class="medilink-login-standalone__title">MediLink</h1>
      </div>
      <LoginForm />
    </div>
    <p class="medilink-login-standalone__copyright">2026 © MediLink By Xiong XinRui.</p>
  </div>
</template>

<script setup lang="ts" name="login">
import { Camera } from "@element-plus/icons-vue";
import LoginForm from "./components/LoginForm.vue";
import SwitchDark from "@/components/SwitchDark/index.vue";
</script>

<style scoped lang="scss">
.medilink-login-standalone {
  position: relative;
  box-sizing: border-box;
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  padding: 32px 16px 56px;
  background: #1a1a1a url("@/assets/images/medilink-ai-bg.png") center / cover no-repeat;
}

.medilink-login-standalone__dark {
  position: absolute;
  top: 20px;
  right: 20px;
  z-index: 2;
}

.medilink-login-standalone__panel {
  box-sizing: border-box;
  width: 100%;
  max-width: 440px;
  padding: 40px 40px 36px;
  background: rgba(255, 255, 255, 0.72);
  border: 1px solid rgba(255, 255, 255, 0.65);
  border-radius: 16px;
  box-shadow: 0 24px 48px rgba(0, 0, 0, 0.25);
  backdrop-filter: blur(14px);
}

.medilink-login-standalone__header {
  display: flex;
  gap: 12px;
  align-items: center;
  justify-content: center;
  margin-bottom: 28px;
}

.medilink-login-standalone__brand {
  color: #1a2b3c;
}

.medilink-login-standalone__title {
  margin: 0;
  font-size: 28px;
  font-weight: 700;
  color: #1a2b3c;
  letter-spacing: 0.02em;
}

.medilink-login-standalone__copyright {
  position: absolute;
  bottom: 20px;
  left: 50%;
  margin: 0;
  font-size: 12px;
  color: rgba(255, 255, 255, 0.82);
  text-align: center;
  text-shadow: 0 1px 4px rgba(0, 0, 0, 0.5);
  transform: translateX(-50%);
}

:deep(.login-form-inner .el-form-item) {
  margin-bottom: 18px;
}

:deep(.login-form-inner .el-input__wrapper) {
  background-color: #ffffff;
  border-radius: 8px;
  box-shadow: 0 0 0 1px rgba(0, 0, 0, 0.06) inset;
}

:deep(.login-form-inner .login-btn) {
  display: flex;
  gap: 14px;
  margin-top: 8px;
}

:deep(.login-form-inner .login-btn .el-button) {
  flex: 1;
  font-weight: 600;
}

:deep(.login-form-inner .login-btn .el-button--default) {
  color: #1a2b3c;
  background: #ffffff;
  border-color: rgba(0, 0, 0, 0.12);
}

:deep(.login-form-inner .login-btn .el-button--primary) {
  color: #ffffff;
  background: #00a18c;
  border-color: #00a18c;
}

:deep(.login-form-inner .login-btn .el-button--primary:hover),
:deep(.login-form-inner .login-btn .el-button--primary:focus) {
  color: #ffffff;
  background: #008f7c;
  border-color: #008f7c;
}
</style>
