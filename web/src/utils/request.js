import Vue from 'vue'
import axios from 'axios'; // 引入axios
import { Message } from 'element-ui';
import { store } from '@/store/index'


const service = axios.create({
    baseURL: process.env.VUE_APP_BASE_API,
    timeout: 99999
})
let acitveAxios = 0
let timer
const showLoading = function () {
    acitveAxios++
    if (timer) {
        clearTimeout(timer)
    }
    timer = setTimeout(function () {
        if (acitveAxios > 0) {
            Vue.bus.emit("showLoading")
        }
    }, 400);
}

const closeLoading = function () {
    acitveAxios--
    if (acitveAxios <= 0) {
        clearTimeout(timer)
        Vue.bus.emit("closeLoading")
    }
}
//http request 拦截器
service.interceptors.request.use(
    function (config) {
        if (!config.donNotShowLoading) {
            showLoading()
        }
        const token = store.getters['user/token']
        const user = store.getters['user/userInfo']
        console.log(user)
        config.data = JSON.stringify(config.data);
        config.headers = {
            'Content-Type': 'application/json',
            'x-token': token
        }
        return config;
    },
    function (error) {
        closeLoading()
        Message({
            showClose: true,
            message: error,
            type: 'error'
        })
        return error;
    }
);


//http response 拦截器
service.interceptors.response.use(
    response => {
        closeLoading()
        if (response.headers["new-token"]) {
            store.commit('user/setToken', response.headers["new-token"])
        }
        if (response.data.code == 200) {
            return response.data
        } else {
            Message({
                showClose: true,
                message: response.data.msg || decodeURI(response.headers.msg),
                type: response.headers.msgtype||'error',
            })
            if (response.data.data && response.data.data.reload) {
                store.commit('user/LoginOut')
            }
            return response.data.msg ? response.data : response
        }
    },
    error => {
        closeLoading()
        Message({
            showClose: true,
            message: error,
            type: 'error'
        })
        return error
    }
)

export default service