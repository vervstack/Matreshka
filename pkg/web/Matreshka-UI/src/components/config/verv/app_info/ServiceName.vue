<script setup lang="ts">
import {ref} from "vue";

import Button from "primevue/button";

import {ConfigValueClass} from "@/models/shared/common.ts";
import ConfigField from "@/components/base/config/fields/ConfigInput.vue";
import {ExtractSourceCodeSystemFromServiceName, PiIconFromSourceCodeSystem} from "@/models/configs/verv/info/AppInfo.ts";

const model = defineModel<ConfigValueClass<string>>({required: true})

const linkIcon = ref<String | undefined>(
    PiIconFromSourceCodeSystem(
        ExtractSourceCodeSystemFromServiceName(model.value.value)))

</script>

<template>
  <div class="ServiceName">
    <div class="ServiceLink" v-show="linkIcon">
      <Button
          :icon="'pi '+linkIcon "
          severity="secondary"
          :link="true"
          as="a"
          :href="'https://'+model.value"
          target="_blank"
      />
    </div>
    <ConfigField
        v-model="model"
        fieldName="Service name"
        :pre-addons="[]"
    />
  </div>
</template>

<style scoped>
.ServiceName {
  display: flex;
  flex-direction: row;
  align-items: center;
  gap: 0.1em;
}

.ServiceLink {
  padding: 0.35em 0 0 0;
}

a {
  text-decoration: none;
}
</style>
