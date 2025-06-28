import { defineStore } from "pinia";
import axiosInstance from "@/services/api";

export const useGithubStore = defineStore("github", {
  state: () => ({
    repos: [],
    commitsByRepo: {},
    repoImgs: {
      "personal-website": "personal-website.png",
      "manga-library-proj": "",
      "Personal-SampleShare": "Personal-SampleShare.png",
    },
  }),

  getters: {
    getGithubRepos(state) {
      return state.repos
    },

    getGithubCommits(state) {
      return state.commitsByRepo
    },

    getRepoImg(state) {
      return state.repoImgs
    }
  },

  actions: {
    async fetchGithubRepos() {
      try {
        const resp = await axiosInstance.get("/github/repos")
        this.repos = resp.data
      }
      catch (error) {
        console.warn("failed to fetch github repo request from backend, possible api/network issues: ", error)
      }
    },

    async fetchGithubCommits(repoName) {
      try {
        const resp = await axiosInstance.get(`/github/commits?repo=${repoName}`)
        this.commitsByRepo[repoName] = resp.data
      }
      catch (error) {
        console.warn("failed to fetch repo commit history request from backend, possible api/network issues: ", error)
      }
    },

    fetchRepoImg(repoName) {
      const imgFile = this.repoImgs[repoName]
      let fileName

      if (imgFile) {
        fileName = imgFile
      } else {
        fileName = "default-image.png"
      }

      try {
        return new URL(`../assets/images/repo-imgs/${fileName}`, import.meta.url).href
      }
      catch (error) {
        console.warn(`Failed to load image for ${repoName}:`, error.message)
        return new URL("../assets/images/repo-imgs/default-image.png", import.meta.url).href
      }
    }
  },
})
