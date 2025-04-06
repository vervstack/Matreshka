<script setup lang="ts">

import ConfigField from "@/components/base/config/fields/ConfigInput.vue";

import {Telegram} from "@/models/configs/verv/Resources/Resource.ts";

const model = defineModel<Telegram>({required: true})
</script>

<template>
  <div class="NodeField">
    <ConfigField
        v-model="model.apiKey"
        field-name="Api Key"
    />
  </div>
</template>

<style scoped>
@import "@/assets/styles/config_display.css";

</style>
