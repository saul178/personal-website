<script setup>
// no script yet
</script>

<template>
  <!-- Container for both cards -->
  <div class="px-6 py-12 flex flex-col md:flex-row gap-16 justify-center">

    <!-- Profile Card -->
    <div class="bg-muted dark:bg-foreground rounded-2xl shadow-lg p-6 w-full max-w-sm text-center">
      <img src="../assets/images/Default_pfp.svg" alt="Profile Picture"
        class="w-32 h-32 rounded-full mx-auto shadow-md" />
      <div class="mt-6">
        <h1 class="text-2xl font-bold dark:text-primary">About Me</h1>
        <p class="mt-2 dark:text-primary">
          Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean massa. Cum
          sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Donec quam felis, ultricies
          nec, pellentesque eu, pretium quis, sem. Nulla consequat massa quis enim. Donec pede justo, fringilla vel,
          aliquet nec, vulputate eget, arcu. In enim justo, rhoncus ut, imperdiet a, venenatis vitae, justo. Nullam
          dictum felis eu pede mollis pretium. Integer tincidunt. Cras dapibus. Vivamus elementum semper nisi. Aenean
          vulputate eleifend tellus. Aenean leo ligula, porttitor eu, consequat vitae, eleifend ac, enim. Aliquam lorem
          ante, dapibus in, viverra quis, feugiat a, tellus. Phasellus viverra nulla ut metus varius laoreet. Quisque
          rutrum. Aenean imperdiet. Etiam ultricies nisi vel augue. Curabitur ullamcorper ultricies nisi. Nam eget dui.
          Etiam rhoncus. Maecenas tempus, tellus eget condimentum rhoncus, sem quam semper libero, sit amet adipiscing
          sem neque sed ipsum. Nam quam nunc, blandit vel, luctus pulvinar, hendrerit id, lorem. Maecenas nec odio et
          ante tincidunt tempus. Donec vitae sapien ut libero venenatis faucibus. Nullam quis ante. Etiam sit amet orci
          eget eros faucibus tincidunt. Duis leo. Sed fringilla mauris sit amet nibh. Donec sodales sagittis magna. Sed
          consequat, leo eget bibendum sodales, augue velit cursus nunc,
        </p>
      </div>
    </div>

    <!-- Resume Card -->
    <div class="bg-muted dark:bg-foreground rounded-2xl shadow-lg p-6 w-full max-w-xl text-center">
      <!-- This will hold resume info later -->
      <h1 class="text-2xl font-bold dark:text-primary">Resume</h1>
      <p class="mt-2 dark:text-primary">
        Coming soon...
      </p>
    </div>
  </div>
</template>
