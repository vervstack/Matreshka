<script setup lang="ts">

import ConfigPickList from "@/components/base/ConfigPickList.vue";
import ConfigToggle from "@/components/base/ConfigToggle.vue";
import ResourcesList from "@/components/resources/minio/ResourcesList.vue"

import {S3Action, Statement} from "@/models/resource_configs/s3/minio/minio.ts";
import SelectActions from "@/components/resources/minio/SelectActions.vue";

const model = defineModel<Statement>({required: true})

</script>

<template>
  <div class="IamStatement">
    <Transition name="blink" mode="out-in">
      <div
          class="Effect"
          :key="model.allow.value ? 'Allow' : 'Deny'">
        {{ model.allow.value ? 'Allow' : 'Deny' }}
      </div>
    </Transition>

    <ConfigToggle
        v-model="model.allow"
    />
    <SelectActions
        v-model="model.action"
        :options="Object.values(S3Action)"
    />
    <ResourcesList
        v-model="model.resources"
    />
  </div>
</template>

<style scoped>
@import '@/assets/styles/config_display.css';

.IamStatement {
  display: flex;
  flex-direction: column;
  gap: 1em;

  border: #05bc05 solid;
  border-radius: 16px;
  padding: 2vh;
}

.Effect {
  font-size: 1.5em;
  display: inline-block;
  width: fit-content;
  text-decoration: underline;
}

/* Transition classes */
.blink-enter-active {
  animation: blink 0.5s forwards;
}


@keyframes blink {
  0%, 100% {
    background: none;
  }

  50% {
    background: var(--warn);
    color: white;
  }
}
</style>
