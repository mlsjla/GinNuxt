<template>
    <div class="flex flex-col p-5 w-full justify-start items-start">
        <div class="mb-5">
            <el-button size="small" type="primary" icon="plus" @click="addRole()"> 添加角色 </el-button>
        </div>
        <el-table :data="roleStore.role" row-key="id">
            <el-table-column align="left" label="id" min-width="220" prop="id" />
            <el-table-column align="left" label="名称" show-overflow-tooltip min-width="160" prop="name" />
            <el-table-column align="left" label="排序" min-width="70" prop="sequence" />
            <el-table-column align="left" label="状态" min-width="100" prop="hidden">
                <template #default="scope">
                    <span>{{ scope.row.status == 1 ? "正常" : "禁用" }}</span>
                </template>
            </el-table-column>
            <el-table-column align="left" fixed="right" label="操作" width="300">
                <template #default="scope">
                    <el-button size="small" text icon="plus" @click="setAcl(scope.row)">设置权限</el-button>
                    <el-button size="small" text icon="plus" @click="addRole(scope.row)">添加子角色</el-button>
                    <el-button size="small" text icon="edit" @click="editRole(scope.row)">编辑</el-button>
                    <el-button size="small" text icon="delete" @click="deleteRole(scope.row)">删除</el-button>
                </template>
            </el-table-column>
        </el-table>

        <el-dialog :destroy-on-close="true" v-model="aclVisible" title="设置权限">
            <el-tabs v-model="settingAclType">
                <el-tab-pane label="API权限" name="api">
                    <div class="pl-6 w-full">
                        <el-checkbox v-model="checkedTreeAll" label="全选" />
                    </div>
                    <el-tree node-key="id" ref="aclTree" :default-checked-keys="checkedApiKeys" default-expand-all :data="apiList" :props="apiProps" show-checkbox :height="400" />
                </el-tab-pane>
                <el-tab-pane label="菜单权限" name="menu">
                    <el-tree node-key="id" ref="menuTree" :default-checked-keys="checkedRoleMenuKeys" default-expand-all :data="menuStore.menu" :props="menuProps" show-checkbox :height="400" />
                </el-tab-pane>
            </el-tabs>
            
            <template #footer>
                <span class="dialog-footer">
                    <el-button v-if="settingAclType == 'api'" type="primary" @click="saveRoleAcl">保存</el-button>
                    <el-button v-if="settingAclType == 'menu'" type="primary" @click="saveRoleMenu">保存</el-button>
                </span>
            </template>
        </el-dialog>

        <el-dialog :destroy-on-close="true" v-model="dialogVisible" title="编辑角色">
            <el-form ref="formDataRef" :model="formData" :rules="rules" label-width="120px" class="demo-formData">
                <el-form-item label="名称" prop="name">
                    <el-input v-show="false" v-model="formData.id"></el-input>
                    <el-input v-model="formData.name"></el-input>
                </el-form-item>
                <el-form-item label="排序" prop="sequence">
                    <el-input-number :controls="false" v-model="formData.sequence"></el-input-number>
                </el-form-item>
                <el-form-item label="是否启用" prop="status">
                    <el-switch :active-value="1" :inactive-value="2" v-model="formData.status"></el-switch>
                </el-form-item>
                <el-form-item label="备注" prop="memo">
                    <el-input v-model="formData.memo" type="textarea"></el-input>
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <el-button type="primary" @click="saveRole">保存</el-button>
                </span>
            </template>
        </el-dialog>
    </div>
</template>

<script setup>
import { useRoleStore } from "@/stores/role";
import { useMenuStore } from "@/stores/menu";

import api from "@/common/api";
import { ElMessage } from "element-plus";

// This will also work in `<script setup>`
definePageMeta({
    layout: "admin",
});

