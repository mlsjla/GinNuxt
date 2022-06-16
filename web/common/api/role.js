import request from "../request"
export default {
    tree: () => {
        return request('/api/v1/roles.tree', {
            method: 'GET'
        })
    },
    get: () => {
        return request('/api/v1/roles', {
            method: 'GET'
        })
    },
    post: (data) => {
        return request('/api/v1/roles', {
            method: 'POST',
            body: data
        })
    },
    put: (id, data) => {
        return request('/api/v1/roles/' + id, {
            method: 'PUT',
            body: data
        })
    },
    delete: (id) => {
        return request('/api/v1/roles/' + id, {
            method: 'DELETE'
        })
    },
}