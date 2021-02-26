<template>
  <div class="container">
    <p>Welcome admin {{username}}!</p>
    <table class="table">
      <thead>
        <tr>
          <th scope="col">#</th>
          <th scope="col">Username</th>
          <th scope="col">Email</th>
          <th scope="col">Handle</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(u, index) in users" :key="u.ID">
          <th scope="row">{{ index + 1 }}</th>
          <td>{{ u.username }}</td>
          <td>{{ u.email }}</td>
          <td>
            <button
              @click="handleDelete(u.ID)"
              type="button"
              class="btn btn-danger"
            >
              Delete
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import axios from '../services/myservices'
import token from "../services/auth-header";

export default {
  data() {
    return {
      username: "",
      users: null,
      msg: "",
      showModal: false,
    };
  },
  mounted() {
    this.handleSubmit();
    this.handleOne();
  },
  components: {},
  methods: {
    handleOne() {
      const _id = localStorage.getItem("_id");
      axios
        .get("/user/query/" + _id, {
          headers: token(),
        })
        .then((res) => {
            this.username = res.data.user.username
          //console.log(res.data.user.username);
        })
        .catch((err) => {
          console.log(err);
        });
    },
    handleSubmit() {
      axios
        .get("/user/query", {
          headers: token(),
        })
        .then((res) => {
          this.users = res.data.users;
        })
        .catch((err) => {
          console.log(err);
        });
    },
    handleDelete(key) {
      axios
        .delete("/user/delete/" + key, {
          headers: token(),
        })
        .then((res) => {
          console.log(res);
          this.handleSubmit();
        })
        .catch((err) => {
          console.log(err);
        });
    },
  },
};
</script>
