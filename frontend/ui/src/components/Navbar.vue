<template>
  <nav class="dark:bg-foreground dark:text-primary shadow-lg">
    <div class="container mx-auto px-4 md:flex items-center justify-between gap-6">

      <!--TODO: This logo -->
      <div class="flex items-center justify-center md:w-auto w-full">
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

      <!-- Desktop view links -->
      <div class="hidden md:flex items-center space-x-4">
        <NavLinks />
      </div>
      <div class="hidden md:flex">
        <ThemeToggle />
      </div>
    </div>

    <!-- mobile nav links -->
    <div v-if="isMobileMenuOpen" class="md:hidden flex flex-col px-4 pb-4 space-y-2
      dark:bg-foreground">
        <NavLinks />
        <ThemeToggle mobile="true" />
    </div>
  </nav>
</template>

<script setup>
import { ref } from 'vue';
import ThemeToggle from './ThemeToggle.vue'
import NavLinks from './NavLinks.vue'

const isMobileMenuOpen = ref(false)

const toggleMobileMenu = () => {
  isMobileMenuOpen.value = !isMobileMenuOpen.value
}

</script>