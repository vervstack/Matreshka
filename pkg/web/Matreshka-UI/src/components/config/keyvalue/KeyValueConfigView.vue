<script setup lang="ts">

import AddNodeIcon from "@/assets/svg/node/add.svg";
import KeyValueConfig from "@/models/configs/keyvalue/KeyValueConfig.ts";
import KeyValueNode from "@/components/config/keyvalue/nodes/KeyValueNode.vue";
import Button from "@/components/base/config/Button.vue";
import { ref } from "vue";

import EnvNode from "@/models/shared/Node.ts";

const contentBlockRef = ref<HTMLElement | null>(null);
const fullSize = ref<string>("100vw");

const model = defineModel<KeyValueConfig>({
  required: true,
});

// Ghosting
const ghostNodeIdx = ref<number>();

function addSubNode() {
  if (ghostNodeIdx.value != undefined &&
    ghostNodeIdx.value < model.value.children.length) {

    const kv = model.value.children[ghostNodeIdx.value];
    if (kv.configValue) {
      kv.configValue.isMuted = false;
    }

    model.value.children[ghostNodeIdx.value] = kv;
  }

  ghostNodeIdx.value = undefined;
  isChildrenFolded.value = false;
  addGhostSubNode();

  model.value.configValue.isMuted = false;
}
function addGhostSubNode() {
  ghostNodeIdx.value = model.value.children.length || 0;

  if (model.value.children.length == 0) {
    model.value.configValue.isMuted = true;
  }

  let newNodeName = "key";

  while (true) {
    const idx = model.value.children.find((child: KeyValueConfig) => {
      return child.configValue.getOriginalName() === newNodeName;
    });

    if (idx === undefined) {
      break;
    }

    newNodeName = "key" + model.value.children.length;
  }

  const kv = new KeyValueConfig(new EnvNode(newNodeName, model.value.configValue.value || "value"));

  if (kv.configValue) {
    kv.configValue.isMuted = true;
    kv.configValue.isNew = true;
  }

  model.value.children.push(kv);
  isChildrenFolded.value = false;
  model.value.configValue.value = "";
}
function removeGhostSubNode() {
  if (ghostNodeIdx.value !== undefined) {
    model.value.children.pop();
  }

  ghostNodeIdx.value = undefined;
  if (model.value.children.length == 0) {
    model.value.configValue.value = model.value.configValue.getOriginalValue();
  }
}

// Folding
const isChildrenFolded = ref<boolean>(false);
function toggleFolding() {
  if (!isChildrenFolded.value) {
    fullSize.value = contentBlockRef.value?.clientWidth + "px" || "100vw";
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

// Rollback functionality
const isPreparingToDeleteNewChildren = ref<boolean>(false);

function prepareRollback() {
  if (model.value.configValue.value !== "") {
    // This is a leaf
    // We simply changed the value - no extra preparations needed
    return;
  }

  isPreparingToDeleteNewChildren.value = true;
}

function unPrepareRollback() {
  isPreparingToDeleteNewChildren.value = false;
  model.value.unmute();
}

function calculateHeight(): string {
  const childrenCount = model.value.countChildren();
  if (childrenCount === 0 || isChildrenFolded.value) {
    return "2.1em";
  }

  return `${2.5 + childrenCount * 2.1 + (childrenCount) * 0.6}em`;
}

</script>

<template>
  <div
    class="KeyValueConfigViewWrapper"
    :style="{
      height: calculateHeight(),
    }"
  >
    <div
      class="ControlButton AddButton"
      title="Add new node"
    >
      <Button
        @click="addSubNode"
        @mouseenter="addGhostSubNode"
        @mouseleave="removeGhostSubNode"
        :label="'+'"
        :icon="AddNodeIcon"
      />
    </div>

    <div
      class="ContentWrapper"
      ref="contentBlockRef"
      :style="{minWidth: isChildrenFolded ? fullSize: ''}"
      :class="{'folded': isChildrenFolded}"
    >
      <div
        class="ContentValue"
      >
        <KeyValueNode
          v-if="model.configValue.getOriginalName() !== ''"
          v-model="model.configValue"
          @rollback="model.rollback()"
          @show-original="prepareRollback"
          @show-actual="unPrepareRollback"
        />
      </div>
      <Transition name="children">
        <div
          class="Children"
          v-if="model.children.length > 0"
        >
          <TransitionGroup name="child">
            <div
              class="Child"
              v-for="(_, idx) in model.children"
              :class="{
                ghosted: idx == ghostNodeIdx || (isPreparingToDeleteNewChildren && model.children[idx].configValue.isNew),
                'to-delete': isPreparingToDeleteNewChildren && model.children[idx].configValue.isNew,
                'to-create': !isPreparingToDeleteNewChildren && model.children[idx].configValue.isNew,
              }"
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
      </Transition>
    </div>

    <div
      class="ControlButton"
      v-if="shouldShowFoldButton()"
      :title="isChildrenFolded ? 'Unfold':'Fold'"
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
.KeyValueConfigViewWrapper {
  width: 100%;
  box-sizing: border-box;

  display: flex;
  flex-direction: row;
  gap: 0.5em;

  border-left: #6b7280 solid 1px;

  border-radius: var(--border-radius);
  transition: height 0.2s ease-in-out;
  overflow: hidden;
}

.ContentWrapper {
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 0.5em;
}

.ContentValue {
  min-height: 2em;
  width: 100%;
}

.Children {
  display: flex;
  flex-direction: column;
  gap: 0.5em;
}

.Child {
  width: 100%;
}

.ghosted {
  padding: 0;
  border: 1px dashed #6b7280;
}

.to-delete {
  border-color: var(--warn);
}

.to-create {
  border-color: var(--good);
}

.folded {
  border-bottom: #6b7280 dashed 1px;
}

.ControlButton {
  width: 1.75em;
  height: 1.75em;
  display: flex;

  justify-content: center;
  align-items: center;

  margin-top: 0.125em;
  margin-left: 0.125em;
}

.AddButton {
  position: sticky;
  top: 0;
}

.child-enter-active,
.child-leave-active {
  transition: all 0.3s ease;
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


.children-enter-active,
.children-leave-active {
  transition: all 0.3s ease;
}

.children-enter-to,
.children-leave-from {
  transform: translateY(0);
  opacity: 1;
}

.children-enter-from,
.children-leave-to {
  transform: translateY(-10%);
  opacity: 0;
}


</style>
