<template>
    <div class="flex flex-col w-full h-full bg-base-100">
        <div class="flex flex-col h-full w-full">
            <ul class="menu p-5 flex-grow">
                <li class="menu-title">
                    <span> 首页 </span>
                </li>
                <li @click="tapMenu(item)" v-for="(item, index) in menuChildren" :key="index">
                    <a> {{ item.name }} </a>
                </li>
            </ul>
        </div>
        <admin-footer class="p-5"></admin-footer>
    </div>
</template>

<script setup>

const router = useRouter();
const route = useRoute();

const menuStore = inject('menuStore');

const tapMenu = (item) => {
    if (item.router) {
        router.push(item.router);
    } else if (item?.children[0] && item?.children[0]?.router) {
        router.push(item?.children[0]?.router);
    }
};

const menuChildren = computed(() => {
    if (!menuStore.menu || !menuStore.menu.length) {
        return "";
    }
    let menu = [];
    menuStore.menu.forEach((item) => {
        if (item.router == route.path) {
            menu = item.children;
            return;
        }
        if (item?.children) {
            item.children.forEach((child) => {
                if (child.router == route.path) {
                    menu = item.children;
                    return;
                }
            });
        }
    });
    if(!menu) {
        return []
    }
    return menu.filter((item) => {
        return item.is_show == 1 && item.status == 1;
    });
});
</script>

<style lang="scss" scoped></style>
