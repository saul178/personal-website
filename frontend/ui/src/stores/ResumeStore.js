import { defineStore } from "pinia";
import axiousInstance from "@/services/api";

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
    pdfBlobUrl: null,
    error: null,
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
        const resp = await axiousInstance.get("/resume")
        this.education = resp.data.education
        this.skills = resp.data.skills
      }
      catch (error) {
        console.warn("failed to fetch resume data, possible backend issue: ", error)
      }
    },

    async fetchResumePdf() {
      this.error = null
      try {
        const resp = await axiousInstance.get("/download-resume", {
          responseType: 'blob',
        })

        const blob = new Blob([resp.data], { type: 'application/pdf' })

        // NOTE: supposedly this helps prevent memory leaks when dealing with blobs
        // garbage collector marks blobs as uncollectable so you have to set to release the
        // protection so that the GC can collect it when it runs
        if (this.pdfBlobUrl) {
          URL.revokeObjectURL(this.pdfBlobUrl)
        }
        this.pdfBlobUrl = URL.createObjectURL(blob)
      }
      catch (error) {
        console.warn("pdf fetch failed, possible issues with backend: ", error)
        this.error = error
      }
    }
  }
})

