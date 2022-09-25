import {
    defineStore
} from 'pinia'
import api from "@/common/api";


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
        async init() {
            console.log("this.token", this.token)
            if(!this.token) {
                return
            }
            const { data: info } = await api.user.current()
            this.user = info.value || {}
        },
        reset() {
            // `this` is the store instance
            this.user = {}
            this.token = ""
        },
    },
})