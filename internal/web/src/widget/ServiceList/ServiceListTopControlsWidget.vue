<script setup lang="ts">

import {ref} from "vue";

import {SortType} from "@godverv/matreshka";
import {Sort} from "@/models/search/search.ts";

import FloatLabel from "primevue/floatlabel";
import InputGroup from "primevue/inputgroup";
import InputText from "primevue/inputtext";
import InputGroupAddon from "primevue/inputgroupaddon";
import Select from "primevue/select";
import Button from "primevue/button";
import ToggleButton from 'primevue/togglebutton';

const emit = defineEmits<{
  (event: 'updateSearchRequest', pattern: string, sort: Sort): void
}>();

// Search
const searchPattern = ref<string>('');

// Sort
const sorting = ref<Sort>({
  type: SortType.default,
  desc: false,
} as Sort)

const sortOptions = ref([
  {name: 'default', code: SortType.default},
  {name: 'name', code: SortType.by_name},
  {name: 'updated at', code: SortType.by_updated_at},
])

function doSearch() {
  emit('updateSearchRequest', searchPattern.value, sorting.value)
}
</script>

<template>
  <div class="TopControls">
    <InputGroup >
      <InputText
          v-model="searchPattern"
          placeholder="Search"
          @input="doSearch"
      />
      <InputGroupAddon>
        <Button
            icon="pi pi-search"
            severity="secondary"
            variant="text"
            @click="doSearch"
        />
      </InputGroupAddon>
    </InputGroup>

    <InputGroup>
      <FloatLabel variant="in">
        <Select
            inputId="select_sort_type"
            v-model="sorting.type"
            :options="sortOptions"
            optionLabel="name"
            optionValue="code"
            @update:modelValue="doSearch()"
        />
        <label for="select_sort_type">Sort by</label>
      </FloatLabel>
        <ToggleButton
            :style="{width: '8em'}"
            v-model="sorting.desc"
            off-label="ASC"
            off-icon="pi pi-arrow-up"
            on-icon="pi pi-arrow-down"
            on-label="DESC"
            @click="doSearch"
        />
    </InputGroup>
  </div>
</template>

<style scoped>
.TopControls {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(15em, 1fr));
  gap: 1em;
}

.TopControls > * {
  height: 100%;
}
</style>
