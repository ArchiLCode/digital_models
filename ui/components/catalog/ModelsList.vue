<template>
  <div>
    <div class="models-list">
      <ModelCard
        v-for="model in props.models"
        :key="model.user_id"
        :id="model.user_id"
        :image="getAvatarUrl(model.avatar_url)"
        :name="model.name"
        :lastName="model.last_name"
      />
    </div>
  </div>
</template>

<script setup>
import ModelCard from './ModelCard.vue'
import { useRuntimeConfig } from '#imports'

const props = defineProps({
  models: {
    type: Array,
    required: true,
  },
})

const STATIC_BASE =
  useRuntimeConfig().public.staticBaseUrl || 'http://localhost:8082'

function getAvatarUrl(avatarUrl) {
  if (!avatarUrl) return ''
  return STATIC_BASE + avatarUrl
}
</script>
<style>
.models-list {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(155px, 1fr));
  gap: 20px;
  width: 100%;
  max-width: 505px;
  margin: 0 auto;
}

.paginate-buttons {
  padding: 5px 10px;
  border-radius: 0;
  border: 1px solid transparent;
  background-color: transparent;
  color: #f8f8f8;
  cursor: pointer;
  transition: background-color 0.3s ease;
  outline: none;
  transition: 0.2s;
}

@media (max-width: 768px) {
  .models-list {
    grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
    gap: 15px;
  }
}

@media (max-width: 480px) {
  .models-list {
    grid-template-columns: repeat(auto-fill, minmax(120px, 1fr));
    gap: 10px;
  }
}
</style>
