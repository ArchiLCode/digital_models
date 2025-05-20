<template>
  <div class="filters-container">
    <div class="filter-group">
      <label>Возраст</label>
      <div class="range-values">
        <span>от {{ ageFrom }}</span>
        <span>до {{ ageTo }}</span>
      </div>
      <div class="double-range-container">
        <div class="range-track"></div>
        <div
          class="range-progress"
          :style="{
            left: ((ageFrom - 18) / (60 - 18)) * 100 + '%',
            right: 100 - ((ageTo - 18) / (60 - 18)) * 100 + '%',
          }"
        ></div>
        <input
          type="range"
          v-model="ageFrom"
          min="18"
          max="60"
          @input="validateAgeRange"
          class="range-slider range-from"
        />
        <input
          type="range"
          v-model="ageTo"
          min="18"
          max="60"
          @input="validateAgeRange"
          class="range-slider range-to"
        />
      </div>
    </div>

    <div class="filter-group">
      <label>Пол</label>
      <select v-model="sex">
        <option value="any">Любой</option>
        <option value="male">Мужской</option>
        <option value="female">Женский</option>
      </select>
    </div>

    <div class="filter-group">
      <label>Рост</label>
      <div class="range-values">
        <span>от {{ heightFrom }}</span>
        <span>до {{ heightTo }}</span>
      </div>
      <div class="double-range-container">
        <div class="range-track"></div>
        <div
          class="range-progress"
          :style="{
            left: ((heightFrom - 150) / (200 - 150)) * 100 + '%',
            right: 100 - ((heightTo - 150) / (200 - 150)) * 100 + '%',
          }"
        ></div>
        <input
          type="range"
          v-model="heightFrom"
          min="150"
          max="200"
          @input="validateHeightRange"
          class="range-slider range-from"
        />
        <input
          type="range"
          v-model="heightTo"
          min="150"
          max="200"
          @input="validateHeightRange"
          class="range-slider range-to"
        />
      </div>
    </div>

    <div class="filter-group">
      <label>Вес</label>
      <div class="range-values">
        <span>от {{ weightFrom }}</span>
        <span>до {{ weightTo }}</span>
      </div>
      <div class="double-range-container">
        <div class="range-track"></div>
        <div
          class="range-progress"
          :style="{
            left: ((weightFrom - 40) / (120 - 40)) * 100 + '%',
            right: 100 - ((weightTo - 40) / (120 - 40)) * 100 + '%',
          }"
        ></div>
        <input
          type="range"
          v-model="weightFrom"
          min="40"
          max="120"
          @input="validateWeightRange"
          class="range-slider range-from"
        />
        <input
          type="range"
          v-model="weightTo"
          min="40"
          max="120"
          @input="validateWeightRange"
          class="range-slider range-to"
        />
      </div>
    </div>

    <button class="apply-btn" @click="applyFilters">Поиск</button>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const ageFrom = ref(18)
const ageTo = ref(60)
const sex = ref('any')
const heightFrom = ref(150)
const heightTo = ref(200)
const weightFrom = ref(40)
const weightTo = ref(120)

const emit = defineEmits(['apply-filters'])

// Функции валидации для каждого диапазона
const validateAgeRange = () => {
  // Убедимся, что "от" не больше чем "до"
  if (Number(ageFrom.value) > Number(ageTo.value)) {
    ageFrom.value = ageTo.value
  }
}

const validateHeightRange = () => {
  if (Number(heightFrom.value) > Number(heightTo.value)) {
    heightFrom.value = heightTo.value
  }
}

const validateWeightRange = () => {
  if (Number(weightFrom.value) > Number(weightTo.value)) {
    weightFrom.value = weightTo.value
  }
}

const applyFilters = () => {
  const filters = {
    ageFrom: ageFrom.value,
    ageTo: ageTo.value,
    sex: sex.value,
    heightFrom: heightFrom.value,
    heightTo: heightTo.value,
    weightFrom: weightFrom.value,
    weightTo: weightTo.value,
  }

  emit('apply-filters', filters)
}

onMounted(() => {
  applyFilters()
})
</script>

<style scoped>
.filters-container {
  max-width: 280px;
  width: 100%;
  background-color: #2a2a2a;
  padding: 30px;
  margin-bottom: 30px;
}

.filter-group {
  margin-bottom: 20px;
}

label {
  display: block;
  margin-bottom: 10px;
  color: #f8f8f8;
  font-size: 18px;
  font-weight: bold;
}

.range-values {
  display: flex;
  justify-content: space-between;
  margin-bottom: 5px;
}

.range-values span {
  font-size: 14px;
  color: #c4c4c4;
}

/* Контейнер для двойного ползунка */
.double-range-container {
  position: relative;
  height: 30px;
  display: flex;
  align-items: center;
}

/* Базовая полоса ползунка */
.range-track {
  position: absolute;
  width: 100%;
  height: 2px;
  background-color: #a5a5a5;
}

/* Полоса прогресса между двумя ползунками */
.range-progress {
  position: absolute;
  height: 2px;
  background-color: #73e221;
  pointer-events: none;
}

/* Стиль для ползунков */
.range-slider {
  position: absolute;
  width: 100%;
  -webkit-appearance: none;
  appearance: none;
  height: 2px;
  background: transparent;
  pointer-events: none;
  z-index: 3;
  margin: 0;
}

/* Стиль для "бегунка" на ползунке */
.range-slider::-webkit-slider-thumb {
  -webkit-appearance: none;
  appearance: none;
  width: 10px;
  height: 10px;
  border-radius: 50%;
  background: #73e221;
  cursor: pointer;
  pointer-events: auto;
  margin-top: 0;
  transform: translateY(-5px);
}

.range-slider::-moz-range-thumb {
  width: 10px;
  height: 10px;
  border-radius: 50%;
  background: #73e221;
  cursor: pointer;
  pointer-events: auto;
}

/* Скрываем стандартный трек ползунка */
.range-slider::-webkit-slider-runnable-track {
  width: 100%;
  height: 0;
  background: transparent;
  border: none;
}

.range-slider::-moz-range-track {
  width: 100%;
  height: 0;
  background: transparent;
  border: none;
}

/* Стиль для ползунка нижнего значения */
.range-from {
  z-index: 4;
}

/* Стиль для ползунка верхнего значения */
.range-to {
  z-index: 5;
}

select {
  width: 100%;
  padding: 8px;
  background-color: #2a2a2a;
  color: #f8f8f8;
  border: 1px solid #f8f8f8;
  cursor: pointer;
}

.apply-btn {
  width: 100%;
  padding: 10px;
  background-color: transparent;
  color: white;
  border: 1px solid #f8f8f8;
  cursor: pointer;
  margin-top: 10px;
  font-weight: bold;
}
</style>
