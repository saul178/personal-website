<script setup>
import { onMounted, computed } from 'vue';
import { useGithubStore } from '@/stores/GithubStore';

const store = useGithubStore()
const getRepos = computed(() => store.getGithubRepos)

// NOTE: this can be used to debug
// watch(() => store.commitsByRepo, (val) => {
//   console.log("commitsByRepo changed", val)
// }, { deep: true })
//
const getCommitsForRepo = (repoTitle) => {
  const data = store.commitsByRepo[repoTitle]

  if (!data) {
    console.log("failed to fetch commit data for repo: ", repoTitle)
    return []
  }

  return data.author.map((author, i) => ({
    author,
    msg: data.commit_msg[i],
    time: data.time[i],
  }))
}

onMounted(() => {
  store.fetchGithubRepos().then(() => {
    store.repos.forEach(repo => {
      store.fetchGithubCommits(repo.title)
    })
  })
})


</script>

<template>
  <div v-for="repo in getRepos" :key="repo.title"
    class="flex flex-col md:flex-row group hover:-translate-y-1 bg-foreground hover:bg-foreground/60 backdrop-blur-sm p-6 rounded-2xl border border-accent/20 transition-all duration-300 shadow-lg hover:shadow-accent/20">
    <!-- project image here -->
    <!--  FIX: each repo shares the same image need to fix it so that a repo has their own image or a default image if not set -->
    <img class="md:w-80 h-48 mb-4 md:mr-4 rounded-md shadow-lg object-cover transition
            duration-500 ease-in-out opacity-25 group-hover:opacity-100"
      src="../assets/images/snapshot_2025-06-10_13-49-07.png" alt="Project 1">

    <div class="flex flex-col justify-between">
      <div>
        <h3 class="text-xl font-semibold dark:text-primary capitalize
                hover:text-accent/60">
          <a :href="repo.url">
            {{ repo.title.replace(/-/g, ' ') }}
          </a>
        </h3>
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
          <span v-for="lang in Object.keys(repo.languages)" :key="lang"
            class="dark:text-primary px-2 py-1 rounded-full text-sm border-2 border-accent/50">{{ lang }}</span>
        </div>

        <!-- this should also be it's own component -->
        <div class="md:flex flex-row hidden mt-6 border-t-2 dark:border-background">
          <h3 class="dark:text-secondary font-semibold mt-2">Commit History</h3>
        </div>
        <div v-if="store.commitsByRepo[repo.title]" class="mt-2">
          <span v-for="commit in getCommitsForRepo(repo.title)" :key="commit.sha" class="dark:text-secondary block">
            {{ commit.author }}: "{{ commit.msg }}" - {{ commit.time }}
          </span>
        </div>
        <div v-else class="mt-2">
          <span class="dark:text-secondary block italic">Loading commit history...</span>
        </div>
      </div>
    </div>
  </div>
</template>
