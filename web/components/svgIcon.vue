<template>
    <img
        @click="click"
        class="flex flex-col w-full h-full"
        :src="svgSrc"
        :mode="mode"
    />
</template>

<script setup lang="ts">

const emit = defineEmits(['close'])


const { src, color='#cdcdcd' } = defineProps<{
    src: string,
    color?: string,
    mode?: string
}>()
const images = new Map([
    ['eye', eye],
    ['chat-text', chatText],
    ['chevron-right', chevronRight],
    ['sort-up', sortUp],
    ['sort-down', sortDown],
    ['hand-thumbs-up', handThumbsUp],
    ['star', star],
    ['share', share],
    ['pencil-square', pencilSquare],
    ['card-image', cardImage],
    ['play-btn', playBtn],
    ['ui-checks-grid', uiChecksGrid]
])


const getColorSVG = (svg: string, colors:string):string => {
    try {
        if (/<svg /.test(svg)) { // 先简单判断是一下否是一个svg
            let newSvg;
            if (/fill=".*?"/.test(svg)) {
                newSvg = svg.replace(/fill=".*?"/, `fill="${colors}"`);  // SVG有默认色
            } else {
                newSvg = svg.replace(/<svg /, `<svg fill="${colors}"`); // 无默认色
            }
            return 'data:image/svg+xml;base64,' + encode(newSvg); // 替换完之后再组合回去
        }
    } catch { }
    return '';
}

const click = () => {
    emit('click')
}

const svgSrc = computed(() => {
    if(!src || !color) return ""
    return getColorSVG(images.get(src) || "", color)
})


</script>

<style scoped>

</style>