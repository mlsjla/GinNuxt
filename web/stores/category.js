import {
    defineStore
} from 'pinia'
import api from "@/common/api";

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
            const {
                data: categoryTree
            } = await api.category.tree();
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