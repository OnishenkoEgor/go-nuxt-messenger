<script setup lang="ts">
import type {NavigationMenuItem} from "@nuxt/ui";
import {useHead, useRouter} from "nuxt/app";
import {type Ref} from "vue";

const items: NavigationMenuItem[] = [
  {
    label: 'Home',
    icon: 'i-lucide-house',
    to: '/'
  },
  {
    label: 'Users',
    icon: 'i-lucide-users',
    to: '/users'
  }, {
    label: 'Roles',
    icon: 'i-lucide-user-cog',
    to: '/roles',
  }, {
    label: 'Permissions',
    icon: 'i-lucide-align-horizontal-justify-center',
    to: '/permissions'
  }
];
const defaultPageTitle: string = 'Default title';
const pageTitle: Ref<string> = ref(defaultPageTitle);
const router = useRouter();
useHead({
  titleTemplate(title) {
    pageTitle.value = title ?? defaultPageTitle;

    return title ?? defaultPageTitle;
  },
});
</script>

<template>
  <UDashboardGroup>
    <UDashboardSidebar open toggle-side="right">
      <UNavigationMenu
          :items="items"
          orientation="vertical"
      />
    </UDashboardSidebar>
    <UDashboardPanel>
      <template #header>
        <UDashboardNavbar :title="pageTitle">
          <template #leading>
            <UButton @click="router.back()"
                     icon="i-lucide-arrow-left"
                     color="neutral"
                     variant="outline">
            </UButton>
          </template>
        </UDashboardNavbar>
      </template>
      <template #body>
        <NuxtPage/>
      </template>
    </UDashboardPanel>
  </UDashboardGroup>
</template>

<style scoped>

</style>