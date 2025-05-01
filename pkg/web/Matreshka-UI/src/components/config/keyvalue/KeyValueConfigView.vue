<script setup lang="ts">

import KeyValueConfig from "@/models/configs/keyvalue/KeyValueConfig.ts";
import KeyValueInput from "@/components/base/config/fields/KeyValueInput.vue";
import Button from "@/components/base/config/Button.vue";
import { ref } from "vue";

const model = defineModel<KeyValueConfig>({
  required: true,
});

function addSubNode() {
  model.value.children.push(new KeyValueConfig({ name: "key", value: "value" }));
}

const isGhostNodePresented = ref<boolean>(false);

function addGhostNode() {
  isGhostNodePresented.value = true
  console.log(123)
}

function removeGhostNode() {
  isGhostNodePresented.value = false
  console.log(321)
}


</script>

<template>
  <div class="wrapper">
    <KeyValueInput
      v-if="model.configValue"
      v-model="model.configValue" />

    <Button
      class="add-button"
      @click="addSubNode"
      @mouseenter="addGhostNode"
      @mouseleave="removeGhostNode"
      :tittle="'Add node'"
    />
    <div
      class="children"
      v-if="model.children.length > 0"
    >
      <div
        class="child"
        v-for="(_, idx) in model.children"
      >
        <KeyValueConfigView
          v-model="model.children[idx]" />
      </div>
    </div>

  </div>
</template>

<style scoped>
.wrapper {
  display: flex;
  flex-direction: column;
  gap: 1em;
}

.children {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  gap: 1em;
  border-left: #6b7280 solid 1px;
}

.child {
  height: fit-content;
  width: 98%;
}

.add-button {
  width: 10em;
}
</style>
