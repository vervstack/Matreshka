<script setup lang="ts">

import CreateVervConfigWidget from "@/widget/CreateVervConfigWidget.vue";
</script>

<template>
    <div class="topHeader">New Verv config constructor</div>
    <CreateVervConfigWidget/>
</template>

<style scoped>
.topHeader {
  padding: 1em 0 2em 0;
  font-weight: bold;

  font-size: 1.25em;
  text-align: center;


  background: linear-gradient(90deg, #ff6347, #1e90ff, #32cd32, #ff69b4, #ffd700);
  background-size: 300%;
  background-clip: text;
  -webkit-background-clip: text; /* Safari/Chrome compatibility */
  color: transparent;
  -webkit-text-fill-color: transparent;
  animation: headerTransfusion 5s infinite linear;
}

@keyframes headerTransfusion {
  0% {
    background-position: 0 50%;
  }
  50% {
    background-position: 100% 50%;
  }
  100% {
    background-position: 0 50%;
  }
}

</style>
