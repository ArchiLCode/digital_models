<template>
  <nav>
    <button
      class="menu-toggle"
      :class="{ active: isMenuOpen }"
      @click="toggleMenu"
    >
      <div class="hamburger"></div>
    </button>
    <ul
      :class="{ open: isMenuOpen, homeNav: isHome }"
      class="navigation desktop"
    >
      <li><NuxtLink to="/" exact-active-class="active">Main</NuxtLink></li>
      <li>
        <NuxtLink to="/registration" exact-active-class="active"
          >Registration</NuxtLink
        >
      </li>
      <li>
        <NuxtLink to="/catalog" exact-active-class="active">Catalog</NuxtLink>
      </li>
      <li>
        <NuxtLink to="/communications" exact-active-class="active"
          >Communications</NuxtLink
        >
      </li>
      <li>
        <NuxtLink to="/career" exact-active-class="active">Career</NuxtLink>
      </li>
      <li>
        <NuxtLink to="/contacts" exact-active-class="active">Contacts</NuxtLink>
      </li>
    </ul>
  </nav>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRoute } from 'vue-router'

const isMenuOpen = ref(false)
const route = useRoute()

const isHome = computed(() => {
  return route.path === '/'
})

const toggleMenu = () => {
  isMenuOpen.value = !isMenuOpen.value
}
</script>

<style scoped>
.homeNav {
  margin-top: 38px;
}

ul {
  display: flex;
  justify-content: center;
  list-style-type: none;
  gap: 30px;
  margin-top: 0;
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

.menu-toggle {
  display: none;
  background: none;
  border: none;
  cursor: pointer;
  position: relative;
  width: 28px;
  height: 20px;
  z-index: 1000;
}

.hamburger,
.hamburger::before,
.hamburger::after {
  content: '';
  position: absolute;
  width: 100%;
  height: 2px;
  display: block;
  background-color: #f8f8f8;
  transition: transform 0.3s ease, opacity 0.3s ease;
}

.hamburger::before {
  transform: translateY(-8px);
}

.hamburger::after {
  transform: translateY(8px);
}

.menu-toggle.active .hamburger {
  transform: rotate(45deg);
}

.menu-toggle.active .hamburger::before {
  transform: rotate(90deg);
  top: 0;
}

.menu-toggle.active .hamburger::after {
  transform: rotate(90deg);
  top: 0;
  opacity: 0;
}

@media (max-width: 918px) {
  .menu-toggle {
    display: block;
  }

  .navigation {
    position: fixed;
    top: 135.8px;
    right: 0;
    height: calc(100% - 135.8px);
    width: 100%;
    background-color: #070707;
    flex-direction: column;
    align-items: center;
    justify-content: flex-start;
    padding-top: 120px;
    transform: translateX(100%);
    transition: transform 0.3s ease-in-out;
    z-index: 999;
  }

  .navigation.open {
    transform: translateX(0);
  }

  ul {
    gap: 20px;
    margin-top: 0;
  }

  li a {
    font-size: 18px;
    text-decoration: none;
    font-weight: 600;
    color: #f8f8f8;
  }
}
</style>
