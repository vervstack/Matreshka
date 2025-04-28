<script setup lang="ts">
defineProps({
  label: {
    type: String,
  },
});
</script>

<template>
  <div class="Label">
    <p class="LabelText">
      {{ label }}
    </p>
  </div>
</template>

<style scoped>
.Label {
  width: 100%;
  height: 100%;

  display: flex;
  justify-content: center;
  overflow: hidden;
}

.LabelText {
  width: 100%;
  height: 100%;

  margin: 0;

  display: flex;
  justify-content: center;
  align-items: center;

  box-sizing: content-box;
  overflow: hidden;
  text-overflow: ellipsis;
}
</style>
