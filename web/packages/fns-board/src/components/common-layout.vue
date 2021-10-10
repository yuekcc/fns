<template>
  <div>
    <MainNavigationMenu style="position: sticky"></MainNavigationMenu>
    <el-container>
      <el-main>
        <router-view></router-view>
      </el-main>
    </el-container>
  </div>
</template>

<style scoped>
::v-deep .el-header {
  padding: 0;
}
</style>

<script>
import { defineComponent } from "vue";
import MainNavigationMenu from "./main-navigation-menu.vue";

export default defineComponent({
  components: {
    MainNavigationMenu,
  },
  setup() {},
});
</script>