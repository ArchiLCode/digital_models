<template>
  <div class="login-form">
    <h2>Login</h2>
    <form @submit.prevent="handleLogin">
      <InputField
        id="email-login"
        type="email"
        v-model="formData.email"
        placeholder="Email"
        required
        :error="emailError"
      />
      <button type="submit" class="submit-button">
        <span class="arrow">></span>
      </button>
      <div class="form-group">
        <input
          id="password-login"
          type="password"
          v-model="formData.password"
          placeholder="Password"
          required
          autocomplete="current-password"
        />
      </div>

      <NuxtLink to="/forgot-password" class="forgot-password">
        Forgot password?
      </NuxtLink>
    </form>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const { signIn, token, refreshToken } = useAuth()

const formData = ref({
  email: '',
  password: '',
})

const emailError = ref('')

const handleLogin = async () => {
  try {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
    if (!emailRegex.test(formData.value.email)) {
      emailError.value = 'Введите корректный email'
      return
    }

    await signIn(
      {
        app_id: 1,
        email: formData.value.email,
        password: formData.value.password,
      },
      {
        redirect: false,
      }
    )

    const tokenCookie = useCookie('auth.token')
    const refreshTokenCookie = useCookie('auth.refresh_token')
    tokenCookie.value = token.value.slice(7)
    refreshTokenCookie.value = refreshToken.value
  } catch (error) {
    console.error('Ошибка авторизации:', error)
    emailError.value = 'Неверный email или пароль'
  }
}
</script>

<style scoped>
.login-form {
  position: absolute;
  left: 0;
  top: 0;
  background-color: transparent;
  margin-top: 60px;
  width: 200px;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

h2 {
  color: #f8f8f8;
  font-size: 20px;
  margin-bottom: 20px;
}

.form-group {
  margin-bottom: 18px;
}

label {
  display: block;
  font-size: 15px;
  font-weight: 600;
  color: #f8f8f8;
  margin-bottom: 5px;
}

input {
  font-size: 15px;
  font-weight: 600;
  border: 1px solid #f8f8f8;
  background-color: #070707;
  padding: 8px;
  color: #f8f8f8;
  width: 100%;
}

.submit-button {
  position: absolute;
  right: -40px;
  top: calc(50% - 10px);
  background-color: transparent;
  border: none;
  color: #f8f8f8;
  width: 40px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
}

.submit-button .arrow {
  font-size: 20px;
}

.forgot-password {
  display: block;
  margin-top: 10px;
  color: #f8f8f8;
  text-decoration: underline;
  font-weight: 600;
  font-size: 14px;
}

.forgot-password:hover {
  text-decoration: underline;
}
</style>
