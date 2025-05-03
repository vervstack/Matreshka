<script setup lang="ts">
import { ref } from "vue";

import { ConfigValue } from "@/models/shared/Values.ts";

import Inputer from "@/components/base/Inputer.vue";
import Button from "@/components/base/config/Button.vue";
import RollbackIcon from "@/assets/svg/general/rollback.svg";

const model = defineModel<ConfigValue<string>>({
  required: true,
});

const emit = defineEmits(["rollback", "showOriginal"]);

const showOriginalRef = ref<Boolean>(model.value.isNameChanged());
const originalValueRef = ref<string>(model.value.getOriginalName());

function rollback() {
  emit("rollback");
}

function showOriginal() {
  emit("showOriginal");
  showOriginalRef.value = !showOriginalRef.value;
}

</script>

<template>
  <div
    class="KeyValueInputer"
  >
    <div
      class="Field show"
    >
      <Inputer
        v-model="model.envName"
      />
    </div>
    <div
      class="Field RollBackButton"
      v-if="model.isNameChanged() || model.isValueChanged()"
      :class="{show: model.isNameChanged() || model.isValueChanged()}"
    >
      <Button
        title="Rollback to original"
        @click="rollback"
        :icon="RollbackIcon"
        borderless
      />
    </div>
    <Inputer
      class="Field RollBackButton"
      :class="{show: model.isNameChanged()}"
      v-if="model.isNameChanged()"
      disabled
      v-model="originalValueRef"
      :style="{
            borderTopLeftRadius: 0,
            borderBottomLeftRadius: 0
        }"
    />
  </div>
</template>

<style scoped>
.KeyValueInputer {
  width: 100%;
  height: 100%;

  display: flex;
  justify-content: center;
  align-items: center;
  gap: 0.5em;
}

.Field {
  width: 100%;
  height: 100%;
  transition: 0.75s ease;
  overflow: hidden;
  flex: 0;
}

.show {
  flex: 1
}

.RollBackButton {
  min-width: 2em;
  max-width: 2em;
  min-height: 100%;

  overflow: hidden;
  border: black solid 1px;
  border-radius: var(--border-radius);
}

</style>
