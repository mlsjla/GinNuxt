<template>
    <div class="flex flex-col flex-grow">
        {{ id }}
    </div>
</template>

<script setup>
import { useRouter } from "vue-router";

const router = useRouter();
const { id } = router.currentRoute.value.params;
</script>

<style lang="scss" scoped></style>
