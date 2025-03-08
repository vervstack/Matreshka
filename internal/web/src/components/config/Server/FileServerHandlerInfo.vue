<script setup lang="ts">

import {FsHandler} from "@/models/AppConfig/Servers/Servers.ts";
import ConfigField from "@/components/base/ConfigField/ConfigInput.vue";

const fsInfo = defineModel<FsHandler>({required: true})

</script>

<template>
  <div class="NodeField">
    <ConfigField
        v-model="fsInfo.dist"
        field-name="Dist folder:"
    />
  </div>
</template>

<style scoped>
@import "@/assets/styles/config_display.css";

</style>
