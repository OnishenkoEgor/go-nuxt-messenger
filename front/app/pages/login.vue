<script setup lang="ts">
import {login} from '~/api/auth'
const { loggedIn, user, fetch: refreshSession } = useUserSession()

const credentials = reactive({
  email: 'test@test.test',
  password: 'test',
})

async function submit() {
  login(credentials.email, credentials.password).then(data => {
    console.log(data)
  }).catch(error => console.log(error));
}
</script>

<template>
  <form @submit.prevent="submit">
    <input
        v-model="credentials.email"
        type="email"
        placeholder="Email"
    >
    <input
        v-model="credentials.password"
        type="password"
        placeholder="Password"
    >
    <button type="submit">
      Login
    </button>
  </form>
</template>