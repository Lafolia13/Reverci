import Vue from "vue";
import VueRouter from "vue-router";
import BootstrapVue from "bootstrap-vue";
import Home from "../views/Home.vue";

Vue.use(VueRouter);
Vue.use(BootstrapVue);

import "bootstrap/dist/css/bootstrap.css";
import "bootstrap-vue/dist/bootstrap-vue.css";

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home
  },
  {
    path: "*",
    redirect: {
      name: "Home"
    }
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;
