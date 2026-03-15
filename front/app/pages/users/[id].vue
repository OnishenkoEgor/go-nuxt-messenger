<script setup lang="ts">
import {type User} from "../../types/User";
import {type Ref} from "vue";
import {UsersApi} from "../../api/UsersApi";
import {useHead} from "nuxt/app";

useHead({
  title: 'Edit user'
})
const route = useRoute()
const user: Ref<User | null> = ref(null);
if (!route.params.id) {
  throw new Error('id missing');
}
const id: number = +route.params.id;
UsersApi.get(id).then(res)
</script>

<template>

</template>

<style scoped>

</style>