<template>
    <div class="flex flex-col p-5 w-full justify-start items-start">
        <div class="mb-5">
            <el-button size="small" type="primary" icon="plus" @click="addThread()"> 添加文章 </el-button>
        </div>
        <el-table :data="threadList?.list" row-key="id">
            <el-table-column align="left" label="id" min-width="220" prop="id" />
            <el-table-column align="left" label="文章名" show-overflow-tooltip min-width="160" prop="title" />
            <el-table-column align="left" label="状态" min-width="100" prop="hidden">
                <template #default="scope">
                    <span>{{ scope.row.status == 1 ? "正常" : "禁用" }}</span>
                </template>
            </el-table-column>
            <el-table-column align="left" fixed="right" label="操作" width="300">
                <template #default="scope">
                    <el-button size="small" text icon="edit" @click="editThread(scope.row)">编辑</el-button>
                    <el-button size="small" text icon="delete" @click="deleteThread(scope.row)">删除</el-button>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>

<script setup>
// This will also work in `<script setup>`
definePageMeta({
    layout: "admin",
});
import api from "@/common/api";
import { useRouter } from "#imports";
import { ElMessage } from "element-plus";
const router = useRouter();

const { data: threadList, refresh } = await api.thread.get();

const formData = $ref({
    email: undefined,
    id: undefined,
    password: undefined,
    realname: undefined,
    status: 0,
    thread_roles: [],
    threadname: undefined,
});
const rules = {
    threadname: [
        {
            required: true,
            message: "请输入文章文章名",
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

const addThread = (row) => {
    const router = useRouter();
    router.push("/admin/topic/post");
    if (row?.parent_id > 0) {
        ElMessage({
            message: "子文章暂不支持添加子文章",
            type: "error",
        });
        return;
    }
};

const editThread = async (row) => {
    if (row.id) {
        const { data: thread } = await api.thread.one(row.id);
        formData = {
            ...thread?.value,
        };
    }
    router.push({
        path: "/admin/topic/post",
        query: {id: row.id}
    });
};

const deleteThread = async (row) => {
    await api.thread.delete(row.id);
    ElMessage({
        message: "删除成功",
        type: "success",
    });
    refresh();
};

const saveThread = async () => {
    if (formData.id) {
        const res = await api.thread.put(formData.id, formData);
    } else {
        const res = await api.thread.post(formData);
    }
    ElMessage({
        message: "保存成功",
        type: "success",
    });
    refresh();
};
</script>

<style lang="scss" scoped></style>
