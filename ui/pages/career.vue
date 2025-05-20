<template>
  <div class="career-page">
    <h2>Career:</h2>
    <p>
      Сегодня в компании ÉLEFÉ работают специалисты высокого уровня, обладающие
      знаниями и опытом в различных областях. Мы уверенны, что основа успеха
      любого проекта - это профессиональная команда.
    </p>
    <p>
      Отдел управления персоналом:<br />
      <a href="tel: +375333665053">+375 33 366 50 53</a><br />
      <a href="mailto: hr@elefe.by">hr@elefe.by</a>
    </p>

    <div v-if="isAdmin" class="admin-controls">
      <button class="admin-button" @click="showAddVacancyForm = true">
        Добавить вакансию
      </button>
    </div>

    <h2>Актуальные вакансии:</h2>
    <div class="vacancies">
      <div
        v-for="vacancy in vacancies"
        :key="vacancy.id"
        class="vacancy"
        :class="{ active: activeVacancy === vacancy.id }"
      >
        <div class="vacancy-header">
          <h3 @click="toggleVacancy(vacancy.id)">{{ vacancy.title }}</h3>
          <div v-if="isAdmin" class="vacancy-actions">
            <button class="action-button edit" @click="editVacancy(vacancy)">
              Редактировать
            </button>
            <button
              class="action-button delete"
              @click="deleteVacancy(vacancy.id)"
            >
              Удалить
            </button>
          </div>
        </div>
        <transition name="slide">
          <div v-if="activeVacancy === vacancy.id" class="description">
            <p><strong>Обязанности:</strong></p>
            <ul>
              <li
                v-for="(duty, index) in vacancy.responsibilities"
                :key="index"
              >
                {{ duty }}
              </li>
            </ul>
            <p><strong>Требования:</strong></p>
            <ul>
              <li
                v-for="(requirement, index) in vacancy.requirements"
                :key="index"
              >
                {{ requirement }}
              </li>
            </ul>
            <p><strong>Условия:</strong></p>
            <ul>
              <li v-for="(condition, index) in vacancy.conditions" :key="index">
                {{ condition }}
              </li>
            </ul>
          </div>
        </transition>
      </div>
    </div>

    <!-- Add Vacancy Modal -->
    <div v-if="showAddVacancyForm" class="modal">
      <div class="modal-content">
        <h3>
          {{ editingVacancy ? 'Редактировать вакансию' : 'Добавить вакансию' }}
        </h3>
        <form @submit.prevent="saveVacancy">
          <div class="form-group">
            <label>Название:</label>
            <input v-model="vacancyForm.title" required />
          </div>

          <div class="form-group">
            <label>Обязанности:</label>
            <div
              v-for="(duty, index) in vacancyForm.responsibilities"
              :key="'duty-' + index"
              class="list-item"
            >
              <input v-model="vacancyForm.responsibilities[index]" required />
              <button
                type="button"
                class="remove-button"
                @click="removeItem('responsibilities', index)"
              >
                Удалить
              </button>
            </div>
            <button
              type="button"
              class="add-button"
              @click="addItem('responsibilities')"
            >
              Добавить обязанность
            </button>
          </div>

          <div class="form-group">
            <label>Требования:</label>
            <div
              v-for="(req, index) in vacancyForm.requirements"
              :key="'req-' + index"
              class="list-item"
            >
              <input v-model="vacancyForm.requirements[index]" required />
              <button
                type="button"
                class="remove-button"
                @click="removeItem('requirements', index)"
              >
                Удалить
              </button>
            </div>
            <button
              type="button"
              class="add-button"
              @click="addItem('requirements')"
            >
              Добавить требование
            </button>
          </div>

          <div class="form-group">
            <label>Условия:</label>
            <div
              v-for="(cond, index) in vacancyForm.conditions"
              :key="'cond-' + index"
              class="list-item"
            >
              <input v-model="vacancyForm.conditions[index]" required />
              <button
                type="button"
                class="remove-button"
                @click="removeItem('conditions', index)"
              >
                Удалить
              </button>
            </div>
            <button
              type="button"
              class="add-button"
              @click="addItem('conditions')"
            >
              Добавить условие
            </button>
          </div>

          <div class="form-actions">
            <button type="submit" class="save-button">Сохранить</button>
            <button type="button" class="cancel-button" @click="closeModal">
              Отмена
            </button>
          </div>
        </form>
      </div>
    </div>

    <!-- Delete Confirmation Modal -->
    <div v-if="showDeleteConfirm" class="modal">
      <div class="modal-content delete-confirm">
        <h3>Подтвердите удаление</h3>
        <p>Вы уверены, что хотите удалить эту вакансию?</p>
        <div class="form-actions">
          <button @click="confirmDelete" class="delete-button">Удалить</button>
          <button @click="showDeleteConfirm = false" class="cancel-button">
            Отмена
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'

