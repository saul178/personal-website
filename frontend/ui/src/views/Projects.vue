<!-- what to show from the github api
name or full name not sure if this is for the repo or user
description box of the repo
created at, updated at or pushed at?
commits??
url hyperlink?
see if i can get what languages are used???

at the top or bottom have total commits?

-->

<template>
  <div class="flex px-8 py-12 gap-8">
    <!-- Left Sidebar -->
    <div class="w-1/4 flex flex-col space-y-4">
      <h2 class="text-4xl font-bold dark:text-primary">Projects</h2>
      <a href="https://github.com/saul178" class="dark:text-primary text-2xl font-bold">
        <font-awesome-icon :icon="['fab', 'github']" class="text-4xl dark:text-primary" />
        View Github →
      </a>
    </div>

    <!-- Right Content Area -->
    <div class="flex-1 space-y-6">
      <!-- Project Card 1 -->
      <div class="bg-foreground p-6 rounded-lg shadow-md">
        <h3 class="text-xl font-semibold dark:text-primary">Project Title Here</h3>
        <p class="mt-2 text-gray-300">
          Description of your project, tech stack used, features, etc.
        </p>
      </div>

      <!-- Project Card 2 -->
      <div class="bg-foreground p-6 rounded-lg shadow-md">
        <h3 class="text-xl font-semibold dark:text-primary">Another Project</h3>
        <p class="mt-2 text-gray-300">
          Another description here.
        </p>
      </div>

      <!-- Add more cards as needed -->
    </div>
  </div>

</template>
