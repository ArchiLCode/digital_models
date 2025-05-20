<template>
  <div class="user-navigation">
    <h3>Welcome, {{ data.first_name }}</h3>
    <ul>
      <li>
        <NuxtLink to="/profile" exact-active-class="active"
          >Мой профиль</NuxtLink
        >
      </li>
      <li>
        <NuxtLink to="/change-password" exact-active-class="active"
          >Сменить пароль</NuxtLink
        >
      </li>
      <li>
        <NuxtLink to="/chats" exact-active-class="active">Сообщения</NuxtLink>
      </li>
      <li>
        <NuxtLink to="/payments" exact-active-class="active">Платежи</NuxtLink>
      </li>
      <li>
        <NuxtLink to="/education" exact-active-class="active"
          >Обучение</NuxtLink
        >
      </li>
    </ul>
    <button @click="logout">Log out</button>
  </div>
</template>

<script setup>
const { data, signOut } = useAuth()

const logout = async () => {
  await signOut({ redirect: false })

  const tokenCookie = useCookie('auth.token')
  const refreshTokenCookie = useCookie('auth.refresh_token')
  tokenCookie.value = ''
  refreshTokenCookie.value = ''
}
</script>

<style scoped>
.user-navigation {
  position: absolute;
  left: 0;
  top: 0;
  background-color: transparent;
  margin-top: 60px;
  width: 220px;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

h3 {
  margin-bottom: 10px;
}

ul {
  list-style: none;
}

li {
  margin-left: 10px;
  margin-bottom: 5px;
}

li a {
  font-size: 15px;
  text-decoration: none;
  font-weight: 600;
  color: #f8f8f8;
}

.active {
  color: #939393;
  text-decoration: underline;
}

button {
  padding: 10px 0;
  background-color: transparent;
  border: none;
  outline: none;
  color: #f8f8f8;
  font-weight: 600;
  font-size: 15px;
  text-align: start;
  text-decoration: underline;
}
</style>
