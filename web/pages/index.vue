<template>
    <div class="flex flex-col justify-start items-center min-h-screen">
        <web-header></web-header>
        <div class="main main-p w-ful flex flex-col justify-start items-center grow">
            <div class="flex hero-content text-center flex flex-grow">
                <div class="flex flex-col h-full leading-relaxed">
                    <h1 class="text-5xl font-bold pb-5">GinNuxt——开源的全栈WEB服务</h1>
                    <p>
                        采用了稍微新的技术Golang/Gorm/Nuxt3
                    </p>
                    <p>
                        基于Gin+Nuxt3+Vue3开发，支持SSR模式渲染，是覆盖前台后台整套开发解决方案
                    </p>
                    <div class="flex justify-center items-center">
                        <button class="btn btn-primary grow-0 mt-5">关注我们</button>
                    </div>
                </div>
            </div>
        </div>
        <web-footer></web-footer>
    </div>
</template>

<script setup></script>

<style lang="scss" scoped>
section {
    @apply p-20;
}
</style>
