import { useUserStore } from '@/stores/user'
import { hash } from 'ohash'


const request = async (url, options) => {    
    // 对服务端请求和客户端请求进行分开处理
    const config = useRuntimeConfig()
    const userStore = useUserStore()

    if(!config) {
        console.error("config error...")
        return
    }
    
    const requestURL  = process.client ? config.baseURL : (config.baseServerURL ? config.baseServerURL : config.baseURL)

    const result = await useFetch(url, {
        ...options,
        key : hash(['api-fetch', url, options.params, options.method]),
        baseURL: requestURL,
        method: options.method,
        body: options.body,
        params: options.params,
        initialCache: false,
        headers: {
            ...options.headers,
            Authorization: 'Bearer ' + userStore?.token?.access_token
        },
        async onResponseError({ request, response, options }) {
            // Log error
            console.log('[fetch response error]', request, response.status, response.body)
            if(response.status == 401) {
                // 登录失效
                userStore && userStore.$patch({
                    token: undefined,
                });
                const router = useRouter();
                router.push('/admin/user/login')
            }
        }
    })

    return result
}
export default request