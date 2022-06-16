
import { useUserStore } from '@/stores/user'

export default defineNuxtRouteMiddleware((to, from) => {
    const userStore = useUserStore()
    if(process.client && !userStore.user?.user_id) {
        console.log("to login...")
        return ('/user/login')
    }
})