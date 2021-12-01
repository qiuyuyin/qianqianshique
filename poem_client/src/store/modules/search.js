const search = {
  namespaced: true,
  state: {
    searchRecords: [],
  },
  mutations: {
    ClearRecords(state) {
      state.searchRecords = []
    },
    appendRecords(state, record) {
      let flag = true
      state.searchRecords.forEach(element => {
        if (element === record) {
          flag = false
        }
      })
      if (flag) {
        state.searchRecords.unshift(record)
        if (state.length > 5) {
          state.searchRecords.pop()
        }
      }
    },
  },
  actions: {
    AppendRecord({ commit }, record) {
      if (record !== '') {
        commit('appendRecords', record)
      }
    },
    ClearRecord({ commit }) {
      commit('ClearRecords')
    },
  },
}

export default search
