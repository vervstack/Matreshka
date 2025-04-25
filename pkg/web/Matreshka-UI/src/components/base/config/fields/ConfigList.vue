<script
    setup
    lang="ts"
    generic="T"
>

import {ConfigValue} from "@/models/shared/common.ts";
import InputText from 'primevue/inputtext';
import Button from 'primevue/button';
import {Nullable} from "@primevue/core";

const model = defineModel<ConfigValue<T[]>>({required: true})

</script>

<template>
  <div class="Node">

    <div class="TopControls">
      <p>{{ model.envName }}</p>
      <Button
          rounded
          icon="pi pi-plus"
      />
    </div>
    <div
        class="Node"
        v-for="(_, idx) in model.value"
    >
      <InputText
          v-model="model.value[idx] as Nullable<string>"
      />
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
