<template>
  <p>Welcome About!</p>
</template>

<script>
import axios from "axios";
import token from "../services/auth-header"

export default {
  data() {
    return {
      username: "",
      email: "",
      password: "",
      msg: "",
      showModal: false,
    };
  },
  mounted() {
    //this.handleSubmit();
  },
  components: {},
  methods: {
    handleSubmit() {
      axios
        .get("http://localhost:3000/user/query", {
          headers : token()
        })
        .then((res) => {
          console.log(res);
        })
        .catch((err) => {
          console.log(err);
        });
    },
  },
};
</script>