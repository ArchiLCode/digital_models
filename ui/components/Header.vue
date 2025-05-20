<template>
  <header :class="headerClass">
    <Logo />
    <Navigation />
    <div class="align"></div>
  </header>
</template>

<script setup>
import { computed } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()

const headerClass = computed(() => {
  return route.path === '/' ? 'home-header' : 'default-header'
})
</script>

<style scoped>
header {
  display: flex;
  align-items: center;
  padding-top: 30px;
  justify-content: space-between;
}

.default-header {
  flex-direction: row;
}

.home-header {
  flex-direction: column;
}

.align {
  width: 175.8px;
  display: none;
}

.default-header .align {
  display: block;
}

@media (max-width: 918px) {
  header {
    padding-top: 55px;
    flex-direction: row;
  }

  .default-header .align {
    display: none;
  }

  .home-header {
    flex-direction: row;
  }
}
</style>
