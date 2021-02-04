import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

const baseRouters = [
    {
        path: '/',
        name: 'home',
        meta: {
            title: '首页',
        },
        component: () =>
            import ('@/components/login/Login.vue')
    },
    {
        path: '/login',
        name: 'login',
        meta: {
            title: '管理员登录',
        },
        component: () =>
            import ('@/components/login/Login.vue')
    }
]

const createRouter = () => new Router({
    routes: baseRouters
})

const router = createRouter()

export default router