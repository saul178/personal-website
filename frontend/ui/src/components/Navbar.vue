<template>
  <nav class="dark:bg-foreground dark:text-primary shadow-lg">
    <div class="container mx-auto px-4 md:flex items-center justify-between gap-6">

      <!-- Logo -->
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
        <RouterLink to="/" class="py-2 px-3 block dark:hover:text-accent transition-colors"
          active-class="dark:text-accent" exact>
          Home
        </RouterLink>

        <RouterLink to="/projects" class="py-2 px-3 block dark:hover:text-accent transition-colors"
          active-class="dark:text-accent" exact>
          Projects
        </RouterLink>

        <RouterLink to="/contact" class="py-2 px-3 block dark:hover:text-accent transition-colors"
          active-class="dark:text-accent" exact>
          Contact
        </RouterLink>

        <RouterLink to="/personal" class="py-2 px-3 block dark:hover:text-accent transition-colors"
          active-class="dark:text-accent" exact>
          Personal
        </RouterLink>
      </div>
      <div class="hidden md:flex">
        <ThemeToggle />
      </div>
    </div>

    <!-- mobile stuff need to make it dry -->
    <div v-if="isMobileMenuOpen" class="md:hidden flex flex-col px-4 pb-4 space-y-2
      dark:bg-foreground">
      <RouterLink to="/" class="py-2 px-3 block dark:hover:text-accent transition-colors"
        active-class="dark:text-accent" exact>
        Home
      </RouterLink>

      <RouterLink to="/projects" class="py-2 px-3 block dark:hover:text-accent transition-colors"
        active-class="dark:text-accent" exact>
        Projects
      </RouterLink>

      <RouterLink to="/contact" class="py-2 px-3 block dark:hover:text-accent transition-colors"
        active-class="dark:text-accent" exact>
        Contact
      </RouterLink>

      <RouterLink to="/personal" class="py-2 px-3 block dark:hover:text-accent transition-colors"
        active-class="dark:text-accent" exact>
        Personal
      </RouterLink>
        <ThemeToggle mobile="true" />
    </div>
  </nav>
</template>

<script setup>
import { ref } from 'vue';
import { RouterLink } from 'vue-router'
import ThemeToggle from './ThemeToggle.vue';

const isMobileMenuOpen = ref(false)

const toggleMobileMenu = () => {
  isMobileMenuOpen.value = !isMobileMenuOpen.value
}

</script>

<style scoped></style>
