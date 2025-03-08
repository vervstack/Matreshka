<script setup lang="ts">

import ConfigField from "@/components/base/ConfigField/ConfigInput.vue";
import {GrpcHandler} from "@/models/AppConfig/Servers/Servers.ts";

const grpcInfo = defineModel<GrpcHandler>({required: true})

</script>

<template>
  <div class="NodeField">
    <ConfigField
        v-model="grpcInfo.module"
        field-name="Module"
    />
  </div>

  <div class="NodeField">
    <ConfigField
        v-model="grpcInfo.gateway"
        field-name="Gateway"
    />
  </div>
</template>

<style scoped>
@import "@/assets/styles/config_display.css";

</style>
