<script setup lang="ts">
import ConfigField from "@/components/base/config/fields/ConfigInput.vue";
import ServiceName from "@/components/config/verv/app_info/ServiceName.vue";
import { AppInfoClass } from "@/models/configs/verv/info/VervConfig.ts";

const info = defineModel<AppInfoClass>({ required: true });
</script>

<template>
  <div v-if="!info">No app info</div>

  <div v-else class="Node">
    <ServiceName v-model="info.name" />

    <ConfigField v-model="info.serviceVersion" field-name="Version" />
  </div>
</template>

<style scoped>
@import "@/assets/styles/config_display.css";
</style>
