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

.KeyValueInputer {
  flex-direction: column;
}

</style>
