<script setup lang="ts">
import Button from "primevue/button";

defineProps({
  tittle: {
    type: String,
    required: true,
  },
  disabled: {
    type: Boolean,
    default: false,
  },
});

const emit = defineEmits(["click"]);

function buttonClicked() {
  emit("click");
}
</script>

<template>
  <Button
    severity="contrast"
    raised
    outlined
    :disabled="disabled"
    icon="pi pi-hammer"
    label="Create"
    :onclick="buttonClicked"
  />
</template>

<style scoped></style>
