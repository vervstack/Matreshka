<script setup lang="ts">
import ConfigField from "@/components/base/config/fields/ConfigInput.vue";
import { Postgres } from "@/models/configs/verv/resources/Resource.ts";

const model = defineModel<Postgres>({
  required: true,
});
</script>

<template>
  <div class="NodeField">
    <ConfigField v-model="model.name" field-name="Database name" />
  </div>

  <div class="NodeField">
    <ConfigField v-model="model.host" field-name="Host" />
  </div>

  <div class="NodeField">
    <ConfigField v-model="model.port" field-name="Port" />
  </div>

  <div class="NodeField">
    <ConfigField v-model="model.user" field-name="Username" />
  </div>

  <div class="NodeField">
    <ConfigField v-model="model.pwd" field-name="Password" />
  </div>
  <div class="NodeField">
    <ConfigField v-model="model.ssl_mode" field-name="Ssl mode" />
  </div>
</template>

<style scoped>
@import "@/assets/styles/config_display.css";
</style>
