<template>
  <login-component />
</template>

<script>
import { defineComponent } from "@vue/composition-api";
import LoginComponent from "../components/Login.vue";

export default defineComponent({
  name: "LoginView",
  components: { LoginComponent },
});
</script>
