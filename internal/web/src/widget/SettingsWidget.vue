<script setup lang="ts">

import {ref} from "vue";

import {useSettingsStore} from "@/app/store/settings.ts";

import {Nullable} from "@primevue/core";

import InputText from 'primevue/inputtext';
import FloatLabel from 'primevue/floatlabel'
import Dialog from "primevue/dialog";

const settingsStore = useSettingsStore()

// API URL
const apiURL = ref<string>(settingsStore.getApiURL)

function changeApiUrl() {
  settingsStore.setApiURL(apiURL.value)
}

const isDialogOpen = ref<boolean>(false)
const dialogHeader = ref<string>("Settings")
const dialogStyle = {
  width: '60vw',
  height: '60vh',
}

// Open Settings
document.addEventListener('keydown', function(event) {
  if ((event.ctrlKey || event.metaKey) && event.key === '.') {
    event.preventDefault();
    isDialogOpen.value = true;
  }
});

</script>

<template>
  <Dialog
      v-model:visible="isDialogOpen"
      modal
      :dismissableMask="true"
      :header="dialogHeader"
      :style="dialogStyle"
  >
    <div class="Settings">
      <FloatLabel
          style="width: 100%"
          variant="on">
        <InputText
            style="width: 100%"
            v-model="apiURL as Nullable<string>"
            :invalid="apiURL == ''"
            @input="changeApiUrl"
        />
        <label> Api Url </label>
      </FloatLabel>
    </div>
  </Dialog>
</template>

<style scoped>
.Settings {
  padding: 1em 0 0 0;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  gap: 1em;
}
</style>
