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
    const res = (await this.$axios.get('/getServices')).data
    if (res.Data == null && res.ErrorCode != 200) {
      return new Promise((resolve, reject) => {
        reject({
          Error: res.Error,
          Code: res.ErrorCode,
        })
      })
    } else {
      context.commit('setServices', res.Data)
      return new Promise((resolve, reject) => {
        resolve(res.Data)
      })
    }
  },
  async getCategories(context) {
    const res = (await this.$axios.get('/getCategories')).data
    if (res.Data == null && res.ErrorCode != 200) {
      return new Promise((resolve, reject) => {
        reject({
          Error: res.Error,
          Code: res.ErrorCode,
        })
      })
    } else {
      context.commit('setCategories', res.Data)
      return new Promise((resolve, reject) => {
        resolve(res.Data)
      })
    }
  },
}
