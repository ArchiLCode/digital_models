<template>
  <div>
    <button :type="type" :disabled="disabled">{{ value }}</button>
  </div>
</template>

<script setup>
const props = defineProps({
  type: String,
  value: String,
  disabled: Boolean,
})
</script>

<style scoped>
button {
  width: 100%;
  padding: 0.75rem;
  font-weight: 600;
  border: 1px solid #f8f8f8;
  color: #f8f8f8;
  background-color: transparent;
  cursor: pointer;
}

button:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}
</style>
