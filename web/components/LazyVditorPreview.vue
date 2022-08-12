<template>
    <div ref="contentRef" class="w-full prose md:prose-lg lg:prose-xl max-w-none"></div>
</template>

<script setup>
import VditorPreview from 'vditor/dist/method.min.js'

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
    preview()
});
onUnmounted(() => {
})

</script>

<style lang="scss" scoped></style>
