// import type {ServerOptions} from 'vite';

export default defineNuxtConfig({
    ssr: false,

    compatibilityDate: '2025-07-15',
    devtools: {enabled: true},

    typescript: {
        typeCheck: true
    },

    // vite: {
    //     server: ({
    //         host: true,
    //         port: 3000,
    //         watch: {
    //             usePolling: true,
    //         },
    //     }) as ServerOptions,
    // },

    modules: [
        'nuxt-auth-utils',
        '@nuxt/ui'
    ],
    css: ['~/assets/css/main.css']
})