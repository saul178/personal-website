<script setup>
import ProjectCard from '@/components/ProjectCard.vue';
</script>

<template>
  <div class="flex flex-col md:flex-row px-8 py-12 gap-8">
    <!-- left area -->
    <div class="flex flex-col space-y-4 gap-2">
      <h2 class="text-4xl font-bold dark:text-primary">Projects</h2>
      <div class="flex justify-start gap-2.5">
        <div class="hover:scale-110 transition-all">
          <a href="https://github.com/saul178" class="dark:text-primary text-2xl font-bold">
            <font-awesome-icon :icon="['fab', 'github']" class="text-4xl hover:text-primary/60 dark:text-primary" />
          </a>
        </div>
        <a href="https://github.com/saul178" class="dark:text-primary text-2xl font-bold
          hover:text-primary/60">
          View them on Github →
        </a>
      </div>
    </div>

    <!-- right area -->
    <div class="md:flex-1 flex-col space-y-6">
      <!-- TODO: project card will be a component to be reused -->
      <!-- call the meat of the project card here -->
      <ProjectCard />
    </div>
  </div>

  <div class="flex flex-col px-8 py-12 gap-8">
    <h1 class="text-3xl font-semibold dark:text-primary">My Github Activity</h1>
    <div class="flex-1 space-y-6">
      <div
        class="group hover:-translate-y-1 bg-foreground hover:bg-foreground/60 backdrop-blur-sm p-6 rounded-2xl border border-accent/20 transition-all duration-300 shadow-lg hover:shadow-accent/20">
        <a href="https://www.github.com/saul178">
          <img class="w-full" src="https://ghchart.rshah.org/072460/saul178.svg" alt="Github Contributations Chart">
        </a>
      </div>
    </div>
  </div>
</template>
