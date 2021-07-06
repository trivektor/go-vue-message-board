<template>
  <form @submit.prevent="onSubmit">
    <h1>Register</h1>
    <v-text-field label="Username" v-model="username" />
    <v-text-field label="Password" type="password" v-model="password" />
    <v-btn outlined type="submit" color="primary">Submit</v-btn>
  </form>
</template>

<script>
import { defineComponent } from "@vue/composition-api";

export default defineComponent({
  setup() {},
  data() {
    return { username: "", password: "" };
  },
  methods: {
    async onSubmit() {
      this.$store.dispatch("registerUser", {
        username: this.username,
        password: this.password,
      });
    },
  },
});
</script>
