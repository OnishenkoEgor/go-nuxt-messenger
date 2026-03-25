<script setup lang="ts">
import {type User} from "../../types/User";
import {type Ref} from "vue";
import {UsersApi} from "../../api/UsersApi";
import {navigateTo, useHead} from "nuxt/app";
import UserForm from "../../components/users/UserForm.vue";

useHead({
  title: 'Edit user'
});
const toast = useToast();

const route = useRoute();
const user: Ref<User | null> = ref(null);

if (!route.params.id) {
  throw new Error('id missing');
}

const id: number = +route.params.id;
try {
  UsersApi.getById(id).then(res => {
    user.value = res;
  });
} catch (e: any) {
  toast.add({
    title: 'Failed get user.',
    description: e instanceof Error ? e.message : '',
    progress: false,
    color: 'error'
  });
}

async function onSubmit(user:User):Promise<void>{
  try {
    await UsersApi.update(id, user);

    navigateTo('/users');
  }catch (e: any){
    toast.add({
      title: 'Failed update user.',
      description: e instanceof Error ? e.message : '',
      progress: false,
      color: 'error'
    });
  }
}
</script>

<template>
  <UserForm v-if="user" :user="user" @submit="onSubmit"></UserForm>
  <div v-else>{{'TODO loading'}}</div>
</template>

<style scoped>

</style>