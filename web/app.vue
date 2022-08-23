<template>
    <NuxtLayout>
        <NuxtErrorBoundary>
            <NuxtPage />
            <template #error="{ error }">
                <p>An error occurred: {{ error }}</p>
            </template>
        </NuxtErrorBoundary>
    </NuxtLayout>
</template>

<script setup>
import { useUserStore } from '@/stores/user'
import { useCategoryStore } from "@/stores/category";


const userStore = useUserStore()
provide("userStore", userStore)

const categoryStore = useCategoryStore()
provide("categoryStore", categoryStore)

onMounted(() => {
    nextTick(() => {
        userStore.init()
        categoryStore.init()
    })
})
</script>

<style lang="scss"></style>
