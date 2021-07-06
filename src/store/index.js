import Vue from "vue";
import Vuex from "vuex";

import actions from "./actions";
import mutations from "./mutations";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    jwtToken: localStorage.getItem("jwtToken"),
    currentUser: null,
    boards: [],
  },
  getters: {
    isLoggedIn: (state) => !!state.jwtToken,
    currentUser: (state) => state.currentUser,
  },
  mutations,
  actions,
  modules: {},
});
