<template>
    <div class="flex flex-col justify-start items-center min-h-screen">
        <web-header></web-header>
        <div class="main w-full flex flex-col justify-start items-center grow">
            <div class="flex-auto w-full min-w-0 lg:static lg:max-h-full lg:overflow-visible">
                <div>
                    <div class="overflow-hidden items-center">
                        <div class="flex flex-wrap justify-center py-0 section">
                            <section
                                class="flex min-h-screen flex-col justify-start items-center w-full px-4 pb-24 sm:pb-36 lg:pb-50 text-center z-20">
                                <div class="flex flex-col justify-start items-center h-60 w-60">
                                    <img class="w-60 h-60" src="/images/home.svg" alt="" />
                                </div>
                                <h1
                                    class="font-normal font-serif text-display-5 xs:text-display-4 md:text-display-3 2xl:text-display-2 mb-6 pt-2">
                                    实现WEB、小程序、APP互通
                                </h1>
                                <h2
                                    class="font-normal text-body-base xs:text-body-lg md:text-body-xl 2xl:text-body-2xl mb-2 px-8 sm:px-0 dark:text-cloud-lighter text-cloud-darker">
                                    七维全平台引擎系统的目的是将本产品打造为一个数据服务中心，通过接口打通WEB、各类小程序、APP...
                                </h2>
                                <p class="font-normal text-body-base xs:text-body-lg md:text-body-xl 2xl:text-body-2xl mb-8 px-8 sm:px-0 dark:text-cloud-lighter text-cloud-darker">无论未来互联网怎么发展，七维全平台引擎系统都将是企业互联网服务坚实的基石。</p>
                                <div
                                    class="flex flex-col sm:flex-row flex-wrap items-center justify-center space-y-3 sm:space-y-0 sm:space-x-3 xl:space-x-4">
                                    <nuxtLink class="cursor-pointer" to=""><i class="bi bi-arrow-right-circle pr-2"></i>联系我们</nuxtLink>
                                </div>
                            </section>
                        </div>
                    </div>
                    <div class="relative dark:bg-secondary-darkest bg-gray-50">
                        <section class="py-24 md:py-40 light:text-secondary-black">
                            <div class="relative w-full d-container-content flex flex-col items-center">
                                <div class="flex flex-col w-full items-center col-span-12">
                                    <h2
                                        class="font-normal text-center font-serif text-display-6 md:text-display-5 2xl:text-display-4 mb-2">
                                        产品特性
                                    </h2>
                                    <p
                                        class="font-normal text-center text-body-base md:text-body-lg 2xl:text-body-xl mb-12">
                                    </p>
                                    <div class="grid grid-cols-1 xs:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
                                        <div to="" aria-label="Lighter"
                                            class="flex flex-col items-center transition duration-200 p-6 rounded-md">
                                            <svg xmlns="http://www.w3.org/2000/svg"
                                                xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true" role="img"
                                                width="32" height="32" preserveAspectRatio="xMidYMid meet"
                                                viewBox="0 0 20 20"
                                                class="iconify iconify--entypo mb-4 transition duration-200 text-primary group-hover:text-primary-400 w-10 h-10">
                                                <path
                                                    d="M4.254 19.567c.307-.982.77-2.364 1.391-4.362c2.707-.429 3.827.341 5.546-2.729c-1.395.427-3.077-.792-2.987-1.321c.091-.528 3.913.381 6.416-3.173c-3.155.696-4.164-.836-3.757-1.067c.939-.534 3.726-.222 5.212-1.669c.766-.745 1.125-2.556.813-3.202c-.374-.781-2.656-1.946-3.914-1.836c-1.258.109-3.231 4.79-3.817 4.754c-.584-.037-.703-2.098.319-4.013c-1.077.477-3.051 1.959-3.67 3.226c-1.153 2.357.108 7.766-.296 7.958c-.405.193-1.766-2.481-2.172-3.694c-.555 1.859-.568 3.721 1.053 6.194c-.611 1.623-.945 3.491-.996 4.441c-.024.759.724.922.859.493z"
                                                    fill="currentColor"></path>
                                            </svg>
                                            <h3
                                                class="relative inline-flex mb-1 text-center text-body-lg items-center lg:text-body-xl font-bold">
                                                轻量
                                                <!---->
                                            </h3>
                                            <p class="text-center text-sm lg:text-base mb-4">
                                                前后端分离，业务全接口化，核心部分基于gin+nuxt3开发
                                            </p>
                                        </div>
                                        <div to="" aria-label="Faster"
                                            class="flex flex-col items-center transition duration-200 p-6 rounded-md">
                                            <svg xmlns="http://www.w3.org/2000/svg"
                                                xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true" role="img"
                                                width="32" height="32" preserveAspectRatio="xMidYMid meet"
                                                viewBox="0 0 28 28"
                                                class="iconify iconify--fluent mb-4 transition duration-200 text-primary group-hover:text-primary-400 w-10 h-10">
                                                <g fill="none">
                                                    <path
                                                        d="M18.184 5.55l6.555 6.554a4.305 4.305 0 0 1-4.038 7.234v.246A2.415 2.415 0 0 1 18.286 22H15.5v-.826A3.173 3.173 0 0 0 12.335 18h-1.59a.75.75 0 0 0 0 1.5h1.59c.917 0 1.666.75 1.666 1.674V22H7.399A2.401 2.401 0 0 1 5 19.584v-3.15c0-.233.015-.463.043-.69A2.875 2.875 0 1 1 7.55 11.82c.831-.519 1.814-.82 2.868-.82h4.849c.729 0 1.424.144 2.06.404c.197-.332.438-.615.72-.905l-2.406-2.406a1.798 1.798 0 1 1 2.543-2.543z"
                                                        fill="currentColor"></path>
                                                </g>
                                            </svg>
                                            <h3
                                                class="relative inline-flex mb-1 text-center text-body-lg items-center lg:text-body-xl font-bold">
                                                快
                                                <!---->
                                            </h3>
                                            <p class="text-center text-sm lg:text-base mb-4">
                                                基于golang语言开发，不只是运行快，开发效率也高
                                            </p>
                                        </div>
                                        <div to="" aria-label="Vite"
                                            class="flex flex-col items-center transition duration-200 p-6 rounded-md">
                                            <svg width="40" height="40" viewBox="0 0 40 40" fill="none"
                                                xmlns="http://www.w3.org/2000/svg"
                                                class="mb-4 transition duration-200 text-primary group-hover:text-primary-400 w-10 h-10">
                                                <path fill-rule="evenodd" clip-rule="evenodd"
                                                    d="M15.29 8.384l-.547 9.368a.436.436 0 00.53.455l3.29-.77a.437.437 0 01.522.516l-.977 4.856a.436.436 0 00.55.51l2.032-.627a.436.436 0 01.551.51l-1.553 7.628c-.097.478.529.738.79.329l.173-.274 9.627-19.495c.16-.327-.117-.7-.47-.63l-3.386.663a.438.438 0 01-.499-.554l.799-2.81 9.253-1.712c.721-.133 1.27.646.912 1.296l-15.963 28.97a.861.861 0 01-1.51.006L3.133 7.646c-.365-.648.182-1.434.906-1.303l11.25 2.04z"
                                                    fill="currentColor"></path>
                                            </svg>
                                            <h3
                                                class="relative inline-flex mb-1 text-center text-body-lg items-center lg:text-body-xl font-bold">
                                                Vite
                                                <!---->
                                            </h3>
                                            <p class="text-center text-sm lg:text-base mb-4">体验快的开发人生</p>
                                        </div>
                                        <div to="" aria-label="Vue 3"
                                            class="flex flex-col items-center transition duration-200 p-6 rounded-md">
                                            <svg width="26" height="26" viewBox="0 0 26 26" fill="none"
                                                xmlns="http://www.w3.org/2000/svg"
                                                class="mb-4 transition duration-200 text-primary group-hover:text-primary-400 w-10 h-10">
                                                <path
                                                    d="M0 2L13 24L26 2H16L13 7.3L10 2H0ZM3.5 4H7.271L13 14L18.729 4H22.5L13 20.1L3.5 4Z"
                                                    fill="currentColor"></path>
                                            </svg>
                                            <h3
                                                class="relative inline-flex mb-1 text-center text-body-lg items-center lg:text-body-xl font-bold">
                                                Vue 3
                                                <!---->
                                            </h3>
                                            <p class="text-center text-sm lg:text-base mb-4">前端基于Vue3</p>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </section>
                    </div>
                    <section class="py-20 text-lg">
                        <div class="relative w-full d-container-content text-justify">
                            <h1
                                class="font-normal font-serif text-display-5 xs:text-display-4 md:text-display-3 2xl:text-display-2 my-4">
                                打通微信、字节、百度，打破壁垒
                            </h1>
                            <p>
                                连通微信公众号、微信开放平台、微信小程序，实现互联互通。同步支持抖音小程序、百度小程序，获取流量渠道更多更广。
                            </p>
                        </div>
                    </section>
                    <section class="dark:bg-secondary-darkest bg-gray-50 py-20 text-justify">
                        <div class="relative w-full d-container-content text-justify">
                            <h1
                                class="font-normal font-serif text-display-5 xs:text-display-4 md:text-display-3 2xl:text-display-2 my-4">
                                采用的技术
                            </h1>
                            <ol class="list-decimal pl-5">
                                <li>Golang Gin <nuxtLink to="https://github.com/LyricTian/gin-admin">gin-admin</nuxtLink></li>
                                <li>Nuxt3 Vue3</li>
                                <li>uniapp</li>
                                <li>Elementui</li>
                                <li>tailwindcss</li>
                            </ol>
                            
                        </div>
                    </section>
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
