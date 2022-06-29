<template>
    <div class="flex flex-col justify-start items-center w-full min-h-screen">
        <!-- <web-header></web-header> -->
        <div class="main main-p w-full flex flex-col justify-center items-center flex-grow">
            <div class="card shadow bg-white w-1/3 mt-5 mb-5">
                <div class="card-body">
                    <h2 class="card-title">
                        <div class="logo flex flex-row justify-center items-center w-full">
                            <h2>用户登录</h2>
                        </div>
                    </h2>

                    <div class="form-control">
                        <label class="label">
                            <span class="label-text">用户名</span>
                        </label>
                        <label class="input-group input-group-xs">
                            <span><i class="bi bi-person" style="font-size: 18px"></i></span>
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
                            <span><i class="bi bi-lock" style="font-size: 18px"></i></span>
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
import { useRouter, useRoute } from "#imports";
import api from "@/common/api";
import { ElMessage } from "element-plus";
import Schema from "async-validator";

definePageMeta({
  layout: "default-nossr",
  // or middleware: 'auth'
})


const config = useRuntimeConfig();
const userStore = useUserStore();

const captcha = $ref({});
const captchaString = $ref("");
// 刷新验证码
const captchaReload = async () => {
    await getCaptcha();
    captchaString = dayjs().unix().toString();
};
const getCaptcha = async () => {
    const { data: res } = await api.captcha.get();
    captcha = res.value;
};

const captcha_id = computed(() => {
    return captcha.captcha_id;
});

const captchaUrl = computed(() => {
    let captcha_id = captcha?.captcha_id;
    return captcha_id
        ? config.baseURL + "/api/v1/pub/login/captcha?id=" + captcha_id + "&reload=" + captchaString
        : undefined;
});

const formData = $ref({
    username: undefined,
    password: undefined,
    captcha_id: undefined, // 验证码id
    captcha_code: undefined,
});

// 登录操作
const login = async () => {
    formData.captcha_id = captcha_id;

    // 表单验证
    const descriptor = {
        username: {
            type: "string",
            required: true,
            message: "用户名必须填写",
        },
        password: {
            type: "string",
            required: true,
            message: "密码不低于8位字符串",
        },
        captcha_id: {
            type: "string",
            required: true,
            message: "验证码参数错误",
        },
        captcha_code: {
            type: "string",
            required: true,
            message: "必须输入验证码",
        },
    };
    const cn = {
        required: "%s 必填",
    };
    const validator = new Schema(descriptor);
    validator.messages(cn);
    await validator.validate(formData, async (errors, fields) => {
        if (errors) {
            return ElMessage({
                message: descriptor[errors[0].field]?.message,
                type: "error",
            });
        }
        // validation passed
        const { data: loginInfo, error } = await api.user.login(formData);
        getCaptcha();
        if (error?.value?.data?.error?.message) {
            let msg = error?.value?.data?.error?.message;
            ElMessage({
                message: msg,
                type: "error",
            });
            return;
        } else {
            ElMessage({
                message: "登录成功",
                type: "success",
            });

            const userStore = useUserStore();
            userStore.$patch({
                token: loginInfo,
            });

            const { data: userInfo } = await api.user.current();
            userInfo &&
                userStore.$patch({
                    user: userInfo,
                });

            const router = useRouter()
            try {
                router.back()
            } catch (e) {
                console.log("router.back()", e)
                router.push("/")
            }
        }
    });
};

onMounted(() => {
    formData = userStore?.user
        ? {
              ...userStore?.user,
          }
        : {};

    nextTick(async () => {
        getCaptcha();
    });
});
</script>

<style scoped></style>
