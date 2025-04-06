<script setup lang="ts">

const value = defineModel<String>()

function doCopy() {
  if (value.value) {
    navigator.clipboard.writeText(value.value.toString())
  }
}
</script>

<template>
  <button @click="doCopy">Copy</button>
</template>

<style scoped>

input {
  border-radius: 0.5vw;
  padding: 0.25em 0.5em 0.25em 0.5em;
  display: inline-block;

  border: cadetblue solid;
  width: fit-content;
  flex: 4;
}

</style>
