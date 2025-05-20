<template>
  <div class="contacts-page">
    <h2>Контакты</h2>
    <div class="contact-info">
      <div>
        <h3>Телефоны:</h3>
        <p><a href="tel:+375333665053">+375333665053</a>, Менеджер</p>
      </div>
      <div>
        <h3>Адрес:</h3>
        <a href="https://yandex.ru/maps/-/CHuN6XIx" target="_blank"
          >213807, Республика Беларусь, Могилевская область, г. Бобруйск, б-р
          Приберезинский 28/24</a
        >
      </div>
      <p>
        <strong>Отдел управления персоналом: </strong>
        <a href="mailto:hr@elefe.by">hr@elefe.by</a>
      </p>
      <p>
        <strong>Отдел рекламы и маркетинга: </strong>
        <a href="mailto:marketing@elefe.by">marketing@elefe.by</a>
      </p>
      <p>
        <strong>Отдел по работе с партнерами: </strong>
        <a href="mailto:partners@elefe.by">partners@elefe.by</a>
      </p>
      <p>
        <strong>Отдел по работе с клиентами: </strong>
        <a href="mailto:clients@elefe.by">clients@elefe.by</a>
      </p>
      <p>
        <strong>Финансовый отдел: </strong>
        <a href="mailto:finance@elefe.by">finance@elefe.by</a>
      </p>
      <p>
        <strong>Свидетельство о государственной регистрации: </strong>№0791347
        от 20.10.2022, выдано Администрацией Первомайского района г. Бобруйска.
      </p>
    </div>
    <div class="feedback-form">
      <h2>Задать вопрос</h2>
      <form @submit.prevent="submitForm">
        <InputField
          id="name"
          label="Ваше имя:"
          type="text"
          v-model="form.name"
          required
        />
        <InputField
          id="phone"
          label="Номер телефона:"
          type="text"
          v-model="form.phone"
        />
        <InputField
          id="name"
          label="E-mail:"
          type="email"
          v-model="form.email"
          :error="errors.email"
          required
          @update:error="(value) => (errors.email = value)"
        />
        <div class="form-group">
          <label for="message">Текст сообщения:</label>
          <textarea
            id="message"
            rows="6"
            v-model="form.message"
            required
          ></textarea>
        </div>
        <OutlinedButton
          type="submit"
          value="Отправить"
          :disabled="!isFormValid"
        ></OutlinedButton>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

definePageMeta({
  layout: 'custom',
})

const form = ref({
  name: '',
  phone: '',
  email: '',
  message: '',
})

const errors = ref({
  email: null,
})

const isFormValid = computed(() => {
  return form.value.name && form.value.email && form.value.message
})

const submitForm = () => {
  // Здесь можно добавить логику для отправки формы, например, через API
  console.log('Форма отправлена:', form.value)
  alert('Ваше сообщение отправлено!')
  form.value = { name: '', phone: '', email: '', message: '' } // Очистка формы после отправки
}
</script>

<style lang="scss" scoped>
.contacts-page {
  max-width: 505px;
  padding-top: 60px;
  margin: 0 auto;
}

h2 {
  font-size: 20px;
  margin-bottom: 20px;
}

h3 {
  font-size: 16px;
  margin-top: 15px;
}

a {
  text-decoration: underline;
  color: #f8f8f8;
  transition: 0.1s;

  &:hover {
    color: #939393;
  }
}

p {
  margin-top: 5px;
}

.contact-info {
  margin-bottom: 40px;
}

form {
  display: flex;
  flex-direction: column;
  gap: 18px;
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

textarea {
  font-size: 15px;
  font-weight: 600;
  border: 1px solid #fff;
  background-color: #070707;
  padding: 8px;
  color: #fff;
  width: 390px;
  min-height: 90px;
  resize: vertical;
}
</style>
