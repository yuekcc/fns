<template>
  <div class="page">
    <el-page-header
      content="Overview"
      style="margin-bottom: 16px"
    ></el-page-header>
    <div class="data-cards">
      <DataCard title="集群"></DataCard>
      <DataCard title="Service"></DataCard>
      <DataCard title="Function"></DataCard>
      <CpuUsageCard></CpuUsageCard>
      <MemUsageCard></MemUsageCard>
    </div>
  </div>
</template>

<style scoped>
.data-cards {
  display: grid;
  gap: 16px;
  grid-template-columns: 1fr 1fr 1fr;
}
</style>

<script>
import { defineComponent } from "vue";
import DataCard from "../../components/data-card.vue";
import CpuUsageCard from "./cpu-usage-card.vue";
import MemUsageCard from './mem-usage-card.vue'

export default defineComponent({
  components: {
    DataCard,
    CpuUsageCard,
    MemUsageCard,
  },
  setup() {},
});
</script>