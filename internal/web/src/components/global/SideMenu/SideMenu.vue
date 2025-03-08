<script setup lang="ts">
import {ref} from "vue";

import CreateVervConfigWidget from "@/widget/CreateVervConfigWidget.vue";

import SpeedDial from "primevue/speeddial";
import {MenuItem} from "primevue/menuitem";
import Dialog from "primevue/dialog";


const isDialogOpen = ref<boolean>(false);

function openCreateVervConfigWidget() {
  isDialogOpen.value = true
}

const buttons: MenuItem[] = [
  {
    label: 'New Verv config',
    icon: 'pi pi-box',
    command(_) {
      openCreateVervConfigWidget()
    },
  },
  {
    visible: false,
    label: 'New environment config',
    icon: 'pi pi-file-plus',
    command(_) {
      // TODO
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
      header="New verv config"
      :dismissableMask="true"
      :pt="{
          root: 'border-none',
          mask: {
            style: 'backdrop-filter: blur(2px)'
          }
        }"
      :style="{
          width: '40vw',
          height: '50vh',
        }"
      position="right"
  >
    <CreateVervConfigWidget/>
  </Dialog>
</template>

<style scoped>

</style>
