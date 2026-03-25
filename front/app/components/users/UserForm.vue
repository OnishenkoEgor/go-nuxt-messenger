<script setup lang="ts">
import {type User} from "../../types/User";
import * as v from "valibot";
import type {FormSubmitEvent} from "@nuxt/ui";
import {computed} from "vue";

const {user: userProp} = defineProps<{
  user: User
}>();

const user = ref(userProp);
const isNew = computed(() => !user.value.hasOwnProperty('id'));

const emit = defineEmits<{
  (e: 'submit', value: User): void
}>()

const schema = v.object({
  login: v.pipe(v.string()),
  password: v.pipe(v.string(), v.minLength(1, 'Must be at least 8 characters'))
});

type Schema = v.InferOutput<typeof schema>

async function onSubmit(_: FormSubmitEvent<Schema>) {
  if (user.value !== null) {
    emit('submit', user.value);
  }
}
</script>

<template>
  <UForm :schema="schema" :state="user" class="space-y-4" @submit="onSubmit">
    <UFormField label="Login" name="login">
      <UInput v-model="user.login"/>
    </UFormField>
    <UFormField label="Password" name="password">
      <UInput v-model="user.password" type="password"/>
    </UFormField>
    <UButton type="submit">
      {{ isNew ? 'Create' : 'Update' }}
    </UButton>
  </UForm>
</template>

<style scoped>

</style>