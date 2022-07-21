<template>
    <div class="flex flex-col justify-start items-center w-full min-h-screen">
        <web-header></web-header>
        <div class="flex flex-col w-full md:flex-row flex-grow mt-5 mb-5 main-p">
            <ul class="menu w-full md:w-56 menu-normal menu-horizontal md:menu-vertical mb-5 bg-base-100 p-2 rounded-box">
                <li v-for="(item, index) in categoryList" :key="index">
                    <NuxtLink :class="{'active': category == item.id ? true : false}" :to="`/article/${item.id == category ? '': item.id}`">{{item.name}}</NuxtLink>
                </li>
            </ul>
            <div class="list flex flex-col justify-start items-start overflow-hidden flex-wrap">
                <NuxtLink :to="`/article/view/${topic.id}`" v-for="(topic, index) in treadList" :key="index" class="item flex flex-col md:flex-row md:pl-5 pb-5 overflow-hidden">
                    <div class="flex flex-col justify-center items-center w-full md:w-40 md:mr-5 pb-4 md:pb-0">
                        <el-image :src="config.fileURL + '/' + topic.cover" fit="scale-down"></el-image>
                    </div>
                    <div class="info h-40 flex flex-col justify-between items-start">
                        <h2 class="text-xl font-bold">{{topic.title}}</h2>
                        <p class="text-sm pt-3 pb-2 flex-grow pr-5 md:pr-0 flex-wrap">{{topic.summary}}</p>
                        <div class="info text-sm flex justify-start flex-wrap w-full">
                            <div class="md:mr-10">
                                <span v-if="topic.username" class="mr-5">作者：{{topic.username}}</span>
                                <span>发布：{{topic.created_at.substr(0, 10)}}</span>
                            </div>
                            <div>
                                <span class="mr-5">评论：{{topic.post_count}}</span>
                                <span class="mr-5">打赏：{{topic.rewarded_count}}</span>
                                <span>浏览：{{topic.view_count}}</span>
                            </div>

                        </div>
                    </div>
                </NuxtLink>
            </div>
        </div>
        <web-footer></web-footer>
    </div>
</template>

<script setup>
import api from "@/common/api";
import { useRouter } from "vue-router";
const router = useRouter();
const config = useRuntimeConfig();

const category = $ref(undefined)
const categoryList = $ref([])
const categoryTree = async () => {
    const { data: categoryTree } = await api.pub.categoryTree();
    categoryList = categoryTree.value?.list ? categoryTree.value?.list : []
}
const treadList = $ref([])
const thread = async () => {
    const { data: thread } = await api.pub.thread({
    });
    treadList = thread.value?.list ? thread.value?.list : []
}

await categoryTree()
await thread()
await categoryTree()

</script>

<style lang="scss" scoped></style>
