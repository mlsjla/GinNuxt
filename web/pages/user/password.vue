<template>
    <div class="flex flex-col justify-start items-center w-full min-h-screen">
        <!-- <web-header></web-header> -->
        <div class="main main-p w-full flex flex-col justify-center items-center flex-grow">
            <div class="card shadow bg-white w-1/3 mt-5 mb-5">
                <div class="card-body">
                    <h2 class="card-title">
                        <div
                            class="logo flex flex-row justify-center items-center w-full"
                        >
                            <h2>密码找回</h2>
                        </div>
                    </h2>

                    <div class="form-control">
                        <label class="label">
                            <span class="label-text">用户名</span>
                        </label>
                        <label class="input-group input-group-xs">
                            <span
                                ><i
                                    class="bi bi-person"
                                    style="font-size: 18px"
                                ></i
                            ></span>
                            <input
                                v-model="formData.username"
                                type="text"
                                placeholder="请输入用户名"
                                class="input input-ghost border-0 flex-grow"
                            />
                        </label>
                    </div>
                    <div class="form-control my-2">
                        <label class="label">
                            <span class="label-text">密码</span>
                        </label>
                        <label class="input-group input-group-xs">
                            <span
                                ><i
                                    class="bi bi-lock"
                                    style="font-size: 18px"
                                ></i
                            ></span>
                            <input
                                v-model="formData.password"
                                type="password"
                                placeholder="请输入密码"
                                class="input input-ghost border-0 flex-grow"
                            />
                        </label>
                    </div>
                    <div class="form-control mb-5">
                        <label class="label">
                            <span class="label-text">验证码</span>
                        </label>
                        <div class="input-group input-group-xs relative">
                            <input
                                v-model="formData.captcha_code"
                                type="text"
                                placeholder="验证码"
                                class="input input-ghost border-0 flex-grow pr-16"
                            />
                            <img
                                @click="captchaReload"
                                :src="captchaUrl"
                                class="absolute top-0 right-0 rounded-l-none w-1/3 h-full"
                            />
                        </div>
                    </div>

                    <button @click="login" class="btn btn-primary">登录</button>
                </div>
            </div>
        </div>
        <web-footer></web-footer>
    </div>
</template>

<script setup>
import { useUserStore } from "@/stores/user";
import { useRouter } from "vue-router";
import api from "@/common/api";
import { ElMessage } from "element-plus";

definePageMeta({
  layout: "default-nossr",
  // or middleware: 'auth'
})


const config = useRuntimeConfig();
const router = useRouter();
const userStore = useUserStore();

const captcha = $ref({})
const captchaString = $ref("");
// 刷新验证码
const captchaReload = async () => {
    await getCaptcha()
    captchaString = dayjs().unix().toString();
}
const getCaptcha = async () => {
    const { data: res } = await api.captcha.get();
    captcha = res.value
}

const captcha_id = computed(() => {
    return captcha.captcha_id;
});

const captchaUrl = computed(() => {
    let captcha_id = captcha?.captcha_id;
    return captcha_id
    ? config.baseURL +
        "/api/v1/pub/login/captcha?id=" +
        captcha_id +
        "&reload=" +
        captchaString
    : undefined;
});

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
        } else {
            console.log("error submit!", fields);
        }
    });
}

onMounted(() => {
    formData = userStore?.user ? {
        ...userStore?.user
    } : {}

    nextTick(async () => {
        getCaptcha()
    })
});
</script>

<style scoped></style>