definePageMeta({
  layout: 'custom',
})

const { data: session } = useAuth()

const isAdmin = computed(() => {
  return session.value?.is_admin === true
})

const activeVacancy = ref(null)
const vacancies = ref([
  {
    id: 1,
    title: 'Scout',
    responsibilities: ['Поиск и подбор перспективных моделей.'],
    requirements: [
      'Опыт работы скаутом от 1 года.',
      'Хорошие коммуникативные навыки.',
      'Стрессоустойчивость, способность работать с большим объемом информации.',
      'Желание и умение учиться.',
    ],
    conditions: [
      'Медицинская страховка после испытательного срока.',
      'Отличные возможности профессионального и карьерного роста.',
      'Высокая заработная плата.',
      'График работы полный рабочий день.',
    ],
  },
  {
    id: 2,
    title: 'Менеджер проекта',
    responsibilities: ['Управление проектами.', 'Координация работы команды.'],
    requirements: [
      'Опыт управления проектами от 2 лет.',
      'Навыки планирования и организации.',
    ],
    conditions: ['Гибкий график работы.', 'Конкурентная заработная плата.'],
  },
  {
    id: 3,
    title: 'Менеджер по работе с партнерами',
    responsibilities: [
      'Поиск и привлечение партнеров.',
      'Поддержание долгосрочных отношений с партнерами.',
    ],
    requirements: [
      'Опыт работы в продажах или партнерском менеджменте.',
      'Навыки ведения переговоров.',
    ],
    conditions: ['Бонусы за выполнение планов.', 'Обучение за счет компании.'],
  },
])

// Form state
const showAddVacancyForm = ref(false)
const showDeleteConfirm = ref(false)
const editingVacancy = ref(null)
const vacancyToDelete = ref(null)
const vacancyForm = ref({
  title: '',
  responsibilities: [''],
  requirements: [''],
  conditions: [''],
})

const toggleVacancy = (id) => {
  activeVacancy.value = activeVacancy.value === id ? null : id
}

// Admin functions
const editVacancy = (vacancy) => {
  editingVacancy.value = vacancy.id
  vacancyForm.value = {
    title: vacancy.title,
    responsibilities: [...vacancy.responsibilities],
    requirements: [...vacancy.requirements],
    conditions: [...vacancy.conditions],
  }
  showAddVacancyForm.value = true
}

const deleteVacancy = (id) => {
  vacancyToDelete.value = id
  showDeleteConfirm.value = true
}

const confirmDelete = () => {
  // Here you would normally make an API call
  vacancies.value = vacancies.value.filter(
    (vacancy) => vacancy.id !== vacancyToDelete.value
  )
  showDeleteConfirm.value = false
  vacancyToDelete.value = null
}

const addItem = (section) => {
  vacancyForm.value[section].push('')
}

const removeItem = (section, index) => {
  if (vacancyForm.value[section].length > 1) {
    vacancyForm.value[section].splice(index, 1)
  }
}

const saveVacancy = () => {
  // Remove any empty entries
  const cleanForm = {
    ...vacancyForm.value,
    responsibilities: vacancyForm.value.responsibilities.filter((item) =>
      item.trim()
    ),
    requirements: vacancyForm.value.requirements.filter((item) => item.trim()),
    conditions: vacancyForm.value.conditions.filter((item) => item.trim()),
  }

  if (editingVacancy.value) {
    // Update existing vacancy
    const index = vacancies.value.findIndex(
      (v) => v.id === editingVacancy.value
    )
    if (index !== -1) {
      vacancies.value[index] = {
        ...vacancies.value[index],
        ...cleanForm,
      }
    }
  } else {
    // Add new vacancy with generated ID
    const newId = Math.max(0, ...vacancies.value.map((v) => v.id)) + 1
    vacancies.value.push({
      id: newId,
      ...cleanForm,
    })
  }

  closeModal()
}

