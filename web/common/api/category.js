import request from "../request"
export default {
    tree: () => {
        return request('/api/v1/categories.tree', {
            method: 'GET'
        })
    },
    post: (data) => {
        return request('/api/v1/categories', {
            method: 'POST',
            body: data
        })
    },
    put: (id, data) => {
        return request('/api/v1/categories/' + id, {
            method: 'PUT',
            body: data
        })
    },
    delete: (id) => {
        return request('/api/v1/categories/' + id, {
            method: 'DELETE'
        })
    },
}