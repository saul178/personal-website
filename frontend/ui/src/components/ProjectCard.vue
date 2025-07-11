<script setup>
import { useGithubStore } from '@/stores/GithubStore'

const store = useGithubStore()
const props = defineProps({
  repo: {
    type: Object,
  }
})

const getRepoImg = (repoTitle) => {
  return store.fetchRepoImg(repoTitle)
}

const getCommitsForRepo = (repoTitle) => {
  const data = store.commitsByRepo[repoTitle]
  if (!data) {
    return []
  }

  return data.author.map((author, i) => ({
    author,
    msg: data.commit_msg[i],
    time: data.time[i],
  }))
}
</script>

<template>
  <div
    class="flex flex-col md:flex-row group hover:-translate-y-1 bg-foreground hover:bg-foreground/60 backdrop-blur-sm p-6 rounded-2xl border border-accent/20 transition-all duration-300 shadow-lg hover:shadow-accent/20">
    <!-- project image here -->
    <img class="md:w-80 h-48 mb-4 md:mr-4 rounded-md shadow-lg object-cover md:transition
      md:duration-500 md:ease-in-out md:opacity-25 md:group-hover:opacity-100" :src="getRepoImg(repo.title)"
      alt="Project Image">

    <div class="flex flex-col justify-between">
      <div>
        <h3 class="text-xl font-semibold dark:text-primary text-shadow-sm capitalize
                hover:text-accent/60">
          <a :href="repo.url">
            {{ repo.title.replace(/-/g, ' ') }}
          </a>
        </h3>
        <div class="hidden md:flex md:flex-row md:gap-4 md:mt-1 md:mb-2">
          <h3 class="text-sm dark:text-primary/90">Created at - {{ repo.created_at.replace(/-/g, '/') }}</h3>
          <h3 class="text-sm dark:text-primary/90">Updated at - {{ repo.updated_at.replace(/-/g, '/') }}</h3>
        </div>
        <div>
          <h2 class="text-sm font-semibold dark:text-primary text-shadow-sm mt-2">Description</h2>
          <p class="mt-1 dark:text-secondary/90">
            {{ repo.desc || 'No description available' }}
          </p>
        </div>

        <!-- tag rows -->
        <div class="flex flex-row flex-wrap gap-4 mt-4">
          <span v-for="lang in Object.keys(repo.languages)" :key="lang" class="dark:text-primary px-2 py-1 rounded-full text-sm border-2 border-accent/50
            text-shadow-sm">{{ lang }}</span>
        </div>

        <!-- commit rows -->
        <div class="md:flex flex-row hidden mt-5 border-t-2 dark:border-background">
          <h3 class="dark:text-primary font-semibold mt-2 text-shadow-sm">Commit History</h3>
        </div>

        <div class="relative hidden md:block">
          <div v-if="store.commitsByRepo[repo.title]" class="mt-2">
            <span v-for="commit in getCommitsForRepo(repo.title)" class="dark:text-secondary/90
              block md:ml-6 ">
              <span class="absolute -ml-6 mt-1.5 w-3 h-3 bg-accent/20 rounded-full group-hover:animate-ping"></span>
              <span class="absolute -ml-6 mt-1.5 w-3 h-3 bg-accent/20 rounded-full"></span>
              {{ commit.author }}: "{{ commit.msg }}" - {{ commit.time.replace(/-/g, '/') }}
            </span>
          </div>
          <div v-else class="mt-2">
            <span class="dark:text-secondary/90 text-shadow-sm block italic">Loading commit history...</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
