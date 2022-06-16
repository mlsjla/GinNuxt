import request from "../request"
export default {
    tree: () => {
        return request('/api/v1/menus.tree', {
            method: 'GET'
        })
    },
    post: (data) => {
        return request('/api/v1/menus', {
            method: 'POST',
            body: data
        })
    },
    put: (id, data) => {
        return request('/api/v1/menus/' + id, {
            method: 'PUT',
            body: data
        })
    },
    delete: (id) => {
        return request('/api/v1/menus/' + id, {
            method: 'DELETE'
        })
    }
}