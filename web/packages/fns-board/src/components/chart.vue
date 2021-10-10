<template>
  <canvas ref="chartContainerRef"></canvas>
</template>

<script>
import { defineComponent, ref, onMounted, watch, onBeforeUnmount } from "vue";
import Chart from "chart.js/auto";
import { get } from "lodash-es";

export default defineComponent({
  props: {
    config: Object,
    updatedData: Object,
  },
  setup(props) {
    const chartContainerRef = ref();
    let theChart = null;

    watch(
      () => props.updatedData,
      (newValue) => {
        if (theChart) {
          const datasets = theChart.data.datasets;
          const labels = theChart.data.labels;

          Object.keys(newValue).forEach((datasetId) => {
            const dataset = datasets.find((it) => it.id === datasetId);

            if (!dataset) {
              return;
            }

            if (dataset.data.length >= labels.length) {
              dataset.data = dataset.data.slice(labels.length - 1);
            }

            dataset.data.push(...newValue[datasetId]);
          });

          theChart.update();
        }
      },
      {
        deep: true,
      }
    );

    onMounted(() => {
      if (chartContainerRef.value) {
        const config = {
          ...props.config,
        };

        theChart = new Chart(chartContainerRef.value, config);
      }
    });

    onBeforeUnmount(() => {
      if (theChart) {
        theChart.destroy();
      }
    });

    return {
      chartContainerRef,
    };
  },
});
</script>