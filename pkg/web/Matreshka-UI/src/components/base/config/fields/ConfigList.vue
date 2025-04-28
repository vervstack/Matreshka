<script setup lang="ts" generic="T">
import { Nullable } from "@primevue/core";
import Button from "primevue/button";
import InputText from "primevue/inputtext";

import { ConfigValue } from "@/models/shared/Values.ts";

const model = defineModel<ConfigValue<T[]>>({ required: true });
</script>

<template>
  <div class="Node">
    <div class="TopControls">
      <p>{{ model.envName }}</p>
      <Button rounded icon="pi pi-plus" />
    </div>
    <div class="Node" v-for="(_, idx) in model.value">
      <InputText v-model="model.value[idx] as Nullable<string>" />
    </div>
  </div>
</template>

<style scoped>
@import "@/assets/styles/config_display.css";

.TopControls {
  display: flex;
  flex-direction: row;
  align-items: center;
  gap: 2em;
}
</style>
