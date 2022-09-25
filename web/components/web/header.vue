<template>
    <div class="main flex w-full h-20  bg-white/95 z-50 shadow-sm">
        <header class="w-full flex main-p">
            <div class="flex w-full justify-between">
                <div class="flex items-center">
                    <NuxtLink to="/" class="nuxt-link-active"
                        ><div class="flex motion-safe:animate-bounce items-start space-x-3 h-8 md:h-9">
                            <i class="bi bi-cone text-3xl col text-blue-700"></i>
                            <span class="text-3xl">GinNuxt</span>
                            <span
                                class="primary text-xs"
                                ></span
                            >
                        </div>
                    </NuxtLink>
                </div>
                <nav class="nav flex items-center h-full">
                    <NuxtLink to="/" class="item font-medium">首页</NuxtLink>
                    <NuxtLink to="/article" class="item font-medium">全部文章</NuxtLink>
                    <NuxtLink to="/package/miniprogram" class="item font-medium">在线打包</NuxtLink>
                    <NuxtLink to="/" class="item font-medium">关于我们</NuxtLink>
                    <!-- <NuxtLink to="/" class="font-medium">资源</NuxtLink> -->
                </nav>
                <div class="flex flex-row items-center">
                    <div class="text-lg mr-3 cursor-pointer" v-if="userStore?.user?.username">{{userStore?.user?.username}}</div>
                    <!-- <nuxt-link class="text-xl mr-3 cursor-pointer" v-else to="/user/login"><i class="bi bi-person"></i></nuxt-link> -->
                    <nuxt-link class="text-xl mr-3 cursor-pointer" to="/user/edit"><i class="bi bi-gear"></i></nuxt-link>
                    <i v-if="userStore?.token" @click="logout" class="text-xl mr-3 cursor-pointer bi bi-box-arrow-right"></i>
                    <i v-else="userStore?.token" @click="toLogin" class="text-xl mr-3 cursor-pointer bi bi-person"></i>
                </div>
            </div>
        </header>
    </div>
</template>

<script setup>
const userStore = inject("userStore")
const router = useRouter();

const logout = () => {
    
    userStore.reset()
    router.push("/")
}
const toLogin = () => {
    router.push("/user/edit")
}
</script>

<style lang="scss" scoped>
.nav {
    .item {
        @apply px-1 md:px-2 lg:px-3 xl:px-4;
    }
}
</style>