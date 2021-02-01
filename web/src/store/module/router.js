import { asyncRouterHandle } from '@/utils/asyncRouter';

const routerList = []

const formatRouter = (routers) => {
    routers && routers.map(item => {
        if ((!item.children || item.children.every(ch => ch.hidden)) && item.name != '404') {
            routerList.push({ label: item.meta.title, value: item.name })
        }
        if (item.children && item.children.length > 0) {
            formatRouter(item.children)
        }
    })
}
export const router = {
    namespaced: true,
    state: {
        asyncRouters: [],
        routerList: routerList
    },
    mutations: {
        setRouterList(state, routerList) {
            state.routerList = routerList
        },
        // 设置动态路由
        setAsyncRouter(state, asyncRouters) {
            state.asyncRouters = asyncRouters
        },
    },
    actions: {
        async SetAsyncRouter({ commit }) {
            const baseRouter = [{
                path: '/layout',
                name: 'layout',
                component: "components/layout/Layout.vue",
                meta: {
                    title: "底层layout"
                },
                children: []
            }]
            const asyncRouter = [
                {
                    path: "404",
                    name: "404",
                    hidden: true,
                    meta: {
                        title: "迷路了*。*",
                    },
                    component: 'components/error/Error.vue'
                },
                {
                    path: 'dashboard',
                    name: 'dashboard',
                    hidden: false,
                    sort: 1,
                    meta: {
                        title: '仪表盘',
                        icon: 'setting'
                    },
                    component: 'components/dashboard/Dashboard.vue'
                }
            ];
            formatRouter(asyncRouter)
            baseRouter[0].children = asyncRouter
            baseRouter.push({
                path: '*',
                redirect: '/layout/404'

            })
            asyncRouterHandle(baseRouter)
            commit('setAsyncRouter', baseRouter)
            commit('setRouterList', routerList)
            return true
        }
    },
    getters: {
        // 获取动态路由
        asyncRouters(state) {
            return state.asyncRouters
        },
        routerList(state) {
            return state.routerList
        }
    }
}
