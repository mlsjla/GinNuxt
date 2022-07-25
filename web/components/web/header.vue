<template>
    <div class="flex w-full h-20  bg-white/95 z-50 shadow-sm">
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
                <nav class="flex items-center h-full">
                    <NuxtLink to="/" class="font-medium px-2 py-1">首页</NuxtLink>
                    <NuxtLink to="/article" class="font-medium px-2 py-1">分类</NuxtLink>
                    <!-- <NuxtLink to="/" class="font-medium px-2 py-1">资源</NuxtLink> -->
                    <NuxtLink to="/" class="font-medium px-2 py-1">关于</NuxtLink>
                </nav>
                <div class="flex flex-row items-center">
                    <div class="text-lg mr-3 cursor-pointer" v-if="userStore?.user?.username">{{userStore?.user?.username}}</div>
                    <!-- <nuxt-link class="text-xl mr-3 cursor-pointer" v-else to="/user/login"><i class="bi bi-person"></i></nuxt-link> -->
                    <nuxt-link class="text-xl mr-3 cursor-pointer" to="/user/edit"><i class="bi bi-gear"></i></nuxt-link>
                    <i class="text-xl mr-3 cursor-pointer bi bi-box-arrow-right"></i>
                </div>
            </div>
        </header>
    </div>
</template>

<script setup>
import { useRouter } from "vue-router";
import { useUserStore } from '@/stores/user'
import { useCategoryStore } from "@/stores/category";

const router = useRouter();
const userStore = useUserStore();
useCategoryStore().init()
</script>

<style lang="scss" scoped>

</style>