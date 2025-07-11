<template>
  <footer
    class="flex items-center justify-between w-full px-6 py-3 mt-auto bg-gradient-to-b dark:from-background to-foreground text-primary shadow-lg rounded-t-lg">
    <div class="text-xs md:text-md font-medium text-primary/80 mt-1">v1.0</div>

    <div
      class="flex items-center gap-2 mt-1 text-sm md:text-md font-medium bg-gradient-to-r dark:from-secondary to-accent bg-clip-text text-transparent">
      <span>Built with</span>
      <font-awesome-icon :icon="['fab', 'golang']" class="text-lg md:text-2xl text-primary/80" />
      <span class="text-lg md:text-2xl">+</span>
      <font-awesome-icon :icon="['fab', 'vuejs']" class="text-lg md:text-2xl text-primary/80" />
      <span class="text-lg md:text-2xl">+</span>
      <span>Saul Gonzalez</span>
    </div>

    <div class="flex items-center gap-4 mt-2">
      <a href="https://www.linkedin.com/in/saul0178">
        <font-awesome-icon :icon="['fab', 'linkedin']"
          class="text-md md:text-2xl md:hover:text-primary/60 transition" />
      </a>
      <a href="https://github.com/saul178">
        <font-awesome-icon :icon="['fab', 'github']" class="text-md md:text-2xl md:hover:text-primary/60 transition" />
      </a>
    </div>
  </footer>
</template>
