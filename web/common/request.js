import { useUserStore } from '@/stores/user'
import { useRouter, useRoute } from "#imports";

const request = async (url, options) => {    
    // 对服务端请求和客户端请求进行分开处理
    if(process.client) {
        console.log("client request...")
    }else {
        console.log("server request...")
    }

    const config = useRuntimeConfig()
    const userStore = useUserStore()

    if(!config) {
        console.error("config error...")
        return
    }
    
    const requestURL  = process.client ? config.baseURL : (config.baseServerURL ? config.baseServerURL : config.baseURL)

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