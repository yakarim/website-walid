import { createWebHistory, createRouter } from "vue-router";
import { store } from '../stores'

import Home from "@/views/Home.vue";
import About from "@/views/About.vue";
import Login from "@/views/Login.vue";
import Signup from "@/views/Signup.vue";
import Logout from "@/views/Logout.vue";
import Admin from "@/views/Admin";

const routes = [
  {
    path: "/",
    name: "home",
    component: Home,
  },
  {
    path: "/about",
    name: "about",
    component: About,
    meta: {
      requiresAuth: true,
    }
  },
  {
    path: "/admin",
    name: "admin",
    component: Admin,
    meta: {
      requiresAuth: true,
    }
  },
  {
    path: "/login",
    name: "login",
    component: Login,
    props: true,
    meta: {
      requiresVisitor: true,
    }
  },
  {
    path: "/signup",
    name: "signup",
    component: Signup,
    meta: {
      requiresVisitor: true,
    }
  },
  {
    path: "/logout",
    name: "logout",
    component: Logout,
  },
  /*
  {
    path: "*",
    name: "Home",
    component: Home,
  },
  */
];




const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (!store.getters.loggedIn) {
      next({
        name: 'login',
      })
    } else {
      next()
    }
  } else if (to.matched.some(record => record.meta.requiresVisitor)) {
    if (store.getters.loggedIn) {
      next({
        name: 'about',
      })
    } else {
      next()
    }
  } else {
    next()
  }
})

export default router;