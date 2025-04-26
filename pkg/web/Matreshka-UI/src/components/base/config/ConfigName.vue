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
  display: flex;
  justify-content: center;
  overflow: hidden;
}

.LabelText {
  box-sizing: content-box;
  width: fit-content;
  overflow: hidden;
  text-overflow: ellipsis;
}
</style>
