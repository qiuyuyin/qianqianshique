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
    dark: false,
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
    setDark(state, dark) {
      state.dark = dark
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
    async LoginIn({ commit, dispatch }, loginInfo) {
      await login(loginInfo).then(res => {
        console.log(res)
        if (res.code === 0) {
          commit('setUserInfo', res.data.user)
          commit('setToken', res.data.token)
          router.push({ name: 'dashboard', replace: true })
        } else {
          dispatch('snackbar/openSnackbar', {
            msg: res.msg,
            color: 'error',
          }, { root: true })
        }
      })
    },
    LoginOut({ commit }) {
      commit('LoginOut')
    },
    SetDark({ commit }, dark) {
      commit('setDark', dark)
    },
  },
}

export default user
