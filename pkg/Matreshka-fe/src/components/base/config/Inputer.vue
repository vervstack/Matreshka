<script setup lang="ts">

const value = defineModel<String | number | undefined>({
  required: true,
})

defineProps({
  disabled: {
    type:  Boolean,
    default: false,
  }
})

</script>

<template>
  <input
      class="Inputer"
      :size="(value?.toString().length ?? 3) + (disabled ? 1 : 0)"
      :disabled="disabled"
      v-model.lazy="value">
</template>

<style scoped>

.Inputer {
  border-radius: 2px;
  padding: 0.125em 1em 0.125em 1em;

  border: none;
  border-bottom: #12c4ca solid;
  width: fit-content;
}

.Inputer:focus {
  outline: none;
}

.Inputer:disabled {
  border-bottom: #4a5353 solid;
}

</style>
