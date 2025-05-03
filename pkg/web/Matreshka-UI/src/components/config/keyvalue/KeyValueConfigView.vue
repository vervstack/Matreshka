<script setup lang="ts">

import AddNodeIcon from "@/assets/svg/node/add.svg";
import KeyValueConfig from "@/models/configs/keyvalue/KeyValueConfig.ts";
import KeyValueNode from "@/components/config/keyvalue/nodes/KeyValueNode.vue";
import RootNode from "@/components/config/keyvalue/nodes/RootNode.vue";
import Button from "@/components/base/config/Button.vue";
import { ref } from "vue";

import EnvNode from "@/models/shared/Node.ts";

const contentBlockRef = ref<HTMLElement | null>(null);
const fullSize = ref<string>("100vw");

const model = defineModel<KeyValueConfig>({
  required: true,
});

const props = defineProps({
  parentPrefix: {
    type: String,
    default: "",
  },
  disable: {
    type: Boolean,
    default: false,
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

const isChildrenFolded = ref<boolean>(false);

function toggleFolding() {
  if (!isChildrenFolded.value) {
    fullSize.value = contentBlockRef.value?.clientWidth + "px" || "100vw";
    console.log(fullSize.value);
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

function shouldShowRoot(): boolean {
  if (model.value.configValue.getOriginalName() == "") {
    // Top root - shouldn't show
    return false;
  }

  return model.value.configValue.value === "";
}

// Rollback functionality
const isPreparingToDeleteChildren = ref<boolean>(false);

function prepareRollback() {
  if (model.value.configValue.value !== "") {
    //   We simply changed the value - no extra preparations needed
    return;
  }

  isPreparingToDeleteChildren.value = true;
}

function unPrepareRollback() {
  isPreparingToDeleteChildren.value = false;
}

</script>

<template>
  <div
    class="Wrapper"
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
      :style="{width: isChildrenFolded ? fullSize: ''}"
      :class="{'folded': isChildrenFolded}"
    >
      <div
        class="ContentValue"
        :class="{changed: model.isChanged()}"
      >
        <RootNode
          v-if="shouldShowRoot()"
          v-model="model.configValue"
          @rollback="model.rollback()"
          @show-original="prepareRollback"
          @show-actual="unPrepareRollback"
        />
        <KeyValueNode
          v-else-if="model.configValue.getOriginalName() !==''"
          v-model="model.configValue"
          :parent-prefix="props.parentPrefix"
          :force-root-mode="model.children.length > 0"
        />
      </div>
      <div
        class="Children"
        v-if="model.children.length > 0"
      >
        <TransitionGroup name="child">
          <div
            class="Child"
            v-for="(_, idx) in model.children"
            :class="{
              ghosted: idx == ghostNodeIdx || isPreparingToDeleteChildren,
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

    </div>

    <div
      class="ControlButton"
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
.Wrapper {
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  height: 100%;
  width: fit-content;
  gap: 0.5em;
}

.ContentWrapper {
  display: flex;
  flex-direction: column;
  gap: 0.75em;
  align-items: center;
  justify-content: center;
}

.ContentValue {
  height: 2em;
  width: 100%;
  display: flex;
}

.changed {
  border-bottom: solid 1px;
  border-color: var(--warn);
}

.Children {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  gap: 0.5em;
}

.Child {
  width: 100%;
  height: 100%;
  flex: 1;
  border-left: #6b7280 solid 1px;
  box-sizing: border-box;

  position: relative;
  border-radius: var(--border-radius);
}

.ghosted {
  border: 1px dashed #6b7280;
}

.folded {
  border-bottom: #6b7280 dashed 1px;
}

.ControlButton {
  width: 1.75em;
  height: 1.75em;
  display: flex;
  padding: 0.125em;
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
