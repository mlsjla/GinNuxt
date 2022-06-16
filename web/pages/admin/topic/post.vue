<template>
    <div class="flex flex-col justify-start items-center w-full h-full grow">
        <el-form
            ref="formDataRef"
            :inline="true"
            :model="formData"
            :rules="rules"
            class="pt-8 flex flex-1 flex-col min-h-screen"
            :size="formSize"
        >
            <el-form-item label="标题" prop="title">
                <el-input v-model="formData.title" />
            </el-form-item>

            <div class="flex flex-row justify-start">
                <el-form-item label="类型" prop="type">
                    <el-select v-model="formData.type" placeholder="请选择类型">
                        <el-option
                            v-for="(item, index) in dictionaryStore.topic.threadType"
                            :key="index"
                            :value="item.value"
                            :label="item.label"
                        ></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="分类" prop="category_id">
                    <el-cascader
                        :props="{ label: 'name', value: 'id', emitPath: false }"
                        placeholder="请选择分类"
                        v-model="formData.category_id"
                        :options="categoryList"
                    />
                </el-form-item>
            </div>

            <div class="flex flex-row justify-start items-end">
                <el-form-item label="封面" prop="cover">
                    <imageUpload :limit="1" v-model:url="formData.cover"></imageUpload>
                </el-form-item>
            </div>

            <el-form-item label="正文" prop="content">
                <LazyVditor @input="vditorInput" v-model:content="formData.content"></LazyVditor>
            </el-form-item>

            <el-form-item label="简介" prop="summary">
                <el-input
                    v-model="formData.summary"
                    :autosize="{ minRows: 5, maxRows: 5 }"
                    type="textarea"
                    placeholder="请输入文章介绍"
                />
            </el-form-item>

            <el-form-item>
                <el-button class="mt-5" type="primary" @click="submitForm(formDataRef)">
                    提交
                </el-button>
            </el-form-item>
        </el-form>
    </div>
</template>

<script setup>
import { useDictionaryStore } from "@/stores/dictionary";
import { useCategoryStore } from "@/stores/category";
import api from "@/common/api";
import { useRouter, useRoute } from "#imports";
import { ElMessage } from "element-plus";
import { LazyVditor } from '#components'


definePageMeta({
    layout: "admin-no-menu",
});

const router = useRouter();

const dictionaryStore = useDictionaryStore();
const categoryStore = useCategoryStore();

const categoryList = computed(() => {
    return categoryStore.category;
});

onMounted(() => {
    categoryStore.init();
    nextTick(async () => {
        fetchThreads();
    });
});

const firstPost = $ref(undefined);
const options = $ref({});
const handleAfter = () => {};
const vditorInput = (content) => {};
const formSize = $ref("default");
const formDataRef = ref("FormInstance");

const rules = {};
const formData = $ref({
    address: undefined,
    attachment_price: undefined,
    category_id: undefined,
    free_words: undefined,
    id: undefined,
    content: "",
    cover: undefined,
    summary: undefined,
    is_anonymous: undefined,
    is_approved: undefined,
    is_display: undefined,
    is_draft: undefined,
    is_essence: undefined,
    is_red_packet: undefined,
    is_site: undefined,
    is_sticky: undefined,
    issue_at: undefined,
    latitude: undefined,
    location: undefined,
    longitude: undefined,
    paid_count: undefined,
    post_count: undefined,
    posted_at: undefined,
    price: undefined,
    rewarded_count: undefined,
    share_count: undefined,
    source: undefined,
    title: undefined,
    type: undefined,
});

const fetchThreads = async () => {
    const route = useRoute();
    if (!route.query?.id) {
        return;
    }
    const { data: thread } = await api.thread.one(route.query.id);
    // 获取content
    const { data: content } = await api.post.get({
        threadId: route.query.id,
    });
    if (content?.value?.list && content?.value?.list[0]) {
        firstPost = content?.value?.list[0];
    }

    formData = {
        ...thread.value,
        content: firstPost.content,
    };
    console.log("formData", formData);
};

const submitForm = async (formEl) => {
    if (!formEl) return;
    await formEl.validate(async (valid, fields) => {
        if (valid) {
            formData.category_id = formData.category_id;
            var res;
            if (formData.id) {
                const { data } = await api.thread.put(formData.id, formData);
                res = data;
            } else {
                const { data } = await api.thread.post(formData);
                res = data;
            }
            console.log("res", res);
            var postRes;
            if (firstPost && firstPost.id) {
                // 修改post
                firstPost.content = formData.content;
                const { data: pData } = await api.post.put(firstPost.id, firstPost);
                postRes = pData;
            } else if (formData.id && (!firstPost || !firstPost.id)) {
                // 修复没有post的thread
                firstPost = firstPost ? firstPost : {};
                firstPost.content = formData.content;
                firstPost.thread_id = formData.id;
                firstPost.is_first = "1";
                const { data: pData } = await api.post.post(firstPost);
                postRes = pData;
            }
            if (res?.value?.id || res?.value?.status == "OK") {
                return ElMessage({
                    message: "编辑成功",
                    type: "success",
                });
            }
            return ElMessage({
                message: "发生错误",
                type: "error",
            });
        }
    });
};
</script>

<style lang="scss" scoped></style>
