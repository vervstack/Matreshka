<script
    setup
    lang="ts"
>

import {ConfigValue} from '@/models/shared/common';
import ToggleSwitch from 'primevue/toggleswitch';

const model = defineModel<ConfigValue<boolean>>({required: true})

</script>

<template>
  <div class="ConfigToggle">
    <div> {{ model.envName }}</div>
    <ToggleSwitch
        v-model="model.value"
    >
      <template #handle="{ checked }">
        <i :class="['!text-xs pi', { 'pi-check': checked, 'pi-times': !checked }]"/>
      </template>
    </ToggleSwitch>
  </div>
</template>

<style scoped>
.ConfigToggle {
  display: flex;
  flex-direction: row;
  width: 100%;
  gap: 1em;
  align-items: center;

}

</style>
