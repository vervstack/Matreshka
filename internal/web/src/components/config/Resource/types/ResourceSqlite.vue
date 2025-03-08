<script setup lang="ts">
import {Sqlite} from "@/models/AppConfig/Resources/Resource.ts";
import ConfigField from "@/components/base/ConfigField/ConfigInput.vue";
import SelectButton from "@/components/base/SelectButton.vue"

import {ref} from "vue";

const model = defineModel<Sqlite>({
  required: true,
})

const inMemoryKey = 'in mem'
const customKey = 'custom'

const selectedOption = ref<string>(customKey)

if (model.value.path.value == "in_memory") {
  selectedOption.value = inMemoryKey;
}

const options = ref<string[]>([inMemoryKey, customKey])

function valueChanged() {
  if (selectedOption.value == inMemoryKey) {
    model.value.path.value = 'in_memory'
  }
}

</script>

<template>
  <div class="Node">
    <div class="InputPath">
      <div class="SqlitePathInputer">
        <ConfigField
            v-model="model.path"
            field-name="File path"
        />
      </div>

      <div class="SqliteTypeSelectButton">
        <SelectButton
            v-model="selectedOption"
            :options="options"
            @update:modelValue="valueChanged"
        />
      </div>
    </div>
  </div>
</template>

<style scoped>
@import "@/assets/styles/config_display.css";

.InputPath {
  display: flex;
  flex-direction: row;
  justify-content: stretch;
  align-items: end;
  width: 100%;
  gap: 1em;
}

.SqlitePathInputer {
  flex: 2;
}

.SqliteTypeSelectButton {
  flex: 1;
}
</style>
