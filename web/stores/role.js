import {
    defineStore
} from 'pinia'
import api from "@/common/api";

// role is the name of the store. It is unique across your application
// and will appear in devtools
export const useRoleStore = defineStore('role', {
    // a function that returns a fresh state
    state: () => ({
        role: [],
    }),
    persist: true,
    // optional getters
    getters: {

    },
    // optional actions
    actions: {
        async init() {
            const {
                data: roleTree
            } = await api.role.get();
            const roleList = $ref([])
            roleList = roleTree?.value?.list
            this.role = roleList || []
            
        },
        reset() {
            // `this` is the store instance
        },
    },
})