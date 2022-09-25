import request from "../request"
export default {
    tree: () => {
        return request('/api/v1/app-logs.tree', {
            method: 'GET'
        })
    },
    get: (params) => {
        return request('/api/v1/app-logs', {
            method: 'GET',
            params: params
        })
    },
    one: (id) => {
        return request('/api/v1/app-logs/' + id, {
            method: 'GET'
        })
    },
    preview: (id) => {
        return request('/api/v1/app-logs.preview/' + id, {
            method: 'POST'
        })
    },
    upload: (id) => {
        return request('/api/v1/app-logs.upload/' + id, {
            method: 'POST'
        })
    },
    post: (data) => {
        return request('/api/v1/app-logs', {
            method: 'POST',
            body: data
        })
    },
    put: (id, data) => {
        return request('/api/v1/app-logs/' + id, {
            method: 'PUT',
            body: data
        })
    },
    delete: (id) => {
        return request('/api/v1/app-logs/' + id, {
            method: 'DELETE'
        })
    },
}