<script setup lang="ts">
import { ConfigValue } from "@/models/shared/Values.ts";
import Inputer from "@/components/base/Inputer.vue";

const model = defineModel<ConfigValue<string>>({
  required: true,
});


</script>

<template>
  <div
    class="KeyValueInputer"
  >
    <div
      class="InputRow"
    >
      <Inputer
        v-model="model.envName"
      />
      <p>:</p>
      <Inputer
        v-model="model.value"
      />
    </div>
  </div>
</template>

<style scoped>
.InputRow, .KeyValueInputer {
  width: 100%;
  height: 100%;

  display: flex;
  justify-content: center;
  align-items: center;
  gap: 0.5em;
}

.InputRow {
  flex-direction: row;
}

.KeyValueInputer {
  flex-direction: column;
}

.edited {
  border: var(--value-changed-outline) solid 1px;
  border-radius: 6px;
}

.InputBox {
  height: 100%;
  width: 100%;
  box-sizing: border-box;

  display: flex;
  flex-direction: column;

  gap: 0.25em;
}

.Input {
  overflow: hidden;
}

</style>
