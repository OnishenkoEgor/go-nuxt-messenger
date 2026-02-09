<script setup lang="ts">
import {h, resolveComponent} from 'vue'
import type {User} from "../../types/User";
import type {TableColumn} from '@nuxt/ui'

const button = resolveComponent('UButton');

const {users} = defineProps({
  users: {
    type: Array<User>,
    required: true
  }
});

const emit = defineEmits({
  delete: (id: string) => {
    return true
  }
});
const columns: TableColumn<User>[] = [
  {
    accessorKey: 'id',
    header: 'ID',
  },
  {
    accessorKey: 'login',
    header: 'Login',
  },
  {
    accessorKey: 'password',
    header: 'Password'
  },
  {
    header: 'Actions',
    cell({row}) {
      return h(button, {
        color: 'error',
        onClick: () => {
          console.log(row.original)
          if (row.original.id) {
            console.log('onClick')
            emit('delete', row.original.id)
          }
        },
      }, () => 'Remove')
    },
  }
];

</script>

<template>
  <UTable :data="users" :columns="columns" class="flex-1"/>
</template>

<style scoped>

</style>