const closeModal = () => {
  showAddVacancyForm.value = false
  editingVacancy.value = null
  vacancyForm.value = {
    title: '',
    responsibilities: [''],
    requirements: [''],
    conditions: [''],
  }
}
</script>

<style scoped>
.career-page {
  max-width: 505px;
  width: 100%;
  margin: 0 auto;
  min-height: calc(100vh - 165px);
  position: relative;
}

h2 {
  margin-top: 60px;
  color: #f8f8f8;
  font-size: 18px;
}

p {
  color: #f8f8f8;
  font-weight: 600;
  font-size: 16px;
}

a {
  font-weight: 600;
  color: #939393;
}

.vacancies {
  margin-top: 20px;
}

.vacancy {
  margin-bottom: 15px;
  transition: background-color 0.3s ease;
  padding-bottom: 10px;
}

.vacancy-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.vacancy h3 {
  font-size: 17px;
  text-decoration: underline;
  transition: 0.2s;
  cursor: pointer;
}

.vacancy.active h3 {
  color: #939393;
}

.description {
  margin-top: 10px;
  padding-top: 10px;
  overflow: hidden;
}

ul {
  list-style-type: none;
  margin-left: 40px;
}

p {
  margin: 10px 0;
}

/* Admin controls */
.admin-controls {
  margin: 20px 0;
}

.admin-button {
  background-color: #4a4a4a;
  color: white;
  border: none;
  padding: 8px 16px;
  border-radius: 4px;
  cursor: pointer;
  font-weight: 600;
  transition: background-color 0.3s;
}

.admin-button:hover {
  background-color: #5a5a5a;
}

.vacancy-actions {
  display: flex;
  gap: 10px;
}

.action-button {
  border: none;
  padding: 4px 8px;
  border-radius: 3px;
  cursor: pointer;
  font-size: 12px;
  font-weight: 600;
}

.edit {
  background-color: #4a4a4a;
  color: white;
}

.delete {
  background-color: #933;
  color: white;
}

/* Modal */
.modal {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.7);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal-content {
  background-color: #2a2a2a;
  border-radius: 8px;
  padding: 20px;
  width: 90%;
  max-width: 500px;
  max-height: 80vh;
  overflow-y: auto;
  color: #f8f8f8;
}

.delete-confirm {
  text-align: center;
  max-width: 400px;
}

.form-group {
  margin-bottom: 15px;
}

.form-group label {
  display: block;
  margin-bottom: 5px;
  font-weight: 600;
}

.form-group input {
  width: 100%;
  padding: 8px;
  background-color: #3a3a3a;
  border: 1px solid #4a4a4a;
  color: #f8f8f8;
  border-radius: 4px;
}

.list-item {
  display: flex;
  gap: 10px;
  margin-bottom: 8px;
}

.add-button,
.remove-button {
  background-color: #4a4a4a;
  color: white;
  border: none;
  padding: 4px 8px;
  border-radius: 3px;
  cursor: pointer;
  font-size: 12px;
}

.remove-button {
  background-color: #933;
}

.form-actions {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
  margin-top: 20px;
}

.save-button,
.cancel-button,
.delete-button {
  padding: 8px 16px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-weight: 600;
}

.save-button {
  background-color: #4a4a4a;
  color: white;
}

.cancel-button {
  background-color: #3a3a3a;
  color: #f8f8f8;
}

.delete-button {
  background-color: #933;
  color: white;
}

/* Анимация для раскрытия/сворачивания */
.slide-enter-active,
.slide-leave-active {
  transition: max-height 0.3s ease-in-out;
}

.slide-enter-from,
.slide-leave-to {
  max-height: 0;
  opacity: 0;
}

.slide-enter-to,
.slide-leave-from {
  max-height: 500px; /* Укажите максимальную высоту, достаточную для контента */
  opacity: 1;
}
</style>
