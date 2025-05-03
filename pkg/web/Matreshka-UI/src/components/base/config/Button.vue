<script setup lang="ts">

defineProps({
  label: {
    type: String,
    default: "",
  },
  icon: {
    type: String,
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
    class="simple-button"
    :class="{'rounded':rounded, 'borderless': borderless}"
    :onclick="buttonClicked"
  >
    <img
      v-if="icon"
      :src="icon"
      alt="no icon"
    />
    <p v-else>
      {{ label }}
    </p>
  </button>
</template>

<style scoped>
.simple-button {
  width: 100%;
  height: 100%;
  background: none;
  border: black solid 1px;
  border-radius: var(--border-radius);
}

.simple-button:hover {
  cursor: pointer;
}

.rounded {
  border-radius: 50%;
}

.borderless {
  border: none;
}
</style>
