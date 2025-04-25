<script setup lang="ts">
import {dateToString} from "@/models/configs/verv/converters/date.ts";

import ConfigName from "@/components/base/config/ConfigName.vue";
import ConfigIcon from "@/components/base/config/ConfigIcon.vue";

import {ConfigBase} from "@/models/configs/config.ts";

defineProps({
  cfgList: {
    type: Object as () => ConfigBase[],
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

          v-for="cfg in cfgList"
          v-tooltip.bottom="`updated at ${cfg.updated_at ? dateToString(cfg.updated_at) : 'NO DATA'}`"

          :key="cfg.name"

          @click="(event: MouseEvent) => { emit('clickService', event, cfg.getMatreshkaName()) }"
      >
        <div class="ConfigTitle">
          <ConfigName :label="cfg.name"/>
        </div>
        <div class="ConfigIcon">
          <ConfigIcon :config-type="cfg.type"/>
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
