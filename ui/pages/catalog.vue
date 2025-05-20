<template>
  <div class="catalog">
    <ModelsFilter class="filter" @apply-filters="fetchFilteredModels" />
    <h2>Catalog:</h2>
    <ModelsList :models="models" />
    <div class="pagination">
      <div class="pagination-controls">
        <select
          id="items-per-page"
          v-model="itemsPerPage"
          @change="onLimitChange"
        >
          <option value="12">12</option>
          <option value="30">30</option>
          <option value="60">60</option>
        </select>
      </div>
      <vue-awesome-paginate
        :total-items="totalItems"
        :items-per-page="itemsPerPage"
        :max-pages-shown="totalPages"
        v-model="currentPage"
      >
        <template #prev-button>
          <span v-html="prevArrow" class="arrow-icon"></span>
        </template>
        <template #next-button>
          <span v-html="nextArrow" class="arrow-icon"></span>
        </template>
      </vue-awesome-paginate>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import ModelsList from '@/components/catalog/ModelsList.vue'
import ModelsFilter from '@/components/catalog/ModelsFilter.vue'
import { VueAwesomePaginate } from 'vue-awesome-paginate'

const { $api } = useNuxtApp()
const models = ref([])
const totalPages = ref(1)

const totalItems = computed(
  () => Number(models.value.length) * Number(totalPages.value)
)
const itemsPerPage = ref(12)
const currentPage = ref(1)

const lastAppliedFilter = ref({})

definePageMeta({
  layout: 'custom',
})

const prevArrow = `<svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 512 512" fill="#f8f8f8" style="transform: rotate(180deg);">
  <path d="M379.5 378.1c-4.7-2.2-8.9-6.4-10.4-10.4-1.5-4-1.4-12.2 0.3-16 0.8-1.8 17.4-19.2 37.5-39.2l36.1-36-214.7-0.5-214.8-0.5-3.8-2.4c-2.1-1.3-5-4.2-6.5-6.4-2.3-3.4-2.7-5.2-2.7-10.7 0-5.5 0.4-7.3 2.7-10.7 1.5-2.2 4.4-5.1 6.5-6.4l3.8-2.4 214.8-0.5 214.7-0.5-36.1-36c-20.1-20-36.7-37.4-37.5-39.2-1.7-3.8-1.8-12-0.3-16 3.6-9.7 16.7-14.9 26.1-10.5 4.3 2 110.4 106.2 114 112 3.4 5.5 3.3 15.4-0.3 20.8-4.2 6.3-109.4 109.6-113.7 111.6-4.8 2.3-10.7 2.2-15.7-0.1z"/>
</svg>`

const nextArrow = `<svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 512 512" fill="#f8f8f8">
  <path d="M379.5 378.1c-4.7-2.2-8.9-6.4-10.4-10.4-1.5-4-1.4-12.2 0.3-16 0.8-1.8 17.4-19.2 37.5-39.2l36.1-36-214.7-0.5-214.8-0.5-3.8-2.4c-2.1-1.3-5-4.2-6.5-6.4-2.3-3.4-2.7-5.2-2.7-10.7 0-5.5 0.4-7.3 2.7-10.7 1.5-2.2 4.4-5.1 6.5-6.4l3.8-2.4 214.8-0.5 214.7-0.5-36.1-36c-20.1-20-36.7-37.4-37.5-39.2-1.7-3.8-1.8-12-0.3-16 3.6-9.7 16.7-14.9 26.1-10.5 4.3 2 110.4 106.2 114 112 3.4 5.5 3.3 15.4-0.3 20.8-4.2 6.3-109.4 109.6-113.7 111.6-4.8 2.3-10.7 2.2-15.7-0.1z"/>
</svg>`

const fetchFilteredModels = async (filters) => {
  try {
    const queryParams = new URLSearchParams()
    lastAppliedFilter.value = filters

    Object.entries(filters).forEach(([key, value]) => {
      if (
        value !== undefined &&
        value !== null &&
        value !== '' &&
        value !== 'any'
      ) {
        queryParams.append(key, value)
      }
    })

    queryParams.append('page', currentPage.value)
    queryParams.append('limit', itemsPerPage.value)

    const result = await $api.get('/catalog', queryParams)
    const data = JSON.parse(result)

    models.value = data.items
    totalPages.value = data.total_pages
  } catch (error) {
    console.error('Error fetching filtered models:', error)
  }
}

const onLimitChange = () => {
  const firstItemIndex = (currentPage.value - 1) * itemsPerPage.value
  currentPage.value = Math.floor(firstItemIndex / itemsPerPage.value) + 1

  if (currentPage.value > totalPages.value && totalPages.value > 0) {
    currentPage.value = totalPages.value
  }
}

watch(
  [currentPage, itemsPerPage],
  () => {
    fetchFilteredModels(lastAppliedFilter.value)
  },
  { immediate: true }
)
</script>

<style>
.catalog {
  max-width: 505px;
  margin: 0 auto;
  padding-top: 60px;
  position: relative;
}

h2 {
  margin-bottom: 20px;
  font-size: 18px;
  color: #f8f8f8;
}

.pagination-controls {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 10px;
  margin-bottom: 20px;
}

.pagination-controls select {
  padding: 5px 8px;
  background-color: #070707;
  border: none;
  border-bottom: 1px solid #f8f8f8;
  color: #f8f8f8;
  border-radius: 0;
  font-size: 14px;
  cursor: pointer;
}

.pagination {
  display: flex;
  justify-content: center;
  gap: 20px;
  margin-top: 20px;
}

.pagination-container {
  display: flex;
  gap: 6px;
  list-style: none;
}

.paginate-buttons {
  padding: 5px;
  border-radius: 0;
  border: 1px solid transparent;
  background-color: transparent;
  color: #f8f8f8;
  cursor: pointer;
  transition: background-color 0.3s ease;
  outline: none;
  transition: 0.2s;
}

.active-page {
  color: #73e221;
  border-bottom: 1px solid #73e221;
}

.filter {
  position: absolute;
  right: -360px;
}

.arrow-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 24px;
  height: 24px;
}

.arrow-icon :deep(svg) {
  width: 100%;
  height: 100%;
}

/* Стили для активного состояния кнопок */
.vue-awesome-paginate .active .arrow-icon {
  color: white;
}

/* Стили для disabled кнопок */
.vue-awesome-paginate .disabled .arrow-icon {
  opacity: 0.5;
  cursor: not-allowed;
}
</style>
