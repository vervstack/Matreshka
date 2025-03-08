<script setup lang="ts">

import {AppInfoClass} from "@/models/AppConfig/Info/AppInfo.ts";
import ConfigField from "@/components/base/ConfigField/ConfigInput.vue";
import ServiceName from "@/components/config/AppInfo/ServiceName.vue";

const info = defineModel<AppInfoClass>({required: true})

</script>

<template>
  <div v-if="!info">No app info</div>

  <div v-else class="Node">
        <ServiceName
            v-model="info.name"
        />

        <ConfigField
            v-model="info.serviceVersion"
            field-name="Version"
        />
  </div>

</template>

<style scoped>
@import "@/assets/styles/config_display.css";

</style>
