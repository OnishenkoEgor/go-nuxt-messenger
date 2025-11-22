<script setup lang="ts">
const { loggedIn, user, fetch: refreshSession } = useUserSession()
const credentials = reactive({
  email: '',
  password: '',
})
</script>

<template>

</template>

<style scoped>

</style>