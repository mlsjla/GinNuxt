import request from "../request"
export default {
    post: (data) => {
        return request('/api/v1/pub/current/sendsms', {
            method: 'POST',
            body: data
        })
    },
}