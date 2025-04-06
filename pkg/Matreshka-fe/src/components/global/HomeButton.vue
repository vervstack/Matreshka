<script setup lang="ts">

import {Pages, router} from "@/app/routes/routes.ts";
import {ref} from "vue";

import Button from 'primevue/button'

const canGoBack = ref<boolean>(window.history.state.back !== null)

function back(event: MouseEvent): void {
  if (event.ctrlKey || event.metaKey) {
    window.open(router.resolve({name: Pages.Home}).href, '_blank')
  } else {
    canGoBack.value ?
        router.back()
        :
        router.push({name: Pages.Home})
  }
}
</script>

<template>
  <Button
      :icon="canGoBack ? 'pi pi-arrow-left': 'pi pi-home'"
      rounded
      @click="back"
  />
</template>

<style scoped>

</style>
