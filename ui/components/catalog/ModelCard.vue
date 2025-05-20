<template>
  <NuxtLink :to="`/model/${id}`" class="model-card">
    <div class="model-image-container">
      <img
        :src="image"
        :alt="name"
        class="model-image"
        loading="lazy"
        @error="onImgError"
        ref="imgRef"
      />
    </div>
    <div class="model-name">
      {{ name + (lastName ? ' ' + lastName : '') }}
    </div>
  </NuxtLink>
</template>

<script setup>
import { ref } from 'vue'
const props = defineProps({
  id: {
    type: [Number, String],
    required: true,
  },
  image: {
    type: String,
    required: true,
  },
  name: {
    type: String,
    required: true,
  },
  lastName: {
    type: String,
  },
})

const imgRef = ref(null)
function onImgError(e) {
  const fallback = '/nouser.png'
  if (imgRef.value) {
    imgRef.value.src = fallback
  } else if (e && e.target) {
    e.target.src = fallback
  }
}
</script>

<style scoped>
.model-card {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 155px;
  overflow: hidden;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  background-color: #2a2a2a;
  transition: transform 0.2s ease, box-shadow 0.2s ease;
  text-decoration: none;
}

.model-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
}

.model-image-container {
  width: 100%;
  aspect-ratio: 1 / 1;
  overflow: hidden;
  display: flex;
  align-items: center;
  justify-content: center;
}

.model-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.model-name {
  padding: 12px;
  font-size: 16px;
  font-weight: 600;
  text-align: center;
  color: #f8f8f8;
}
</style>
