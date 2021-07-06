import router from "../../router";

const auth = {
  async login(_, { username, password }) {
    const response = await fetch("/api/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
      },
      body: JSON.stringify({
        username,
        password,
      }),
    });
    const json = await response.json();
    console.log(json);
  },

  logout({ commit }) {
    commit("logout");
  },

  async getCurrentUser() {},

  async registerUser({ commit }, { username, password }) {
    // TODO: error handling
    const response = await fetch("/api/register", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
      },
      body: JSON.stringify({
        username,
        password,
      }),
    });
    const json = await response.json();
    localStorage.setItem("jwtToken", json.token);
    localStorage.setItem("currentUser", json.user);
    commit("setCurrentUser", json.user);
    router.push("/boards");
  },
};

export default auth;
