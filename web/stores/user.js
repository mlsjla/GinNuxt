import {
    defineStore
} from 'pinia'

// user is the name of the store. It is unique across your application
// and will appear in devtools
export const useUserStore = defineStore('user', {
    // a function that returns a fresh state
    state: () => {
        return {
            user: {},
            token: undefined,
        }
    },
    persist: true,
    // optional getters
    getters: {

    },
    // optional actions
    actions: {
        reset() {
            // `this` is the store instance
        },
    },
})