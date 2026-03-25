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
const toast = useToast();

let users: Ref<User[]> = ref([]);

async function loadUsers(): Promise<void> {
  UsersApi.getAll().then((loadedUsers: User[]) => {
    console.log(loadedUsers);
    users.value = loadedUsers;
  }).catch(e => {
    toast.add({
      title: 'Failed get users.',
      description: e instanceof Error ? e.message : '',
      progress: false,
      color: 'error'
    });
  });
}

async function removeUser(id: number): Promise<void> {
  try {
    await UsersApi.delById(id);

    console.log('user removed');
    await loadUsers();
  } catch (e: any) {
    toast.add({
      title: 'Failed delete user.',
      description: e instanceof Error ? e.message : '',
      progress: false,
      color: 'error'
    });
  }
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