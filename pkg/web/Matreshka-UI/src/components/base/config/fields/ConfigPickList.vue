<script setup lang="ts" generic="T extends any, B extends ConfigValue<T[]>">
import PickList from "primevue/picklist";
import { ref } from "vue";

import { ConfigValue } from "@/models/shared/common.ts";

const model = defineModel<B>({ required: true });

const props = defineProps({
  options: {
    type: Array as () => T[],
    required: true,
  },
});
const options = props.options.filter((v) => !model.value.value.includes(v));
const fromToList = ref<any[][]>([options, model.value.value]);
</script>

<template>
  <div class="Node">
    <div>{{ model.envName }}</div>
    <div class="Node">
      <PickList
        v-model="fromToList"
        :responsive="false"
        :show-source-controls="false"
        :show-target-controls="false"
      />
    </div>
  </div>
</template>

<style scoped>
@import "@/assets/styles/config_display.css";
</style>
