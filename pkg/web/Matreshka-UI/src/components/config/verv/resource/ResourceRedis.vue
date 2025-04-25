<script setup lang="ts">

import ConfigField from "@/components/base/config/fields/ConfigInput.vue";
import {Redis} from "@/models/configs/verv/resources/Resource.ts";

const model = defineModel<Redis>(
    {
      required: true,
    })

</script>

<template>
  <div class="NodeField">
    <ConfigField
        v-model="model.host"
        field-name="Host"
    />
  </div>

  <div class="NodeField">
    <ConfigField
        v-model="model.port"
        field-name="Port"
    />
  </div>

  <div class="NodeField">
    <ConfigField
        v-model="model.user"
        field-name="User"
    />
  </div>

  <div class="NodeField">
    <ConfigField
        v-model="model.pwd"
        field-name="Password"
    />
  </div>

  <div class="NodeField">
    <ConfigField
        v-model="model.db"
        field-name="Db Number"
    />
  </div>
</template>

<style scoped>
@import "@/assets/styles/config_display.css";

</style>
