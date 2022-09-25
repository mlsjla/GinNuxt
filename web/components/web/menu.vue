<template>
    <ul class="menu menu-vertical menu-horizontal bg-base-100 rounded-box">
        <li @click="toItem(item.path)" v-for="(item, index) in tabList" :key="index">
            <h5
                :class="{ active: currentAction == item.value ? true : false }"
                >{{ item.label }}</h5
            >
        </li>
    </ul>
</template>

<script setup lang="ts">
const { currentAction } = defineProps(["currentAction"]);

const tabList = $ref([
    {
        label: "总览信息",
        value: "index",
        path: "/package/miniprogram",
    },
    {
        label: "基础设置",
        value: "edit",
        path: "/package/miniprogram/edit",
    }
]);

const toItem = (path) => {
    const router = useRouter()
    router.push(path)
}
</script>

<style lang="scss" scoped>

</style>