<template>
    <div ref="contentRef" class="w-full vditor-content flex-col leading-relaxed"></div>
</template>

<script setup>
import VditorPreview from "vditor/dist/method.min";
const { content } = defineProps(["content"]);
const contentRef = $ref();
onMounted(() => {
    console.log('preview content', content)
    nextTick(() => {
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
    });
});
</script>

<style lang="scss" scoped></style>
