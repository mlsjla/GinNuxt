import {
    defineStore
} from 'pinia'
import api from "@/common/api";

// dictionary is the name of the store. It is unique across your application
// and will appear in devtools
export const useDictionaryStore = defineStore('dictionary', {
    // a function that returns a fresh state
    state: () => ({
        setting: {
            smsType: [
                {
                    value: 0,
                    label: "阿里云短信" 
                }
            ],
            smsTemplate: [
                {
                    label: "身份验证",
                    value: "register",
                    templateCode: "",
                    template: ""
                },
                {
                    label: "修改密码",
                    value: "password",
                    templateCode: "",
                    template: ""
                }
            ]
        },
        topic: {
            threadType: [{
                    value: 0,
                    label: "普通"
                },
                {
                    value: 1,
                    label: "长文"
                },
                {
                    value: 2,
                    label: "视频"
                },
                {
                    value: 3,
                    label: "图片"
                },
                {
                    value: 4,
                    label: "微信文章"
                }
            ]

        },
    }),
    persist: false,
    // optional getters
    getters: {

    },
    // optional actions
    actions: {
        async init() {},
        reset() {
            // `this` is the store instance
        },
    },
})