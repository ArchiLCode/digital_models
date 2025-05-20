<template>
  <div v-if="status === 'unauthenticated'" class="unauth-warning">
    <h2>Для просмотра профиля необходимо авторизоваться</h2>
  </div>
  <div v-else-if="model" class="model-profile">
    <div class="profile-header">
      <h2 class="model-name">
        {{ model.name + (model.last_name ? ' ' + model.last_name : '') }}
        <span v-if="model.verified" class="verified-badge">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            version="1.0"
            width="512.000000pt"
            height="512.000000pt"
            viewBox="0 0 512.000000 512.000000"
            preserveAspectRatio="xMidYMid meet"
          >
            <g
              transform="translate(0.000000,512.000000) scale(0.100000,-0.100000)"
              fill="#7fff00"
              stroke="none"
            >
              <path
                d="M2365 5114 c-348 -35 -633 -113 -920 -253 -675 -329 -1170 -931 -1360 -1655 -63 -238 -79 -369 -79 -646 0 -205 4 -267 22 -380 88 -538 319 -1005 685 -1386 404 -421 906 -679 1497 -770 86 -14 174 -18 350 -18 176 0 264 4 350 18 284 43 522 117 765 235 675 329 1170 931 1360 1655 63 238 79 369 79 646 0 205 -4 267 -22 380 -88 538 -319 1005 -685 1386 -398 414 -896 674 -1467 765 -120 19 -473 33 -575 23z m1534 -1456 c50 -16 127 -91 148 -146 25 -66 22 -147 -8 -207 -34 -67 -1675 -1697 -1750 -1738 -40 -21 -64 -27 -113 -27 -83 1 -127 23 -233 118 -465 420 -823 751 -845 783 -62 93 -55 222 18 308 78 93 242 108 339 31 33 -26 505 -456 653 -594 l53 -50 747 746 c411 410 761 753 779 762 69 35 131 39 212 14z"
              />
            </g></svg
        ></span>
      </h2>
      <div class="model-details">
        <div class="details-column">
          <div class="detail-item">
            <span class="detail-label">Age:</span>
            <span class="detail-value">{{ model.age }} y.o.</span>
          </div>
          <div class="detail-item">
            <span class="detail-label">Gender:</span>
            <span class="detail-value">{{ model.sex }}</span>
          </div>
          <div class="detail-item">
            <span class="detail-label">City:</span>
            <span class="detail-value">{{ model.city }}</span>
          </div>
        </div>

        <div class="details-column">
          <div class="detail-item">
            <span class="detail-label">Height:</span>
            <span class="detail-value">{{ model.height }} см</span>
          </div>
          <div class="detail-item">
            <span class="detail-label">Weight:</span>
            <span class="detail-value">{{ model.weight }} kg</span>
          </div>
          <div class="detail-item">
            <span class="detail-label">Parameters:</span>
            <span class="detail-value">{{ model.parameters }}</span>
          </div>
        </div>
        <div class="detail-item space-between">
          <div class="detail-item">
            <span class="detail-label">Profile:</span>
            <a :href="profileUrl" target="_blank" class="profile-link"
              >{{}} <span class="link-icon"></span
            ></a>
            <button
              class="copy-button"
              @click="copyProfileLink"
              :title="copyButtonTooltip"
            >
              <span class="copy-icon">
                <svg
                  v-if="isCopied"
                  viewBox="0 0 24 24"
                  width="24"
                  height="24"
                  fill="none"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <g id="SVGRepo_bgCarrier" stroke-width="0"></g>
                  <g
                    id="SVGRepo_tracerCarrier"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  ></g>
                  <g id="SVGRepo_iconCarrier">
                    <rect width="20" height="20" fill="none"></rect>
                    <path
                      d="M5 13.3636L8.03559 16.3204C8.42388 16.6986 9.04279 16.6986 9.43108 16.3204L19 7"
                      stroke="#7fff00"
                      stroke-linecap="round"
                      stroke-linejoin="round"
                    ></path>
                  </g>
                </svg>
                <svg
                  v-else
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 24 24"
                  width="24"
                  height="24"
                  fill="none"
                  preserveAspectRatio="xMidYMid meet"
                >
                  <g
                    transform="translate(0, 24) scale(0.005, -0.005)"
                    fill="#7fff00"
                    stroke="none"
                  >
                    <path
                      d="M1890 4683 c-141 -24 -294 -106 -401 -216 -140 -143 -209 -321 -209 -538 l0 -89 -89 0 c-216 0 -392 -68 -536 -207 -98 -94 -154 -190 -207 -348 -10 -32 -13 -256 -13 -1155 l0 -1115 28 -79 c80 -227 247 -394 472 -473 l80 -28 1120 0 1120 0 72 26 c123 45 202 96 293 188 71 71 94 103 133 181 26 52 54 127 63 166 l16 72 136 4 c116 4 149 9 216 32 225 77 394 245 473 471 l28 80 0 1225 0 1225 -28 79 c-80 227 -246 393 -473 473 l-79 28 -1095 1 c-602 1 -1106 0 -1120 -3z m2190 -444 c80 -37 115 -72 152 -147 l33 -67 0 -1145 0 -1145 -33 -67 c-57 -116 -145 -168 -299 -176 l-93 -5 0 842 c0 922 0 920 -60 1063 -70 168 -228 324 -399 394 -131 53 -153 54 -942 54 l-729 0 0 84 c0 110 21 174 78 237 49 55 101 86 165 99 23 4 499 7 1057 6 l1015 -1 55 -26z m-832 -863 c60 -33 101 -76 134 -139 l23 -42 0 -1060 0 -1060 -27 -50 c-34 -65 -101 -127 -164 -151 -47 -18 -90 -19 -1084 -19 l-1035 0 -48 22 c-80 37 -123 78 -159 151 l-33 67 0 1035 c0 994 1 1037 19 1084 32 83 115 159 206 187 14 4 495 7 1070 6 l1045 -2 53 -29z"
                    />
                  </g>
                </svg>
              </span>
            </button>
          </div>

          <div class="hype-section">
            <div class="hype-icon">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                version="1.1"
                xmlns:xlink="http://www.w3.org/1999/xlink"
                width="22"
                height="22"
                x="0"
                y="0"
                viewBox="0 0 22 22"
                xml:space="preserve"
                class=""
              >
                <g transform="scale(0.04)">
                  <path
                    d="M396.619 194.325c-1.381-1.73-3.055-1.39-3.932-1.037-.735.298-2.419 1.25-2.186 3.636.28 2.865.437 5.786.467 8.683.125 12.018-4.696 23.792-13.226 32.303-8.476 8.456-19.624 13.012-31.494 12.88-16.214-.207-29.662-8.664-38.891-24.457-7.631-13.059-4.277-29.902-.726-47.735 2.078-10.438 4.227-21.232 4.227-31.505 0-79.989-53.774-126.137-85.828-146.526-.663-.421-1.294-.569-1.853-.569-.909 0-1.629.392-1.984.632-.688.466-1.789 1.528-1.435 3.408 12.252 65.062-24.292 104.193-62.982 145.621-39.88 42.703-85.081 91.104-85.081 178.396C71.695 429.483 154.212 512 255.64 512c83.512 0 157.143-58.224 179.055-141.59 14.942-56.843-.716-129.251-38.076-176.085zm-136.39 278.412c-25.398 1.158-49.552-7.951-68.001-25.591-18.251-17.452-28.719-41.807-28.719-66.821 0-46.941 17.948-81.401 66.222-127.149.79-.749 1.599-.986 2.304-.986.639 0 1.193.195 1.574.378.803.387 2.123 1.345 1.945 3.417-1.726 20.085-1.696 36.756.088 49.552 4.56 32.685 28.487 54.646 59.542 54.646 15.226 0 29.729-5.73 40.838-16.135a3.643 3.643 0 0 1 3.281-.937c.731.158 1.71.606 2.223 1.843 4.606 11.121 6.96 22.926 6.996 35.085.147 48.925-39.461 90.51-88.293 92.698z"
                    fill="#7fff00"
                    opacity="1"
                    data-original="#000000"
                    class=""
                  ></path>
                </g>
              </svg>
            </div>
            <div class="hype-text">Hype</div>
            <div>-</div>
            <div class="hype-value">{{ model.hype }}</div>
          </div>
        </div>
      </div>
    </div>
    <div class="model-portfolio">
      <div class="main-image-container" @click="openMainPhoto">
        <img
          :src="getPhotoUrl(galleryPhotos[0]?.image)"
          :alt="model.name"
          class="main-image"
        />
      </div>
      <div class="thumbnail-grid">
        <div
          v-for="(photo, index) in galleryPhotos.slice(1)"
          :key="index"
          class="thumbnail-item"
          @click="openGallery(index)"
        >
          <img
            :src="getPhotoUrl(photo.image)"
            :alt="`${model.name} - фото ${index + 1}`"
            class="thumbnail-image"
          />
        </div>
      </div>
    </div>

    <div class="booking-section">
      <template v-if="data.user_id !== route.params.id">
        <OutlinedButton
          class="booking-button"
          @click="openChat"
          value="Booking"
        />
      </template>
    </div>

    <!-- Модальное окно для фотогалереи -->
    <div v-if="isGalleryOpen" class="gallery-modal" @click="closeGallery">
      <div class="gallery-content" @click.stop>
        <button class="gallery-close" @click="closeGallery">×</button>
        <button
          v-if="galleryPhotos.length > 1"
          class="gallery-nav gallery-prev"
          @click="prevPhoto"
        >
          ❮
        </button>
        <div class="gallery-image-container">
          <img
            :src="getPhotoUrl(getGalleryPhoto.image)"
            :alt="model.name"
            class="gallery-image"
          />
        </div>
        <button
          v-if="galleryPhotos.length > 1"
          class="gallery-nav gallery-next"
          @click="nextPhoto"
        >
          ❯
        </button>
        <div class="gallery-counter">
          {{ currentPhotoIndex + 1 }} / {{ galleryPhotos.length }}
        </div>
      </div>
    </div>
  </div>
  <div v-else class="not-found">
    <h2>Модель не найдена</h2>
    <NuxtLink to="/catalog" class="back-to-catalog"
      >Вернуться к каталогу</NuxtLink
    >
  </div>
