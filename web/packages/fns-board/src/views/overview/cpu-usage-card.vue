<template>
  <DataCard title="CPU 使用率">
    <Chart :config="config" :updated-data="updatedData"></Chart>
  </DataCard>
</template>

<script>
import {
  defineComponent,
  reactive,
  onMounted,
  onBeforeMount,
  onBeforeUnmount,
} from "vue";
import Chart from "../../components/chart.vue";
import DataCard from "../../components/data-card.vue";

export default defineComponent({
  components: {
    Chart,
    DataCard,
  },
  setup() {
    const config = {
      type: "line",
      data: {
        labels: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10],
        datasets: [
          {
            id: 'cpu_usage',
            label: "CPU 使用率",
            data: [0],
            borderColor: "#ff6384",
            backgroundColor: "#ffb1c1",
          },
        ],
      },
    };

    const updatedData = reactive({
      cpu_usage: []
    })

    let timer;
    onMounted(() => {
      timer = setInterval(() => {
        const num = Math.random() * 100;
        console.log("update cpu usage --> ", num);
        updatedData.cpu_usage = [num]
      }, 1000);
    });

    onBeforeUnmount(() => {
      clearInterval(timer);
    });

    return {
      config,
      updatedData,
    };
  },
});
</script>