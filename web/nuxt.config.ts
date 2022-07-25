import { defineNuxtConfig } from 'nuxt'
import { ElementPlusResolver } from '@mlsjla/unplugin-vue-components/resolvers'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'

const lifecycle = process.env.npm_lifecycle_event

const autoImportOpts = {
    dirs: ['components', 'pages', 'layouts'],
    resolvers: [ElementPlusResolver({ ssr: true, nuxt: true, importStyle: 'sass'})],
    // Allow for components to override other components w})],
    dts: true // or a custom path
}
const vueComponentsOpts = {
    dirs: ['components', 'pages', 'layouts'],
    directoryAsNamespace: true,
    allowOverrides: false,
    resolvers: [ElementPlusResolver({ ssr: true, nuxt: true, importStyle: 'sass'})],
    dts: true // or a custom path
}

var config = {
    baseURL: process.env.baseURL,
    baseServerURL: process.env.baseServerURL,
    fileURL: process.env.fileURL
}
if(process.env.runMode == 'dev') {
    config = {
        baseURL: process.env.baseDevURL,
        baseServerURL: process.env.baseDevServerURL,
        fileURL: process.env.fileDevURL
    }
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
        '@formkit/nuxt',
    ],
    app: {
    },
    css: ["~/assets/scss/main.scss", "~/assets/scss/index.scss", "bootstrap-icons/font/bootstrap-icons.css"],
    components: true,
    publicRuntimeConfig: {
        ...config
    },
    experimental: {
        reactivityTransform: true
    },
    vite: {
        plugins: [
            AutoImport(autoImportOpts),
            Components(vueComponentsOpts),
        ]
    },
    build: {
        transpile: [
            ...(lifecycle === 'build' || lifecycle === 'generate'
                ? ['element-plus']
                : []),
            'element-plus/es',
        ],
    }
})
