<template>
  <div class="container">
    <Header />
    <NuxtLayout>
      <NuxtPage />
    </NuxtLayout>
    <NuxtRouteAnnouncer />
    <Footer />
  </div>
</template>

<script setup>
useHead({
  title: 'ÉLEFÉ',
  meta: [{ name: 'description', content: 'Описание моего сайта' }],
})

const { status, getSession } = useAuth()

onMounted(async () => {
  if (status.value === 'unauthenticated') {
    try {
      await getSession({ force: true })
    } catch (error) {
      console.error(error)
    }
  }
})
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  outline: none;
  -webkit-tap-highlight-color: transparent;
}

html {
  overflow-y: scroll;
  overflow-y: overlay;
}

body {
  font-family: 'Calibri', sans-serif;
  color: #f8f8f8;
  background-color: #070707;
  line-height: 1.4;
}

::-webkit-scrollbar {
  width: 8px;
}

::-webkit-scrollbar-track {
  background-color: transparent;
}

::-webkit-scrollbar-thumb {
  background-color: #666;
  border-radius: 6px;
}

::-webkit-scrollbar-thumb:hover {
  background-color: #444;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
}

@media (max-width: 918px) {
  .container {
    margin: 0 20px;
  }
}
</style>
