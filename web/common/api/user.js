import request from "../request"
export default {
    get: (params) => {
        return request('/api/v1/users', {
            method: 'GET',
            params: params
        })
    },
    one: (id) => {
        return request('/api/v1/users/' + id, {
            method: 'GET'
        })
    },
    post: (data) => {
        return request('/api/v1/users', {
            method: 'POST',
            body: data
        })
    },
    put: (id, data) => {
        return request('/api/v1/users/' + id, {
            method: 'PUT',
            body: data
        })
    },
    delete: (id) => {
        return request('/api/v1/users/' + id, {
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
        return request('/api/v1/pub/current/user', {
            method: 'GET'
        })
    },
    editUser: (data) => {
        return request('/api/v1/pub/current/user', {
            body: data,
            method: 'PUT'
        })
    },
    menutree: () => {
        return request('/api/v1/pub/current/menutree', {
            method: 'GET'
        })
    },
}