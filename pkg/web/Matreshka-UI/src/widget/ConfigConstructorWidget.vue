<script setup lang="ts">
import { ConfigTypePrefix } from "@vervstack/matreshka";
import { useToast } from "primevue/usetoast";
import { ref, watch } from "vue";

import { RouteToConfigDisplay } from "@/app/routes/Routes.ts";
import Button from "@/components/base/config/Button.vue";
import DropdownSelector from "@/components/base/DropdownSelector.vue";
import Inputer, { IInputer } from "@/components/base/Inputer.vue";
import { CreateConfig } from "@/processes/api/ApiService.ts";
import handleGrpcError from "@/processes/api/ErrorCodes.ts";
import { isConfigNameValid } from "@/processes/config/Validators.ts";

const toastApi = useToast();

const configType = ref<string>("");
const configName = ref("");
const isInputerShown = ref<boolean>(false);

const options = [
  ConfigTypePrefix.verv,
  ConfigTypePrefix.pg,
  ConfigTypePrefix.nginx,
  ConfigTypePrefix.minio,
];

watch(configName, (val: string) => {
  configName.value = val.replace(/[- ]/g, "_");
});

function createConfig() {
  CreateConfig(configName.value, configType.value as ConfigTypePrefix)
    .then((resp) => {
      toastApi.add({
        closable: true,
        life: 2_000,
        severity: "success",
        summary: "Service created. Check it out",
      });

      if (resp.name) {
        RouteToConfigDisplay(resp.name);
      }
    })
    .catch(handleGrpcError(toastApi));
}

// UI
const inputerRef = ref<IInputer>();

function focusOnInput() {
  isInputerShown.value = true;
  setTimeout(() => {
    inputerRef.value?.doFocus();
  }, 50);
}

const emPerSymbol = 0.5;
</script>

<template>
  <div class="widget-container">
    <div class="basic-inputers">
      <DropdownSelector
        :style="{
          width:
            configType.length !== 0 ? 3 + emPerSymbol * configType.length * 1.5 + 'em' : '100%',
        }"
        v-model="configType"
        :options="options"
        @optionSelected="focusOnInput"
        with-clear-button
      />
      <Inputer
        v-if="configType.length !== 0"
        ref="inputerRef"
        v-model="configName"
        floatingLabel="Config name"
      />
    </div>
    <div class="button-row">
      <Button
        :disabled="!isConfigNameValid(configType, configName)"
        class="create-button"
        label="Create"
        @click="createConfig"
      />
    </div>
  </div>
</template>

<style scoped>
.widget-container {
  display: flex;
  gap: 1em;

  padding: 1em 0 0 0;

  flex-direction: column;

  width: 100%;
  height: 100%;
}

.basic-inputers {
  width: 100%;
  height: 5vh;

  display: flex;
  flex-direction: row;
  gap: 0.25em;
}

.button-row {
  width: 100%;
  display: flex;
  justify-content: flex-end;
}

.create-button {
  width: 20vw;
  height: 5vh;
}
</style>
