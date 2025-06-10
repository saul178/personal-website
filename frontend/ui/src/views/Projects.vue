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
      <div class="flex group hover:-translate-y-1 bg-foreground backdrop-blur-sm p-6 rounded-2xl border
        border-accent/20 transition-all duration-300 shadow-lg hover:shadow-accent/20">
        <img class="h-48 w-80 mr-4 rounded-md shadow-lg object-cover"
          src="../assets/images/snapshot_2025-06-10_13-49-07.png" alt="Project 1">

        <div class="flex flex-col justify-between">
          <div>
            <h3 class="text-xl font-semibold dark:text-primary">Project Title Here</h3>
            <div>
              <p class="mt-2 dark:text-secondary"> <!-- this need to change for lightmode -->
                PROJECT DESCRIPTION HERE She's asked the question so many times that she barely listened to the answers
                anymore. The answers were always the same. Well, not exactly the same, but the same in a general sense.
                A more accurate
                description was the answers never surprised her. So, she asked for the 10,000th time, "What's your
                favorite animal?"
                But this time was different. When she heard the young boy's answer, she wondered if she had heard him
                correctly.
              </p>
            </div>

            <!-- tag rows (this will probably be a vue component) -->
            <div class="flex flex-row flex-wrap gap-4 mt-4">
              <span class="dark:text-primary px-2 py-1 rounded-full text-sm border-2 border-accent">Tag 1</span>
              <span class="dark:text-primary px-2 py-1 rounded-full text-sm border-2 border-accent">Tag 2</span>
              <span class="dark:text-primary px-2 py-1 rounded-full text-sm border-2 border-accent">Tag 3</span>
            </div>

            <div class="md:flex flex-row hidden mt-6 border-t-2 dark:border-background">
              <h3 class="dark:text-secondary font-semibold mt-2">Commit History Example</h3>
            </div>
            <div class="mt-2">
              <span class="dark:text-secondary block">Latest commit here!</span>
              <span class="dark:text-secondary block">Latest commit here!</span>
              <span class="dark:text-secondary block">Latest commit here!</span>
            </div>
          </div>
        </div>
      </div>

      <!-- Project Card 2 -->
      <div class="group hover:-translate-y-1 bg-foreground backdrop-blur-sm p-6 rounded-2xl border
        border-accent/20 transition-all duration-300 shadow-lg hover:shadow-accent/20">
        <div>
          <h3 class="text-xl font-semibold dark:text-primary">Project 2 Title Here</h3>
        </div>
        <div>
          <p class="mt-2 dark:text-secondary"> <!-- this need to change for lightmode -->
            Another PROJECT DESCRIPTION HERE She's asked the question so many times that she barely listened to the
            answers
            anymore. The answers were
            always the same. Well, not exactly the same, but the same in a general sense. A more accurate description
            was the answers never surprised her. So, she asked for the 10,000th time, "What's your favorite animal?" But
            this time was different. When she heard the young boy's answer, she wondered if she had heard him correctly.
          </p>
        </div>
        <!-- tag rows (this will probably be a vue component) -->
        <div class="flex flex-row flex-wrap gap-4 mt-4">
          <span class="dark:text-primary px-2 py-1 rounded-full text-sm border-2 border-accent">Tag 1</span>
          <span class="dark:text-primary px-2 py-1 rounded-full text-sm border-2 border-accent">Tag 2</span>
          <span class="dark:text-primary px-2 py-1 rounded-full text-sm border-2 border-accent">Tag 3</span>
        </div>

        <div class="md:flex flex-row hidden mt-6 border-t-2 dark:border-background">
          <h3 class="dark:text-secondary font-semibold mt-2">Commit History Example</h3>
        </div>
        <div class="mt-4">
          <span class="dark:text-secondary">Latest commit here!</span>
        </div>
      </div>

      <!-- Project Card 3 -->
      <div class="group hover:-translate-y-1 bg-foreground backdrop-blur-sm p-6 rounded-2xl border
        border-accent/20 transition-all duration-300 shadow-lg hover:shadow-accent/20">
        <div>
          <h3 class="text-xl font-semibold dark:text-primary">Project 3 Title Here</h3>
        </div>
        <div>
          <p class="mt-2 dark:text-secondary"> <!-- this need to change for lightmode -->
            PROJECT DESCRIPTION HERE She's asked the question so many times that she barely listened to the answers
            anymore. The answers were
            always the same. Well, not exactly the same, but the same in a general sense. A more accurate description
            was the answers never surprised her. So, she asked for the 10,000th time, "What's your favorite animal?" But
            this time was different. When she heard the young boy's answer, she wondered if she had heard him correctly.
          </p>
        </div>
        <!-- tag rows (this will probably be a vue component) -->
        <div class="flex flex-row flex-wrap gap-4 mt-4">
          <span class="dark:text-primary px-2 py-1 rounded-full text-sm border-2 border-accent">Tag 1</span>
          <span class="dark:text-primary px-2 py-1 rounded-full text-sm border-2 border-accent">Tag 2</span>
          <span class="dark:text-primary px-2 py-1 rounded-full text-sm border-2 border-accent">Tag 3</span>
        </div>

        <div class="md:flex flex-row hidden mt-6 border-t-2 dark:border-background">
          <h3 class="dark:text-secondary font-semibold mt-2">Commit History Example</h3>
        </div>
        <div class="mt-4">
          <span class="dark:text-secondary">Latest commit here!</span>
        </div>
      </div>
    </div>
  </div>

  <div class="flex px-8 py-12 gap-8">
    <div class="w-1/4 flex flex-col space-y-4">
      <h1 class="dark:text-primary font-bold text-3xl">My Github Activity</h1>
    </div>

    <div class="flex-1 space-y-6">
      <div class="group hover:-translate-y-1 bg-foreground backdrop-blur-sm p-6 rounded-2xl border
          border-accent/20 transition-all duration-300 shadow-lg hover:shadow-accent/20">
        <a href="https://www.github.com/saul178">
          <img class="w-full" src="https://ghchart.rshah.org/072460/saul178.svg" alt="Github Contributations Chart">
        </a>
      </div>
    </div>
  </div>
</template>
