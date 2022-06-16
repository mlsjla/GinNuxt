import {
    defineStore
} from 'pinia'
import api from "@/common/api";

// menu is the name of the store. It is unique across your application
// and will appear in devtools
export const useMenuStore = defineStore('menu', {
    // a function that returns a fresh state
    state: () => ({
        menu: [],
        userMenu: []
    }),
    persist: true,
    // optional getters
    getters: {

    },
    // optional actions
    actions: {
        async init() {
            const {
                data: menuTree
            } = await api.menu.tree();
            const menuList = $ref([])
            if(menuTree.value) {
                menuList = menuTree.value.list
                this.menu = menuList
            }

            const {data:userMenuTree} = api.user.menutree()
        },
        reset() {
            // `this` is the store instance
        },
    },
})