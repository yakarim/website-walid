<template>
  <div id="app">
    <div class="container">
      <ul class="nav">
        <li><router-link :to="{ name: 'home' }">Home</router-link></li>
        <li v-if="loggedIn">
          <router-link :to="{ name: 'about' }">About</router-link>
        </li>
        <li v-if="loggedIn">
          <router-link :to="{ name: 'admin' }">Admin</router-link>
        </li>
        <li v-if="!loggedIn">
          <router-link :to="{ name: 'login' }">Login</router-link>
        </li>
        <li v-if="!loggedIn">
          <router-link :to="{ name: 'signup' }">Signup</router-link>
        </li>
        <li v-if="loggedIn">
          <router-link :to="{ name: 'logout' }">Logout</router-link>
        </li>
      </ul>
    </div>

    <div class="container">
      <transition
        name="router-animation"
        enter-active-class="animated fadeIn"
        leave-active-class="animated fadeOut"
        mode="out-in"
      >
        <router-view></router-view>
      </transition>
    </div>
  </div>
</template>

<script>
export default {
  name: "App",
  computed: {
    loggedIn() {
      return this.$store.getters.loggedIn;
    },
  },
  components: {},
};
</script>

<style lang="scss">
@import "~bootstrap/dist/css/bootstrap.css";
@import url("https://cdnjs.cloudflare.com/ajax/libs/animate.css/3.5.2/animate.min.css");

* {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  font-size: 24px;
  height: 100vh;
}
.flex-center {
  display: flex;
  justify-content: center;
}
.nav {
  display: flex;
  list-style: none;
  padding: 15px 0;
  margin: 0;
  justify-content: flex-end;
  background: #f5f8fa;
  border-bottom: 1px solid lightgrey;
  margin-bottom: 24px;
}
.nav a {
  color: #636b6f;
  padding: 0 25px;
  font-size: 14px;
  font-weight: 600;
  letter-spacing: 0.1rem;
  text-decoration: none;
  text-transform: uppercase;
}
// Auth Pages
label {
  display: block;
  margin-bottom: 4px;
}
.login-heading {
  margin-bottom: 16px;
}
.form-control {
  margin-bottom: 24px;
}
.mb-more {
  margin-bottom: 42px;
}
.login-form {
  max-width: 500px;
  margin: auto;
}
.login-input {
  width: 100%;
  font-size: 16px;
  padding: 12px 16px;
  outline: 0;
  border-radius: 3px;
  border: 1px solid lightgrey;
}
.btn-submit {
  width: 100%;
  padding: 14px 12px;
  font-size: 18px;
  font-weight: bold;
  background: #60bd4f;
  color: white;
  border-radius: 3px;
  cursor: pointer;
  &:hover {
    background: darken(#60bd4f, 10%);
  }
}
</style>

