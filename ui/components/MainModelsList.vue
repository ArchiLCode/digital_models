<template>
  <div class="main-models-list">
    <h2 class="title">Discover Models</h2>
    <div class="grid-container">
      <div v-for="model in models" :key="model.user_id">
        <NuxtLink :to="`/model/${model.user_id}`" class="grid-item">
          <div class="model-photo">
            <img
              :src="getAvatarUrl(model.avatar_url)"
              :alt="model.name"
              @error="onImgError"
              ref="imgRef"
            />
          </div>
          <div class="model-info">
            <p>{{ model.name }}</p>
            <p>{{ model.last_name }}</p>
          </div>
        </NuxtLink>
      </div>
    </div>
    <NuxtLink to="/catalog" class="about-discover">Show all models</NuxtLink>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRuntimeConfig } from '#imports'

const models = ref([])
const imgRef = ref(null)

const STATIC_BASE =
  useRuntimeConfig().public.staticBaseUrl || 'http://localhost:8082'

async function fetchModels() {
  try {
    const response = await fetch(
      'http://localhost:8080/api/catalog?page=1&limit=8'
    )
    const data = await response.json()
    models.value = data.items || []
  } catch (error) {
    console.error('Ошибка при получении данных:', error)
    models.value = []
  }
}

function getAvatarUrl(avatarUrl) {
  if (!avatarUrl) return '/nouser.png'
  return STATIC_BASE + avatarUrl
}

function onImgError(e) {
  const fallback = '/nouser.png'
  if (e && e.target) {
    e.target.src = fallback
  }
}

onMounted(() => {
  fetchModels()
})
</script>

<style scoped>
.main-models-list {
  margin: 0 auto;
  width: 70%;
}

.title {
  font-size: 16px;
  text-align: left;
  margin-bottom: 20px;
  padding-bottom: 15px;
  border-bottom: 1px solid #f8f8f8;
}

.grid-container {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 30px;
}

.grid-item {
  display: flex;
  align-items: center;
  justify-content: flex-start;
  gap: 20px;
  padding: 10px;
}

.model-photo img {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  background-color: #2a2a2a;
}

a {
  display: inline-block;
  margin-top: 10px;
  color: #f8f8f8;
  font-weight: 600;
  font-size: 15px;
}
</style>
