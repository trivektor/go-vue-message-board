<template>
  <boards />
</template>

<script>
import { defineComponent } from "@vue/composition-api";
import Boards from "../components/Boards.vue";

export default defineComponent({
  components: {
    Boards,
  },
});
</script>
