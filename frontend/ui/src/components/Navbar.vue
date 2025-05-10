<template>
  <nav class="bg-gray-900 text-white shadow-lg">
    <div class="container mx-auto px-4 md:flex items-center gap-6">
      <!-- Logo -->
      <div class="flex items-center justify-between md:w-auto w-full">
        <RouterLink to="/" class="py-5 px-2 text-white flex-1 font-bold">
          ICON-HERE
        </RouterLink>
        <div class="md:hidden flex items-center">
          <button @click="toggleMobileMenu" class="md:hidden mobile-menu-button">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
              stroke="currentColor" class="size-6">
              <path stroke-linecap="round" stroke-linejoin="round" d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5" />
            </svg>
          </button>
        </div>
        <!-- links -->
        <div :class="['flex flex-col md:flex-row md:space-x-1 pb-3 md:pb-0 w-full md:w-auto', isMobileMenuOpen ?
          'block' : 'hidden', 'md:flex']" class="navigation-menu px-4">
          <RouterLink to="/" class="py-2 px-3 block hover:text-blue-400 transition-colors" active-class="text-blue-400"
            exact>
            Home
          </RouterLink>

          <RouterLink to="/projects" class="py-2 px-3 block hover:text-blue-400 transition-colors"
            active-class="text-blue-400">
            Projects
          </RouterLink>

          <RouterLink to="/contact" class="py-2 px-3 block hover:text-blue-400 transition-colors"
            active-class="text-blue-400">
            Contact
          </RouterLink>

          <RouterLink to="/personal" class="py-2 px-3 block hover:text-blue-400 transition-colors"
            active-class="text-blue-400">
            Personal
          </RouterLink>
        </div>
      </div>
    </div>
  </nav>
</template>

<script setup>
import { ref } from 'vue';
import { RouterLink } from 'vue-router'
const isMobileMenuOpen = ref(false)

const toggleMobileMenu = () => {
  isMobileMenuOpen.value = !isMobileMenuOpen.value
}
</script>

<style></style>
