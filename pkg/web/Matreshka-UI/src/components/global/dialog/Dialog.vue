<script setup lang="ts">
import Dialog from "primevue/dialog";
import { Component as VueComponent, PropType } from "vue";

defineProps({
  child: {
    type: Object as PropType<VueComponent>,
    required: true,
  },
  childProps: {
    type: Object as PropType<any>,
  },
  dialogHeader: {
    type: String,
    required: true,
  },
  dialogPt: {
    default: {
      root: "border-none",
      mask: {
        style: "backdrop-filter: blur(2px)",
      },
    },
  },
  dialogStyle: {
    default: {
      width: "80vw",
      height: "95vh",
    },
  },
  dialogPosition: {
    type: String as PropType<"center" | "right">,
  },
});

const isOpen = defineModel<boolean>({
  required: true,
});
</script>

<template>
  <Dialog
    v-model:visible="isOpen"
    modal
    :dismissableMask="true"
    :header="dialogHeader"
    :pt="dialogPt"
    :style="dialogStyle"
    :position="dialogPosition"
  >
    <component :is="child" v-bind="childProps" />
  </Dialog>
</template>

<style scoped></style>
