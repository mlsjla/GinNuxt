<template>
    <div class="flex flex-col justify-start item-content w-full h-full grow mt-5 bg-white">
        <el-form
            ref="formDataRef"
            label-width="auto"
            :model="formData"
            :rules="rules"
            class="p-5 max-w-lg"
            :size="formSize"
        >
            <el-form-item label="启用验证码" prop="useCaptcha">
                <el-switch v-model="formData.useCaptcha" />
            </el-form-item>
            <el-form-item label="短信通道" prop="type">
                <el-select class="w-full" v-model="formData.smsType" placeholder="请选择类型">
                    <el-option
                        v-for="(item, index) in dictionaryStore.setting.smsType"
                        :key="index"
                        :value="item.value"
                        :label="item.label"
                    ></el-option>
                </el-select>
            </el-form-item>
            <el-form-item label="阿里云accessKeyId" prop="accessKeyId">
                <el-input v-model="formData.accessKeyId" placeholder="请输入阿里云短信服务accessKeyId" />
            </el-form-item>
            <el-form-item label="阿里云accessKeySecret" prop="accessKeySecret">
                <el-input v-model="formData.accessKeySecret" placeholder="请输入阿里云短信服务accessKeySecret" />
            </el-form-item>
            <el-form-item label="短信签名" prop="signName">
                <el-input v-model="formData.signName" placeholder="请输入短信签名" />
            </el-form-item>

            <div v-for="(item, index) in dictionaryStore.setting.smsTemplate" :key="index">
                <el-form-item :label="item.label + '短信模板ID'" prop="templateCode">
                    <el-input v-model="formData.smsTemplate[index].templateCode" placeholder="" />
                    <el-input v-show="false" v-model="formData.smsTemplate[index].value" placeholder="" />
                </el-form-item>
                <el-form-item :label="item.label + '短信模板详情'" prop="template">
                    <el-input
                        type="textarea"
                        v-model="formData.smsTemplate[index].template"
                        placeholder="请输入短信签名"
                    />
                </el-form-item>
            </div>

            <el-form-item>
                <el-button type="primary" @click="submitForm(formDataRef)">创建</el-button>
                <el-button @click="dialogFormVisible = !dialogFormVisible">测试</el-button>
                
                <el-dialog v-model="dialogFormVisible" title="测试短信">
                    <el-form :model="smsForm">
                        <el-form-item class="mt-5" label="短信模板" :label-width="150">
                            <el-select v-model="smsForm.value" placeholder="选择发送短信模板">
                                <el-option v-for="(item, index) in formData.smsTemplate" :key="index" :label="item.label" :value="item.value" />
                            </el-select>
                        </el-form-item>
                        <el-form-item class="mt-5" label="手机号" :label-width="150">
                            <el-input v-model="smsForm.mobile" autocomplete="off" />
                        </el-form-item>
                    </el-form>
                    <template #footer>
                        <span class="dialog-footer">
                            <el-button @click="dialogFormVisible = false">取消</el-button>
                            <el-button type="primary" @click="sendSms">发送</el-button>
                        </span>
                    </template>
                </el-dialog>
            </el-form-item>
        </el-form>
    </div>
</template>

<script setup>
definePageMeta({
    layout: "admin",
});
import { ElMessage } from "element-plus";
import api from "@/common/api";

const dictionaryStore = inject('dictionaryStore');
const categoryStore = inject("categoryStore")
const categoryList = computed(() => {
    return categoryStore.category;
});

const content = $ref("");
const options = $ref({});
const handleAfter = () => {};
const vditorInput = (content) => {};
const formSize = $ref("default");
const formDataRef = ref("FormInstance");

const formData = $ref({
    useCaptcha: undefined,
    smsType: undefined,
    accessKeyId: undefined,
    accessKeySecret: undefined,
    signName: undefined,
    smsTemplate: dictionaryStore.setting.smsTemplate,
});

let params = {
    pageSize: 1,
    key: "sms",
    tag: "sms",
};

const rules = reactive({});

const smsSetting = {
    id: undefined,
    key: undefined,
    value: undefined,
    tag: "sms",
};

onMounted(async () => {
    nextTick(async () => {
        console.log("nextTick...")
        const { data: settingRes } = await api.setting.get(params);
        if (settingRes?.value?.list && settingRes?.value?.list[0]) {
            smsSetting.id = settingRes?.value?.list[0].id;
            formData = JSON.parse(settingRes?.value?.list[0].value);
        }
    })
    console.log("onMounted...")
})

const smsForm = $ref({
    value: undefined,
    mobile: undefined
})
const dialogFormVisible = $ref(false)
const sendSms = async () => {
    let {data} = await api.sms.post(smsForm)
}

async function submitForm(formEl) {
    if (!formEl) return;
    await formEl.validate(async (valid, fields) => {
        if (valid) {
            smsSetting.key = "sms";
            smsSetting.tag = "sms";
            smsSetting.value = JSON.stringify(formData);

            let res;
            if (!smsSetting.id) {
                res = await api.setting.post(smsSetting);
            } else {
                res = await api.setting.put(smsSetting.id, smsSetting);
            }
            const data = res.data;
            console.log("data=>", data);

            if (data?.value?.id) {
                ElMessage({
                    message: "配置成功",
                    type: "success",
                });
                smsSetting.id = data.value?.id;
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
</script>

<style lang="scss" scoped></style>
