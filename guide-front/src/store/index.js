import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    user: {},
    pid: {}
  },
  getters: {
    getuser: function (state) {
      return state.user
    },
    getPid: function (state) {
      return state.pid
    }
  },
  mutations: {
    changeUser (state, target) {
      state.user = target
    },
    changePid (state, target) {
      state.pid = target
    }
  },
  actions: {
    changeFun (context, target) {
      context.commit('changeUser', target)
    },
    changePidFun (context, target) {
      context.commit('changePid', target)
    }
  }
})
export default store
