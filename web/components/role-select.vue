<template>
    <div>
        <ClientOnly>
            <el-select @change="roleChange" v-model="roles" multiple placeholder="Select">
                <el-option v-for="item in rolestore.role" :key="item.id" :label="item.name" :value="item.id"> </el-option>
            </el-select>
        </ClientOnly>
    </div>
</template>

<script setup>
import { useRoleStore } from "@/stores/role";
const rolestore = useRoleStore();
const { user_roles } = defineProps(["user_roles"]);
const emit = defineEmits(["change", "update:user_roles"]);

const roles = $ref([])
roles = user_roles.map(item => {
    return item.role_id
})

const roleChange = () => {
    console.log('user_roles', user_roles, roles)
    let newArr = []
    for (let i = 0; i <roles.length; i++) {
        let id = roles[i]
        let item = user_roles.find(item => {
            return item.role_id == id
        })
        if(item) {
            newArr.push(item)
        }else {
            newArr.push({
                role_id: id
            })
        }
    }
    emit('change', newArr)
    emit('update:user_roles', newArr)
    
}
</script>

<style lang="scss" scoped></style>
