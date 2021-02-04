import { loginServer } from '@/api/user'
import router from '@/router/index'

export const user = {
    namespaced: true,
    state: {
        userInfo: {
            username: '',
            roleId: 0
        },
        token: ''
    },
    mutations: {
        setUserInfo(state, userInfo) {
            state.userInfo = userInfo
        },
        setToken(state, token) {
            state.token = token
        },
        Logout(state) {
            state.userInfo = {}
            state.token = ''
            sessionStorage.clear()
            router.push({ name: 'home', replace: true })
        }
    },
    actions: {
        async LoginAction({ commit, dispatch, rootState, getters }, loginInfo) {
            const res = await loginServer(loginInfo)
            if (res.code == 200) {
                commit('setToken', res.data.token)
                commit('setUserInfo', {
                    username: res.data.username,
                    roleId: 0
                })
                console.log(getters["userInfo"])
                await dispatch('router/SetAsyncRouter', {}, { root: true })
                const asyncRouters = rootState.router.asyncRouters
                console.log(asyncRouters)
                router.addRoutes(asyncRouters)
                router.push({ name: 'dashboard' })
                return true
            }
        },
        async LogoutAction({ commit }) {
            alert('Logout action')
            commit('Logout')
        }
    },
    getters: {
        userInfo(state) {
            return state.userInfo
        },
        token(state) {
            return state.token
        }
    }
}