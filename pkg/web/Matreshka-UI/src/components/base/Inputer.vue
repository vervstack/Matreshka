<script setup lang="ts">
import {ref, watch} from 'vue'

export interface IInputer {
  doFocus: () => void
}

defineProps({
  placeholder: {
    type: String,
    default: ''
  },
  floatingLabel: {
    type: String,
    default: ''
  },
})

const model = defineModel<string>({
  required: true,
})

const inputElementRef = ref<HTMLElement>()

function doFocus() {
  inputElementRef.value?.focus()
}

defineExpose<IInputer>({doFocus})

</script>

<template>
  <div class="input-wrap">
    <label v-if="floatingLabel" class="floating-label">{{ floatingLabel }}</label>
    <input
        ref="inputElementRef"
        v-model="model"
        type="text"
        class="input"
        :placeholder="placeholder"
    />
  </div>
</template>

<style scoped>
.input-wrap {
  display: flex;
  flex-direction: column;
  width: 100%;
  height: 100%;
  position: relative;
}

.input {
  width: 100%;
  height: 100%;
  padding: 0.5rem 0.75rem;
  font-size: 16px;
  border: 1px solid #d1d5db;
  border-radius: 6px;
  outline: none;
}

.input:focus {
  border-color: var(--focus);
  box-shadow: 0 0 0 2px var(--focus);
}

.floating-label {
  position: absolute;
  font-size: 0.75em;

  color: #6b7280;
  margin-bottom: 4px;
  top: -10px;
  left: 10px;
  background-color: white;
  padding: 0 4px;
  transform: translateY(-12.5%);
  pointer-events: none;
}
</style>