</template>

<script setup>
import { useRoute } from 'vue-router'
import { ref, computed, onMounted } from 'vue'
import { useRuntimeConfig } from '#imports'
import { useRouter } from 'vue-router'

const { status, data } = useAuth()

definePageMeta({
  layout: 'custom',
})

const route = useRoute()
const router = useRouter()
const { $api } = useNuxtApp()

const model = ref(null)
const profileUrl = ref(null)
const isCopied = ref(false)
const isGalleryOpen = ref(false)
const currentPhotoIndex = ref(0)

const STATIC_BASE =
  useRuntimeConfig().public.staticBaseUrl || 'http://localhost:8082'

function getPhotoUrl(photo) {
  if (!photo) return ''
  if (typeof photo === 'string') {
    if (/^https?:\/\//.test(photo)) return photo
    return STATIC_BASE + photo
  }
  if (photo.image) {
    if (/^https?:\/\//.test(photo.image)) return photo.image
    return STATIC_BASE + photo.image
  }
  return ''
}

const fetchModelData = async () => {
  try {
    const queryParams = new URLSearchParams()
    queryParams.append('user_id', route.params.id)
    const response = await $api.get('/profile', queryParams)
    const calculateAge = (dateSeconds) => {
      if (!dateSeconds) return null
      const birthDate = new Date(dateSeconds * 1000)
      const today = new Date()
      let age = today.getFullYear() - birthDate.getFullYear()
      const m = today.getMonth() - birthDate.getMonth()
      if (m < 0 || (m === 0 && today.getDate() < birthDate.getDate())) {
        age--
      }
      return age
    }

    model.value = {
      ...response,
      verified: true,
      profile: `www.elefe.by/id${route.params.id}`,
      hype: '989',
      mainImage: response.avatar_url,
      gallery: (response.photos?.slice(1) || []).map((photo) => ({
        image: photo,
      })),
      image: response.avatar_url,
      parameters: `${response.chest_size || '-'}x${
        response.waist_size || '-'
      }x${response.hip_size || '-'}`,
      age: calculateAge(response.date_of_birth?.seconds),
      sex:
        response.sex?.charAt(0).toUpperCase() + response.sex?.slice(1) || '-',
    }

    profileUrl.value = window.location.host + window.location.pathname
  } catch (error) {
    console.error('Ошибка при загрузке данных модели:', error)
    model.value = null
  }
}

const copyButtonTooltip = computed(() =>
  isCopied.value ? 'Скопировано!' : 'Копировать ссылку'
)

const copyProfileLink = () => {
  if (!profileUrl?.value) return

  navigator.clipboard
    .writeText(profileUrl.value)
    .then(() => {
      isCopied.value = true
      setTimeout(() => {
        isCopied.value = false
      }, 2000)
    })
    .catch((err) => {
      console.error('Не удалось скопировать ссылку:', err)
    })
}

const galleryPhotos = computed(() => {
  if (!model.value) return []
  // Создаем массив, начиная с mainImage (если есть), и добавляем остальные фото
  const allPhotos = [
    { image: model.value.mainImage || model.value.avatar_url },
    ...(model.value.photos?.slice(1).map((photo) => ({ image: photo })) || []),
  ]
  return allPhotos
})

// Обновляем обработчик для основного фото и галереи
const getGalleryPhoto = computed(() => {
  if (!model.value || currentPhotoIndex.value === -1) {
    return { image: model.value?.mainImage || model.value?.avatar_url }
  }
  return galleryPhotos.value[currentPhotoIndex.value]
})

// Функция для открытия основного фото
const openMainPhoto = () => {
  currentPhotoIndex.value = 0
  isGalleryOpen.value = true
  document.body.style.overflow = 'hidden'
}

const openGallery = (index) => {
  currentPhotoIndex.value = index
  isGalleryOpen.value = true
  document.body.style.overflow = 'hidden'
}

const closeGallery = () => {
  isGalleryOpen.value = false
  document.body.style.overflow = ''
}

const nextPhoto = () => {
  currentPhotoIndex.value =
    (currentPhotoIndex.value + 1) % galleryPhotos.value.length
}

const prevPhoto = () => {
  currentPhotoIndex.value =
    (currentPhotoIndex.value - 1 + galleryPhotos.value.length) %
    galleryPhotos.value.length
}

const openBookingModal = async () => {
  // Если пользователь не авторизован или пытается открыть чат с собой — ничего не делаем
  if (!data.value?.user_id || data.value.user_id === route.params.id) return

  // Проверяем, есть ли уже чат с этим пользователем
  let chat = null
  try {
    const response = await $api.get(`/users/${data.value.user_id}/chats`)
    chat = response.find(
      (c) =>
        c.is_private &&
        c.participants &&
        c.participants.includes(Number(route.params.id))
    )
  } catch (e) {
    // fallback: если participants нет, ищем по имени/ID
    // или просто не нашли — создаём новый
  }

  // Если чата нет — создаём
  if (!chat) {
    const newChat = await $api.post('/chats', {
      name: `Chat with user ${route.params.id}`,
      is_private: true,
      participants: [Number(data.value.user_id), Number(route.params.id)],
    })
    chat = newChat
  }

  // Переходим на страницу чата
  router.push({ path: '/chats', query: { chatId: chat.id } })
}

const openChat = () => {
  // Переходим в чат с model_id
  router.push({
    path: '/chats',
    query: { model_id: route.params.id },
  })
}

// Загружаем данные при монтировании компонента
onMounted(() => {
  fetchModelData()

  // Обработка клавиш для галереи
  if (process.client) {
    window.addEventListener('keydown', (e) => {
      if (!isGalleryOpen.value) return

      if (e.key === 'Escape') {
        closeGallery()
      } else if (e.key === 'ArrowRight') {
        nextPhoto()
      } else if (e.key === 'ArrowLeft') {
        prevPhoto()
      }
    })
  }
})
</script>

<style scoped>
.model-profile {
  max-width: 505px;
  margin: 0 auto;
  padding-top: 60px;
}

.profile-header {
  position: relative;
  margin-bottom: 15px;
}

.model-name {
  font-size: 20px;
  font-weight: 600;
  margin-bottom: 15px;
  display: flex;
  align-items: center;
  gap: 10px;
}

.verified-badge {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 22px;
  height: 22px;
}

.model-details {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 15px;
  margin-bottom: 10px;
}

.details-column {
  display: flex;
  flex-direction: column;
  gap: 3px;
}

.detail-item {
  display: flex;
  align-items: center;
  gap: 8px;
}

.detail-label {
  min-width: 90px;
  font-weight: 500;
}

.profile-item {
  align-items: flex-start;
}

.profile-actions {
  display: flex;
  align-items: center;
  gap: 8px;
}

.profile-link {
  color: #7fff00;
  text-decoration: none;
  display: flex;
  align-items: center;
}

.link-icon {
  margin-left: 5px;
  font-size: 12px;
}

.copy-button {
  background: none;
  border: none;
  cursor: pointer;
  color: #7fff00;
  font-size: 16px;
  padding: 2px 5px;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: opacity 0.2s;
}

.copy-icon {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 22px;
  height: 22px;
}

.space-between {
  display: flex;
  justify-content: space-between;
}

.hype-section {
  display: flex;
  align-items: center;
  gap: 8px;
  color: #7fff00;
}

.hype-icon {
  height: 20px;
}

.model-portfolio {
  display: flex;
  gap: 10px;
  margin-bottom: 30px;
}

.main-image-container {
  flex: 1;
  aspect-ratio: 3/4;
  overflow: hidden;
  cursor: pointer;
}

.main-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.2s;
}

.main-image-container:hover .main-image {
  transform: scale(1.03);
}

.thumbnail-grid {
  display: flex;
  flex-direction: column;
  gap: 10px;
  width: 125px;
}

.thumbnail-item {
  width: 100%;
  aspect-ratio: 1;
  overflow: hidden;
  cursor: pointer;
}

.thumbnail-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.2s ease;
}

