<template>
  <div :class="logoClass">
    <NuxtLink to="/">
      <h1>ÉLEFÉ</h1>
      <h3>You are not a model, you are muse.</h3>
    </NuxtLink>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()

const logoClass = computed(() => {
  return route.path === '/' ? 'home-logo' : 'default-logo'
})
</script>

<style scoped>
h1 {
  font-size: 40px;
  text-align: center;
  letter-spacing: 4px;
}
h3 {
  font-size: 12px;
  text-align: center;
}

a {
  text-decoration: none;
  color: #f8f8f8;
}

.home-logo h1,
h3 {
  text-align: center;
}

.default-logo h1,
h3 {
  text-align: start;
}
</style>
