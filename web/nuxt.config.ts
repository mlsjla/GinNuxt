import { defineNuxtConfig } from 'nuxt'


// https://v3.nuxtjs.org/docs/directory-structure/nuxt.config
export default defineNuxtConfig({
    meta: {
        title: 'GinNuxt——开源的全栈程序（预览版）',
        description: '基于Gin+Nuxt3+Vue3开发，支持SSR模式渲染，是覆盖前台后台整套开发解决方案'
    },
    ssr: true,
    buildModules: [
        '@vueuse/nuxt',
        '@pinia/nuxt',
    ],
    app: {
    },
    css: ["~/assets/scss/tailwind.scss", "~/assets/scss/main.scss", "~/assets/scss/index.scss", "bootstrap-icons/font/bootstrap-icons.css"],
    components: true,
    publicRuntimeConfig: {
        baseURL: process.env.baseURL,
        baseServerURL: process.env.baseServerURL,
        fileURL: process.env.fileURL
    },
    experimental: {
        reactivityTransform: true
    },
    vite: {
    },
    build: {
        postcss: {
            postcssOptions: {
                plugins: {
                    tailwindcss: {},
                    autoprefixer: {},
                    ...(process.env.NODE_ENV === 'production' ? { cssnano: {} } : {})
                }
            }
        },
    }
})
