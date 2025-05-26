<script setup>
</script>

<template>
  <!-- Container for both cards -->
  <div class="px-6 py-12 flex flex-col md:flex-row space-x-32 space-y-16 justify-center">

    <!-- Profile Card -->
    <div class="dark:bg-foreground rounded-2xl shadow-lg p-6 w-full max-w-2xl text-center">
      <img src="../assets/images/me.jpeg" alt="Profile Picture" class="w-52 h-52 rounded-full
        mx-auto shadow-xl" />
      <div class="mt-6">
        <h1 class="text-2xl font-bold dark:text-primary">Welcome to my site!</h1>
        <p class="mt-2 dark:text-primary text-left">
          My name is Saul Gonzalez, at a young age I have always loved messing with computers and
          being the families in house IT person. In my
          journey to complete my Computer Science degree, I rediscovered a passion that I lost through time. I fell in
          love with building computers, programming slowly became a new passion, and I fell in love with Linux. I always
          strive to break my skill cap by learning from others or challenging myself to learn new things.

          I have worked in retail for years, while I appreciate and value all the skills that I have learned from my
          job. I always knew that I could do more for myself and strive for higher levels. The completion of my degree
          is proof that I can do anything I set my mind to and I plan to reach higher levels through out my journey in
          life.
        </p>
      </div>
    </div>

    <!-- Resume Card -->
    <div class="dark:bg-foreground rounded-2xl shadow-lg p-6 w-full max-w-2xl text-center">
      <h1 class="text-2xl font-bold dark:text-primary">My Skills</h1>
      <!-- golang backend will fill up this information with a metadata json file -->
      <!-- OS Grid -->
      <h1 class="text-lg font-medium mt-4 dark:text-primary"> Proficient with these OS's </h1>
      <div class="flex items-start gap-4 mt-3">
        <div class="hidden md:flex flex-col items-center pt-8 px-3">
          <font-awesome-icon :icon="['fas', 'desktop']" class="dark:text-primary text-4xl mb-2" />
        </div>
        <div class="grid grid-cols-2 md:grid-cols-3 gap-4 w-full">
          <div class="dark:bg-background p-3 rounded-lg shadow-md shadow-satin-sheen-gold-500 text-center">
            <font-awesome-icon :icon="['fab', 'linux']" class="dark:text-primary text-5xl" />
            <span class="dark:text-primary block mt-1">Linux</span>
          </div>
          <div class="dark:bg-background p-3 rounded-lg shadow-md shadow-satin-sheen-gold-500 text-center">
            <font-awesome-icon :icon="['fab', 'windows']" class="dark:text-primary text-5xl" alt="Windows" />
            <span class="dark:text-primary block mt-1"> Windows </span>
          </div>
          <div class="dark:bg-background p-3 rounded-lg shadow-md shadow-satin-sheen-gold-500 text-center">
            <font-awesome-icon :icon="['fab', 'apple']" class="dark:text-primary text-5xl" alt="Apple" />
            <span class="dark:text-primary block mt-1"> Apple </span>
          </div>
        </div>
      </div>

      <!-- language grid -->
      <h1 class="text-lg font-medium mt-4 pt-2 dark:text-primary"> Proficient in these languages </h1>
      <div class="flex items-start gap-4 mt-3">
        <div class="hidden md:flex flex-col items-center pt-12 px-3">
          <font-awesome-icon :icon="['fas', 'code']"
            class="dark:text-primary text-4xl drop-shadow-[0_0_3px_rgba(255,215,0,0.8)]" alt="Code" />
        </div>
        <div class="grid grid-cols-2 md:grid-cols-4 gap-4 mt-4 w-full">
          <div class="dark:bg-background p-3 rounded-lg shadow-md shadow-satin-sheen-gold-500 text-center">
            <font-awesome-icon :icon="['fab', 'golang']" class="dark:text-primary text-5xl" alt="Golang" />
            <span class="dark:text-primary block mt-1"> Golang </span>
          </div>
          <div class="dark:bg-background p-3 rounded-lg shadow-md shadow-satin-sheen-gold-500 text-center">
            <font-awesome-icon :icon="['fab', 'python']" class="dark:text-primary text-5xl" alt="Python" />
            <span class="dark:text-primary block mt-1"> Python </span>
          </div>
          <div class="dark:bg-background p-3 rounded-lg shadow-md shadow-satin-sheen-gold-500 text-center">
            <font-awesome-icon :icon="['fab', 'java']" class="dark:text-primary text-5xl" alt="Java" />
            <span class="dark:text-primary block mt-1"> Java </span>
          </div>
          <div class="dark:bg-background p-3 rounded-lg shadow-md shadow-satin-sheen-gold-500 text-center">
            <font-awesome-icon :icon="['fab', 'js']" class="dark:text-primary text-5xl" alt="JavaScript" />
            <span class="dark:text-primary block mt-1"> JavaScript </span>
          </div>
        </div>
      </div>

      <!-- tools and frameworks -->
      <h1 class="text-lg font-medium mt-4 pt-2 dark:text-primary"> Proficient with these
        technologies </h1>
      <div class="flex items-start gap-4 mt-3">
        <div class="hidden md:flex flex-col items-center pt-12 px-3">
          <font-awesome-icon :icon="['fas', 'screwdriver-wrench']" class="dark:text-primary
            text-4xl" alt="Tools" />
        </div>
        <div class="grid grid-cols-2 md:grid-cols-4 gap-4 mt-4 w-full">
          <div class="dark:bg-background p-3 rounded-lg shadow-md shadow-satin-sheen-gold-500 text-center">
            <font-awesome-icon :icon="['fab', 'docker']" class="dark:text-primary text-5xl" alt="docker" />
            <span class="dark:text-primary block mt-1"> Docker </span>
          </div>

          <div class="dark:bg-background p-3 rounded-lg shadow-md shadow-satin-sheen-gold-500 text-center">
            <font-awesome-icon :icon="['fab', 'git-alt']" class="dark:text-primary text-5xl" alt="Git" />
            <span class="dark:text-primary block mt-1"> Git </span>
          </div>
          <div class="dark:bg-background p-3 rounded-lg shadow-md shadow-satin-sheen-gold-500 text-center">
            <font-awesome-icon :icon="['fab', 'github']" class="dark:text-primary text-5xl" alt="Github" />
            <span class="dark:text-primary block mt-1"> Github </span>
          </div>
        </div>
      </div>

      <!-- framework grid -->
      <h1 class="text-lg font-medium mt-4 pt-2 dark:text-primary"> Proficient with these frameworks </h1>
      <div class="flex items-start gap-4 mt-3">
        <div class="hidden md:flex flex-col items-center pt-12 px-3">
          <font-awesome-icon :icon="['fas', 'layer-group']" class="dark:text-primary text-4xl" alt="Frameworks" />
        </div>
        <div class="grid grid-cols-2 md:grid-cols-4 gap-4 mt-4 w-full">
          <div class="dark:bg-background p-3 rounded-lg shadow-md shadow-satin-sheen-gold-500 text-center">
            <font-awesome-icon :icon="['fab', 'vuejs']" class="dark:text-primary text-5xl" alt="Vuejs" />
            <span class="dark:text-primary block mt-1"> Vuejs </span>
          </div>
          <div class="dark:bg-background p-3 rounded-lg shadow-md shadow-satin-sheen-gold-500 text-center">
            <svg viewBox="0 0 15 15" fill="none" class="w-12 h-12 mx-auto dark:text-primary"
              xmlns="http://www.w3.org/2000/svg">
              <g id="SVGRepo_bgCarrier" stroke-width="0"></g>
              <g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g>
              <g id="SVGRepo_iconCarrier">
                <path fill-rule="evenodd" clip-rule="evenodd"
                  d="M7.5 3C5.63333 3 4.46667 4 4 5.99999C4.7 4.99999 5.51667 4.625 6.45 4.87499C6.98252 5.01763 7.36314 5.43155 7.78443 5.88974C8.47074 6.63613 9.26506 7.49999 11 7.49999C12.8667 7.49999 14.0333 6.49999 14.5 4.5C13.8 5.49999 12.9833 5.87499 12.05 5.62499C11.5175 5.48235 11.1369 5.06844 10.7156 4.61025C10.0293 3.86386 9.23494 3 7.5 3ZM4 7.49999C2.13333 7.49999 0.966667 8.49998 0.5 10.5C1.2 9.49998 2.01667 9.12498 2.95 9.37498C3.48252 9.51762 3.86314 9.93154 4.28443 10.3897C4.97074 11.1361 5.76506 12 7.5 12C9.36667 12 10.5333 11 11 8.99998C10.3 9.99998 9.48333 10.375 8.55 10.125C8.01748 9.98234 7.63686 9.56843 7.21557 9.11023C6.52926 8.36385 5.73494 7.49999 4 7.49999Z"
                  stroke="currentColor" stroke-linejoin="round"></path>
              </g>
            </svg>
            <span class="dark:text-primary block mt-1"> Tailwindcss </span>
          </div>
          <div class="dark:bg-background p-3 rounded-lg shadow-md shadow-satin-sheen-gold-500 text-center">
            <svg fill="#000000" viewBox="0 0 32 32" class="dark:text-primary w-12 h-12 mx-auto"
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
                  0.745-1.495 0.984-3.229 0.984-7.781v-10.328zM23.188 0.026h5.229v5.37h-5.229z" stroke="currentColor">
                </path>
              </g>
            </svg>
            <span class="dark:text-primary block mt-1"> Django </span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
