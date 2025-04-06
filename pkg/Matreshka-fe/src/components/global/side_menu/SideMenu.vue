<script setup lang="ts">
import {Component as VueComponent, ref, shallowRef, watch} from "vue";

import CreateConfigWidget from "@/widget/CreateConfigWidget.vue";

import {MenuItem} from "primevue/menuitem";
import SpeedDial from "primevue/speeddial";
import Dialog from "primevue/dialog";

const isDialogOpen = ref<boolean>(false);
const newConfigDialog = shallowRef<VueComponent | undefined>();

watch(isDialogOpen, ()=>{
  if (!isDialogOpen.value) {
    newConfigDialog.value = undefined
  }
})

watch(newConfigDialog, ()=>{
  if (newConfigDialog.value !== undefined) {
    isDialogOpen.value = true
  }
})

const buttons: MenuItem[] = [
  {
    label: 'New config',
    icon: "pi pi-box",
    command(_) {
      newConfigDialog.value = CreateConfigWidget
    },
  },
]

</script>

<template>
  <!-- Help button at the bottom -->
  <SpeedDial
      :style="{ position: 'absolute', bottom: '2%', right: '2%' }"
      :tooltipOptions="{ event: 'hover', position: 'left' }"
      :model="buttons"
      direction="up"
      :radius="100"
  />

  <Dialog
      v-model:visible="isDialogOpen"
      modal
      header="New config"
      :dismissableMask="true"
      :pt="{
          root: 'border-none',
          mask: {
            style: 'backdrop-filter: blur(2px)'
          }
        }"
      :style="{
          width: 'clamp(40em, 40vw, 60vh)',
          height: '50vh',
        }"
      position="right"
  >
    <Component :is="newConfigDialog"/>
  </Dialog>
</template>

<style scoped>
</style>
