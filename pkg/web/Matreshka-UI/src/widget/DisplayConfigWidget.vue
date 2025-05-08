<script setup lang="ts">
import Button from "primevue/button";
import InputGroup from "primevue/inputgroup";
import SelectButton from "primevue/selectbutton";
import { useToast } from "primevue/usetoast";
import { ref } from "vue";

import Config from "@/models/configs/Config.ts";
import { GetConfigNodes, PatchConfig } from "@/processes/api/ApiService.ts";
import handleGrpcError from "@/processes/api/ErrorCodes.ts";
import ConfigIcon from "@/components/base/config/ConfigIcon.vue";
import ConfigName from "@/components/base/config/ConfigName.vue";

const toastApi = useToast();

const props = defineProps({
  configName: {
    type: String,
    required: true,
  },
});

const configData = ref<Config>(new Config(props.configName));

async function fetchConfig() {
  GetConfigNodes(props.configName, configData.value.selectedVersion)
    .then((d) => {
      configData.value = d;
    })
    .catch(handleGrpcError(toastApi));
}

async function save() {
  if (!configData.value) return;

  PatchConfig(configData.value)
    .then((d) => (configData.value = d))
    .catch(handleGrpcError(toastApi));
}

fetchConfig().then(fetchConfig);
</script>

<template>
  <div v-if="!configData">No App config data</div>

  <div v-else class="Display">
    <div class="LogoAndTittle">
      <div class="Logo">
        <ConfigIcon
          :config-type="configData.type" />
      </div>
      <div class="Tittle">
        <ConfigName
          :label="configData.name" />
      </div>
    </div>
    <SelectButton
      v-if="configData.versions.length > 1"
      v-model="configData.selectedVersion"
      :options="configData.versions"
      @update:modelValue="fetchConfig"
    />
    <!--TODO Add "New Version" button?-->

    <component
      :is="configData.getComponent()"
      v-model="configData.content" />

    <div
      class="Footer"
    >
      <Transition name="BottomControls">
        <InputGroup
          v-if="configData?.isChanged()"
          :style="{
            display: 'flex',
            justifyContent: 'center',
          }"
        >
          <Button @click="save" label="Save" icon="pi pi-check" iconPos="right" severity="danger" />
          <Button
            @click="configData?.rollback()"
            label="Rollback"
            icon="pi pi-refresh"
            iconPos="right"
          />
        </InputGroup>


      </Transition>
    </div>
  </div>
</template>

<style scoped>
@import "@/assets/styles/config_display.css";

.Display {
  display: flex;
  flex-direction: column;
  gap: 1em;
}

.LogoAndTittle {
  width: 100%;
  display: flex;
  flex-direction: row;
  gap: 1em;
}

.Logo {
  width: 4vw;
}

.Tittle {
  width: fit-content;
  display: flex;
  justify-content: flex-start;
  font-size: 2em;
}

.Footer {
  width: 100%;
  height: 10vh;
}

.BottomControls-enter-active,
.BottomControls-leave-active {
  transition: 0.25s;
}

.BottomControls-enter-to,
.BottomControls-leave-from {
  transform: translateY(0);
  opacity: 1;
}

.BottomControls-enter-from,
.BottomControls-leave-to {
  transform: translateY(-100%);
  opacity: 0;
}

</style>
