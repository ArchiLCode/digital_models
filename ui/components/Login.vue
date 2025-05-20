<template>
  <div class="auth-form">
    <h2>Авторизация</h2>
    <form @submit.prevent="login">
      <div class="form-group">
        <label for="login-email">Email</label>
        <input type="email" id="login-email" v-model="loginEmail" required />
      </div>
      <div class="form-group">
        <label for="login-password">Пароль</label>
        <input
          type="password"
          id="login-password"
          v-model="loginPassword"
          required
        />
      </div>
      <button type="submit">Войти</button>
    </form>
  </div>
</template>
<script setup>
import { ref } from 'vue'

const loginEmail = ref('')
const loginPassword = ref('')

const login = () => {
  // Add your login logic here
  console.log('Email:', loginEmail.value)
  console.log('Password:', loginPassword.value)
}
</script>

<style scoped>
.auth-form {
  max-width: 400px;
  margin: 0 auto;
  padding: 1rem;
  border: 1px solid #ccc;
  border-radius: 4px;
}

.form-group {
  margin-bottom: 1rem;
}

label {
  display: block;
  margin-bottom: 0.5rem;
}

input {
  width: 100%;
  padding: 0.5rem;
  border: 1px solid #ccc;
  border-radius: 4px;
}

button {
  width: 100%;
  padding: 0.75rem;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:hover {
  background-color: #0056b3;
}
</style>
