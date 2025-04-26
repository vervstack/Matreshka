<script setup lang="ts">
import { ConfigTypePrefix } from "@vervstack/matreshka";
import { PropType } from "vue";

import getConfigIcon from "@/models/configs/Icons.ts";

defineProps({
  configType: {
    type: String as PropType<ConfigTypePrefix>,
    required: true,
  },
});
</script>

<template>
  <img class="config-icon" :src="getConfigIcon(configType)" :alt="configType" />
</template>

<style scoped>
.config-icon {
  height: 100%;
  width: 100%;
}
</style>
