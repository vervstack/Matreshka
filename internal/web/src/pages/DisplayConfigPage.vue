<script setup lang="ts">

import DisplayConfigDisplay from "@/widget/DisplayConfigWidget.vue";

import HomeButton from "@/components/global/HomeButton.vue";

defineProps({
  name: {
    type: String,
    required: true,
  },
})

</script>

<template>
  <div class="DisplayConfigPage">
    <div class="topSection">
      <HomeButton/>
    </div>
    <DisplayConfigDisplay
        :service-name="name"/>
  </div>
</template>

<style scoped>
.DisplayConfigPage {
  display: flex;
  flex-direction: column;
  gap: 2em;
}

.topSection {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}
</style>
