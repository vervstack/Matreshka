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
      :class="{edited: model.isNameChanged()}"
    />
    :
    <Inputer
      v-model="model.value"
      :class="{edited: model.isValueChanged()}"
    />
  </div>
</template>

<style scoped>
.KeyValueInputer {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  gap: 0.5em;
}

.edited {
  border: red solid 1px;
  border-radius: 6px;
}
</style>
