<template>
  <div class="profile-page">
    <h2>Profile:</h2>
    <form @submit.prevent="saveProfile">
      <h3 @click="toggleBaseInfo" :class="{ expanded: isBaseInfoExpanded }">
        Base info
      </h3>
      <div class="base" :class="{ expanded: isBaseInfoExpanded }">
        <InputField
          id="name-profile"
          label="Your name"
          type="text"
          placeholder="Name"
          v-model="profile.name"
          :error="nameError"
          required
        />
        <InputField
          id="lastname-profile"
          label="Your lastname"
          type="text"
          placeholder="Lastname"
          v-model="profile.last_name"
          :error="lastNameError"
          required
        />
      </div>
      <template v-if="isModelUser">
        <div class="form-group">
          <label for="gender">Your gender</label>
          <select
            name="gender"
            id="gender"
            class="gender"
            v-model="profile.sex"
            required
          >
            <option selected="" disabled value="">Make your choice</option>
            <option value="female">Female</option>
            <option value="male">Male</option>
          </select>
        </div>
        <InputField
          id="birthday-profile"
          label="Birthday Date"
          type="date"
          v-model="profile.date_of_birth"
          :error="dateError"
        />
        <div class="form-group">
          <label>Height / Weight</label>
          <div class="height-weight">
            <InputField
              id="height-profile"
              type="number"
              placeholder="X"
              inputClass="half-number-input"
              v-model="profile.height"
              :min="0"
              :max="500"
              :error="heightError"
            />
            <InputField
              id="weight-profile"
              type="number"
              placeholder="X"
              inputClass="half-number-input"
              :min="0"
              :max="500"
              v-model="profile.weight"
              :error="weightError"
            />
          </div>
        </div>

        <div class="form-group">
          <label>Body sizes</label>
          <div class="body-sizes">
            <InputField
              id="breast-profile"
              type="number"
              placeholder="X"
              inputClass="third-number-input"
              :min="0"
              :max="500"
              :error="breastSizeError"
              v-model="profile.chest_size"
            />
            <InputField
              id="waist-profile"
              type="number"
              placeholder="X"
              inputClass="third-number-input"
              :min="0"
              :max="500"
              :error="waistSizeError"
              v-model="profile.waist_size"
            />
            <InputField
              id="hip-profile"
              type="number"
              placeholder="X"
              inputClass="third-number-input"
              :min="0"
              :max="500"
              :error="hipSizeError"
              v-model="profile.hip_size"
            />
          </div>
        </div>
        <InputField
          id="city-profile"
          type="text"
          placeholder="City"
          label="Your city"
          :error="cityError"
          v-model="profile.city"
        />
        <div>
          <label>Your Photos (max 6 items):</label>
          <div
            class="drop-zone"
            @dragover.prevent="handleDragOver"
            @drop.prevent="handleDrop"
            @click="openFilePicker"
          >
            Drop photos here or click for select
            <input
              type="file"
              multiple
              accept="image/png, image/jpeg"
              ref="fileInput"
              @change="handleFileUpload"
              :disabled="profile.photos.length >= 6"
            />
          </div>
          <div v-if="profile.photos.length > 0" class="photo-preview">
            <div
              v-for="(photo, index) in profile.photos"
              :key="index"
              class="photo-item"
              draggable="true"
              @dragstart="handleDragStart(index)"
              @dragover="handleDragOverItem(index)"
              @drop="handleDropItem(index)"
            >
              <img :src="getPhotoUrl(photo)" alt="Фото профиля" />
              <button type="button" @click="removePhoto(index)">
                <svg
                  version="1.1"
                  id="Capa_1"
                  width="18px"
                  xmlns="http://www.w3.org/2000/svg"
                  xmlns:xlink="http://www.w3.org/1999/xlink"
                  x="0px"
                  y="0px"
                  viewBox="0 0 612 612"
                  style="enable-background: new 0 0 612 612"
                  xml:space="preserve"
                >
                  <g>
                    <g id="cross">
                      <g>
                        <polygon
                          points="612,36.004 576.521,0.603 306,270.608 35.478,0.603 0,36.004 270.522,306.011 0,575.997 35.478,611.397 306,341.411 576.521,611.397 612,575.997 341.459,306.011 "
                          fill="#000000"
                          style="fill: rgb(224, 36, 36)"
                        ></polygon>
                      </g>
                    </g>
                  </g>
                </svg>
              </button>
            </div>
          </div>
        </div>
      </template>
      <OutlinedButton type="submit" value="Save" :disabled="!isFormValid" />
    </form>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const { $api } = useNuxtApp()
