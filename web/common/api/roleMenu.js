import request from "../request"
export default {
    get: (params) => {
        return request('/api/v1/role-menus', {
            method: 'GET',
            params: params
        })
    },
    post: (data) => {
        return request('/api/v1/role-menus', {
            method: 'POST',
            body: data
        })
    },
    put: (id, data) => {
        return request('/api/v1/role-menus/' + id, {
            method: 'PUT',
            body: data
        })
    },
    delete: (id) => {
        return request('/api/v1/role-menus/' + id, {
            method: 'DELETE'
        })
    },
}