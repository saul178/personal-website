<script setup>
import { RouterView } from 'vue-router'
import Navbar from '@/components/Navbar.vue'
</script>

<template>
  <div class="bg-linear-to-b dark:from-background to-foreground min-h-screen flex flex-col font-rubik">
    <Navbar />
    <main class="flex-1 p-6">
      <RouterView />
    </main>
  </div>
</template>
