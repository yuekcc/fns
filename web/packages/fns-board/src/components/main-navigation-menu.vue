<template>
  <div
    class="flex border-bottom bg-white"
    style="align-items: center; padding: 0 32px"
  >
    <div class="flex-initial" style="margin-right: 16px">
      <img class="logo" src="/logo.png" />
    </div>
    <div>
      <el-menu default-active="overview" mode="horizontal" @select="onSelect">
        <el-submenu index="quick_start">
          <template #title>快速开始</template>
          <el-menu-item index="create_service">创建 Service</el-menu-item>
          <el-menu-item index="create_function">创建 Function</el-menu-item>
        </el-submenu>
        <el-menu-item index="overview">总览</el-menu-item>
        <el-menu-item index="service_list">服务</el-menu-item>
        <el-menu-item index="function_list">函数</el-menu-item>
        <el-menu-item index="system_setup">设置</el-menu-item>
      </el-menu>
    </div>
  </div>
</template>

<style scoped>
.logo {
  width: 32px;
  height: 32px;
}

::v-deep .el-menu.el-menu--horizontal {
  border: 0;
}
</style>

<script>
import { router } from "../router";

const routesMapping = {
  overview: "overview",
  service_list: "services",
  create_service: "create_service",
  function_list: "functions",
  create_function: "create_function",
  system_setup: "system_setup",
};

export default {
  setup() {
    function onSelect(menuIndex) {
      router.push({ name: menuIndex });
    }

    return {
      onSelect,
    };
  },
};
</script>