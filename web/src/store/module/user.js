import { loginServer } from '@/api/user'

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
        }
    },
    actions: {
        async LoginAction({ commit }, loginInfo) {
            const res = await loginServer(loginInfo)
            if (res.code == 200) {
                commit('setToken', res.data.token)
                commit('setUserInfo', {
                    username: res.data.username,
                    roleId: 0
                })
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