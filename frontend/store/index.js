// import Vue from 'vue'
// import Vuex from 'vuex'

// Vue.use(Vuex)

export const state = {
  services: new Array(),
  categories: new Array(),
}

export const mutations = {
  setServices(state, services) {
    state.services = services
  },
  setCategories(state, categories) {
    state.categories = categories
  },
}

export const actions = {
  async getServices(context) {
    try {
      const res = (await this.$axios.get('http://localhost:8000/getServices'))
        .data
      context.commit('setServices', res.Data)
    } catch (error) {
      console.log(error)
    }
  },
  async getCategories(context) {
    try {
      const res = (await this.$axios.get('http://localhost:8000/getCategories'))
        .data
      context.commit('setCategories', res.Data)
    } catch (error) {
      console.log(error)
    }
  },
}
