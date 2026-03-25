<script setup lang="ts">
import {useHead} from "nuxt/app";
import {type Ref} from "vue";
import {type User} from "../../types/User";
import UserForm from "../../components/users/UserForm.vue";
import {UsersApi} from "../../api/UsersApi";

definePageMeta({
  layout: 'default'
});
useHead({
  title: 'Create user'
});
const toast = useToast();

const user: Ref<User> = ref<User>({
  login: '',
  password: ''
});

async function onSubmit(user: User): Promise<void> {
  try {
    await UsersApi.create(user);

    navigateTo('/users');
  } catch (e: any) {
    toast.add({
      title: 'Failed create user.',
      description: e instanceof Error ? e.message : '',
      progress: false,
      color: 'error'
    });
  }
}
</script>

<template>
  <UserForm :user="user" @submit="onSubmit"></UserForm>
</template>

<style scoped>

</style>