const roleStore = useRoleStore();
roleStore.init()
const menuStore = useMenuStore();
const dialogVisible = $ref(false);
const formData = $ref({
    id: undefined,
    name: undefined,
    icon: undefined,
    status: 1,
    sequence: undefined,
    memo: undefined,
    creator: undefined
});
const rules = {
    name: [
        {
            required: true,
            message: "请输入角色名称",
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

// 计算属性， 计算一级角色
const parentRoles = computed(() => {
    return roleStore.role.map((item) => {
        return {
            label: item.name,
            value: item.id,
            id: item.id,
        };
    });
});

const apiList = $ref([])
const checkedTreeAll = $ref(false)
const checkedApiKeys = $ref([])
const checkedRoleMenuKeys = $ref([])
const settingAclType = $ref("api")
const aclTree = ref()
const menuTree = ref()

watch(() => checkedTreeAll, () => {
    checkedApiKeys = []
    if(!checkedTreeAll) {
        aclTree.value.setCheckedKeys([], false)
        return
    }
    apiList.forEach(item => {
        checkedApiKeys.push(item.id)
    })
})

const getApi = async () => {
    const { data } = await api.casbin.getApi()
    if(!data.value || data.value.length == 0) {
        return
    }
    let list = data.value.map(item => {
        return {
            id: item.path + "---" + item.method,
            value: item.path + "---" + item.method,
            label: item.path + "---" + item.method,
            method: item.method,
            path: item.path,
            children: []
        }
    })

    const formatApi = (list) => {
        let output = []
        if(!list || list.length == 0) {
            return []
        }
        list.forEach((item) => {
            if(output.length == 0) {
                output.push({
                    label: item.path,
                    id: item.path,
                    value: item.path,
                    method: undefined,
                    path: item.path,
                    children: [
                        item
                    ]
                })
                return
            }else if(item.path.indexOf('/api/v1/pub') != -1) {
                return
            }else if(item.path.indexOf('/api/v1/upload') != -1) {
                let current = output[output.length -1]
                if(current.path != '/api/v1/upload') {
                    output.push({
                        label: '/api/v1/upload',
                        id: '/api/v1/upload',
                        value: '/api/v1/upload',
                        method: undefined,
                        path: '/api/v1/upload',
                        children: [
                            item
                        ]
                    })
                }else {
                    current.children.push(item)
                }
                return
            }else {
                let current = output[output.length -1]
                if(item.path.indexOf(current.path) != -1) {
                    current.children.push(item)
                    return
                }
                output.push({
                    label: item.path,
                    id: item.path,
                    value: item.path,
                    method: undefined,
                    path: item.path,
                    children: [
                        item
                    ]
                })
                return
            }
        })
        return output
    }

    apiList = formatApi(list)
}
onMounted( async() => {
    console.log('onMounted..')
    await getApi()
})

const aclVisible = $ref(false)
const apiProps = {
  value: 'id',
  label: 'label',
  children: 'children',
}
const menuProps = {
  id: 'id',
  label: 'name',
  children: 'children',
}
const currentRoleId = $ref(undefined)
const setAcl = async (row) => {
    currentRoleId = row.id
    await getUserCasbin()
    await getApi()
    await getRoleMenu()

    aclVisible = true
}

// 菜单权限处理
const currentRoleMenuList = $ref([])
const getRoleMenu = async () => {
    const { data } = await api.roleMenu.get({
        role_id: currentRoleId
    })
    currentRoleMenuList = data?.value?.list ? data.value.list : []
    checkedRoleMenuKeys = currentRoleMenuList.map(item => {
        return item.menu_id.toString()
    })
    console.log("checkedRoleMenuKeys", checkedRoleMenuKeys, currentRoleMenuList)
}



const currentAclList = $ref([])
const getUserCasbin = async () => {
    if(!currentRoleId) {
        console.error('currentRoleId error')
        return
    }
    const {data} = await api.casbin.get({
        pageSize: 1000,
        v0: currentRoleId.toString()
    })
    if(data?.value?.list) {
        let keys = []
        data.value.list.forEach(item => {
            keys.push(item.v1 + "---" + item.v2)
        })
        // aclTree.value.setCheckedKeys(keys, false)
        checkedApiKeys = keys
        currentAclList = data.value.list
    }
}

const addRole = (row) => {
    if (row?.parent_id > 0) {
        ElMessage({
            message: "子角色暂不支持添加子角色",
            type: "error",
        });
        return;
    }
    dialogVisible = true;
    formData = {
        status: 1
    }
};

const editRole = (row) => {
    dialogVisible = true;
    formData = {
        ...row,
    };
};

const deleteRole = async (row) => {
    await api.role.delete(row.id);
    roleStore.init();
    ElMessage({
        message: "删除成功",
        type: "success",
    });
};

const  saveRoleMenu = async () => {
    let checkItems = menuTree.value.getCheckedNodes(true, false)
    console.log("checkItems", checkItems)
    for (let i = 0; i < checkItems.length; i++) {
        let item = checkItems[i]

        // 如果当前已经有此权限，跳过
        let id = checkedRoleMenuKeys.find((id) => {
            return id == item.id
        })
        if(id) {
            console.log('已经有此权限，跳过...')
            continue
        }
        
        const { data:r } = api.roleMenu.post({
            role_id: currentRoleId,
            menu_id: item.id
        })
    }

    // 删除多余的
    for (let j=0; j<currentRoleMenuList.length; j++) {
        let item = currentRoleMenuList[j]
        let index = checkItems.find(menu => {
            return menu.id == item.menu_id
        })
        if(index) {
            continue
        }
        // 多余的
        const {data:r} = await api.roleMenu.delete(item.id)
    }
    aclVisible = false
}

const saveRoleAcl = async () => {
    let checkItems = aclTree.value.getCheckedNodes(true, false)
    for (let i = 0; i < checkItems.length; i++) {
        let item = checkItems[i]

        // 如果当前已经有此权限，跳过
        let id = checkedApiKeys.find((id) => {
            return id == item.id
        })
        if(id) {
            console.log('已经有此权限，跳过...')
            continue
        }
        
        const { data:r } = api.casbin.post({
            p_type: "g",
            v0: currentRoleId,
            v1: item.path,
            v2: item.method
        })
        console.log("r", r)
    }

    // 删除多余的
    for (let j=0; j<currentAclList.length; j++) {
        let item = currentAclList[j]
        let id1 = item.v1 + "---" + item.v2
        let index = checkItems.find(acl => {
            return acl.id == id1
        })
        if(index) {
            continue
        }
        // 多余的
        const {data:r} = await api.casbin.delete(item.id)
    }
    aclVisible = false
}

const saveRole = async () => {
    if (formData.id) {
        const res = await api.role.put(formData.id, formData);
    } else {
        const res = await api.role.post(formData);
    }
    roleStore.init();
    dialogVisible = false;
    ElMessage({
        message: "保存成功",
        type: "success",
    });
};
</script>

<style lang="scss" scoped></style>
