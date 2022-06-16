import request from "../request"
export default {
    get: () => {
        return request('/api/v1/threads', {
            method: 'GET'
        })
    },
    one: (id) => {
        return request('/api/v1/threads/' + id, {
            method: 'GET'
        })
    },
    post: (data) => {
        return request('/api/v1/threads', {
            method: 'POST',
            body: data
        })
    },
    put: (id, data) => {
        return request('/api/v1/threads/' + id, {
            method: 'PUT',
            body: data
        })
    },
    delete: (id) => {
        return request('/api/v1/threads/' + id, {
            method: 'DELETE'
        })
    },
    login: (data) => {
        return request('/api/v1/pub/login', {
            method: 'POST',
            body: data
        })
    },
    current: () => {
        return request('/api/v1/pub/current/thread', {
            method: 'GET'
        })
    }
}