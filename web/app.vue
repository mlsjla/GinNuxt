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

onMounted(() => {
    nextTick(() => {
        useUserStore().init()
    })
})
</script>

<style lang="scss"></style>
