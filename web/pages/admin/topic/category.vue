<template>
    <div class="flex flex-col p-5 w-full justify-start items-start">
        <div class="mb-5">
            <el-button size="small" type="primary" icon="plus" @click="addCategory()"> 添加根分类 </el-button>
        </div>
        <el-table default-expand-all :data="categoryStore.category" row-key="id">
            <el-table-column align="left" label="id" min-width="220" prop="id" />
            <el-table-column align="left" label="名称" show-overflow-tooltip min-width="160" prop="name" />
            <el-table-column align="left" label="类型" min-width="100" prop="property">
                <template #default="scope">
                    <span>{{ scope.row.property == 0 ? "普通" : "显示首页" }}</span>
                </template>
            </el-table-column>
            <el-table-column align="left" label="父节点" min-width="200" prop="parentid" />
            <el-table-column align="left" label="主题数" min-width="100" prop="thread_count" />
            <el-table-column align="left" label="版主" min-width="200" prop="moderators" />
            <el-table-column align="left" label="排序" min-width="70" prop="sort" />
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
                    <el-button size="small" text icon="plus" @click="addCategory(scope.row)">添加子分类</el-button>
                    <el-button size="small" text icon="edit" @click="editCategory(scope.row)">编辑</el-button>
                    <el-button size="small" text icon="delete" @click="deleteCategory(scope.row)">删除</el-button>
                </template>
            </el-table-column>
        </el-table>
        
        <el-dialog :destroy-on-close="true" v-model="dialogVisible" title="编辑分类">
            <el-form ref="formDataRef" :model="formData" :rules="rules" label-width="120px" class="demo-formData">
                <el-form-item label="名称" prop="name">
                    <el-input v-show="false" v-model="formData.id"></el-input>
                    <el-input v-model="formData.name"></el-input>
                </el-form-item>
                <el-form-item label="描述" prop="description">
                    <el-input v-model="formData.description" type="textarea"></el-input>
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
                <el-form-item label="属性" prop="property">
                    <el-select v-model="formData.property" placeholder="选择属性">
                        <el-option v-for="item in propertyOptions" :key="item.value" :label="item.label" :value="item.value">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="版主" prop="moderators">
                    <el-input v-model="formData.moderators"></el-input>
                </el-form-item>
                <el-form-item label="父分类">
                    <el-select @change="parentChange" v-model="formData.parentid" placeholder="选择父分类">
                        <el-option v-for="item in parentCategorys" :key="item.value" :label="item.label" :value="item.value">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="排序" prop="sort">
                    <el-input-number :controls="false" v-model="formData.sort"></el-input-number>
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <el-button type="primary" @click="saveCategory">保存</el-button>
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
import { useCategoryStore } from "@/stores/category";
import iconsData from "@/common/icons";
import api from "@/common/api";
import { ElMessage } from "element-plus";
const propertyOptions = $ref([
    {
        label: "普通",
        value: 0,
    },
    {
        label: "首页展示",
        value: 1,
    }
])

const categoryStore = useCategoryStore();
categoryStore.init()
const dialogVisible = $ref(false);
const iconsList = reactive(iconsData);
const formData = $ref({
    id: undefined,
    name: undefined,
    description: undefined,
    icon: undefined,
    parentid: undefined,
    thread_count: undefined,
    sort: undefined
});
const rules = {
    name: [
        {
            required: true,
            message: "请输入分类名称",
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

// 计算属性， 计算一级分类
const parentCategorys = computed(() => {
    if(!categoryStore?.category) {
        return []
    }
    return categoryStore.category.map((item) => {
        return {
            label: item.name,
            value: item.id,
            id: item.id
        };
    });
});

const parentChange = (parentid) => {
    if(parentid == formData.id) {
        ElMessage({
            message: '不允许选择自己为父分类',
            type: 'error'
        })
        formData.parentid = undefined
    }else if(parentid){
        let parent = parentCategorys.value.find((item) => {
            return parentid == item.id
        })
    }
}

const addCategory = (row) => {
    if(row?.parentid > 0) {
        ElMessage({
            message: "子分类暂不支持添加子分类",
            type: "error",
        });
        return;
    }
    dialogVisible = true;
    formData = {
    };
    if(row?.id) {
        formData.parentid = row.id;
    }
    console.log('formData', formData, row)
    
};

const editCategory = (row) => {
    dialogVisible = true;
    formData = {
        ...row,
    };
};

const deleteCategory = async (row) => {
    await api.category.delete(row.id);
    categoryStore.init();
    ElMessage({
        message: "删除成功",
        type: "success",
    });
};

const saveCategory = async () => {
    if (formData.id) {
        const res = await api.category.put(formData.id, formData);
    } else {
        const res = await api.category.post(formData);
    }
    categoryStore.init();
    dialogVisible = false;
    ElMessage({
        message: "保存成功",
        type: "success",
    });
};
</script>

<style lang="scss" scoped></style>