.thumbnail-item:hover .thumbnail-image {
  transform: scale(1.05);
}

.booking-section {
  display: flex;
  gap: 10px;
  margin-top: 20px;
}

.booking-button,
.message-button {
  flex: 1;
  background-color: transparent;
  width: 100%;
  font-size: 16px;
  cursor: pointer;
  transition: all 0.2s ease;
  text-transform: uppercase;
}

.booking-button:hover {
  background-color: rgba(255, 255, 255, 0.1);
}

/* Стили для модального окна галереи */
.gallery-modal {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.9);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.gallery-content {
  position: relative;
  width: 90%;
  max-width: 1000px;
  height: 80vh;
  display: flex;
  justify-content: center;
  align-items: center;
}

.gallery-image-container {
  height: 90%;
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}

.gallery-image {
  max-height: 100%;
  max-width: 100%;
  object-fit: contain;
}

.gallery-close {
  position: absolute;
  top: -40px;
  right: 0;
  font-size: 36px;
  background: none;
  border: none;
  color: #fff;
  cursor: pointer;
  z-index: 1001;
}

.gallery-nav {
  background: none;
  border: none;
  color: #fff;
  font-size: 36px;
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  cursor: pointer;
  padding: 15px;
  background-color: rgba(0, 0, 0, 0.2);
  transition: background-color 0.2s;
}

