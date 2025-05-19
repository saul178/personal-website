<template>
    <RouterLink v-for="(item) in items" 
    :key="item.to" 
    :to="item.to" 
    class="py-2 px-3 block dark:hover:text-accent transition-colors" 
    active-class="dark:text-accent"
    exact>
        {{ item.label }}
    </RouterLink>
</template>

<script setup>
import { RouterLink } from 'vue-router'

const items = [
    {to: '/', label: 'Home'},
    {to: '/projects', label: 'Projects'},
    {to: '/contact', label: 'Contact'},
    {to: '/personal', label: 'Personal'},
]
</script>