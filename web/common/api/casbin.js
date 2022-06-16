import request from "../request"
export default {
    get: (params) => {
        return request('/api/v1/casbin-rules', {
            method: 'GET',
            params: params
        })
    },
    post: (data) => {
        return request('/api/v1/casbin-rules', {
            method: 'POST',
            body: data
        })
    },
    put: (id, data) => {
        return request('/api/v1/casbin-rules/' + id, {
            method: 'PUT',
            body: data
        })
    },
    delete: (id) => {
        return request('/api/v1/casbin-rules/' + id, {
            method: 'DELETE'
        })
    },
    getApi: () => {
        return request('/api/v1/casbin-rules/api', {
            method: 'GET'
        })
    },
}