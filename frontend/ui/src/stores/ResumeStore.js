import { defineStore } from "pinia";
import instance from "@/services/api";

export const useResumeStore = defineStore("resume", {
  state: () => ({
    education: [],
    skills: {
      os: {},
      languages: {},
      tools: {},
      frameworks: {},
      databases: {},
    },

  }),
  getters: {
    getEducation(state) {
      return state.education
    },

    getSkills(state) {
      return state.skills
    }
  },
  actions: {
    async fetchResumeData() {
      try {
        const resp = await instance.get("/resume")
        this.education = resp.data.education
        this.skills = resp.data.skills
      }
      catch (error) {
        alert("failed to fetch resume data: ", error)
        console.log(error)
      }
    }
  }
})

