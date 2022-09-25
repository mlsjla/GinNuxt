<template>
    <div class="flex flex-col">
        <el-upload name="file[]" :action="uploadURL" :headers="headers" class="flex" 
            :on-success="uploadSuccess" :auto-upload="true" 
            :show-file-list="false"
        >
            <el-button type="primary">点击上传</el-button>
            <template #tip>
                <div class="el-upload__tip">
                    
                </div>
            </template>
        </el-upload>
    </div>
</template>
<script setup>
import { ref } from "vue";
import { Delete, Download, Plus, ZoomIn } from "@element-plus/icons-vue";
import { genFileId } from "element-plus";

const { url, limit } = defineProps(["url", "limit"]);
const emit = defineEmits(["change", "esc", "blur", "focus", "input", "update:url"]);
const config = useRuntimeConfig();

const userStore = inject("userStore")
const upload = ref();
const uploadURL = config.baseURL + "/api/v1/upload/upload"
const headers = {
    Authorization: "Bearer " + userStore?.token?.access_token,
}


const dialogVisible = $ref(false);
const disabled = $ref(false);

const previewUrl = computed(() => {
    if(!url) {
        return undefined
    }
    return config.fileURL + "/" + url
})

const uploadSuccess = (res) => {
    if(res && res[0]?.file_path) {
        emit('update:url', res[0]?.file_path)
        emit('change', res[0]?.file_path)
    }
}


</script>
<style lang="scss" scoped>
:deep(.image-uploader) {
    .image {
        width: 128px;
        height: 128px;
        display: flex;
    }

    .el-upload {
        border: 1px dashed var(--el-color-primary);
        border-radius: 6px;
        cursor: pointer;
        position: relative;
        overflow: hidden;
        transition: var(--el-transition-duration-fast);
    }

    .el-upload:hover {
        border-color: var(--el-color-primary);
    }

    .el-icon.image-uploader-icon {
        font-size: 28px;
        color: #8c939d;
        width: 128px;
        height: 128px;
        text-align: center;
    }
}
</style>
