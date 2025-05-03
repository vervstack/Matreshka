<script setup lang="ts">
import { ref, watch } from "vue";

import { ConfigValue } from "@/models/shared/Values.ts";

import Inputer from "@/components/base/Inputer.vue";
import Button from "@/components/base/config/Button.vue";
import EyeIcon from "@/assets/svg/node/eye.svg";
import RollbackIcon from "@/assets/svg/general/rollback.svg";

const model = defineModel<ConfigValue<string>>({
  required: true,
});

const showOriginalRef = ref<Boolean>(model.value.isNameChanged());

function rollback() {
  showOriginalRef.value = false
  model.value.rollbackName()
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
      class="Field Original"
      :class="{show: showOriginalRef}"
    >
      <div
        class="RollBackButton"
        title="Rollback to original"
      >
      <Button
        @click="rollback"
        :icon="RollbackIcon"
        :style="{
            borderTopRightRadius: 0,
            borderBottomRightRadius: 0
        }"
      />
      </div>
      <Inputer
        disabled
        v-model="model.envName"
        :style="{
            borderTopLeftRadius: 0,
            borderBottomLeftRadius: 0
        }"
      />
    </div>

    <div
      class="ShowButton"
      title="Show changes"
      v-show="model.isChanged()"
    >
      <Button
        @click="showOriginalRef=!showOriginalRef"
        :icon="EyeIcon"
      />
    </div>
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
  transition: 0.75s ease;
  overflow: hidden;
  flex: 0;
}

.Original {
  display: flex;
  flex-direction: row;
}

.show {
  flex: 1
}

.ShowButton {
  width: 2em;
  height: 2em;
  background: var(--warn);
  border-radius: var(--border-radius);
}

.RollBackButton {
  min-width: 2em;
  height: 2em;
}

</style>
