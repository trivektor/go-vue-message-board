const auth = {
  async login() {},

  logout({ commit }) {
    commit("logout");
  },

  async getCurrentUser({ commit }) {
    commit("setCurrentUser", null);
  },
};

export default auth;
