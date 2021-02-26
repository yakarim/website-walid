<template>
  <div class="container">
    <div class="row">
      <div class="col-md-4 mx-auto">
        <div class="myform form">
          <form @submit.prevent="submit">
            <div class="form-group">
              <input
                v-model="email"
                type="email"
                name="email"
                class="form-control my-input"
                placeholder="Email"
              />
            </div>
            <br />
            <div class="form-group">
              <input
                v-model="password"
                type="password"
                name="password"
                class="form-control my-input"
                placeholder="Password"
              />
            </div>
            <br />
            <div class="text-center">
              <button type="submit" class="btn btn-primary">
                <div class="lds-ring-container" v-if="loading">
                  <div class="lds-ring">
                    <div></div>
                    <div></div>
                    <div></div>
                    <div></div>
                  </div>
                </div>
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
      <template v-slot="{ params }">
        Hi {{ params.Msg }}
        <button class="modal-close" @click="showModal = false">x</button>
      </template>
    </vue-final-modal>
  </div>
</template>

<script>
import { inject } from "vue";

export default {
  setup() {
    const $vfm = inject("$vfm");
    return {
      $vfm: $vfm,
    };
  },
  data() {
    return {
      email: "",
      password: "",
      msg: "",
      showModal: false,
      loading: false,
    };
  },
  components: {},
  methods: {
    submit() {
      this.loading = true;
      this.$store
        .dispatch('signIn', {
          email: this.email,
          password: this.password,
        })
        .then(() => {
          this.loading = false;
          this.$router.push({ path: "/admin" });
        })
        .catch((err) => {
          this.loading = false
          this.msg = err.response.data.msg;
          this.$vfm.show("alert", { Msg: this.msg });
        });
    },
  },
};
</script>

<style lang='scss'>
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

// CSS Spinner
.lds-ring-container {
  position: absolute;
  right: 50%;
}
.lds-ring {
  display: inline-block;
  position: relative;
  width: 64px;
  height: 64px;
}
.lds-ring div {
  box-sizing: border-box;
  display: block;
  position: absolute;
  width: 25px;
  height: 25px;
  // margin: 6px;
  border: 3px solid #fff;
  border-radius: 50%;
  animation: lds-ring 1.2s cubic-bezier(0.5, 0, 0.5, 1) infinite;
  border-color: #fff transparent transparent transparent;
}
.lds-ring div:nth-child(1) {
  animation-delay: -0.45s;
}
.lds-ring div:nth-child(2) {
  animation-delay: -0.3s;
}
.lds-ring div:nth-child(3) {
  animation-delay: -0.15s;
}
@keyframes lds-ring {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}
</style>
