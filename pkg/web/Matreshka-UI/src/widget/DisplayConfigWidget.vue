<script setup lang="ts">
import {ref} from "vue";

import {ListConfigsRequest} from "@vervstack/matreshka";
import {GetConfigNodes, ListServices, PatchConfig} from "@/processes/api/api.ts";
import {handleGrpcError} from "@/processes/api/error_codes.ts";

import {AppConfigClass} from "@/models/configs/verv/AppConfig.ts";

import Button from 'primevue/button';
import InputGroup from "primevue/inputgroup";
import SelectButton from 'primevue/selectbutton';
import {useToast} from "primevue/usetoast";

import AppInfo from "@/components/config/verv/app_info/AppInfo.vue";
import ResourcesInfo from "@/components/config/verv/resource/AppResources.vue";
import ServersInfo from "@/components/config/verv/server/ServersInfo.vue";

const toastApi = useToast();

const props = defineProps({
  serviceName: {
    type: String,
    required: true
  },
})

const configData = ref<AppConfigClass>();
const versions = ref<string[]>([]);
const selectedVersion = ref<string>('');

function setData(c: AppConfigClass) {
  configData.value = c
}

function rollbackAll() {
  configData.value?.rollback()
}

async function fetchConfig() {
  configData.value = undefined

  GetConfigNodes(props.serviceName, selectedVersion.value)
      .then(setData)
      .catch(handleGrpcError(toastApi))
}

async function fetchVersions() {
  const listReq = {
    paging: {
      limit: 1
    },
    searchPattern: props.serviceName
  } as ListConfigsRequest;

  ListServices(listReq)
      .then(res => {
          versions.value = res.servicesInfo[0].versions
          selectedVersion.value = versions.value[0]
      })
      .catch(handleGrpcError(toastApi))
}

async function save() {
  if (!configData.value) return

  const changes = configData.value.getChanges()
  PatchConfig(props.serviceName, selectedVersion.value, changes)
      .then(setData)
      .catch(handleGrpcError(toastApi))
}

fetchVersions()
    .then(fetchConfig)

</script>

<template>
  <div v-if="!configData">No App config data</div>

  <div v-else class="Display">
    <div class="Tittle">{{ configData.appInfo.name.value }}</div>
    <SelectButton
        v-if="versions.length > 1"
        v-model="selectedVersion"
        :options="versions"
        @update:modelValue="fetchConfig"
    />

    <div class="Content">
      <div
          class="ContentBlock"
          :style="{
            borderColor: configData.appInfo.isChanged() ? 'var(--warn)':'var(--good)',
          }"
      >
        <AppInfo
            v-model="configData.appInfo"/>
      </div>
      <div
          class="ContentBlock"
          :style="{
            borderColor: configData?.getChangedDataSourcesNames().length != 0 ? 'var(--warn)':'var(--good)',
          }"
      >
        <ResourcesInfo
            v-model="configData.dataSources"/>
      </div>
      <div
          class="ContentBlock"
          :style="{
            borderColor: configData?.getChangedServersNames().length != 0 ? 'var(--warn)':'var(--good)',
          }"
      >
        <ServersInfo
            v-model="configData.servers"/>
      </div>
    </div>

    <Transition name="BottomControls">
      <InputGroup
          v-show="configData?.isChanged()"
          :style="{
              display: 'flex',
              justifyContent: 'center'
            }"
      >
        <Button
            @click="save"
            label="Save"
            icon="pi pi-check" iconPos="right"
            severity="danger"
        />
        <Button
            @click="rollbackAll"
            label="Rollback"
            icon="pi pi-refresh" iconPos="right"
        />
      </InputGroup>
    </Transition>
  </div>
</template>

<style scoped>
@import "@/assets/styles/config_display.css";

.Display {
  overflow-x: hidden;
  display: flex;
  flex-direction: column;
  gap: 1em;
}

.Content {
  display: flex;
  justify-content: center;
  flex-direction: column;
  gap: 1em;
}

.ContentBlock {
  border: solid;
  /*border: var(--good) solid;*/
  border-radius: 16px;
  padding: 1em;
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
