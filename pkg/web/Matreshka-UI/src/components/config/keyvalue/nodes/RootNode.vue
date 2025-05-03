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
    :class="{
      changed: model.isChanged(),
    }"
  >
      <Inputer
        v-model="model.envName"
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

.changed {
  border: var(--value-changed-outline) solid 1px;
  border-radius: var(--border-radius);
}
</style>
