<script setup lang="ts">
import { Nullable } from "@primevue/core";
import Dialog from "primevue/dialog";
import FloatLabel from "primevue/floatlabel";
import InputText from "primevue/inputtext";
import { ref } from "vue";

import { useSettingsStore } from "@/app/store/Settings.ts";

const settingsStore = useSettingsStore();

// API
const apiURL = ref<string>(settingsStore.getApiURL);
const pass = ref<string>(settingsStore.getPass);

function changeApiUrl() {
  settingsStore.setApiURL(apiURL.value);
}

function changePass() {
  settingsStore.setPass(pass.value);
}

const isDialogOpen = ref<boolean>(false);
const dialogHeader = ref<string>("Settings");
const dialogStyle = {
  width: "60vw",
  height: "60vh",
};


// Open Settings
document.addEventListener("keydown", function (event) {
  if ((event.ctrlKey || event.metaKey) && event.key === ".") {
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
      <FloatLabel style="width: 100%" variant="on">
        <InputText style="width: 100%" v-model="apiURL as Nullable<string>" @input="changeApiUrl" />
        <label> {{ apiURL ? "Api Url" : "Requests will be routed to root" }} </label>
      </FloatLabel>

      <FloatLabel style="width: 100%" variant="on">
        <InputText style="width: 100%" v-model="pass as Nullable<string>" @input="changePass" />
        <label> {{ "Password is NOT reset on refresh" }} </label>
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
