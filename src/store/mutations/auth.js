const auth = {
  setCurrentUser(state, currentUser) {
    state.currentUser = currentUser;
  },

  logout(state) {
    state.jwtToken = undefined;
    state.isLoggedIn = false;
    state.currentUser = undefined;
  },
};

export default auth;
