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
  <div class="flex flex-col md:flex-row px-8 py-12 gap-8">
    <!-- Left Sidebar -->
    <div class="flex flex-col space-y-4 gap-2">
      <h2 class="text-4xl font-bold dark:text-primary">Projects</h2>
      <div class="flex justify-start gap-2.5">
        <a href="https://github.com/saul178" class="dark:text-primary text-2xl font-bold">
          <font-awesome-icon :icon="['fab', 'github']" class="text-4xl dark:text-primary" />
        </a>
        <a href="https://github.com/saul178" class="dark:text-primary text-2xl font-bold">
          View My Github →
        </a>
      </div>
    </div>

    <!-- Right Content Area -->
    <div class="md:flex-1 flex-col space-y-6">
      <!-- project card will be a component to be reused -->
      <div
        class="flex flex-col md:flex-row group hover:-translate-y-1 bg-foreground hover:bg-foreground/60 backdrop-blur-sm p-6 rounded-2xl border border-accent/20 transition-all duration-300 shadow-lg hover:shadow-accent/20">
        <img class="md:w-80 h-48 mb-4 md:mr-4 rounded-md shadow-lg object-cover"
          src="../assets/images/snapshot_2025-06-10_13-49-07.png" alt="Project 1">

        <div class="flex flex-col justify-between">
          <div>
            <h3 class="text-xl font-semibold dark:text-primary">Project Title Here Should be a link
              to project</h3>
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
              <span class="dark:text-primary px-2 py-1 rounded-full text-sm border-2
                border-accent/50">Github</span>
              <span class="dark:text-primary px-2 py-1 rounded-full text-sm border-2
                border-accent/50">api</span>
              <span class="dark:text-primary px-2 py-1 rounded-full text-sm border-2
                border-accent/50">here</span>
            </div>

            <div class="md:flex flex-row hidden mt-6 border-t-2 dark:border-background">
              <h3 class="dark:text-secondary font-semibold mt-2">Commit History Example</h3>
            </div>
            <div class="mt-2">
              <span class="dark:text-secondary block">saul178: committed last week -- "wip: i think i now know how to
                lay this out."</span>
              <span class="dark:text-secondary block">github api will be used here!</span>
              <span class="dark:text-secondary block">Latest commit here!</span>
            </div>
          </div>
        </div>
      </div>
      <!-- project card component end -->

      <!-- other projects -->
    </div>
  </div>

  <div class="flex flex-col px-8 py-12 gap-8">
    <h1 class="text-3xl font-semibold dark:text-primary">My Github Activity</h1>
    <div class="flex-1 space-y-6">
      <div
        class="group hover:-translate-y-1 bg-foreground hover:bg-foreground/60 backdrop-blur-sm p-6 rounded-2xl border border-accent/20 transition-all duration-300 shadow-lg hover:shadow-accent/20">
        <a href="https://www.github.com/saul178">
          <img class="w-full" src="https://ghchart.rshah.org/072460/saul178.svg" alt="Github Contributations Chart">
        </a>
      </div>
    </div>
  </div>
</template>
