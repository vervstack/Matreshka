<script setup lang="ts">

import AddNodeIcon from "@/assets/svg/node/add.svg";
import KeyValueConfig from "@/models/configs/keyvalue/KeyValueConfig.ts";
import KeyValueInput from "@/components/base/config/fields/KeyValueInput.vue";
import Button from "@/components/base/config/Button.vue";
import { ref } from "vue";

import { Node } from "@vervstack/matreshka";

const rootRef = ref<HTMLElement | null>(null);
const childrenRef = ref<HTMLElement | null>(null);

const model = defineModel<KeyValueConfig>({
  required: true,
});

const ghostNodeIdx = ref<number>();

function addSubNode() {
  if (ghostNodeIdx.value != undefined
    &&
    ghostNodeIdx.value < model.value.children.length) {

    const kv = model.value.children[ghostNodeIdx.value];
    if (kv.configValue) {
      kv.configValue.isMuted = false;
    }
    model.value.children[ghostNodeIdx.value] = kv;
  }


  ghostNodeIdx.value = undefined;
  isChildrenFolded.value = false;
  addGhostNode();
}

function addGhostNode() {
  ghostNodeIdx.value = model.value.children.length || 0;
  const kv = new KeyValueConfig({ name: "key", value: "value" } as Node);

  if (kv.configValue) {
    kv.configValue.isMuted = true;
    kv.configValue.isNew = true;
  }

  model.value.children.push(kv);
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
  <div
    class="wrapper"
    ref="rootRef"
    :style="{width: isChildrenFolded ? rootRef?.clientWidth +'px': ''}"
  >

    <div
      class="top-wrapper"
      :class="{'folded': isChildrenFolded}"
    >
      <div
        class="button add-button"
        title="Add new node"
      >
        <Button
          @click="addSubNode"
          @mouseenter="addGhostNode"
          @mouseleave="removeGhostNode"
          :label="'+'"
          :icon="AddNodeIcon"
        />
      </div>
      <KeyValueInput
        v-if="model.configValue"
        v-model="model.configValue" />
      <div
        class="button"
        :title="isChildrenFolded ? 'Unfold':'Fold'"
        v-if="shouldShowFoldButton()"
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
      ref="childrenRef"
    >
      <TransitionGroup name="child">
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
  justify-content: center;
  height: 100%;
  width: fit-content;
}

.children {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  gap: 0.25em;
  padding: 0 0 0 2em;
}

.child {
  width: 98%;
  max-width: 98%;
  flex: 1;
  border-left: #6b7280 solid 1px;
}

.top-wrapper {
  display: flex;
  flex-direction: row;
  gap: 1em;
  align-items: center;
  justify-content: space-between;
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

.add-button {
  position: sticky;
  top: 0;
}

.ghost {
  border: #6b7280 dashed 1px;
  border-radius: 6px;
}

.child-enter-active,
.child-leave-active {
  transition: 0.3s;
}

.child-enter-to,
.child-leave-from {
  transform: translateY(0);
  opacity: 1;
}

.child-enter-from,
.child-leave-to {
  transform: translateY(-10%);
  opacity: 0;
}

</style>
