<template>
  <div class="container">
    <div class="row">
      <div class="col-md-4 mx-auto">
        <div class="myform form">
          <form @submit.prevent="handleSubmit">
            <div class="form-group">
              <input
                v-model="username"
                type="text"
                name="name"
                class="form-control my-input"
                id="name"
                placeholder="Name"
              />
            </div>
            <br />
            <div class="form-group">
              <input
                v-model="email"
                type="email"
                name="email"
                class="form-control my-input"
                id="email"
                placeholder="Email"
              />
            </div>
            <br />
            <div class="form-group">
              <input
                v-model="password"
                type="password"
                min="4"
                name="password"
                id="password"
                class="form-control my-input"
                placeholder="Password"
              />
            </div>
            <br />
            <div class="text-center">
              <button type="submit" class="btn btn-primary">
                Create Your Free Account
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>

     <vue-final-modal
      v-model="showModal"
      name="alert"
      classes="modal-container"
      content-class="modal-content"
    >
      <template v-slot="{ params }"> Hi {{ params.Msg }} 
      <button class="modal-close" @click="showModal = false">x</button>
      </template>
    </vue-final-modal>
  </div>
</template>



<script>
import axios from "../services/myservices";
import { mapState } from "vuex";

export default {
  name: "Signup",
  data() {
    return {
      username: "",
      email: "",
      password: "",
      msg: "",
      showModal: false,
    };
  },
   computed: mapState([
    'http'
  ]),
  components: {},
  methods: {
    handleSubmit() {
      const data = {
        username: this.username,
        email: this.email,
        password: this.password,
      };

      axios
        .post("./user/create", data, {
          "Content-Type": "application/x-www-form-urlencoded",
        })
        .then((res) => {
          this.msg = res.data.msg;
          this.$vfm.show("alert", { Msg: this.msg });
        })
        .catch((err) => {
          this.msg = err.response.data.msg;
          this.$vfm.show("alert", { Msg: this.msg });
        });
    },
  },
};
</script>

<style>
.modal-container {
  display: flex;
  justify-content: center;
  align-items: center;
}
.modal-content {
  position: relative;
  width: 50%;
  max-height: 300px;
  padding: 16px;
  overflow: auto;
  background-color: #fff;
  border-radius: 4px;
}
.modal-close {
  position: absolute;
  top: 0;
  right: 0;
  display: flex;
  justify-content: center;
  align-items: center;
  width: 32px;
  height: 32px;
  margin: 8px 8px 0 0;
  cursor: pointer;
}
.modal-close::hover {
  color: gray;
}
</style>
