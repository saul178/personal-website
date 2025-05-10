import Home from '@/views/Home.vue'
import Personal from '@/views/Personal.vue'
import Projects from '@/views/Projects.vue'
import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: '/', name: 'Home', component: Home },
    { path: '/projects', name: 'Projects', component: Projects },
    { path: '/personal', name: 'Personal', component: Personal },

  ],
})

export default router
