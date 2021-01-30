import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

const baseRouters = [
    {
        path: '/',
        redirect: '/login'
    },
    {
        path: '/login',
        name: 'login',
        component: () =>
            import ('@/components/login/Login.vue')
    }
]

const createRouter = () => new Router({
    routes: baseRouters
})

const router = createRouter()

export default router