<template>
    <div class="flex flex-col justify-start items-center relative min-h-screen">
        <web-header></web-header>
        <div v-if="thread" class="w-full main flex flex-col flex-start justify-start items-center flex-grow">
            <h2 class="w-full text-3xl font-bold text-left mb-5">{{ thread.title }}</h2>
            <div class="w-full flex flex-row justify-between items-center cursor-pointer">
                <div class="flex flex-row justify-start items-center">
                    <div class="w-12 rounded-full mr-3">
                        <!-- <img class="rounded-full" src="http://127.0.0.1:10088/uploads/file/2382c315a6ba396be1932dd8dafaff52_20220523204927.png" /> -->
                    </div>
                    <div class="user flex flex-col justify-center items-start text-sm">
                        <span class="flex">{{ thread.username }}</span>
                        <div class="flex">发布：{{ thread?.created_at?.substr(0, 10) }}</div>
                    </div>
                </div>
                <button class="btn btn-sm">
                    <i class="bi bi-plus-lg text-lg"></i>
                    关注
                </button>
            </div>
            <div v-if="isBot" class="w-full prose lg:prose-xl flex-col leading-relaxed" v-html="thread.content"></div>
            <ClientOnly v-else-if="thread.content">
                <LazyVditorPreview v-model:content="thread.content"></LazyVditorPreview>
            </ClientOnly>
        </div>
        <web-footer></web-footer>
    </div>
</template>

<script setup>
import { useRouter } from "vue-router";
import api from "@/common/api";
import MarkdownIt from "markdown-it";
import isbot from "isbot";
import { LazyVditorPreview } from "#components";

const config = useRuntimeConfig();
const headers = useRequestHeaders(["user-agent"]);

const router = useRouter();
const { id } = router.currentRoute.value.params;
const isBot = useState("isBot", () => {});
if (!isBot.value) {
    isBot.value = isbot(headers["user-agent"]);
}
const thread = useState("thread", () => {});

const postList = $ref([]);

const fetchThreads = async () => {
    if (!id) {
        console.error("id error", id);
        return;
    }
    // const [{ data: res }, { data: resList }] = await Promise.all([
    //     api.thread.one(id),
    //     api.post.get({
    //         threadId: id,
    //     }),
    // ]);

    const { data: res } = await api.thread.one(id);
    const { data: resList } = await api.post.get({
        threadId: id,
    });

    console.log("request...", res.value, resList.value);
    thread.value = res.value = res.value ? res.value : {};
    // 格式化 cover
    thread.value.cover = thread.value.cover ? config.fileURL + "/" + thread.value.cover : undefined;

    if (resList?.value?.list && resList?.value?.list[0]) {
        if (isBot.value) {
            const md = new MarkdownIt();
            thread.value.content = md.render(resList?.value?.list[0].content);
        } else {
            thread.value.content = resList?.value?.list[0].content;
        }
    }
    console.log("thread.value.content", resList, thread.value.content);
};

if (isBot.value) {
    await fetchThreads();
}

onMounted(async () => {
    console.log("onMounted...");
    if (!isBot.value) {
        // 进行md渲染
        await fetchThreads();
    }
});
</script>

<style lang="scss" scoped>
.main {
    max-width: 760px;
}
</style>
