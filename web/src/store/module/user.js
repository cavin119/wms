import { loginServer } from '@/api/user'

export const user = {
    namespaced: true,
    state: {
        userInfo: {
            username: 'Tom',
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
        }
    },
    action: {
        async LoginAction({ commit }, loginInfo) {
            const res = await loginServer()
            console.log(res)
            console.log(loginInfo)
            alert('login action')
            commit('setToken', 'token123456')
            return true
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