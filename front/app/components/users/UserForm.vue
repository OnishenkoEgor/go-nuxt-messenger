<script setup lang="ts">
import {type User} from "../../types/User";
import {type PropType} from "vue";

const {user} = defineProps({
  user: {
    type: Object as PropType<User>,
    required: false
  },
});
console.log(user)
</script>

<template>

</template>

<style scoped>

</style>