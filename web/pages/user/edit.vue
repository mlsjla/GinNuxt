<template>
    <div class="flex flex-col justify-start items-center w-full min-h-screen">
        <web-header></web-header>
        <div class="main main-p w-full flex flex-row flex-center justify-start items-left flex-grow">
            <ul
                class="menu w-full md:w-56 menu-normal menu-horizontal md:menu-vertical mb-5 bg-base-100 p-2 rounded-box"
            >
                <li v-for="(item, index) in tabList" :key="index">
                    <NuxtLink
                        :class="{ active: currentAction == item.value ? true : false }"
                        :to="`/user/${item.value}`"
                        >{{ item.label }}</NuxtLink
                    >
                </li>
            </ul>
            <el-form
                ref="formDataRef"
                :inline="true"
                :model="formData"
                :rules="rules"
                class="ml-5 flex flex-1 flex-col"
                label-width="100px"
                :size="formSize"
            >
                <el-form-item label="头像" prop="avatar">
                    <imageUpload :limit="1" v-model:url="formData.avatar"></imageUpload>
                </el-form-item>

                <el-form-item label="昵称" prop="nickname">
                    <el-input v-model="formData.nickname" />
                </el-form-item>
                <el-form-item label="真实名称" prop="realname">
                    <el-input v-model="formData.realname" />
                </el-form-item>
                <el-form-item label="简介" prop="introduce">
                    <el-input
                        v-model="formData.introduce"
                        :autosize="{ minRows: 5, maxRows: 5 }"
                        type="textarea"
                        placeholder="个人介绍"
                    />
                </el-form-item>

                <el-form-item label=" ">
                    <el-button class="mt-5" type="primary" @click="submitForm(formDataRef)"> 提交 </el-button>
                </el-form-item>
            </el-form>
        </div>
        <web-footer></web-footer>
    </div>
</template>

<script setup>
import { useRouter } from "vue-router";
import api from "@/common/api";
import { ElMessage } from "element-plus";

definePageMeta({
  middleware: ["auth"],
  layout: "default-nossr",
  // or middleware: 'auth'
})


const router = useRouter();
const userStore = inject("userStore")
const currentAction = 'edit'

const tabList = $ref([
    {
        label: "个人设置",
        value: "edit",
    },
    {
        label: "修改密码",
        value: "password",
    },
]);

const formSize = $ref("default");
const rules = {};
const formData = $ref({});
const formDataRef = ref("FormInstance");

const  submitForm = async (formEl) => {
    if (!formEl) return;
    await formEl.validate(async (valid, fields) => {
        if (valid) {
            const { data } = await api.user.editUser(formData)
            console.log('editUser', data.value)
            if (data?.value?.id) {
                ElMessage({
                    message: "修改成功",
                    type: "success",
                });
            } else if (data?.value?.status == "OK") {
                ElMessage({
                    message: "更新成功",
                    type: "success",
                });
            }

            await userStore.init()
            formData = userStore?.user ? {
                ...userStore?.user
            } : {}
        } else {
            console.log("error submit!", fields);
        }
    });
}

onMounted(() => {
    formData = userStore?.user ? {
        ...userStore?.user
    } : {}
});
</script>

<style scoped></style>
