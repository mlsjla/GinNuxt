import request from "../request"
export default {
    get: (params) => {
        return request('/api/v1/settings', {
            method: 'GET',
            params: params
        })
    },
    post: (data) => {
        return request('/api/v1/settings', {
            method: 'POST',
            body: data
        })
    },
    put: (id, data) => {
        return request('/api/v1/settings/' + id, {
            method: 'PUT',
            body: data
        })
    },
    delete: (id) => {
        return request('/api/v1/settings/' + id, {
            method: 'DELETE'
        })
    },
}