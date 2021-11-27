import router from '@/router/index'
import { login } from '@/api/user'

const user = {
  namespaced: true,
  state: {
    userInfo: {
      uuid: '',
      nickName: '',
      headerImg: '',
    },
    token: '',
  },
  mutations: {
    setUserInfo(state, userInfo) {
      // 这里的 `state` 对象是模块的局部状态
      state.userInfo = userInfo
    },
    setToken(state, token) {
      // 这里的 `state` 对象是模块的局部状态
      state.token = token
    },
    LoginOut(state) {
      state.userInfo = {}
      state.token = ''
      sessionStorage.clear()
      console.log(state.token === '')
      router.push({ name: 'dashboard', replace: true })
      window.location.reload()
    },
  },
  actions: {
    LoginIn({ commit }, loginInfo) {
      login(loginInfo).then(res => {
        console.log(res)
        if (res.code === 0) {
          commit('setUserInfo', res.data.user)
          commit('setToken', res.data.token)
          router.push({ name: 'dashboard', replace: true })
          window.location.reload()

          return true
        }

        return false
      })

      return false
    },
    LoginOut({ commit }) {
      commit('LoginOut')
    },
  },
}

export default user
