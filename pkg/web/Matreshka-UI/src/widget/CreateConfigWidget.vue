<script setup lang="ts">

import {ref, watch} from "vue";

import Select, {SelectChangeEvent} from "primevue/select";
import InputText from "primevue/inputtext";
import FloatLabel from "primevue/floatlabel";
import InputGroup from "primevue/inputgroup";

import CreateConfigButton from "@/components/base/config/CreateConfigButton.vue";

import {ConfigTypePrefix} from "@vervstack/matreshka";

const configTypeInput = ref<string>('');
const configType = ref<ConfigTypePrefix | null>(null);

const serviceName = ref<string>('');

const configTypeOptions: ConfigTypePrefix[] = [
  ConfigTypePrefix.verv,
  ConfigTypePrefix.pg,
  ConfigTypePrefix.minio,
  ConfigTypePrefix.nginx,
]

watch(configTypeInput, async (newValue, oldValue) => {
  if (configType.value) return

  if (!newValue.endsWith("_")) return

  const match = configTypeOptions.find(opt => opt.startsWith(oldValue));
  if (match) {
    configType.value = match;
    configTypeInput.value = newValue.slice(0, -1);
  }
}, {flush: 'post'})

watch(serviceName, (newValue)=> {
  serviceName.value = newValue.replace(" ","_")
})

function onManualSelect(e: SelectChangeEvent) {
  if (e.originalEvent.type != 'click') return

  configType.value = e.value;
  configTypeInput.value = e.value;
}

</script>

<template>
  <div class="container">
    <div class="name-container">
      <Select
          v-model="configTypeInput"
          :options="configTypeOptions"
          placeholder="Select a config type"
          class="type-select"
          filterMatchMode="contains"
          v-on:change="onManualSelect"
          editable
      />
      <div
          v-show="configType"
          class="name-inputer"
      >
        <InputGroup>
          <FloatLabel variant="on">
            <InputText
                v-model="serviceName"
            />
            <label>Service name</label>
          </FloatLabel>
        </InputGroup>
      </div>
    </div>

    <div :style="{display: 'flex', justifyContent: 'end'}">
      <CreateConfigButton
          :serviceName="serviceName"/>
    </div>
  </div>
</template>

<style scoped>

.container {
  display: flex;
  gap: 1em;

  padding: 1em 0 0 0;

  flex-direction: column;

  width: 100%;
  height: 100%;
}

.name-container {
  display: flex;
  width: 100%;
  gap: 0;
}

.type-select {
  flex: 1;
  width: 100%;
}

.name-inputer {
  flex: 4;
}


</style>
