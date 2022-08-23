<template>
    <div class="flex flex-col p-5 w-full justify-start items-start">
        <div class="mb-5">
            <el-button size="small" type="primary" icon="plus" @click="addMenu()"> 添加根菜单 </el-button>
        </div>
        <el-table default-expand-all :data="menuStore.menu" row-key="id">
            <el-table-column align="left" label="id" min-width="220" prop="id" />
            <el-table-column align="left" label="名称" show-overflow-tooltip min-width="160" prop="name" />
            <el-table-column align="left" label="路由" show-overflow-tooltip min-width="160" prop="router" />
            <el-table-column align="left" label="状态" min-width="100" prop="hidden">
                <template #default="scope">
                    <span>{{ scope.row.is_show == 1 ? "显示" : "隐藏" }}</span>
                </template>
            </el-table-column>
            <el-table-column align="left" label="父节点" min-width="200" prop="parent_id" />
            <el-table-column align="left" label="排序" min-width="70" prop="sequence" />
            <el-table-column align="left" label="图标" min-width="140" prop="icon">
                <template #default="scope">
                    <div v-if="scope.row?.icon" class="icon-column">
                        <el-icon>
                            <component :is="scope.row?.icon" />
                        </el-icon>
                        <span>{{ scope.row?.icon }}</span>
                    </div>
                </template>
            </el-table-column>
            <el-table-column align="left" fixed="right" label="操作" width="300">
                <template #default="scope">
                    <el-button size="small" text icon="plus" @click="addMenu(scope.row)">添加子菜单</el-button>
                    <el-button size="small" text icon="edit" @click="editMenu(scope.row)">编辑</el-button>
                    <el-button size="small" text icon="delete" @click="deleteMenu(scope.row)">删除</el-button>
                </template>
            </el-table-column>
        </el-table>

        <el-dialog :destroy-on-close="true" v-model="dialogVisible" title="编辑菜单">
            <el-form ref="formDataRef" :model="formData" :rules="rules" label-width="120px" class="demo-formData">
                <el-form-item label="名称" prop="name">
                    <el-input v-show="false" v-model="formData.id"></el-input>
                    <el-input v-model="formData.name"></el-input>
                </el-form-item>
                <el-form-item label="图标" prop="icon">
                    <el-select-v2 :options="iconsList" v-model="formData.icon" filterable placeholder="">
                        <template #default="{ item }">
                            <div class="flex flex-row justify-between items-center">
                                <i :class="'bi-' + item.value" style="font-size: 24px"></i>
                                <span class="">
                                    {{ item.value }}
                                </span>
                            </div>
                        </template>
                    </el-select-v2>
                    <i v-if="formData.icon" class="pl-5" :class="'bi-' + formData.icon" style="font-size: 24px"></i>
                </el-form-item>
                <el-form-item label="路由" prop="router">
                    <el-input v-model="formData.router"></el-input>
                </el-form-item>
                <el-form-item label="父菜单">
                    <el-select @change="parentChange" v-model="formData.parent_id" clearable placeholder="选择父菜单">
                        <el-option v-for="item in parentMenus" :key="item.value" :label="item.label" :value="item.value">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="父菜单路由" prop="parent_router">
                    <el-input :disabled="true" v-model="formData.parent_router"></el-input>
                </el-form-item>
                <el-form-item label="是否显示" prop="is_show">
                    <el-switch :active-value="1" :inactive-value="2" v-model="formData.is_show"></el-switch>
                </el-form-item>
                <el-form-item label="是否启用" prop="status">
                    <el-switch :active-value="1" :inactive-value="2" v-model="formData.status"></el-switch>
                </el-form-item>
                <el-form-item label="排序" prop="sequence">
                    <el-input-number :controls="false" v-model="formData.sequence"></el-input-number>
                </el-form-item>
                <el-form-item label="备注" prop="memo">
                    <el-input v-model="formData.memo" type="textarea"></el-input>
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <el-button type="primary" @click="saveMenu">保存</el-button>
                </span>
            </template>
        </el-dialog>
    </div>
</template>

<script setup>
import iconsData from "@/common/icons";
import api from "@/common/api";
import { ElMessage } from "element-plus";

// This will also work in `<script setup>`
definePageMeta({
    layout: "admin",
});

const menuStore = inject('menuStore');
const dialogVisible = $ref(false);
const iconsList = reactive(iconsData);
const formData = $ref({
    id: undefined,
    name: undefined,
    icon: undefined,
    router: undefined,
    parent_id: undefined,
    parent_router: undefined,
    is_show: 1,
    status: 1,
    sequence: undefined,
    memo: undefined,
    creator: undefined,
});
const rules = {
    name: [
        {
            required: true,
            message: "请输入菜单名称",
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

// 计算属性， 计算一级菜单
const parentMenus = computed(() => {
    return menuStore.menu.map((item) => {
        return {
            label: item.name,
            value: item.id,
            id: item.id,
            router: item.router
        };
    });
});

const parentChange = (parent_id) => {
    if(parent_id == formData.id) {
        ElMessage({
            message: '不允许选择自己为父菜单',
            type: 'error'
        })
        formData.parent_id = undefined
    }else if(parent_id){
        let parent = parentMenus.value.find((item) => {
            return parent_id == item.id
        })
        formData.parent_router = parent.router
    }
}

const addMenu = (row) => {
    if(row?.parent_id > 0) {
        ElMessage({
            message: "子菜单暂不支持添加子菜单",
            type: "error",
        });
        return;
    }
    dialogVisible = true;
    formData = {
        is_show: 1,
        status: 1
    };
    if(row?.id) {
        formData.parent_id = row.id;
        formData.parent_router = row.router;
    }
    
};

const editMenu = (row) => {
    dialogVisible = true;
    formData = {
        ...row,
    };
};

const deleteMenu = async (row) => {
    await api.menu.delete(row.id);
    menuStore.init();
    ElMessage({
        message: "删除成功",
        type: "success",
    });
};

const saveMenu = async () => {
    if (formData.id) {
        const res = await api.menu.put(formData.id, formData);
    } else {
        const res = await api.menu.post(formData);
    }
    menuStore.init();
    dialogVisible = false;
    ElMessage({
        message: "保存成功",
        type: "success",
    });
};
</script>

<style lang="scss" scoped></style>
