<script setup lang="ts">
import * as v from "valibot";
import {useHead} from "nuxt/app";
import {reactive} from "vue";
import {type User} from "../../types/User";
import type { FormSubmitEvent } from '@nuxt/ui'
import {UsersApi} from "../../api/UsersApi";

definePageMeta({
  layout: 'default'
});
useHead({
  title: 'Create user'
});

const schema = v.object({
  login: v.pipe(v.string()),
  password: v.pipe(v.string(), v.minLength(1, 'Must be at least 8 characters'))
})
type Schema = v.InferOutput<typeof schema>

const user: User = reactive<User>({
  login: 'test',
  password: '123'
});

async function onSubmit(event: FormSubmitEvent<Schema>) {
  UsersApi.create(user).then(res=>{
    console.log('component res')
    console.log(res);
  })
}
</script>

<template>
  <UForm :schema="schema" :state="user" class="space-y-4" @submit="onSubmit">
    <UFormField label="Login" name="login">
      <UInput v-model="user.login"/>
    </UFormField>
    <UFormField label="Password" name="password">
      <UInput v-model="user.password" type="password" />
    </UFormField>
    <UButton type="submit">
      Create
    </UButton>
  </UForm>
</template>

<style scoped>

</style>