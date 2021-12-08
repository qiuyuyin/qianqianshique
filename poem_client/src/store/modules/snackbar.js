const snackbar = {
  //
  namespaced: true,
  state: {
    msg: '',
    color: '',
    visible: false,
    showClose: true,
    timeout: 4000,

  },

  // 逻辑处理,同步函数
  mutations: {
    OPEN_SNACKBAR(state, options) {
      state.visible = true
      state.msg = options.msg
      state.color = options.color
    },
    CLOSE_SNACKBAR(state) {
      state.visible = false
    },
    setShowClose(state, isShow) {
      state.showClose = isShow
    },
    setTimeout(state, timeout) {
      state.timeout = timeout
    },
  },

  // 逻辑处理,异步函数
  actions: {
    openSnackbar(context, options) {
      const { timeout } = context.state
      context.commit('OPEN_SNACKBAR', {
        msg: options.msg,
        color: options.color,
      })
      setTimeout(() => {
        context.commit('CLOSE_SNACKBAR')
      }, timeout)
    },
  },
}
export default snackbar
