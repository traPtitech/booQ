import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    drawer: null,
    color: 'success',
    sidebarBackgroundColor: 'rgba(27, 27, 27, 0.74)'
  },
  mutations: {
    setDrawer (state, data) {
      state.drawer = data
    },
    setColor (state, data) {
      state.color = data
    },
    toggleDrawer (state) {
      state.drawer = !state.drawer
    }
  },
  actions: {
    //
  }
})
