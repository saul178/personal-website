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
        <h1 class="text-2xl font-bold dark:text-primary">About Me</h1>
        <p class="mt-2 dark:text-primary text-left">
          At a young age I have always loved messing with computers and being the families in house IT person. In my
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
      <!-- This will hold resume info later -->
      <h1 class="text-2xl font-bold dark:text-primary">Resume</h1>
      <p class="mt-2 dark:text-primary">
        Dummy data for now.
      </p>
    </div>
  </div>
</template>
