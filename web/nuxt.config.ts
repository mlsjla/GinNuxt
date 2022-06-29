import { defineNuxtConfig } from 'nuxt'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'

const lifecycle = process.env.npm_lifecycle_event

const autoImportOpts = {
    // imports: [
    //     {},
    // ],
    resolvers: [ElementPlusResolver({ ssr: true,})],
}
const vueComponentsOpts = {
    resolvers: [ElementPlusResolver({ ssr: true, directives: false})],
}

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
    modules: [
        '@nuxtjs/tailwindcss',
    ],
    app: {
    },
    css: ["~/assets/scss/main.scss", "~/assets/scss/index.scss", "bootstrap-icons/font/bootstrap-icons.css"],
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
        plugins: [
            
        ],
        build: {
        }
    },
    build: {
        transpile: [
            ...(lifecycle === 'build' || lifecycle === 'generate'
                ? ['element-plus']
                : []),
            'element-plus/es',
        ],
        plugins: [
            AutoImport({
                resolvers: [ElementPlusResolver()],
            }),
            Components({
                resolvers: [ElementPlusResolver()],
            }),
        ]
    }
})
