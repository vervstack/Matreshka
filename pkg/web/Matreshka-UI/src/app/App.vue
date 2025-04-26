<script setup lang="ts">
import Toast from "primevue/toast";

import SideMenu from "@/components/global/side_menu/SideMenu.vue";
import SettingsWidget from "@/widget/SettingsWidget.vue";
</script>

<template>
  <div class="root">
    <Toast />
    <RouterView />
    <SettingsWidget />
    <SideMenu />
  </div>
</template>

<style scoped>
.root {
  padding: 2em;
  width: 100%;
  height: 100%;
}
</style>
