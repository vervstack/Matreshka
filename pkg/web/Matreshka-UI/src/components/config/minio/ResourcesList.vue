<script
    setup
    lang="ts"
>

import {ConfigValue} from "@/models/shared/common.ts";
import Button from 'primevue/button';
import InputGroup from 'primevue/inputgroup';
import InputGroupAddon from 'primevue/inputgroupaddon';
import InputText from 'primevue/inputtext';
import {Nullable} from "@primevue/core";

const model = defineModel<ConfigValue<string[]>>({required: true})


function addResource() {
  model.value.value.push('')
}

function deleteResource(deleteIdx: number) {
  model.value.value.splice(deleteIdx, 1)
}

</script>

<template>
  <div class="Node">
    <div class="TopControls">
      <p>{{ model.envName }}</p>
      <Button
          rounded
          icon="pi pi-plus"
          v-tooltip.right="'Add resource'"
          @click="addResource"
      />
    </div>

    <div class="Node">
      <InputGroup
          v-for="(_, idx) in model.value"
      >
        <InputGroupAddon>
          <div>arn:aws:s3:::</div>
        </InputGroupAddon>
        <InputText
            v-model="model.value[idx] as Nullable<string>"
        />
        <Button
            icon="pi pi-trash"
            v-tooltip.top="'Delete resource'"
            @click="()=>deleteResource(idx)"
        />
      </InputGroup>
    </div>
  </div>
</template>

<style>
@import "@/assets/styles/config_display.css";

.TopControls {
  display: flex;
  flex-direction: row;
  align-items: center;
  gap: 2em;
}

</style>
