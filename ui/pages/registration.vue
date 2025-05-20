<template>
  <div class="container">
    <div class="registration-form">
      <h2>Registration:</h2>
      <form @submit.prevent="register">
        <InputField
          id="name"
          label="Your name"
          type="text"
          v-model="user.name"
          :error="errors.name"
          placeholder="Name"
          required
          @update:error="(value) => (errors.name = value)"
        />
        <InputField
          id="lastname"
          label="Your lastname"
          type="text"
          v-model="user.lastName"
          :error="errors.lastName"
          placeholder="Lastname"
          required
          @update:error="(value) => (errors.lastName = value)"
        />
        <InputField
          id="email"
          label="Email"
          type="email"
          v-model="user.email"
          :error="errors.email"
          placeholder="example@email.com"
          required
          @update:error="(value) => (errors.email = value)"
        />
        <div class="form-group">
          <label for="goals">Your goals</label>
          <select
            name="goals"
            id="goals"
            class="goals"
            v-model="user.goal"
            required
          >
            <option selected="" disabled value="">Make your choice</option>
            <option value="model">model registration</option>
            <option value="scout">model search</option>
          </select>
        </div>
        <div class="form-group">
          <label for="password">Password</label>
          <input
            type="password"
            id="password"
            v-model="password"
            minlength="8"
            required
            autocomplete="new-password"
          />
        </div>
        <div class="form-group">
          <label for="confirm-password">Confirm password</label>
          <input
            type="password"
            id="confirm-password"
            v-model="confirmPassword"
            minlength="8"
            required
            autocomplete="new-password"
          />
          <FieldErrorMsg :error="passwordError" />
        </div>
        <div class="form-group">
          <label class="checkbox-label">
            <input
              type="checkbox"
              id="agree"
              v-model="isAgreed"
              required
              class="custom-checkbox"
              checked
            />
            <span class="checkmark"></span>
            <p>
              Согласен с
              <a href="/docs/agreement.pdf" target="_blank"
                >условиями договора</a
              >
              и
              <a href="/docs/politics.pdf" target="_blank"
                >политикой конфиденциальности</a
              >.
            </p>
          </label>
        </div>
        <OutlinedButton
          type="submit"
          :disabled="!isFormValid"
          value="Registration"
        />
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import InputField from './components/InputField.vue'

definePageMeta({
  layout: 'custom',
})

const { signIn } = useAuth()
const { $api } = useNuxtApp()

const user = ref({
  name: '',
  lastName: '',
  goal: '',
  email: '',
})

const errors = ref({
  name: '',
  lastName: '',
  email: '',
})

const isAgreed = ref(true)
const password = ref('')
const confirmPassword = ref('')
const passwordError = ref('')

const formData = ref({
  email: user.value.email,
  goal: user.value.goal,
  last_name: user.value.lastName,
  name: user.value.name,
  password: password.value,
})

const validatePassword = () => {
  if (password.value !== confirmPassword.value) {
    passwordError.value = 'Пароли не совпадают'
  } else {
    passwordError.value = ''
  }
}

const isFormValid = computed(() => {
  return (
    user.value.name &&
    !errors.value.name &&
    user.value.email &&
    !errors.value.email &&
    password.value &&
    confirmPassword.value &&
    !passwordError.value &&
    isAgreed.value &&
    (user.value.goal === 'model' || user.value.goal === 'scout')
  )
})

const register = async () => {
  try {
    await $api.post('auth/register', formData.value)
    await signIn({
      app_id: 1,
      email: formData.value.email,
      password: formData.value.password,
    })
  } catch (error) {
    console.error(error)
  }
}

watch(
  [user, password],
  () => {
    formData.value = {
      email: user.value.email,
      goal: user.value.goal,
      last_name: user.value.lastName,
      name: user.value.name,
      password: password.value,
    }
  },
  { deep: true }
)

watch([confirmPassword], validatePassword)
</script>

<style scoped>
.container {
  padding-top: 60px;
  height: calc(100vh - 102px);
}

h2 {
  font-size: 20px;
  margin-bottom: 20px;
}

form {
  display: flex;
  flex-direction: column;
  gap: 18px;
  width: 100%;
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

.country-code {
  font-size: 14px;
  font-weight: 600;
  border: 1px solid #f8f8f8;
  background-color: #070707;
  color: #f8f8f8;
  width: 80px;
  position: absolute;
  cursor: pointer;
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

.error {
  color: red;
  font-size: 12px;
  margin-top: 5px;
}

a {
  font-weight: 600;
  color: #939393;
}

label.checkbox-label {
  margin: 10px 0;
  max-width: 100vw;
}

.checkbox-label {
  display: flex;
  align-items: center;
  cursor: pointer;
  font-size: 15px;
  gap: 10px;
}

.custom-checkbox {
  position: absolute;
  opacity: 0;
  cursor: pointer;
  height: 0;
  width: 0;
}

.checkmark {
  position: relative;
  display: inline-block;
  width: 24px;
  height: 24px;
  background-color: #070707;
  border: 1px solid #f8f8f8;
}

.checkmark::before,
.checkmark::after {
  content: '';
  position: absolute;
  display: none;
  background-color: #fff;
}

.checkmark::before {
  left: 50%;
  top: 50%;
  width: 14px;
  height: 2px;
  transform: translate(-50%, -50%) rotate(45deg);
}

.checkmark::after {
  left: 50%;
  top: 50%;
  width: 14px;
  height: 2px;
  transform: translate(-50%, -50%) rotate(-45deg);
}

.custom-checkbox:checked ~ .checkmark::before,
.custom-checkbox:checked ~ .checkmark::after {
  display: block;
}

@media (max-width: 918px) {
  input,
  select {
    width: 230px;
  }

  .container {
    height: calc(100vh - 125px);
  }
}
</style>
