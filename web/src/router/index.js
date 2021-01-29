import {createWebHistory, createRouter} from 'vue-router'

const routers = [
    {
        path: "/",
        name: "hello",
        component: () => import('@/components/HelloWorld')
    },
    {
        path: "/login",
        name: "login",
        component: () => import('@/components/login/login')
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes: routers,
})

export default router


