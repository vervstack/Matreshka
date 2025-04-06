<script setup lang="ts">
import Inputer from "@/components/base/config/Inputer.vue";
import {keyMap} from "@/models/shared/common.ts";

const model = defineModel<keyMap>({required: true})

const flatValues: (string | number)[][] = []
const innerNodes: keyMap[] = []

for (const key in model.value.val) {
  const val = model.value.val[key];

  const valType = typeof val;
  switch (valType) {
    case "object":
      const ur = {
        fieldName: key
      } as keyMap

      for (const innerKey in val) {
        ur[innerKey] = val[innerKey]
      }

      innerNodes.push(ur)
      break;
    case "string":
      flatValues.push([key, val])
      break;
    case "number":
      flatValues.push([key, val])
  }

}
</script>

<template>
  <div
      class="Node"
      v-for="(_, idx) in flatValues" :key="flatValues[idx][0]">

    <div class="InfoTableRow">
      <div class="InfoTableColumn">
        <Inputer
            :disabled="true"
            v-model="flatValues[idx][0]"/>
      </div>

      <div class="Separator">:</div>
      <div class="InfoTableColumn">
        <Inputer v-model="flatValues[idx][1]"/>
      </div>
    </div>
  </div>
  <div
      class="Node"
      v-for="(_, i) in innerNodes" :key="innerNodes[i].resource_name"
  >
    <Inputer
        :disabled="true"
        v-model="innerNodes[i].fieldName"></Inputer>

    <KeyMapComponent
        v-model="innerNodes[i]"/>
  </div>
</template>

<style scoped>
@import "@/assets/styles/config_display.css";
.Separator {
  display: flex;
  align-items: center;
}
</style>
