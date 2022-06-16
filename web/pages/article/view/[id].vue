<template>
    <div class="flex flex-col justify-start items-center relative min-h-screen">
        <web-header></web-header>
        <div class="main flex flex-col flex-start justify-start items-center flex-grow">
            <div class="flex flex-col">
                <div class="hero-content flex-col" v-if="thread">
                    <h2 class="text-3xl font-bold w-full text-left mb-5">{{ thread.title }}</h2>
                    <div class="info flex flex-row justify-between items-center w-full cursor-pointer">
                        <div class="flex flex-row justify-start items-center">
                            <div class="w-12 rounded-full mr-3">
                                <img class="rounded-full" src="http://127.0.0.1:10088/uploads/file/2382c315a6ba396be1932dd8dafaff52_20220523204927.png" />
                            </div>
                            <div class="user flex flex-col justify-center items-start text-sm">
                                <span class="flex">{{thread.username}}</span>
                                <div class="flex">发布：{{thread?.created_at?.substr(0, 10)}}</div>
                            </div>
                        </div>
                        <button class="btn btn-sm">
                            <i class="bi bi-plus-lg text-lg"></i>
                            关注
                        </button>
                    </div>
                    <div v-if="isBot" class="w-full vditor-content flex-col leading-relaxed" v-html="thread.content"></div>
                    <ClientOnly v-else-if="thread.content">
                        <LazyVditorPreview v-model:content="thread.content"></LazyVditorPreview>
                    </ClientOnly>
                </div>
            </div>
        </div>
        <web-footer></web-footer>
    </div>
</template>

<script setup>
import { useRouter } from "vue-router";
import api from "@/common/api";
import MarkdownIt from 'markdown-it'
import isbot from 'isbot'
import { LazyVditorPreview } from '#components'


const config = useRuntimeConfig();
const headers = useRequestHeaders(['user-agent'])

const router = useRouter();
const { id } = router.currentRoute.value.params;
const isBot = useState('isBot', () => false)
const thread = useState('thread', () => {})

const postList = $ref([]);

const fetchThreads = async () => {
    if (!id) {
        console.error("id error", id)
        return;
    }
    const [{ data: res }, { data: resList }] = await Promise.all([
        api.thread.one(id),
        api.post.get({
            threadId: id,
        }),
    ]);
    thread.value = res.value = res.value ? res.value : {}
    // 格式化 cover
    thread.value.cover = thread.value.cover ? config.fileURL + '/' + thread.value.cover : undefined

    if (resList?.value?.list && resList?.value?.list[0]) {
        if(isBot.value) {
            const md = new MarkdownIt()
            thread.value.content = md.render(resList?.value?.list[0].content)
        }else {
            thread.value.content = resList?.value?.list[0].content
        }
    }
};

if(process.server) {
    if(isbot(headers['user-agent'])) {
        isBot.value = true
    }
    await fetchThreads()
}


onMounted(() => {
    if(!isBot.value) {
        nextTick(async () => {
            // 进行md渲染
            await fetchThreads()
            console.log("md render...", thread.value.content)
        })
    }
})

</script>

<style lang="scss" scoped>
.main {
    max-width: 690px;
}

:deep(.vditor-content) {
    h1 {
        @apply text-6xl leading-relaxed;
    }
    h2 {
        @apply text-4xl leading-relaxed;
    }
    h3 {
        @apply text-3xl leading-relaxed;
    }
    h4 {
        @apply text-2xl leading-relaxed;
    }
    h5 {
        @apply text-xl leading-relaxed;
    }
    h6 {
        @apply text-lg leading-relaxed;
    }
    p,div,span,li {
        @apply leading-relaxed;
    }
}


</style>
