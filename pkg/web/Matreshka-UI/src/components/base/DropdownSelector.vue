<script setup lang="ts">
import { computed, ref, watch } from "vue";

const props = defineProps({
  options: {
    type: Array<string>,
    default: [],
  },
  withClearButton: {
    type: Boolean,
    default: false,
  },
});

const emit = defineEmits(["optionSelected"]);

const model = defineModel<string>({
  required: true,
});

const isDropdownOpen = ref(false);
const highlightedIndex = ref(0);
const searchQuery = ref("");

const filteredOptions = computed(() =>
  props.options.filter((opt) => opt.toLowerCase().includes(searchQuery.value.toLowerCase()))
);

function selectOption(option: string) {
  model.value = option;
  searchQuery.value = option;
  isDropdownOpen.value = false;
  emit("optionSelected", option);
}

function handleBlur() {
  setTimeout(() => {
    isDropdownOpen.value = false;
  }, 150);
}

function handleKeydown(e: KeyboardEvent) {
  if (!filteredOptions.value.length) return;

  if (e.key === "ArrowDown") {
    e.preventDefault();
    highlightedIndex.value = (highlightedIndex.value + 1) % filteredOptions.value.length;
  } else if (e.key === "ArrowUp") {
    e.preventDefault();
    highlightedIndex.value =
      (highlightedIndex.value - 1 + filteredOptions.value.length) % filteredOptions.value.length;
  } else if (e.key === "Enter" && highlightedIndex.value >= 0) {
    selectOption(filteredOptions.value[highlightedIndex.value]);
  }
}

function handleClear() {
  model.value = "";
  searchQuery.value = "";
  isDropdownOpen.value = false;
}

watch(searchQuery, (newV) => {
  if (!props.options.includes(newV)) {
    model.value = "";
  }
});
</script>

<template>
  <div class="container">
    <div class="select-container">
      <!--      INPUT AND DESELECTOR-->
      <div class="input-container">
        <input
          v-model="searchQuery"
          @focus="isDropdownOpen = true"
          @blur="handleBlur"
          @keydown="handleKeydown"
          type="text"
          class="input"
          placeholder="Select a config type"
          :class="{ invalid: !filteredOptions.length }"
        />
        <button
          title="Undo selection"
          v-show="withClearButton && searchQuery !== ''"
          @click="handleClear"
          class="clear-button"
        >
          x
        </button>
      </div>
      <!--      DROP DOWN MENU-->
      <div v-show="isDropdownOpen && filteredOptions.length" class="dropdown">
        <div
          v-for="(option, index) in filteredOptions"
          :key="index"
          @mousedown.prevent="selectOption(option)"
          class="dropdown-item"
          :class="{ highlighted: index === highlightedIndex }"
        >
          {{ option }}
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.container {
  width: 100%;
  height: 100%;
  box-sizing: border-box;
}

.select-container {
  width: 100%;
  height: 100%;
  position: relative;
  box-sizing: border-box;
}

.input-container {
  width: 100%;
  height: 100%;
  position: relative;
}

.input {
  width: 100%;
  height: 100%;

  font-size: inherit;
  padding: 0.5em 0.75em;
  border: 1px solid #d1d5db;
  border-radius: 6px;
  outline: none;
}

.input:focus {
  border-color: var(--focus);
  box-shadow: 0 0 0 2px var(--focus);
}

.dropdown {
  position: absolute;
  top: 100%;
  left: 0;
  right: 0;
  background: white;
  border: 1px solid #d1d5db;
  border-radius: 6px;
  margin-top: 4px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  z-index: 10;
  max-height: 200px;
  overflow-y: auto;
}

.dropdown-item {
  padding: 0.5em 0.75em;
  cursor: pointer;
}

.dropdown-item:hover {
  background-color: #f3f4f6;
}

.highlighted {
  background-color: #e5e7eb; /* Light gray */
}

.invalid {
  border-color: rgba(150, 50, 50, 0.2);
  box-shadow: 0 0 0 2px rgba(255, 50, 50, 0.2);
}

.input.invalid:focus {
  border-color: rgba(150, 50, 50, 0.5);
  box-shadow: 0 0 0 2px rgba(255, 50, 50, 0.4);
}

.clear-button {
  position: absolute;
  top: 50%;
  right: 0.5em;
  font-size: 1em;

  transform: translateY(-50%);
  cursor: pointer;
  border-radius: 12.5%;
  background: none;
  border: none;
}
</style>
