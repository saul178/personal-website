import { defineStore } from "pinia";
import instance from "@/services/api";

export const useGithubStore = defineStore("github", {
  state: () => ({
    repos: [],
    commitsByRepo: {},
  }),
  getters: {
    getGithubRepos(state) {
      return state.repos
    },

    getGithubCommits(state) {
      return state.commits
    }
  },
  actions: {
    async fetchGithubRepos() {
      try {
        const resp = await instance.get("/github/repos")
        this.repos = resp.data
      }
      catch (error) {
        alert(error)
        console.log(error)
      }
    },

    async fetchGithubCommits(repoName) {
      try {
        const resp = await instance.get(`/github/commits?repo=${repoName}`)
        this.commitsByRepo[repoName] = resp.data
      }
      catch (error) {
        alert(error)
        console.log(error)
      }
    }
  },
})
