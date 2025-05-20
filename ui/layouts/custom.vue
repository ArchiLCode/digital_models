<template>
  <div class="layout">
    <!-- Форма авторизации -->
    <LoginForm v-if="!isAuthenticated" />
    <UserNavigation v-else />

    <!-- Содержимое страницы -->
    <slot />
  </div>
</template>

<script setup>
import LoginForm from '~/components/login/LoginForm.vue'

const { status } = useAuth()

const isAuthenticated = computed(() => {
  return status.value === 'authenticated'
})
</script>

<style scoped>
.layout {
  display: flex;
  position: relative;
}
</style>
