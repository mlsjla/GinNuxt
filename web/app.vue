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

<script setup></script>

<style lang="scss"></style>
