<script setup>
import { ref } from 'vue';
import ThemeToggle from './ThemeToggle.vue'
import NavLinks from './NavLinks.vue'

const isMobileMenuOpen = ref(false)

const toggleMobileMenu = () => {
  isMobileMenuOpen.value = !isMobileMenuOpen.value
}
</script>

<template>
  <nav class="shadow-md rounded-b-lg dark:bg-foreground dark:text-primary">
    <div class="flex px-4 justify-between">

      <!-- home logo -->
      <div class="flex md:w-auto w-full px-4">
        <RouterLink to="/" class="py-5 px-2 flex-1 md:flex-col dark:hover:text-accent/60">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
            stroke="currentColor" class="size-6">
            <path stroke-linecap="round" stroke-linejoin="round"
              d="m2.25 12 8.954-8.955c.44-.439 1.152-.439 1.591 0L21.75 12M4.5 9.75v10.125c0 .621.504 1.125 1.125 1.125H9.75v-4.875c0-.621.504-1.125 1.125-1.125h2.25c.621 0 1.125.504 1.125 1.125V21h4.125c.621 0 1.125-.504 1.125-1.125V9.75M8.25 21h8.25" />
          </svg>
        </RouterLink>

        <!-- burger icon into X icon -->
        <button @click="toggleMobileMenu" class="md:hidden mobile-menu-button">
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

      <!-- desktop links -->
      <div class="hidden md:flex items-center justify-center space-x-4">
        <NavLinks />
      </div>
      <div class="hidden md:flex mr-2">
        <ThemeToggle />
      </div>
    </div>

    <!-- mobile nav links -->
    <div v-if="isMobileMenuOpen" class="md:hidden flex flex-col px-4 pb-4 space-y-2
      bg-foreground">
      <NavLinks />
      <ThemeToggle mobile="true" />
    </div>
  </nav>
</template>
