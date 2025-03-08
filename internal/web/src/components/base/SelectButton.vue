<script setup lang="ts">

import SelectButton from "primevue/selectbutton";

const model = defineModel<string>()
const props = defineProps({
  options: {
    type: Array<string>,
    default: []
  }
})

</script>

<template>
  <SelectButton
      v-model="model"
      :options="props.options"

  />
</template>

<style>
.p-togglebutton {
  font-size: 1em;
}
</style>
