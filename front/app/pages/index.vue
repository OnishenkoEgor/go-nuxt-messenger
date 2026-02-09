<script setup lang="ts">

import {useHead} from "nuxt/app";

definePageMeta({
  layout: 'default',
  // middleware: ['authenticated'],
});
useHead({
  title: 'Home'
});
</script>

<template>
  <p>main</p>
</template>
