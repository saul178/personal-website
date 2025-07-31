<script setup>
import { useResumeStore } from '@/stores/ResumeStore';
import { computed, onMounted } from 'vue';

const store = useResumeStore()

const getEducation = computed(() => store.getEducation)
const getSkills = computed(() => store.getSkills)
const getPdfUrl = computed(() => store.pdfBlobUrl)

onMounted(() => {
  store.fetchResumeData()
  store.fetchResumePdf()
})
</script>

<template>
  <!-- Container for both cards -->
  <Transition name="slide-fade" appear>
    <div class="px-6 py-12 flex flex-col md:flex-row md:space-x-14 space-y-12 justify-around">
      <!-- Profile Card -->
      <div class="flex flex-col space-y-6 w-full max-w-2xl">
        <div>
          <img src="../assets/images/me.jpeg" alt="Profile Picture" class="w-52 h-52 rounded-full
            mx-auto shadow-xl" />
        </div>

        <div class="group bg-foreground hover:bg-foreground/60 backdrop-blur-sm
          rounded-2xl shadow-lg p-6 border border-accent/20 transition-all duration-300
          hover:shadow-accent/20">
          <div class="mt-2">
            <h1 class="text-2xl font-bold dark:text-primary text-center text-shadow-sm">Hi, I'm Saul Gonzalez</h1>
            <p class="mt-2 dark:text-secondary/90 text-shadow-2xs text-left">
              I'm a recent Computer Science graduate from Metropolitan State University of
              Denver, and this is my personal website. The tech stack consists of Vue.js, Tailwind CSS, and a Golang
              backend powered by the Gin Framework. I've deployed this project on my personal server using
              Nginx Proxy Manager and Docker to host all the services used.
              <br>
              <br>
              I've built this not only to showcase what I'm capable of, but also to explore and answer many of the
              questions I've had about development: what is a backend? How does the frontend communicate with it? How
              do HTML and CSS work together? How does the world wide web work? What does it take to actually release
              and self host a product to the cold outskirts of the internet? And many more. Feel free to explore my
              other projects, I made it so that it actively shows what I'm working on.
              <br>
              <br>
              As for why I chose the tools that I did for this project, it's honestly really simple. I'm a big fan of
              Golang, Vue seemed like a good fit for a project of this size, and Tailwind CSS just looked fun to learn.
              There were plenty of ups and downs building this, but I plan to maintain this site and eventually turn it
              into more than just a place to show off my resume.
              <br>
              <br>
              You might notice the personal section is currently empty, but that's where I plan to share more about my
              personal life outside of tech.
            </p>
          </div>
        </div>

        <!-- education card -->
        <div
          class="group bg-foreground/50 backdrop-blur-sm border border-accent/20 rounded-2xl shadow-lg p-6 transition-transform hover:shadow-accent/20">
          <div class="flex items-center gap-4 mb-2">
            <div class="w-12 h-12 bg-accent/10 flex items-center justify-center rounded-full text-accent text-2xl">
              <font-awesome-icon :icon="['fas', 'graduation-cap']" />
            </div>
            <div>
              <h2 class="text-xl font-bold dark:text-primary text-shadow-sm">Education</h2>
            </div>
          </div>
          <div class="relative border-l-2 border-background group-hover:border-accent/20 pl-12 space-y-4">
            <div v-for="(school, i) in getEducation" :key="i" class="relative mt-2">
              <span class="absolute -left-8 top-1 w-4 h-4 bg-accent/20 rounded-full group-hover:animate-ping"></span>
              <span class="absolute -left-8 top-1 w-4 h-4 bg-accent/20 rounded-full"></span>
              <h3 class="text-xl font-semibold dark:text-primary text-shadow-sm">{{ school.degree }}</h3>
              <p class="text-sm dark:text-secondary/90 text-shadow-sm">
                {{ school.school }} ({{ school.date_completed }})
              </p>
            </div>
          </div>
        </div>

        <!-- contact card -->
        <div class="flex justify-center items-center">
          <div class="group bg-foreground hover:bg-foreground/60 backdrop-blur-sm
          rounded-2xl shadow-lg p-6 border border-accent/20 transition-all duration-300
          hover:shadow-accent/20">
            <div class="dark:text-primary text-2xl font-semibold text-center mb-2">Connect With Me</div>
            <div class="flex justify-center gap-4">
              <div class="flex items-center justify-center hover:scale-110
              transition-transform">
                <a href="https://www.linkedin.com/in/saul0178">
                  <font-awesome-icon :icon="['fab', 'linkedin']"
                    class="text-4xl dark:text-primary hover:text-primary/60" />
                </a>
              </div>
              <div class="flex items-center justify-center hover:scale-110
              transition-transform">
                <a href="https://github.com/saul178">
                  <font-awesome-icon :icon="['fab', 'github']"
                    class="text-4xl dark:text-primary hover:text-primary/60" />
                </a>
              </div>
            </div>
            <div class="flex justify-center mt-4">
              <div class="text-xl font-semibold text-center dark:text-primary">Download Resume
                <div class="flex justify-center">
                  <a :href="getPdfUrl" download="resume">
                    <div class="w-12 h-12 bg-accent/10 flex items-center justify-center rounded-full
                    text-accent text-2xl hover:scale-110 transition-transform cursor-pointer">
                      <font-awesome-icon :icon="['fas', 'file-arrow-down']" />
                    </div>
                  </a>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>


      <!-- 2nd col with skill cards -->
      <div class="px-6 py-12 flex flex-wrap justify-center items-center gap-8
      md:justify-end md:items-end">
        <!-- OS grid card -->
        <div class="flex flex-col space-y-6">
          <div class="border-b dark:border-b-accent/20">
            <div class="bg-gradient-to-r dark:from-primary to-accent bg-clip-text text-transparent
            text-4xl text-center font-semibold mb-3 text-shadow-md">
              My Skills
            </div>
          </div>
          <div class="group w-full max-w-2xl hover:-translate-y-1 bg-foreground hover:bg-foreground/60 backdrop-blur-sm
          rounded-2xl shadow-lg p-6 border border-accent/20 transition-all duration-300
          hover:shadow-accent/20">
            <div class="flex items-start gap-4">
              <div class="hidden w-16 h-16 bg-accent/10 rounded-full md:flex flex-col items-center justify-center group-hover:scale-110
              transition-transform pt-2 px-3">
                <font-awesome-icon :icon="['fas', 'desktop']" class="dark:text-primary text-4xl mb-2" />
              </div>
              <div class="flex flex-col gap-2 items-center md:items-start">
                <div>
                  <div class="bg-gradient-to-r dark:from-primary to-accent bg-clip-text text-transparent text-3xl font-semibold text-center
                  md:text-start text-shadow-sm">OS's
                  </div>
                  <div class="bg-gradient-to-r dark:from-primary to-accent bg-clip-text text-transparent text-xl text-center
                  md:text-start">Linux lover, but proficient with these as well
                  </div>
                </div>

                <div class="flex flex-wrap gap-2 items-center justify-center md:items-start
                  md:justify-start border-t-2 dark:border-background hover:border-accent/50 transition-colors">
                  <span class="flex items-center px-8 py-2 dark:bg-background/60 rounded-xl dark:text-primary border
                  border-accent/20 hover:border-accent/50 transition-colors text-center gap-2 mt-2">
                    <font-awesome-icon :icon="['fab', 'linux']" class="dark:text-primary text-4xl" />
                    <span class="dark:text-primary/90 text-lg">{{ getSkills.os.linux }}</span>
                  </span>
                  <span class="flex items-center px-8 py-2 dark:bg-background/60 rounded-xl dark:text-primary border
                  border-accent/20 hover:border-accent/50 transition-colors text-center gap-2 mt-2">
                    <font-awesome-icon :icon="['fab', 'windows']" class="dark:text-primary text-4xl" alt="Windows" />
                    <span class="dark:text-primary/90 text-lg"> {{ getSkills.os.windows }} </span>
                  </span>
                  <span class="flex items-center px-8 py-2 dark:bg-background/60 rounded-xl dark:text-primary border
                  border-accent/20 hover:border-accent/50 transition-colors text-center gap-2 mt-2">
                    <font-awesome-icon :icon="['fab', 'apple']" class="dark:text-primary text-4xl" alt="Apple" />
                    <span class="dark:text-primary/90 text-lg"> {{ getSkills.os.apple }}</span>
                  </span>
                </div>
              </div>
            </div>
          </div>

          <!-- language grid -->
          <div class="flex flex-col space-y-6">
            <div class="group w-full max-w-2xl hover:-translate-y-1 bg-foreground hover:bg-foreground/60
            backdrop-blur-sm rounded-2xl shadow-lg p-6 border border-accent/20 transition-all
            duration-300 hover:shadow-accent/20">
              <div class="flex items-start gap-4">
                <div class="hidden w-16 h-16 bg-accent/10 rounded-full md:flex flex-col items-center justify-center group-hover:scale-110
                transition-transform pt-2 px-3">
                  <font-awesome-icon :icon="['fas', 'code']" class="dark:text-primary text-4xl mb-2" />
                </div>
                <div class="flex flex-col gap-2 items-center md:items-start">
                  <div>
                    <div class="bg-gradient-to-r dark:from-primary to-accent bg-clip-text text-transparent text-3xl font-semibold text-center
                  md:text-start text-shadow-sm">Languages
                    </div>
                    <div class="bg-gradient-to-r dark:from-secondary to-accent bg-clip-text text-transparent text-xl text-center
                  md:text-start">Proficient with these languages
                    </div>
                  </div>
                  <div class="flex flex-wrap gap-2 items-center justify-center md:items-start
                  md:justify-start border-t-2 dark:border-background hover:border-accent/50 transition-colors">
                    <span class="flex items-center px-8 py-2 dark:bg-background/60 rounded-xl dark:text-primary border
                    border-accent/20 hover:border-accent/50 transition-colors text-center gap-2 mt-2">
                      <font-awesome-icon :icon="['fab', 'golang']" class="dark:text-primary text-4xl" alt="Golang" />
                      <span class="dark:text-primary/90 text-lg"> {{ getSkills.languages.golang }} </span>
                    </span>
                    <span class="flex items-center px-8 py-2 dark:bg-background/60 rounded-xl dark:text-primary border
                    border-accent/20 hover:border-accent/50 transition-colors text-center gap-2 mt-2">
                      <font-awesome-icon :icon="['fab', 'python']" class="dark:text-primary text-4xl" alt="Python" />
                      <span class="dark:text-primary/90 text-lg"> {{ getSkills.languages.python }} </span>
                    </span>
                    <span class="flex items-center px-8 py-2 dark:bg-background/60 rounded-xl dark:text-primary border
                    border-accent/20 hover:border-accent/50 transition-colors text-center gap-2 mt-2">
                      <font-awesome-icon :icon="['fab', 'java']" class="dark:text-primary text-4xl" alt="Java" />
                      <span class="dark:text-primary/90 text-lg"> {{ getSkills.languages.java }} </span>
                    </span>
                    <span class="flex items-center px-8 py-2 dark:bg-background/60 rounded-xl dark:text-primary border
                    border-accent/20 hover:border-accent/50 transition-colors text-center gap-2 mt-2">
                      <font-awesome-icon :icon="['fab', 'js']" class="dark:text-primary text-4xl" alt="JavaScript" />
                      <span class="dark:text-primary/90 text-lg"> {{ getSkills.languages.javascript }} </span>
                    </span>
                    <span class="flex items-center px-8 py-2 dark:bg-background/60 rounded-xl dark:text-primary border
                    border-accent/20 hover:border-accent/50 transition-colors text-center gap-2 mt-2">
                      <font-awesome-icon :icon="['fab', 'html5']" class="dark:text-primary text-4xl" alt="HTML" />
                      <span class="dark:text-primary/90 text-lg">{{ getSkills.languages.html }}</span>
                    </span>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- tools -->
          <div class="flex flex-col space-y-6">
            <div class="group w-full max-w-2xl hover:-translate-y-1 bg-foreground hover:bg-foreground/60
            backdrop-blur-sm rounded-2xl shadow-lg p-6 border border-accent/20 transition-all
            duration-300 hover:shadow-accent/20">
              <div class="flex items-start gap-4">
                <div class="hidden w-16 h-16 bg-accent/10 rounded-full md:flex flex-col items-center justify-center group-hover:scale-110
                transition-transform pt-2 px-3">
                  <font-awesome-icon :icon="['fas', 'screwdriver-wrench']" class="dark:text-primary
                  text-4xl mb-2" alt="Tools" />
                </div>
                <div class="flex flex-col gap-2 items-center md:items-start">
                  <div>
                    <div class="bg-gradient-to-r dark:from-primary to-accent bg-clip-text text-transparent text-3xl font-semibold text-center
                  md:text-start text-shadow-sm">Tools
                    </div>
                    <div class="bg-gradient-to-r dark:from-secondary to-accent bg-clip-text text-transparent text-xl text-center
                  md:text-start">Proficient with these tools
                    </div>
                  </div>
                  <div class="flex flex-wrap gap-2 items-center justify-center md:items-start
                    md:justify-start border-t-2 dark:border-background hover:border-accent/50 transition-colors">
                    <span class="flex items-center px-8 py-2 dark:bg-background/60 rounded-xl dark:text-primary border
                      border-accent/20 hover:border-accent/50 transition-colors text-center gap-2 mt-2">
                      <font-awesome-icon :icon="['fab', 'docker']" class="dark:text-primary text-4xl" alt="docker" />
                      <span class="dark:text-primary/90 text-lg"> {{ getSkills.tools.docker }} </span>
                    </span>
                    <span class="flex items-center px-8 py-2 dark:bg-background/60 rounded-xl dark:text-primary border
                      border-accent/20 hover:border-accent/50 transition-colors text-center gap-2 mt-2">
                      <font-awesome-icon :icon="['fab', 'git-alt']" class="dark:text-primary text-4xl" alt="Git" />
                      <span class="dark:text-primary/90 text-lg"> {{ getSkills.tools.git }} </span>
                    </span>
                    <span class="flex items-center px-8 py-2 dark:bg-background/60 rounded-xl dark:text-primary border
                      border-accent/20 hover:border-accent/50 transition-colors text-center gap-2 mt-2">
                      <font-awesome-icon :icon="['fab', 'github']" class="dark:text-primary text-4xl" alt="Github" />
                      <span class="dark:text-primary/90 text-lg"> {{ getSkills.tools.github }} </span>
                    </span>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- framework grid -->
          <div class="flex flex-col space-y-6">
            <div class="group w-full max-w-2xl hover:-translate-y-1 bg-foreground hover:bg-foreground/60
            backdrop-blur-sm rounded-2xl shadow-lg p-6 border border-accent/20 transition-all
            duration-300 hover:shadow-accent/20">
              <div class="flex items-start gap-4">
                <div class="hidden md:w-16 md:h-16 md:bg-accent/10 md:rounded-full md:flex flex-col items-center justify-center group-hover:scale-110
                transition-transform pt-2 px-3">
                  <font-awesome-icon :icon="['fas', 'layer-group']" class="dark:text-primary text-4xl"
                    alt="Frameworks" />
                </div>
                <div class="flex flex-col gap-2 items-center md:items-start">
                  <div>
                    <div class="bg-gradient-to-r dark:from-primary to-accent bg-clip-text text-transparent text-3xl font-semibold text-center
                  md:text-start text-shadow-sm">Frameworks
                    </div>
                    <div class="bg-gradient-to-r dark:from-secondary to-accent bg-clip-text text-transparent text-xl text-center
                  md:text-start">Proficient with these frameworks
                    </div>
                  </div>
                  <div class="flex flex-wrap gap-2 items-center justify-center md:items-start
                    md:justify-start border-t-2 dark:border-background hover:border-accent/50 transition-colors">
                    <span class="flex items-center px-8 py-2 dark:bg-background/60 rounded-xl dark:text-primary border
                      border-accent/20 hover:border-accent/50 transition-colors text-center gap-2 mt-2">
                      <font-awesome-icon :icon="['fab', 'vuejs']" class="dark:text-primary text-4xl" alt="Vuejs" />
                      <span class="dark:text-primary/90 text-lg"> {{ getSkills.frameworks.vue }} </span>
                    </span>
                    <span class="flex items-center px-8 py-2 dark:bg-background/60 rounded-xl dark:text-primary border
                      border-accent/20 hover:border-accent/50 transition-colors text-center gap-2 mt-2">
                      <svg viewBox="0 0 15 15" fill="none" class="w-9 h-9 mx-auto dark:text-primary"
                        xmlns="http://www.w3.org/2000/svg">
                        <g id="SVGRepo_bgCarrier" stroke-width="0"></g>
                        <g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g>
                        <g id="SVGRepo_iconCarrier">
                          <path fill-rule="evenodd" clip-rule="evenodd"
                            d="M7.5 3C5.63333 3 4.46667 4 4 5.99999C4.7 4.99999 5.51667 4.625 6.45 4.87499C6.98252 5.01763 7.36314 5.43155 7.78443 5.88974C8.47074 6.63613 9.26506 7.49999 11 7.49999C12.8667 7.49999 14.0333 6.49999 14.5 4.5C13.8 5.49999 12.9833 5.87499 12.05 5.62499C11.5175 5.48235 11.1369 5.06844 10.7156 4.61025C10.0293 3.86386 9.23494 3 7.5 3ZM4 7.49999C2.13333 7.49999 0.966667 8.49998 0.5 10.5C1.2 9.49998 2.01667 9.12498 2.95 9.37498C3.48252 9.51762 3.86314 9.93154 4.28443 10.3897C4.97074 11.1361 5.76506 12 7.5 12C9.36667 12 10.5333 11 11 8.99998C10.3 9.99998 9.48333 10.375 8.55 10.125C8.01748 9.98234 7.63686 9.56843 7.21557 9.11023C6.52926 8.36385 5.73494 7.49999 4 7.49999Z"
                            stroke="currentColor" stroke-linejoin="round"></path>
                        </g>
                      </svg>
                      <span class="dark:text-primary/90 text-lg"> {{ getSkills.frameworks.tailwind }} </span>
                    </span>
                    <span class="flex items-center px-8 py-2 dark:bg-background/60 rounded-xl dark:text-primary border
                      border-accent/20 hover:border-accent/50 transition-colors text-center gap-2 mt-2">
                      <svg fill="" viewBox="0 0 32 32" class="dark:text-primary w-9 h-9 mx-auto"
                        xmlns="http://www.w3.org/2000/svg">
                        <g id="SVGRepo_bgCarrier" stroke-width="0"></g>
                        <g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g>
                        <g id="SVGRepo_iconCarrier">
                          <path d="M14.859 0h5.234v24.219c-2.682 0.51-4.656 0.714-6.797 0.714-6.385
                            0-9.714-2.885-9.714-8.427 0-5.333 3.531-8.797 9-8.797 0.849 0 1.495 0.068 2.276
                            0.271zM14.859 12.193c-0.568-0.193-1.167-0.286-1.766-0.276-2.651 0-4.177 1.63-4.177
                            4.49 0 2.786 1.458 4.313 4.146 4.313 0.578 0 1.052-0.031
                            1.797-0.135v-8.396zM28.417 8.078v12.13c0 4.177-0.302 6.188-1.219 7.917-0.849
                            1.667-1.974 2.719-4.281 3.875l-4.859-2.313c2.307-1.089 3.432-2.036 4.146-3.5
                            0.745-1.495 0.984-3.229 0.984-7.781v-10.328zM23.188 0.026h5.229v5.37h-5.229z"
                            stroke="currentColor">
                          </path>
                        </g>
                      </svg>
                      <span class="dark:text-primary/90 text-lg"> {{ getSkills.frameworks.django }} </span>
                    </span>
                    <span class="flex items-center px-8 py-2 dark:bg-background/60 rounded-xl dark:text-primary border
                      border-accent/20 hover:border-accent/50 transition-colors text-center gap-2 mt-2">
                      <font-awesome-icon :icon="['fab', 'bootstrap']" class="dark:text-primary text-4xl" />
                      <span class="dark:text-primary/90 text-lg">{{ getSkills.frameworks.bootstrap }}</span>
                    </span>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- database grid card -->
          <div class="flex flex-col space-y-6">
            <div class="group w-full max-w-2xl hover:-translate-y-1 bg-foreground hover:bg-foreground/60
            backdrop-blur-sm rounded-2xl shadow-lg p-6 border border-accent/20 transition-all
            duration-300 hover:shadow-accent/20">
              <div class="flex items-start gap-4">
                <div class="hidden w-16 h-16 bg-accent/10 rounded-full md:flex flex-col items-center justify-center group-hover:scale-110
                transition-transform pt-2 px-3">
                  <font-awesome-icon :icon="['fas', 'database']" class="dark:text-primary
                  text-4xl" alt="Database" />
                </div>
                <div class="flex flex-col gap-2 items-center md:items-start">
                  <div>
                    <div class="bg-gradient-to-r dark:from-primary to-accent bg-clip-text text-transparent text-3xl font-semibold text-center
                  md:text-start text-shadow-sm">Databases
                    </div>
                    <div class="bg-gradient-to-r dark:from-secondary to-accent bg-clip-text text-transparent text-xl text-center
                  md:text-start">Proficient with these relational databases
                    </div>
                  </div>
                  <div class="flex flex-wrap gap-2 items-center justify-center md:items-start
                      md:justify-start border-t-2 dark:border-background hover:border-accent/50 transition-colors">
                    <span class="flex items-center px-8 py-2 dark:bg-background/60 rounded-xl dark:text-primary border
                        border-accent/20 hover:border-accent/50 transition-colors text-center gap-2 mt-2">
                      <svg fill="currentColor" viewBox="0 0 24 24" class="dark:text-primary w-9 h-9 mx-auto"
                        xmlns="http://www.w3.org/2000/svg" stroke="currentColor" stroke-width="0.00024000000000000003">
                        <g id="SVGRepo_bgCarrier" stroke-width="0"></g>
                        <g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g>
                        <g id="SVGRepo_iconCarrier">
                          <path
                            d="m24.129 23.412-.508-.484c-.251-.331-.518-.624-.809-.891l-.005-.004q-.448-.407-.931-.774-.387-.266-1.064-.641c-.371-.167-.661-.46-.818-.824l-.004-.01-.048-.024c.212-.021.406-.06.592-.115l-.023.006.57-.157c.236-.074.509-.122.792-.133h.006c.298-.012.579-.06.847-.139l-.025.006q.194-.048.399-.109t.351-.109v-.169q-.145-.217-.351-.496c-.131-.178-.278-.333-.443-.468l-.005-.004q-.629-.556-1.303-1.076c-.396-.309-.845-.624-1.311-.916l-.068-.04c-.246-.162-.528-.312-.825-.435l-.034-.012q-.448-.182-.883-.399c-.097-.048-.21-.09-.327-.119l-.011-.002c-.117-.024-.217-.084-.29-.169l-.001-.001c-.138-.182-.259-.389-.355-.609l-.008-.02q-.145-.339-.314-.651-.363-.702-.702-1.427t-.651-1.452q-.217-.484-.399-.967c-.134-.354-.285-.657-.461-.942l.013.023c-.432-.736-.863-1.364-1.331-1.961l.028.038c-.463-.584-.943-1.106-1.459-1.59l-.008-.007c-.509-.478-1.057-.934-1.632-1.356l-.049-.035q-.896-.651-1.96-1.282c-.285-.168-.616-.305-.965-.393l-.026-.006-1.113-.278-.629-.048q-.314-.024-.629-.024c-.148-.078-.275-.171-.387-.279-.11-.105-.229-.204-.353-.295l-.01-.007c-.605-.353-1.308-.676-2.043-.93l-.085-.026c-.193-.113-.425-.179-.672-.179-.176 0-.345.034-.499.095l.009-.003c-.38.151-.67.458-.795.84l-.003.01c-.073.172-.115.371-.115.581 0 .368.13.705.347.968l-.002-.003q.544.725.834 1.14.217.291.448.605c.141.188.266.403.367.63l.008.021c.056.119.105.261.141.407l.003.016q.048.206.121.448.217.556.411 1.14c.141.425.297.785.478 1.128l-.019-.04q.145.266.291.52t.314.496c.065.098.147.179.241.242l.003.002c.099.072.164.185.169.313v.001c-.114.168-.191.369-.217.586l-.001.006c-.035.253-.085.478-.153.695l.008-.03c-.223.666-.351 1.434-.351 2.231 0 .258.013.512.04.763l-.003-.031c.06.958.349 1.838.812 2.6l-.014-.025c.197.295.408.552.641.787.168.188.412.306.684.306.152 0 .296-.037.422-.103l-.005.002c.35-.126.599-.446.617-.827v-.002c.048-.474.12-.898.219-1.312l-.013.067c.024-.063.038-.135.038-.211 0-.015-.001-.03-.002-.045v.002q-.012-.109.133-.206v.048q.145.339.302.677t.326.677c.295.449.608.841.952 1.202l-.003-.003c.345.372.721.706 1.127 1.001l.022.015c.212.162.398.337.566.528l.004.004c.158.186.347.339.56.454l.01.005v-.024h.048c-.039-.087-.102-.157-.18-.205l-.002-.001c-.079-.044-.147-.088-.211-.136l.005.003q-.217-.217-.448-.484t-.423-.508q-.508-.702-.969-1.467t-.871-1.555q-.194-.387-.375-.798t-.351-.798c-.049-.099-.083-.213-.096-.334v-.005c-.006-.115-.072-.214-.168-.265l-.002-.001c-.121.206-.255.384-.408.545l.001-.001c-.159.167-.289.364-.382.58l-.005.013c-.141.342-.244.739-.289 1.154l-.002.019q-.072.641-.145 1.318l-.048.024-.024.024c-.26-.053-.474-.219-.59-.443l-.002-.005q-.182-.351-.326-.69c-.248-.637-.402-1.374-.423-2.144v-.009c-.009-.122-.013-.265-.013-.408 0-.666.105-1.308.299-1.91l-.012.044q.072-.266.314-.896t.097-.871c-.05-.165-.143-.304-.265-.41l-.001-.001c-.122-.106-.233-.217-.335-.335l-.003-.004q-.169-.244-.326-.52t-.278-.544c-.165-.382-.334-.861-.474-1.353l-.022-.089c-.159-.565-.336-1.043-.546-1.503l.026.064c-.111-.252-.24-.47-.39-.669l.006.008q-.244-.326-.436-.617-.244-.314-.484-.605c-.163-.197-.308-.419-.426-.657l-.009-.02c-.048-.097-.09-.21-.119-.327l-.002-.011c-.011-.035-.017-.076-.017-.117 0-.082.024-.159.066-.223l-.001.002c.011-.056.037-.105.073-.145.039-.035.089-.061.143-.072h.002c.085-.055.188-.088.3-.088.084 0 .165.019.236.053l-.003-.001c.219.062.396.124.569.195l-.036-.013q.459.194.847.375c.298.142.552.292.792.459l-.018-.012q.194.121.387.266t.411.291h.339q.387 0 .822.037c.293.023.564.078.822.164l-.024-.007c.481.143.894.312 1.286.515l-.041-.019q.593.302 1.125.641c.589.367 1.098.743 1.577 1.154l-.017-.014c.5.428.954.867 1.38 1.331l.01.012c.416.454.813.947 1.176 1.464l.031.047c.334.472.671 1.018.974 1.584l.042.085c.081.154.163.343.234.536l.011.033q.097.278.217.57.266.605.57 1.221t.57 1.198l.532 1.161c.187.406.396.756.639 1.079l-.011-.015c.203.217.474.369.778.422l.008.001c.368.092.678.196.978.319l-.047-.017c.143.065.327.134.516.195l.04.011c.212.065.396.151.565.259l-.009-.005c.327.183.604.363.868.559l-.021-.015q.411.302.822.57.194.145.651.423t.484.52c-.114-.004-.249-.007-.384-.007-.492 0-.976.032-1.45.094l.056-.006c-.536.072-1.022.203-1.479.39l.04-.014c-.113.049-.248.094-.388.129l-.019.004c-.142.021-.252.135-.266.277v.001c.061.076.11.164.143.26l.002.006c.034.102.075.19.125.272l-.003-.006c.119.211.247.393.391.561l-.004-.005c.141.174.3.325.476.454l.007.005q.244.194.508.399c.161.126.343.25.532.362l.024.013c.284.174.614.34.958.479l.046.016c.374.15.695.324.993.531l-.016-.011q.291.169.58.375t.556.399c.073.072.137.152.191.239l.003.005c.091.104.217.175.36.193h.003v-.048c-.088-.067-.153-.16-.184-.267l-.001-.004c-.025-.102-.062-.191-.112-.273l.002.004zm-18.576-19.205q-.194 0-.363.012c-.115.008-.222.029-.323.063l.009-.003v.024h.048q.097.145.244.326t.266.351l.387.798.048-.024c.113-.082.2-.192.252-.321l.002-.005c.052-.139.082-.301.082-.469 0-.018 0-.036-.001-.054v.003c-.045-.044-.082-.096-.108-.154l-.001-.003-.081-.182c-.053-.084-.127-.15-.214-.192l-.003-.001c-.094-.045-.174-.102-.244-.169z">
                          </path>
                        </g>
                      </svg>
                      <span class="dark:text-primary/90 block mt-1"> {{ getSkills.databases.mysql }}</span>
                    </span>
                    <span class="flex items-center px-8 py-2 dark:bg-background/60 rounded-xl dark:text-primary border
                        border-accent/20 hover:border-accent/50 transition-colors text-center gap-2 mt-2">
                      <svg fill="currentColor" viewBox="0 0 32 32" class="dark:text-primary w-9 h-9 mx-auto"
                        xmlns="http://www.w3.org/2000/svg">
                        <g id="SVGRepo_bgCarrier" stroke-width="0"></g>
                        <g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g>
                        <g id="SVGRepo_iconCarrier">
                          <path
                            d="M22.839 0c-1.245 0.011-2.479 0.188-3.677 0.536l-0.083 0.027c-0.751-0.131-1.516-0.203-2.276-0.219-1.573-0.027-2.923 0.353-4.011 0.989-1.073-0.369-3.297-1.016-5.641-0.885-1.629 0.088-3.411 0.583-4.735 1.979-1.312 1.391-2.009 3.547-1.864 6.485 0.041 0.807 0.271 2.124 0.656 3.837 0.38 1.709 0.917 3.709 1.589 5.537 0.672 1.823 1.405 3.463 2.552 4.577 0.572 0.557 1.364 1.032 2.296 0.991 0.652-0.027 1.24-0.313 1.751-0.735 0.249 0.328 0.516 0.468 0.755 0.599 0.308 0.167 0.599 0.281 0.907 0.355 0.552 0.14 1.495 0.323 2.599 0.135 0.375-0.063 0.771-0.187 1.167-0.359 0.016 0.437 0.032 0.869 0.047 1.307 0.057 1.38 0.095 2.656 0.505 3.776 0.068 0.183 0.251 1.12 0.969 1.953 0.724 0.833 2.129 1.349 3.739 1.005 1.131-0.24 2.573-0.677 3.532-2.041 0.948-1.344 1.375-3.276 1.459-6.412 0.020-0.172 0.047-0.312 0.072-0.448l0.224 0.021h0.027c1.208 0.052 2.521-0.12 3.62-0.631 0.968-0.448 1.703-0.901 2.239-1.708 0.131-0.199 0.281-0.443 0.319-0.86 0.041-0.411-0.199-1.063-0.595-1.364-0.791-0.604-1.291-0.375-1.828-0.26-0.525 0.115-1.063 0.176-1.599 0.192 1.541-2.593 2.645-5.353 3.276-7.792 0.375-1.443 0.584-2.771 0.599-3.932 0.021-1.161-0.077-2.187-0.771-3.077-2.177-2.776-5.235-3.548-7.599-3.573-0.073 0-0.145 0-0.219 0zM22.776 0.855c2.235-0.021 5.093 0.604 7.145 3.228 0.464 0.589 0.6 1.448 0.584 2.511s-0.213 2.328-0.573 3.719c-0.692 2.699-2.011 5.833-3.859 8.652 0.063 0.047 0.135 0.088 0.208 0.115 0.385 0.161 1.265 0.296 3.025-0.063 0.443-0.095 0.767-0.156 1.105 0.099 0.167 0.14 0.255 0.349 0.244 0.568-0.020 0.161-0.077 0.317-0.177 0.448-0.339 0.509-1.009 0.995-1.869 1.396-0.76 0.353-1.855 0.536-2.817 0.547-0.489 0.005-0.937-0.032-1.319-0.152l-0.020-0.004c-0.147 1.411-0.484 4.203-0.704 5.473-0.176 1.025-0.484 1.844-1.072 2.453-0.589 0.615-1.417 0.979-2.537 1.219-1.385 0.297-2.391-0.021-3.041-0.568s-0.948-1.276-1.125-1.719c-0.124-0.307-0.187-0.703-0.249-1.235-0.063-0.531-0.104-1.177-0.136-1.911-0.041-1.12-0.057-2.24-0.041-3.365-0.577 0.532-1.296 0.88-2.068 1.016-0.921 0.156-1.739 0-2.228-0.12-0.24-0.063-0.475-0.151-0.693-0.271-0.229-0.12-0.443-0.255-0.588-0.527-0.084-0.156-0.109-0.337-0.073-0.509 0.041-0.177 0.145-0.328 0.287-0.443 0.265-0.215 0.615-0.333 1.14-0.443 0.959-0.199 1.297-0.333 1.5-0.496 0.172-0.135 0.371-0.416 0.713-0.828 0-0.015 0-0.036-0.005-0.052-0.619-0.020-1.224-0.181-1.771-0.479-0.197 0.208-1.224 1.292-2.468 2.792-0.521 0.624-1.099 0.984-1.713 1.011-0.609 0.025-1.163-0.281-1.631-0.735-0.937-0.912-1.688-2.48-2.339-4.251s-1.177-3.744-1.557-5.421c-0.375-1.683-0.599-3.037-0.631-3.688-0.14-2.776 0.511-4.645 1.625-5.828s2.641-1.625 4.131-1.713c2.672-0.151 5.213 0.781 5.724 0.979 0.989-0.672 2.265-1.088 3.859-1.063 0.756 0.011 1.505 0.109 2.24 0.292l0.027-0.016c0.323-0.109 0.651-0.208 0.984-0.28 0.907-0.215 1.833-0.324 2.76-0.339zM22.979 1.745h-0.197c-0.76 0.009-1.527 0.099-2.271 0.26 1.661 0.735 2.916 1.864 3.801 3 0.615 0.781 1.12 1.64 1.505 2.557 0.152 0.355 0.251 0.651 0.303 0.88 0.031 0.115 0.047 0.213 0.057 0.312 0 0.052 0.005 0.105-0.021 0.193 0 0.005-0.005 0.016-0.005 0.021 0.043 1.167-0.249 1.957-0.287 3.072-0.025 0.808 0.183 1.756 0.235 2.792 0.047 0.973-0.072 2.041-0.703 3.093 0.052 0.063 0.099 0.125 0.151 0.193 1.672-2.636 2.88-5.547 3.521-8.032 0.344-1.339 0.525-2.552 0.541-3.509 0.016-0.959-0.161-1.657-0.391-1.948-1.792-2.287-4.213-2.871-6.24-2.885zM16.588 2.088c-1.572 0.005-2.703 0.48-3.561 1.193-0.887 0.74-1.48 1.745-1.865 2.781-0.464 1.224-0.625 2.411-0.688 3.219l0.021-0.011c0.475-0.265 1.099-0.536 1.771-0.687 0.667-0.157 1.391-0.204 2.041 0.052 0.657 0.249 1.193 0.848 1.391 1.749 0.939 4.344-0.291 5.959-0.744 7.177-0.172 0.443-0.323 0.891-0.443 1.349 0.057-0.011 0.115-0.027 0.172-0.032 0.323-0.025 0.572 0.079 0.719 0.141 0.459 0.192 0.771 0.588 0.943 1.041 0.041 0.12 0.072 0.244 0.093 0.38 0.016 0.052 0.027 0.109 0.027 0.167-0.052 1.661-0.048 3.323 0.015 4.984 0.032 0.719 0.079 1.349 0.136 1.849 0.057 0.495 0.135 0.875 0.188 1.005 0.171 0.427 0.421 0.984 0.875 1.364 0.448 0.381 1.093 0.631 2.276 0.381 1.025-0.224 1.656-0.527 2.077-0.964 0.423-0.443 0.672-1.052 0.833-1.984 0.245-1.401 0.729-5.464 0.787-6.224-0.025-0.579 0.057-1.021 0.245-1.36 0.187-0.344 0.479-0.557 0.735-0.672 0.124-0.057 0.244-0.093 0.343-0.125-0.104-0.145-0.213-0.291-0.323-0.432-0.364-0.443-0.667-0.937-0.891-1.463-0.104-0.22-0.219-0.439-0.344-0.647-0.176-0.317-0.4-0.719-0.635-1.172-0.469-0.896-0.979-1.989-1.245-3.052-0.265-1.063-0.301-2.161 0.376-2.932 0.599-0.688 1.656-0.973 3.233-0.812-0.047-0.141-0.072-0.261-0.151-0.443-0.359-0.844-0.828-1.636-1.391-2.355-1.339-1.713-3.511-3.412-6.859-3.469zM7.735 2.156c-0.167 0-0.339 0.005-0.505 0.016-1.349 0.079-2.62 0.468-3.532 1.432-0.911 0.969-1.509 2.547-1.38 5.167 0.027 0.5 0.24 1.885 0.609 3.536 0.371 1.652 0.896 3.595 1.527 5.313 0.629 1.713 1.391 3.208 2.12 3.916 0.364 0.349 0.681 0.495 0.968 0.485 0.287-0.016 0.636-0.183 1.063-0.693 0.776-0.937 1.579-1.844 2.412-2.729-1.199-1.047-1.787-2.629-1.552-4.203 0.135-0.984 0.156-1.907 0.135-2.636-0.015-0.708-0.063-1.176-0.063-1.473 0-0.011 0-0.016 0-0.027v-0.005l-0.005-0.009c0-1.537 0.272-3.057 0.792-4.5 0.375-0.996 0.928-2 1.76-2.819-0.817-0.271-2.271-0.676-3.843-0.755-0.167-0.011-0.339-0.016-0.505-0.016zM24.265 9.197c-0.905 0.016-1.411 0.251-1.681 0.552-0.376 0.433-0.412 1.193-0.177 2.131 0.233 0.937 0.719 1.984 1.172 2.855 0.224 0.437 0.443 0.828 0.619 1.145 0.183 0.323 0.313 0.547 0.391 0.745 0.073 0.177 0.157 0.333 0.24 0.479 0.349-0.74 0.412-1.464 0.375-2.224-0.047-0.937-0.265-1.896-0.229-2.864 0.037-1.136 0.261-1.876 0.277-2.751-0.324-0.041-0.657-0.068-0.985-0.068zM13.287 9.355c-0.276 0-0.552 0.036-0.823 0.099-0.537 0.131-1.052 0.328-1.537 0.599-0.161 0.088-0.317 0.188-0.463 0.303l-0.032 0.025c0.011 0.199 0.047 0.667 0.063 1.365 0.016 0.76 0 1.728-0.145 2.776-0.323 2.281 1.333 4.167 3.276 4.172 0.115-0.469 0.301-0.944 0.489-1.443 0.541-1.459 1.604-2.521 0.708-6.677-0.145-0.677-0.437-0.953-0.839-1.109-0.224-0.079-0.457-0.115-0.697-0.109zM23.844 9.625h0.068c0.083 0.005 0.167 0.011 0.239 0.031 0.068 0.016 0.131 0.037 0.183 0.073 0.052 0.031 0.088 0.083 0.099 0.145v0.011c0 0.063-0.016 0.125-0.047 0.183-0.041 0.072-0.088 0.14-0.145 0.197-0.136 0.151-0.319 0.251-0.516 0.281-0.193 0.027-0.385-0.025-0.547-0.135-0.063-0.048-0.125-0.1-0.172-0.157-0.047-0.047-0.073-0.109-0.084-0.172-0.004-0.061 0.011-0.124 0.052-0.171 0.048-0.048 0.1-0.089 0.157-0.12 0.129-0.073 0.301-0.125 0.5-0.152 0.072-0.009 0.145-0.015 0.213-0.020zM13.416 9.849c0.068 0 0.147 0.005 0.22 0.015 0.208 0.032 0.385 0.084 0.525 0.167 0.068 0.032 0.131 0.084 0.177 0.141 0.052 0.063 0.077 0.14 0.073 0.224-0.016 0.077-0.048 0.151-0.1 0.208-0.057 0.068-0.119 0.125-0.192 0.172-0.172 0.125-0.385 0.177-0.599 0.151-0.215-0.036-0.412-0.14-0.557-0.301-0.063-0.068-0.115-0.141-0.157-0.219-0.047-0.073-0.067-0.156-0.057-0.24 0.021-0.14 0.141-0.219 0.256-0.26 0.131-0.043 0.271-0.057 0.411-0.052zM25.495 19.64h-0.005c-0.192 0.073-0.353 0.1-0.489 0.163-0.14 0.052-0.251 0.156-0.317 0.285-0.089 0.152-0.156 0.423-0.136 0.885 0.057 0.043 0.125 0.073 0.199 0.095 0.224 0.068 0.609 0.115 1.036 0.109 0.849-0.011 1.896-0.208 2.453-0.469 0.453-0.208 0.88-0.489 1.255-0.817-1.859 0.38-2.905 0.281-3.552 0.016-0.156-0.068-0.307-0.157-0.443-0.267zM14.787 19.765h-0.027c-0.072 0.005-0.172 0.032-0.375 0.251-0.464 0.52-0.625 0.848-1.005 1.151-0.385 0.307-0.88 0.469-1.875 0.672-0.312 0.063-0.495 0.135-0.615 0.192 0.036 0.032 0.036 0.043 0.093 0.068 0.147 0.084 0.333 0.152 0.485 0.193 0.427 0.104 1.124 0.229 1.859 0.104 0.729-0.125 1.489-0.475 2.141-1.385 0.115-0.156 0.124-0.391 0.031-0.641-0.093-0.244-0.297-0.463-0.437-0.52-0.089-0.043-0.183-0.068-0.276-0.084z">
                          </path>
                        </g>
                      </svg>
                      <span class="dark:text-primary/90 block mt-1"> {{ getSkills.databases.postgresql }}</span>
                    </span>
                    <span class="flex items-center px-8 py-2 dark:bg-background/60 rounded-xl dark:text-primary border
                        border-accent/20 hover:border-accent/50 transition-colors text-center gap-2 mt-2">
                      <svg fill="currentColor" class="dark:text-primary w-9 h-9 mx-auto" viewBox="0 0 32 32"
                        version="1.1" xmlns="http://www.w3.org/2000/svg">
                        <g id="SVGRepo_bgCarrier" stroke-width="0"></g>
                        <g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g>
                        <g id="SVGRepo_iconCarrier">
                          <title>sqlite</title>
                          <path
                            d="M4.884 2.334c-1.265 0.005-2.289 1.029-2.293 2.294v20.754c0.004 1.264 1.028 2.288 2.293 2.292h11.769c-0.056-0.671-0.088-1.452-0.088-2.241 0-0.401 0.008-0.801 0.025-1.198l-0.002 0.057c-0.008-0.077-0.014-0.176-0.020-0.25q-0.229-1.498-0.591-2.972c-0.119-0.504-0.277-0.944-0.478-1.36l0.017 0.039c-0.080-0.126-0.127-0.279-0.127-0.443 0-0.034 0.002-0.068 0.006-0.101l-0 0.004c0.003-0.173 0.020-0.339 0.049-0.501l-0.003 0.019c0.088-0.523 0.19-0.963 0.314-1.394l-0.022 0.088 0.271-0.035c-0.021-0.044-0.018-0.081-0.039-0.121l-0.051-0.476q0.224-0.751 0.477-1.492l0.25-0.024c-0.010-0.020-0.012-0.047-0.023-0.066l-0.054-0.395c1.006-4.731 3.107-8.864 6.029-12.272l-0.031 0.037c0.082-0.086 0.166-0.16 0.247-0.242zM28.094 1.655c-1.29-1.15-2.849-0.687-4.39 0.68q-0.356 0.319-0.684 0.669c-2.8 3.294-4.843 7.319-5.808 11.747l-0.033 0.18c0.261 0.551 0.494 1.201 0.664 1.876l0.016 0.075q0.115 0.436 0.205 0.878s-0.024-0.089-0.12-0.37l-0.062-0.182q-0.019-0.050-0.041-0.1c-0.172-0.4-0.647-1.243-0.857-1.611-0.179 0.529-0.337 1.022-0.47 1.47 0.413 0.863 0.749 1.867 0.959 2.917l0.014 0.083s-0.031-0.124-0.184-0.552c-0.342-0.739-0.664-1.338-1.015-1.919l0.050 0.089c-0.185 0.464-0.292 1.001-0.292 1.564 0 0.1 0.003 0.199 0.010 0.297l-0.001-0.013c0.219 0.426 0.401 0.921 0.519 1.439l0.008 0.043c0.357 1.375 0.606 3.049 0.606 3.049l0.021 0.28c-0.015 0.342-0.023 0.744-0.023 1.147 0 0.805 0.034 1.602 0.101 2.39l-0.007-0.103c0.058 1.206 0.283 2.339 0.651 3.406l-0.026-0.086 0.194-0.105c-0.346-1.193-0.545-2.564-0.545-3.981 0-0.344 0.012-0.684 0.035-1.022l-0.003 0.046c0.221-3.782 0.964-7.319 2.158-10.641l-0.083 0.264c1.655-4.9 4.359-9.073 7.861-12.417l0.012-0.011c-2.491 2.249-5.863 9.535-6.873 12.232-0.963 2.42-1.798 5.294-2.365 8.263l-0.048 0.305c0.664-1.639 1.914-2.926 3.483-3.622l0.042-0.017s1.321-1.63 2.864-3.956c-1.195 0.25-2.184 0.521-3.15 0.843l0.199-0.057c-0.75 0.314-0.952 0.421-0.952 0.421 1.288-0.791 2.777-1.515 4.337-2.092l0.178-0.058c2.867-4.515 5.991-10.929 2.845-13.736z">
                          </path>
                        </g>
                      </svg>
                      <span class="dark:text-primary/90 block mt-1"> {{ getSkills.databases.sqlite }}</span>
                    </span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </Transition>
</template>
