<script setup lang="ts">

import AddNodeIcon from "@/assets/svg/node/add.svg";
import KeyValueConfig from "@/models/configs/keyvalue/KeyValueConfig.ts";
import KeyValueNode from "@/components/config/keyvalue/nodes/KeyValueNode.vue";
import Button from "@/components/base/config/Button.vue";
import { ref } from "vue";

import EnvNode from "@/models/shared/Node.ts";

const contentBlockRef = ref<HTMLElement | null>(null);
const fullSize = ref<string>('100vw');

const model = defineModel<KeyValueConfig>({
  required: true,
});

const props = defineProps({
  parentPrefix: {
    type: String,
    default: "",
  },
});

const ghostNodeIdx = ref<number>();

function addSubNode() {
  if (ghostNodeIdx.value != undefined &&
    ghostNodeIdx.value < model.value.children.length) {

    const kv = model.value.children[ghostNodeIdx.value];
    if (kv.configValue) {
      kv.configValue.isMuted = false;
    }

    model.value.children[ghostNodeIdx.value] = kv;

    model.value.configValue.isRoot = true;
    model.value.configValue.value = "";
  }

  ghostNodeIdx.value = undefined;
  isChildrenFolded.value = false;
  addGhostNode();
}

function addGhostNode() {
  ghostNodeIdx.value = model.value.children.length || 0;
  const kv = new KeyValueConfig(new EnvNode("key", model.value.configValue.value));
  model.value.configValue.isRoot = true;

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
  model.value.configValue.isRoot = model.value.configValue.getOriginalIsRoot();
}

const isChildrenFolded = ref<boolean>(false);

function toggleFolding() {
  if (!isChildrenFolded.value) {
    fullSize.value = contentBlockRef.value?.clientWidth+'px' || '100vw'
    console.log(fullSize.value)
  }

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

    <div
      class="content-wrapper"
      ref="contentBlockRef"
      :class="{'folded': isChildrenFolded}"
      :style="{width: isChildrenFolded ? fullSize: ''}"
    >
      <div
        class="config-value"
      >
        <KeyValueNode
          v-if="model.configValue.getOriginalName() !== ''"
          v-model="model.configValue"
          :parent-prefix="props.parentPrefix"
          :force-root-mode="model.children.length > 0"
        />
      </div>
      <div
        class="children"
        v-if="model.children.length > 0"
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
              :parent-prefix="model.configValue.envName"
            />
          </div>
        </TransitionGroup>
      </div>

    </div>

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
</template>

<style scoped>
.wrapper {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  height: 100%;
  width: fit-content;
}

.config-value {
  height: 2em;
  width: 100%;
  display: flex;
}

.children {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  gap: 0.25em;
}

.child {
  width: 100%;
  flex: 1;
  border-left: #6b7280 solid 1px;
}

.content-wrapper {
  display: flex;
  flex-direction: column;
  gap: 1em;
  align-items: center;
  justify-content: center;
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
