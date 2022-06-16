<template>
    <div class="flex flex-col p-5 w-full justify-start items-start">
        <div class="mb-5">
            <el-button size="small" type="primary" icon="plus" @click="addUser()"> 添加用户 </el-button>
        </div>
        <el-table :data="userList?.list" row-key="id">
            <el-table-column align="left" label="id" min-width="220" prop="id" />
            <el-table-column align="left" label="用户名" show-overflow-tooltip min-width="160" prop="username" />
            <el-table-column align="left" label="状态" min-width="100" prop="hidden">
                <template #default="scope">
                    <span>{{ scope.row.status == 1 ? "正常" : "禁用" }}</span>
                </template>
            </el-table-column>
            <el-table-column align="left" fixed="right" label="操作" width="300">
                <template #default="scope">
                    <el-button size="small" text icon="edit" @click="editUser(scope.row)">编辑</el-button>
                    <el-button size="small" text icon="delete" @click="deleteUser(scope.row)">删除</el-button>
                </template>
            </el-table-column>
        </el-table>

        <el-dialog :destroy-on-close="true" v-model="dialogVisible" title="编辑用户">
            <el-form ref="formDataRef" :model="formData" :rules="rules" label-width="120px" class="demo-formData">
                <el-form-item label="用户名" prop="username">
                    <el-input v-show="false" v-model="formData.id"></el-input>
                    <el-input v-model="formData.username"></el-input>
                </el-form-item>
                <el-form-item label="密码" prop="password">
                    <el-input v-model="formData.password"></el-input>
                </el-form-item>
                <el-form-item label="邮箱" prop="email">
                    <el-input v-model="formData.email"></el-input>
                </el-form-item>
                <el-form-item label="手机号" prop="mobile">
                    <el-input v-model="formData.mobile"></el-input>
                </el-form-item>
                <el-form-item label="真实名字" prop="realname">
                    <el-input v-model="formData.realname"></el-input>
                </el-form-item>
                <el-form-item label="是否启用" prop="status">
                    <el-switch :active-value="1" :inactive-value="2" v-model="formData.status"></el-switch>
                </el-form-item>
                <el-form-item label="角色" prop="user_roles">
                    <role-select v-model:user_roles="formData.user_roles"></role-select>
                    <!-- <el-input v-model="formData.user_roles" type="textarea"></el-input> -->
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <el-button type="primary" @click="saveUser">保存</el-button>
                </span>
            </template>
        </el-dialog>
    </div>
</template>

<script setup>
// This will also work in `<script setup>`
definePageMeta({
    layout: "admin",
});
import api from "@/common/api";
import { ElMessage } from "element-plus";
const dialogVisible = $ref(false);
const {data: userList, refresh} = await api.user.get()

const formData = $ref({
    email: undefined,
    id: undefined,
    password: undefined,
    realname: undefined,
    status: 0,
    user_roles: [
    ],
    username: undefined,
});
const rules = {
    username: [
        {
            required: true,
            message: "请输入用户用户名",
            trigger: "blur",
        },
        {
            min: 1,
            max: 10,
            message: "字符串长度为1到10",
            trigger: "blur",
        },
    ],
};


const addUser = (row) => {
    if (row?.parent_id > 0) {
        ElMessage({
            message: "子用户暂不支持添加子用户",
            type: "error",
        });
        return;
    }
    dialogVisible = true;
    formData = {
        status: 1,
        user_roles: [],
    };
};

const editUser = async (row) => {
    if(row.id) {
        const {data: user} = await api.user.one(row.id)
        formData = {
            ...user?.value
        }
    }
    dialogVisible = true;
};

const deleteUser = async (row) => {
    await api.user.delete(row.id);
    ElMessage({
        message: "删除成功",
        type: "success",
    });
    refresh()
};

const saveUser = async () => {
    if (formData.id) {
        const res = await api.user.put(formData.id, formData);
    } else {
        const res = await api.user.post(formData);
    }
    dialogVisible = false;
    ElMessage({
        message: "保存成功",
        type: "success",
    });
    refresh()
};
</script>

<style lang="scss" scoped></style>
