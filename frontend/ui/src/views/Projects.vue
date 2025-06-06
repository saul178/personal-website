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
        <div>
          <h3 class="text-xl font-semibold dark:text-primary">Project Title Here</h3>
        </div>
        <div>
          <p class="mt-2 text-gray-300"> <!-- this need to change for lightmode -->
            She's asked the question so many times that she barely listened to the answers anymore. The answers were
            always the same. Well, not exactly the same, but the same in a general sense. A more accurate description
            was the answers never surprised her. So, she asked for the 10,000th time, "What's your favorite animal?" But
            this time was different. When she heard the young boy's answer, she wondered if she had heard him correctly.
          </p>
        </div>
        <!-- tag rows -->
        <div class="flex flex-row flex-wrap gap-4 mt-4">
          <span class="dark:text-primary px-2 py-1 rounded-full text-sm border-2 border-accent">Tag
            1</span>
          <div class="inline-block px-2 py-1 rounded-full text-sm dark:text-primary border
            border-accent w-fit">Tag 2</div>
          <div>tag1</div>
        </div>
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
