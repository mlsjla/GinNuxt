import request from "../request"
export default {
    tree: () => {
        return request('/api/v1/apps.tree', {
            method: 'GET'
        })
    },
    get: (params) => {
        return request('/api/v1/apps', {
            method: 'GET',
            params: params
        })
    },
    one: (id) => {
        return request('/api/v1/apps/' + id, {
            method: 'GET'
        })
    },
    post: (data) => {
        return request('/api/v1/apps', {
            method: 'POST',
            body: data
        })
    },
    put: (id, data) => {
        return request('/api/v1/apps/' + id, {
            method: 'PUT',
            body: data
        })
    },
    delete: (id) => {
        return request('/api/v1/apps/' + id, {
            method: 'DELETE'
        })
    },
}