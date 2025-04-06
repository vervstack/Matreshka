<script setup lang="ts">
import {AppInfoClass} from "@/models/configs/verv/info/AppInfo.ts";
import {dateToString} from "@/models/configs/verv/converters/date.ts";

import ConfigName from "@/components/base/config/ConfigName.vue";
import ConfigIcon from "@/components/base/config/ConfigIcon.vue";

import {ConfigTypePrefix} from "@godverv/matreshka";

defineProps({
  servicesList: {
    type: Object as () => AppInfoClass[],
  }
})

const emit = defineEmits<{
  (event: 'clickService', mouseEvent: MouseEvent, name: string): void
}>()

</script>

<template>
  <div class="Wrapper">
    <div class="ListServices">
      <div
          class="ListItem"

          v-for="service in servicesList"
          v-tooltip.bottom="`updated at ${service.updated_at ? dateToString(service.updated_at) : 'NO DATA'}`"

          :key="service.name.value"

          @click="(event: MouseEvent) => { emit('clickService', event, service.name.value) }"
      >
        <div class="ConfigTitle">
          <ConfigName :label="service.name.value"/>
        </div>
        <div class="ConfigIcon">
          <ConfigIcon :config-type="ConfigTypePrefix.verv"/>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.Wrapper {
  display: flex;
  justify-content: center;
}

.ListServices {
  width: 100%;
  height: 100%;

  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(30vw, 1fr));
  gap: 2em;
  place-self: center;
}

.ListItem {
  max-height: 6em;
  overflow: hidden;
  border: var(--good) solid;

  border-radius: 16px;

  padding: 1em 0.75em 1em 0.75em;

  display: flex;
  gap: 1em;
  justify-content: space-around;
  align-items: center;

  cursor: pointer;
}

.ConfigTitle {
  flex: 3;
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
}

.ConfigIcon {
  flex: 1;
  width: 100%;
  height: 100%;
}

</style>
