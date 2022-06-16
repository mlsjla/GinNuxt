import request from "../request"

export default {
    get: () => {
        return request('/api/v1/pub/login/captchaid', {
            method: 'GET'
        })
    },
}