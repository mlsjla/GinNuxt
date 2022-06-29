<template>
    <div class="flex flex-col justify-center item-content w-full h-full grow">
        <div ref="editorRef" class="flex flex-col w-full h-full"></div>
    </div>
</template>

<script setup>
import "vditor/dist/index.css";
import { useUserStore } from "@/stores/user";
import Vditor from 'vditor'

const editorRef = $ref();
var instance;

const { content } = defineProps(["content"]);
const emit = defineEmits(["change", "esc", "blur", "focus", "input", "update:content"]);
const config = useRuntimeConfig();
const userStore = useUserStore();

const initEditor = () => {
    instance = new Vditor(editorRef, {
        height: "auto",
        width: "100%",
        minHeight: "600px",
        toolbarConfig: {
            pin: false,
        },
        upload: {
            url: config.baseURL + "/api/v1/upload/upload",
            max: 1024 * 1024 * 50,
            headers: {
                Authorization: "Bearer " + userStore?.token?.access_token,
            },
            multiple: true,
            success: (element, res) => {
                let files = JSON.parse(res);
                const config = useRuntimeConfig();
                console.log("config", config.fileURL);
                let insertValue = "";
                if (files?.length > 0) {
                    files.forEach((item) => {
                        if (item?.file_path) {
                            insertValue += `![${item.file_name}](${config.fileURL}/${item.file_path})\n`;
                        }
                    });
                }
                /**
                    处理四种情况的附件
                    图片、视频、音频 其他
                 **/
                instance.insertValue(insertValue);
            },
        },
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
        cache: {
            enable: false,
        },
        after: () => {
            instance.setValue(content);
        },
        input: () => {
            emit("input", instance.getValue());
            emit("update:content", instance.getValue());
        },
        focus: () => {
            emit("focus");
        },
        blur: () => {
            emit("blur");
        },
        esc: () => {
            emit("esc");
        },
        select: () => {},
    });
};

onMounted(() => {
    nextTick(() => {
        initEditor();
    });
});

const stop = watch(
    () => content,
    (newVal, oldVal) => {
        if(newVal) {
            console.log("watch content...")
            instance.setValue(content);
            stop()
        }
    }
);
</script>

<style lang="scss" scoped>
:deep(.vditor-content) {
    min-height: 400px;
}
</style>
