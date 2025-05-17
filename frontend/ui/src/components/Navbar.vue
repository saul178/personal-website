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

      <!-- links -->
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
      <!-- dark and light mode icon -->
      <button @click="toggleDark()" class="hidden md:flex py-2 px-3 dark:hover:text-accent transition-colors">
        <span v-if="isDark">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
            stroke="currentColor" class="size-6">
            <path stroke-linecap="round" stroke-linejoin="round"
              d="M21.752 15.002A9.72 9.72 0 0 1 18 15.75c-5.385 0-9.75-4.365-9.75-9.75 0-1.33.266-2.597.748-3.752A9.753 9.753 0 0 0 3 11.25C3 16.635 7.365 21 12.75 21a9.753 9.753 0 0 0 9.002-5.998Z" />
          </svg>
        </span>
        <span v-else>
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
            stroke="currentColor" class="size-6">
            <path stroke-linecap="round" stroke-linejoin="round"
              d="M12 3v2.25m6.364.386-1.591 1.591M21 12h-2.25m-.386 6.364-1.591-1.591M12 18.75V21m-4.773-4.227-1.591 1.591M5.25 12H3m4.227-4.773L5.636 5.636M15.75 12a3.75 3.75 0 1 1-7.5 0 3.75 3.75 0 0 1 7.5 0Z" />
          </svg>
        </span>
      </button>
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

      <button @click="toggleDark()" class="py-2 px-3 dark:hover:text-accent transition-colors">
        <span v-if="isDark">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
            stroke="currentColor" class="size-6">
            <path stroke-linecap="round" stroke-linejoin="round"
              d="M21.752 15.002A9.72 9.72 0 0 1 18 15.75c-5.385 0-9.75-4.365-9.75-9.75 0-1.33.266-2.597.748-3.752A9.753 9.753 0 0 0 3 11.25C3 16.635 7.365 21 12.75 21a9.753 9.753 0 0 0 9.002-5.998Z" />
          </svg>
        </span>
        <span v-else>
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
            stroke="currentColor" class="size-6">
            <path stroke-linecap="round" stroke-linejoin="round"
              d="M12 3v2.25m6.364.386-1.591 1.591M21 12h-2.25m-.386 6.364-1.591-1.591M12 18.75V21m-4.773-4.227-1.591 1.591M5.25 12H3m4.227-4.773L5.636 5.636M15.75 12a3.75 3.75 0 1 1-7.5 0 3.75 3.75 0 0 1 7.5 0Z" />
          </svg>
        </span>
      </button>
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
