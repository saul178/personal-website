<script setup>
import { RouterLink } from 'vue-router'

const items = [
  { to: '/', label: 'Home' },
  { to: '/projects', label: 'Projects' },
  { to: '/personal', label: 'Personal' },
]
</script>

<template>
  <RouterLink v-for="(item) in items" :key="item.to" :to="item.to"
    class="py-2 px-3 block border-b-1 border-transparent dark:hover:text-accent transition-colors"
    active-class="dark:border-accent dark:text-accent" exact>
    {{ item.label }}
  </RouterLink>
</template>
