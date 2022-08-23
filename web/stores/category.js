import {
    defineStore
} from 'pinia'
import api from "@/common/api";
import { useUserStore } from './user.js'

// category is the name of the store. It is unique across your application
// and will appear in devtools
export const useCategoryStore = defineStore('category', {
    // a function that returns a fresh state
    state: () => ({
        category: [],
    }),
    persist: true,
    // optional getters
    getters: {

    },
    // optional actions
    actions: {
        async init() {
            let userStore = useUserStore()
            const fetchCategory = async () => {
                if(userStore.token) {
                    return await api.category.tree()
                }else {
                    return await api.category.openTree()
                }
            }
            const {
                data: categoryTree
            } = await fetchCategory()
            
            const categoryList = $ref([])
            if(categoryTree.value) {
                categoryList = categoryTree.value.list
                this.category = categoryList
            }
            
        },
        reset() {
            // `this` is the store instance
        },
    },
})