<template>
  <div class="card bg-white border">
    <div class="flex" style="justify-content: space-between">
      <div class="flex-initial">{{ title }}</div>
      <div class="flex-initial">
        <slot name="action"></slot>
      </div>
    </div>
    <div class="flex">
      <slot></slot>
    </div>
  </div>
</template>

<style scoped>
.card {
  height: 380px;
  padding: 16px;
  border-radius: 8px;
}
</style>

<script>
import { defineComponent } from "vue";

export default defineComponent({
  props: {
    title: String
  },
  setup() {},
});
</script>