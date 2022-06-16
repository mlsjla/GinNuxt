import request from "../request"
export default {
    categoryTree: (data) => {
        return request('api/v1/pub/open/categories.tree', {
            method: 'GET'
        })
    },
    thread: (data) => {
        return request('api/v1/pub/open/thread', {
            method: 'GET',
            params: data
        })
    },
    post: (data) => {
        return request('api/v1/pub/open/post', {
            method: 'GET',
            params: data
        })
    }
}