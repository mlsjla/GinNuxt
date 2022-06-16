import { useUserStore } from '@/stores/user'
import { useRouter, useRoute } from "#imports";

const request = async (url, options) => {
    console.log('request...', url)
    const config = useRuntimeConfig()
    const userStore = useUserStore()
    return await useFetch(url, {
        ...options,
        baseURL: config.baseURL,
        method: options.method,
        body: options.body,
        params: options.params,
        initialCache: false,
        server: true,
        headers: {
            ...options.headers,
            Authorization: 'Bearer ' + userStore?.token?.access_token
        },
        async onResponseError({ request, response, options }) {
            // Log error
            console.log('[fetch response error]', request, response.status, response.body)
            if(response.status == 401) {
                // 登录失效
                const router = useRouter();
                router.push('/admin/user/login')
            }
        }
    })
}
export default request