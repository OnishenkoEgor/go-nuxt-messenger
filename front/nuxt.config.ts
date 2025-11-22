import type {ServerOptions} from 'vite';

export default defineNuxtConfig({
  compatibilityDate: '2025-07-15',
  devtools: {enabled: true},

  typescript: {
      typeCheck: true
  },

  vite: {
      server: ({
          host: true,
          port: 3000,
          watch: {
              usePolling: true,
          },
      }) as ServerOptions,
  },

  modules: [
      'nuxt-auth-utils',
      '@nuxt/ui'
  ]
})