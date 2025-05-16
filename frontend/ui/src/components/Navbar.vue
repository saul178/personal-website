<template>
  <nav class="bg-polynesian-blue-900 text-royal-blue-100 shadow-lg">
    <div class="container mx-auto px-4 md:flex items-center justify-between gap-6">
      <!-- Logo -->
      <div class="flex items-center justify-between md:w-auto w-full">
        <RouterLink to="/" class="py-5 px-2 flex-1 font-bold">
          ICON-HERE
        </RouterLink>
        <!-- burger icon into X icon -->
        <button @click="toggleMobileMenu" class="md:hidden cursor-pointer mobile-menu-button">
          <span v-if="!isMobileMenuOpen">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
              stroke="currentColor" class="size-6">
              <path stroke-linecap="round" stroke-linejoin="round" d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5" />
            </svg>
          </span>

          <span v-else>
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
              stroke="currentColor" class="size-6">
              <path stroke-linecap="round" stroke-linejoin="round" d="M6 18 18 6M6 6l12 12" />
            </svg>
          </span>
        </button>
      </div>

      <!-- links -->
      <div :class="['flex flex-col md:flex-row md:space-x-1 pb-3 md:pb-0 w-full md:w-auto', isMobileMenuOpen ?
        'block' : 'hidden', 'md:flex']" class="navigation-menu px-4">
        <RouterLink to="/" class="py-2 px-3 block hover:text-powder-blue-400 transition-colors"
          active-class="text-powder-blue-400" exact>
          Home
        </RouterLink>

        <RouterLink to="/projects" class="py-2 px-3 block hover:text-powder-blue-400 transition-colors"
          active-class="text-powder-blue-400" exact>
          Projects
        </RouterLink>

        <RouterLink to="/contact" class="py-2 px-3 block hover:text-powder-blue-400 transition-colors"
          active-class="text-powder-blue-400" exact>
          Contact
        </RouterLink>

        <RouterLink to="/personal" class="py-2 px-3 block hover:text-powder-blue-400 transition-colors"
          active-class="text-powder-blue-400" exact>
          Personal
        </RouterLink>
      </div>
      <div>
        <span class="dark:primary-text">
          {{ isDark ? "Dark" : "Light" }} Mode
        </span>
        <button @click="toggleDark()" class="py-2 px-4 bg-background primary-text">

        </button>
      </div>
    </div>
  </nav>
</template>

<script setup>
import { ref } from 'vue';
import { RouterLink } from 'vue-router'
import { useDark, useToggle } from '@vueuse/core'

const isMobileMenuOpen = ref(false)

const toggleMobileMenu = () => {
  isMobileMenuOpen.value = !isMobileMenuOpen.value
}

const isDark = useDark()
const toggleDark = useToggle(isDark)
</script>

<style scoped></style>