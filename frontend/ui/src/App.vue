<script setup>
import { RouterView } from 'vue-router'
import Navbar from '@/components/Navbar.vue'
import Footer from './components/Footer.vue'
</script>

<template>
  <div class="bg-linear-to-b dark:from-background to-foreground min-h-screen flex flex-col font-rubik">
    <Navbar />
    <main class="flex-1 p-6">
      <RouterView />
    </main>
    <Footer />
  </div>
</template>

<style>
.slide-fade-enter-active,
.slide-fade-leave-active {
  transition: all 1.0s ease;
  transition-delay: var(--delay, 0ms);
}

.slide-fade-enter-from {
  opacity: 0;
  transform: translateX(50px);
}

.slide-fade-leave-to {
  opacity: 0;
  transform: translateX(-50px);
}
</style>