const { getSession, data } = useAuth()

definePageMeta({
  layout: 'custom',
})

const profile = ref({
  name: '',
  last_name: '',
  sex: '',
  date_of_birth: '',
  chest_size: null,
  waist_size: null,
  hip_size: null,
  height: null,
  weight: null,
  city: '',
  photos: [], // Массив для хранения загруженных фотографий
})

const isFormValid = computed(() => {
  return (
    profile.value.name &&
    profile.value.last_name &&
    profile.value.sex &&
    profile.value.photos.length <= 6
  )
})

const isBaseInfoExpanded = ref(false)

let draggedIndex = null

const fileInput = ref(null)

const isModelUser = computed(() => data.value?.role === 'model')

const toggleBaseInfo = () => {
  isBaseInfoExpanded.value = !isBaseInfoExpanded.value
}

// Открытие диалога выбора файлов
const openFilePicker = () => {
  if (profile.value.photos.length < 6) {
    fileInput.value.click()
  }
}

// Upload photo to the server
const uploadPhoto = async (file) => {
  try {
    const formData = new FormData()
    formData.append('photo', file)

    // $api.post уже возвращает JSON
    const data = await $api.post('http://localhost:3000/upload', formData)
    // Возвращаем массив путей для всех файлов
    if (data.files && Array.isArray(data.files)) {
      return data.files.map((f) => f.path)
    }
    return []
  } catch (error) {
    console.error('Error uploading photo:', error)
    return []
  }
}

// Обработка загрузки файлов
const handleFileUpload = async (event) => {
  const files = event.target.files
  if (files.length + profile.value.photos.length > 6) {
    alert('You can upload up to 6 photos')
    return
  }

  for (let i = 0; i < files.length; i++) {
    const file = files[i]
    const reader = new FileReader()

    reader.onload = async (e) => {
      const localUrl = e.target.result
      const serverPaths = await uploadPhoto(file)
      // serverPaths — массив путей (обычно один, но может быть больше)
      serverPaths.forEach((serverPath) => {
        if (serverPath) {
          profile.value.photos.push({
            url: localUrl, // для превью
            path: serverPath, // путь с бэка
          })
        }
      })
    }

    reader.readAsDataURL(file)
  }

  event.target.value = ''
}

// Удаление фотографии
const removePhoto = (index) => {
  profile.value.photos.splice(index, 1)
}

// Drag-and-drop для загрузки файлов
const handleDragOver = (event) => {
  event.preventDefault()
}

const handleDrop = (event) => {
  event.preventDefault()
  const files = event.dataTransfer.files
  handleFileUpload({ target: { files } })
}

// Drag-and-drop для изменения порядка фотографий
const handleDragStart = (index) => {
  draggedIndex = index
}

const handleDragOverItem = (index) => {
  event.preventDefault()
  if (draggedIndex !== index) {
    const draggedItem = profile.value.photos[draggedIndex]
    profile.value.photos.splice(draggedIndex, 1)
    profile.value.photos.splice(index, 0, draggedItem)
    draggedIndex = index
  }
}

const handleDropItem = (index) => {
  event.preventDefault()
  draggedIndex = null
}

function formatDateToISO(dateString) {
  console.log(dateString)

  const [year, month, day] = dateString.split('-')

  const date = `${year}-${month}-${day}T00:00:00Z`

  return date
}

