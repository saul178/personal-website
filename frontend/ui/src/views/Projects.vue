<!-- what to show from the github api
name or full name not sure if this is for the repo or user
description box of the repo
created at, updated at or pushed at?
commits??
url hyperlink?
see if i can get what languages are used???

at the top or bottom have total commits?
-->

<template>
  <div class="flex flex-col md:flex-row px-8 py-12 gap-8">
    <!-- left area -->
    <div class="flex flex-col space-y-4 gap-2">
      <h2 class="text-4xl font-bold dark:text-primary">Projects</h2>
      <div class="flex justify-start gap-2.5">
        <a href="https://github.com/saul178" class="dark:text-primary text-2xl font-bold">
          <font-awesome-icon :icon="['fab', 'github']" class="text-4xl dark:text-primary" />
        </a>
        <a href="https://github.com/saul178" class="dark:text-primary text-2xl font-bold">
          View them on Github →
        </a>
      </div>
    </div>

    <!-- right area -->
    <div class="md:flex-1 flex-col space-y-6">
      <!-- project card will be a component to be reused -->
      <div v-for="repo in getRepos" :key="repo.title"
        class="flex flex-col md:flex-row group hover:-translate-y-1 bg-foreground hover:bg-foreground/60 backdrop-blur-sm p-6 rounded-2xl border border-accent/20 transition-all duration-300 shadow-lg hover:shadow-accent/20">

        <!-- project image here -->
        <img class="md:w-80 h-48 mb-4 md:mr-4 rounded-md shadow-lg object-cover transition
          duration-500 ease-in-out opacity-25 group-hover:opacity-100"
          src="../assets/images/snapshot_2025-06-10_13-49-07.png" alt="Project 1">

        <div class="flex flex-col justify-between">
          <!--TODO: fix the alignment here and make it presentable -->
          <div>
            <h3 class="text-xl font-semibold dark:text-primary first-letter:uppercase">{{ repo.title }}</h3>
            <div class="flex flex-row gap-4 mt-1 mb-2">
              <h3 class="text-sm dark:text-primary">Created at - {{ repo.created_at }}</h3>
              <h3 class="text-sm dark:text-primary">Updated at - {{ repo.updated_at }}</h3>
            </div>
            <div>
              <h2 class="text-sm font-semibold dark:text-primary">Description</h2>
              <p class="mt-2 dark:text-secondary"> <!-- this need to change for lightmode -->
                {{ repo.desc || 'No description available' }}
              </p>
            </div>

            <!-- tag rows (this will probably be a vue component) -->
            <div class="flex flex-row flex-wrap gap-4 mt-4">
              <span v-for="lang in Object.keys(repo.languages)" :key="lang" class="dark:text-primary px-2 py-1 rounded-full text-sm border-2
                border-accent/50">{{ lang }}</span>
            </div>

            <div class="md:flex flex-row hidden mt-6 border-t-2 dark:border-background">
              <h3 class="dark:text-secondary font-semibold mt-2">Commit History</h3>
            </div>
            <div v-for="(author, i) in store.commitsByRepo[repo.title]?.author || []" :key="i" class="mt-2">
              <span class="dark:text-secondary block">
                {{ author }}: {{ store.commitsByRepo[repo.title].time[i] }} — "{{
                  store.commitsByRepo[repo.title].commit_msg[i] }}"
              </span>
            </div>
          </div>
        </div>
      </div>
      <!-- project card component end -->

      <!-- other projects -->
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

<script setup>
import { onMounted, computed } from 'vue';
import { useGithubStore } from '@/stores/GithubStore';

const store = useGithubStore()
const getRepos = computed(() => {
  return store.getGithubRepos
})

onMounted(() => {
  store.fetchGithubRepos().then(() => {
    store.repos.forEach(repo => {
      console.log("Fetching commits for:", repo.title)
      store.fetchGithubCommits(repo.title)
    })
  })
})

</script>
