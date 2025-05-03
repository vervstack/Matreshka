<script setup lang="ts">
import { ref } from "vue";

import { ConfigValue } from "@/models/shared/Values.ts";

import Inputer from "@/components/base/Inputer.vue";
import Button from "@/components/base/config/Button.vue";
import RollbackIcon from "@/assets/svg/general/rollback.svg";

const model = defineModel<ConfigValue<string>>({
  required: true,
});

const emit = defineEmits(["rollback", "showOriginal", "showActual"]);

const originalNameRef = ref<string>(model.value.getOriginalName());
const originalValueRef = ref<string>(model.value.getOriginalValue());

const isPreparingToRevert = ref<boolean>(false);

function rollback() {
  emit("rollback");
}

function showOriginal() {
  emit("showOriginal");
  isPreparingToRevert.value = true;
}

function showActual() {
  emit("showActual");
  isPreparingToRevert.value = false;
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

    <p v-if="!model.isNameChanged() && isPreparingToRevert">:</p>

    <div
      class="Field"
      :class="{show: isPreparingToRevert}"
    >
      <Inputer
        v-if="model.isNameChanged()"
        disabled
        v-model="originalNameRef"
      />
      <Inputer
        v-else
        disabled
        v-model="originalValueRef"
      />
    </div>

    <Button
      v-if="model.isNameChanged() || model.isValueChanged() || isPreparingToRevert"
      class="RollBackButton"
      title="Rollback to original"
      @click="rollback"
      :icon="RollbackIcon"

      @mouseenter="showOriginal"
      @mouseleave="showActual"
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
  width: 1.75em;
  height: 1.75em;

  overflow: hidden;
  border: black solid 1px;
  border-radius: var(--border-radius);
}

</style>