// Сохранение профиля
const saveProfile = async () => {
  // Используем path вместо server_url
  const photoPaths = profile.value.photos.map((photo) => photo.path)

  const formData = {
    ...profile.value,
    date_of_birth: formatDateToISO(profile.value.date_of_birth),
    photos: photoPaths,
  }

  try {
    const response = await $api.put(
      '/profile?user_id=' + data.value.user_id,
      formData
    )
    console.log('Profile updated successfully:', response)
  } catch (error) {
    console.error('Error updating profile:', error)
  }
}

onMounted(async () => {
  const response = await $api.get('/profile?user_id=' + data.value.user_id)

  // Преобразуем дату, если она пришла в формате {seconds: ...}
  let date_of_birth = response.date_of_birth
  if (
    date_of_birth &&
    typeof date_of_birth === 'object' &&
    'seconds' in date_of_birth
  ) {
    const dateObj = new Date(date_of_birth.seconds * 1000)
    // Формат YYYY-MM-DD
    const yyyy = dateObj.getFullYear()
    const mm = String(dateObj.getMonth() + 1).padStart(2, '0')
    const dd = String(dateObj.getDate()).padStart(2, '0')
    date_of_birth = `${yyyy}-${mm}-${dd}`
  }

  profile.value = { ...profile.value, ...response, date_of_birth }
})

const STATIC_BASE =
  useRuntimeConfig().public.staticBaseUrl || 'http://localhost:8082'

function getPhotoUrl(photo) {
  if (photo.url) return photo.url
  if (photo.path) return STATIC_BASE + photo.path
  if (typeof photo === 'string') return STATIC_BASE + photo
  return ''
}
</script>

<style scoped>
.profile-page {
  max-width: 505px;
  width: 100%;
  margin: 0 auto;
  padding-top: 60px;
  min-height: calc(100vh - 165px);
}

h2 {
  font-size: 20px;
  margin-bottom: 25px;
}

h3 {
  position: relative;
  cursor: pointer;
}

h3::after {
  content: '▼';
  font-size: 14px;
  position: absolute;
  right: 0;
  transition: transform 0.2s ease;
}

h3.expanded::after {
  transform: rotate(180deg);
}

form {
  display: flex;
  flex-direction: column;
  gap: 18px;
  width: 100%;
}

.base {
  display: flex;
  flex-direction: column;
  gap: 18px;
  width: 100%;
  height: 0;
  overflow: hidden;
  opacity: 0;
  transition: height 0.3s ease, opacity 0.3s ease;
}

.base.expanded {
  height: auto;
  opacity: 1;
  padding-top: 10px;
  padding-bottom: 5px;
}

.form-group {
  display: flex;
  position: relative;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
  gap: 20px;
}

label {
  max-width: 95px;
  font-size: 15px;
  font-weight: 600;
}

input,
select {
  font-size: 15px;
  font-weight: 600;
  border: 1px solid #fff;
  background-color: #070707;
  padding: 8px;
  color: #fff;
  width: 390px;
}

select {
  cursor: pointer;
}

select option {
  font-size: 13px;
  color: #f8f8f8;
  background-color: #070707;
  font-weight: 600;
}

input,
select:focus {
  outline: none;
}

.drop-zone {
  border: 1px dashed #f8f8f8;
  padding: 20px;
  text-align: center;
  cursor: pointer;
  margin-top: 10px;
}

.drop-zone input {
  display: none;
}

.photo-preview {
  display: flex;
  align-items: center;
  justify-content: space-between;
  flex-wrap: wrap;
  gap: 20px;
  margin-top: 10px;
}

.photo-item {
  position: relative;
  width: 140px;
  aspect-ratio: 1;
  cursor: grab;
}

.photo-item img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.photo-item button {
  position: absolute;
  top: 0;
  right: 0;
  background: transparent;
  border: none;
  cursor: pointer;
}

.body-sizes,
.height-weight {
  display: flex;
  flex-direction: row;
  width: 390px;
  align-items: center;
  justify-content: center;
}
</style>
