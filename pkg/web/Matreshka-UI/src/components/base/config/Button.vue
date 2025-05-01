<script setup lang="ts">

defineProps({
  label: {
    type: String,
    required: true,
  },
  disabled: {
    type: Boolean,
    default: false,
  },
  rounded: {
    type: Boolean,
    default: false,
  },
  borderless: {
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
  <button
    class="button"
    :class="{'rounded':rounded, 'borderless': borderless}"
    :onclick="buttonClicked"
  >
    {{ label }}
  </button>
</template>

<style scoped>
.button {
  width: 100%;
  height: 100%;
  background: none;
  border: black solid 1px;
  border-radius: 4px;
}
.button:hover {
  cursor: pointer;
}

.rounded {
  border-radius: 50%;
}

.borderless {
  border: none;
}
</style>
