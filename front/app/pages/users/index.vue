<script setup lang="ts">
import {useHead} from "nuxt/app";
import {UsersApi} from "../../api/UsersApi";
import {type User} from "../../types/User";
import {type Ref} from "vue";

definePageMeta({
  layout: 'default',
});
useHead({
  title: 'Users'
});

let users: Ref<User[]> = ref([]);

async function loadUsers(): Promise<void> {
  UsersApi.getAll().then((loadedUsers: User[]) => {
    users.value = loadedUsers;
    console.log(users.value);
  });
}

async function removeUser(id: number): Promise<void>{
  console.log(`remove user ${id}`);
  const deleted = await UsersApi.remove(id.toString());
}

loadUsers();
</script>

<template>
  <div class="">
    <UButton to="/users/create" curs>Create</UButton>
    <UsersList :users="users" @delete="removeUser"/>
    <NuxtPage/>
  </div>
</template>

<style scoped>

</style>