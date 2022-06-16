import request from "../request"
export default {
    get: (params) => {
        return request('/api/v1/posts', {
            method: 'GET',
            params: params
        })
    },
    one: (id) => {
        return request('/api/v1/posts/' + id, {
            method: 'GET'
        })
    },
    post: (data) => {
        return request('/api/v1/posts', {
            method: 'POST',
            body: data
        })
    },
    put: (id, data) => {
        return request('/api/v1/posts/' + id, {
            method: 'PUT',
            body: data
        })
    },
    delete: (id) => {
        return request('/api/v1/posts/' + id, {
            method: 'DELETE'
        })
    }
}