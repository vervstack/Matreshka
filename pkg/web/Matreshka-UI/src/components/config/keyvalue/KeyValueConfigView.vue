<script setup lang="ts">

import KeyValueConfig from "@/models/configs/keyvalue/KeyValueConfig.ts";
import KeyValueInput from "@/components/base/config/fields/KeyValueInput.vue";
import Button from "@/components/base/config/Button.vue";
import { ref } from "vue";

import { Node } from "@vervstack/matreshka";

const model = defineModel<KeyValueConfig>({
  required: true,
});

const ghostNodeIdx = ref<number>();

function addSubNode() {
  ghostNodeIdx.value = undefined;
  isChildrenFolded.value = false;
  addGhostNode()
}

function addGhostNode() {
  ghostNodeIdx.value = model.value.children.length || 0;
  model.value.children.push(new KeyValueConfig({ name: "key", value: "value" } as Node));
  isChildrenFolded.value = false;
}

function removeGhostNode() {
  if (ghostNodeIdx.value !== undefined) {
    model.value.children.pop();
  }

  ghostNodeIdx.value = undefined;
}

const isChildrenFolded = ref<boolean>(false);

function toggleFolding() {
  isChildrenFolded.value = !isChildrenFolded.value;
}

function shouldShowFoldButton(): boolean {
  if (model.value.children.length < 2) {
    return false;
  }

  if (model.value.children.length > 2) {
    return true;
  }

  return ghostNodeIdx.value == undefined;
}
</script>

<template>
  <div class="wrapper">
    <div
      class="top-wrapper"
     :class="{'folded': isChildrenFolded}"
    >
      <div
        class="button"
      >
        <Button
          @click="addSubNode"
          @mouseenter="addGhostNode"
          @mouseleave="removeGhostNode"
          :label="'+'"
        />
      </div>
      <KeyValueInput
        v-if="model.configValue"
        v-model="model.configValue" />
      <div
        class="button"
        v-show="shouldShowFoldButton()"
      >
        <Button
          borderless
          :label="isChildrenFolded ? '▲' : '▼'"
          @click="toggleFolding"
        />
      </div>
    </div>
    <div
      class="children"
    >
      <TransitionGroup name="ghost">
        <div
          class="child"
          v-for="(_, idx) in model.children"
          :class="{'ghost': idx == ghostNodeIdx}"
          :key="idx"
          v-show="!isChildrenFolded"
        >
          <KeyValueConfigView
            v-model="model.children[idx]"
          />
        </div>
      </TransitionGroup>
    </div>
  </div>
</template>

<style scoped>
.wrapper {
  display: flex;
  flex-direction: column;
  gap: 1em;
  height: 100%;
}

.children {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  gap: 1em;
  border-left: #6b7280 solid 1px;
  transition: height 0.4s ease;
}

.child {
  height: fit-content;
  width: 98%;
}

.top-wrapper {
  display: flex;
  flex-direction: row;
  gap: 1em;
  align-items: center;
}

.folded {
 border-bottom: #6b7280 dashed 1px;
}

.button {
  width: 2.5em;
  height: 2.5em;
  display: flex;
  justify-content: center;
  align-items: center;
}

.ghost {
  border: #6b7280 dashed 1px;
  border-radius: 6px;
  padding: 2px;
}

.ghost-enter-active,
.ghost-leave-active {
  transition: 0.25s;
}

.ghost-enter-to,
.ghost-leave-from {
  transform: translateY(0);
  opacity: 1;
}

.ghost-enter-from,
.ghost-leave-to {
  transform: translateY(-10%);
  opacity: 0;
}

</style>