.gallery-nav:hover {
  background-color: rgba(0, 0, 0, 0.5);
}

.gallery-prev {
  left: -60px;
}

.gallery-next {
  right: -60px;
}

.gallery-counter {
  position: absolute;
  bottom: -30px;
  left: 50%;
  transform: translateX(-50%);
  color: #fff;
  font-size: 14px;
}

.not-found {
  width: 100%;
  text-align: center;
  padding: 60px 0;
}

.not-found h2 {
  font-size: 24px;
  margin-bottom: 20px;
}

.back-to-catalog {
  display: inline-block;
  padding: 10px 20px;
  border: 1px solid #f8f8f8;
  text-decoration: none;
  color: #f8f8f8;
  transition: all 0.2s ease;
}

.back-to-catalog:hover {
  background-color: rgba(255, 255, 255, 0.1);
}

.unauth-warning {
  width: 100%;
  text-align: center;
  padding: 60px 0;
  color: #ff5252;
}

@media (max-width: 768px) {
  .model-portfolio {
    flex-direction: column;
  }

  .thumbnail-grid {
    flex-direction: row;
    width: 100%;
    height: 100px;
  }

  .thumbnail-item {
    height: 100%;
    width: auto;
    aspect-ratio: initial;
  }

  .gallery-prev {
    left: 10px;
  }

  .gallery-next {
    right: 10px;
  }
}

@media (max-width: 480px) {
  .profile-header {
    text-align: center;
  }

  .model-details {
    align-items: center;
  }

  .detail-item {
    flex-direction: column;
    align-items: center;
  }

  .profile-actions {
    flex-direction: row;
    justify-content: center;
  }

  .thumbnail-grid {
    height: 80px;
  }

  .gallery-nav {
    font-size: 24px;
    padding: 10px;
  }
}
</style>
