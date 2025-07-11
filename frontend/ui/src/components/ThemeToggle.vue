<script setup>
import { useDark, useToggle } from '@vueuse/core'

const isDark = useDark()
const toggleDark = useToggle(isDark)
</script>

<template>
  <button @click="toggleDark()" class="relative ml-3 px-2 w-18 h-8 rounded-full md:mt-4 transition-colors duration-300 ease-in-out
    dark:bg-background cursor-pointer">
    <span class="text-xs font-medium transition-opacity duration-300 ease-in-out"
      :class="isDark ? 'dark:text-primary/20 opacity-100' : 'opacity-0'">dark</span>
    <span class="text-xs font-medium transition-opacity duration-300 ease-in-out"
      :class="isDark ? 'opacity-0' : 'dark:text-primary/20 opacity-100'">light</span>
    <div
      class="toggle-slider absolute top-1 left-1 w-6 h-6 rounded-full bg-foreground shadow-md flex items-center justify-center"
      :style="{ transform: isDark ? 'translateX(2.5rem)' : 'translateX(0.25px)' }">
      <span v-if="isDark">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor"
          class="w-4 h-4">
          <path stroke-linecap="round" stroke-linejoin="round"
            d="M21.752 15.002A9.72 9.72 0 0 1 18 15.75c-5.385 0-9.75-4.365-9.75-9.75 0-1.33.266-2.597.748-3.752A9.753 9.753 0 0 0 3 11.25C3 16.635 7.365 21 12.75 21a9.753 9.753 0 0 0 9.002-5.998Z" />
        </svg>
      </span>
      <span v-else>
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor"
          class="w-4 h-4">
          <path stroke-linecap="round" stroke-linejoin="round"
            d="M12 3v2.25m6.364.386-1.591 1.591M21 12h-2.25m-.386 6.364-1.591-1.591M12 18.75V21m-4.773-4.227-1.591 1.591M5.25 12H3m4.227-4.773L5.636 5.636M15.75 12a3.75 3.75 0 1 1-7.5 0 3.75 3.75 0 0 1 7.5 0Z" />
        </svg>
      </span>
    </div>
  </button>
</template>

<style scoped>
.toggle-slider {
  transition: transform 0.2s ease-in-out !important;
}
</style>
