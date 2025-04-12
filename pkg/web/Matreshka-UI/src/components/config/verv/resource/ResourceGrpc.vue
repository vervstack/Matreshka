<script setup lang="ts">

import {GrpcClient} from "@/models/configs/verv/Resources/Resource.ts";

import ConfigField from "@/components/base/config/fields/ConfigInput.vue";

const model = defineModel<GrpcClient>({required: true})
</script>

<template>
    <div class="NodeField">
      <ConfigField
          v-model="model.connectionString"
          field-name="Service url"
      />
    </div>

    <div class="NodeField">
      <ConfigField
          v-model="model.module"
          field-name="Module"
      />
    </div>
</template>

<style scoped>
@import "@/assets/styles/config_display.css";
</style>
