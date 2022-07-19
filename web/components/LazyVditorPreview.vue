<template>
    <div ref="contentRef" class="w-full prose md:prose-lg lg:prose-xl max-w-none"></div>
</template>

<script setup>
import VditorPreview from 'vditor/dist/method.min'

const { content } = defineProps(["content"]);
const contentRef = $ref();
const preview = async() => {
    VditorPreview.preview(contentRef, content, {
        cdn: "/vditor",
        IPreviewOptions: {
            cdn: "/vditor",
        },
        emojiPath: {
            cdn: "/vditor",
        },
        themes: {
            cdn: "/vditor",
        },
        mode: "ir",
        hljs: false,
    });
}

onMounted(async() => {
    console.log("渲染...")
    preview()
});
onUnmounted(() => {
    console.log("onUnmounted...")
})

</script>

<style lang="scss" scoped></style>
