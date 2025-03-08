<script setup lang="ts">


import {S3Action} from "@/models/resource_configs/s3/minio/minio.ts";
import {ConfigValueClass} from "@/models/shared/common.ts";

import Chip from 'primevue/chip';
import Select from 'primevue/select';
import {ref} from "vue";

const model = defineModel<ConfigValueClass<S3Action[]>>({required: true})
const lastSelected = ref<string>()

function select(v: S3Action) {
  lastSelected.value = ""

  if (model.value.value.includes(v)) {
    return
  }

  model.value.value.push(v)
}

function remove(idx: number) {
  model.value.value.splice(idx, 1)
}


</script>

<template>
  <div class="Node">
    Actions
    <Select
        v-model="lastSelected"
        :options="Object.values(S3Action)"
        @update:modelValue="select"
    />
    <div class="TopChips">
      <Chip
          v-for="(_, idx) in model.value"
          :key="model.value[idx].toString()"
          :label="model.value[idx].toString()"
          @remove="()=>remove(idx)"
          removable
      />
    </div>

  </div>


</template>

<style scoped>
.TopChips {
  width: 100%;
  max-height: 100%;

  display: flex;
  flex-direction: row;
  overflow-x: scroll;

  gap: 0.5em;
}
</style>
