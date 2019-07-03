import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    user: {}
  },
  getters: {
    getuser: function (state) {
      return state.user
    }
  },
  mutations: {
    changeUser (state, target) {
      state.user = target
    }
  },
  actions: {
    changeFun (context, target) {
      context.commit('changeUser', target)
    }
  }
})
export default store
