<script
    setup
    lang="ts"
    generic="T extends any, B  extends ConfigValue<T[]>">

import MultiSelect from 'primevue/multiselect';
import FloatLabel from "primevue/floatlabel";
import {ConfigValue} from "@/models/shared/common.ts";
import InputGroup from "primevue/inputgroup";

const model = defineModel<B>({required: true})

const props = defineProps({
  options: {
    type: Array as () => T[],
    required: true
  },
})

</script>

<template>
  <div class="MultiSelect">
    <InputGroup>
      <FloatLabel variant="on">
        <MultiSelect
            v-model="model.value"
            display="chip"
            :options="props.options"
        />
        <label>{{ model.envName }}</label>
      </FloatLabel>
    </InputGroup>
  </div>
</template>

<style scoped>
.MultiSelect {
  display: flex;
  width: 100%;
}

label {
  display: inline-block;
  overflow: hidden;
}

</style>
