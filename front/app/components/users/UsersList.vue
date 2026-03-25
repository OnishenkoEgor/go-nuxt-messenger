<script setup lang="ts">

import type {User} from "../../types/User";
import type {TableColumn} from '@nuxt/ui'

const {users} = defineProps({
  users: {
    type: Array<User>,
    required: true
  }
});

const emit = defineEmits<{
  (e: 'delete', id: number): void
}>();

function del(id: number): void {
  emit('delete', id);
}

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
    id: 'actions',
    header: 'Actions',
  }
];

</script>

<template>
  <UTable :data="users" :columns="columns" class="flex-1">
    <template #actions-cell="{row}">
      <UFieldGroup>
        <UButton v-if="row.original.id !== undefined"
                 :to="`/users/${row.original.id}`">
          Edit
        </UButton>
        <UButton v-if="row.original.id !== undefined"
                 @click="del(row.original.id)"
                 color="error">
          Delete
        </UButton>
      </UFieldGroup>
    </template>
  </UTable>
</template>

<style scoped>

</style>