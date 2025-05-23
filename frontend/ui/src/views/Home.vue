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
      <h1 class="text-lg font-medium mt-4 pt-2 dark:text-primary"> Proficient with these tools and frameworks </h1>
      <div class="flex items-start gap-4 mt-3">
        <div class="hidden md:flex flex-col items-center pt-12 px-3">
          <font-awesome-icon :icon="['fas', 'screwdriver-wrench']" class="dark:text-primary
            text-4xl" alt="Tools" />
        </div>
        <div class="grid grid-cols-2 md:grid-cols-4 gap-4 mt-4 w-full">
          <div class="dark:bg-background p-3 rounded-lg shadow-md shadow-satin-sheen-gold-500 text-center">
            <font-awesome-icon :icon="['fab', 'vuejs']" class="dark:text-primary text-5xl" alt="Vuejs" />
            <span class="dark:text-primary block mt-1"> Vuejs </span>
          </div>
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
    </div>
  </div>
</template>
