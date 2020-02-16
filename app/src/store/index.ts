import Vue from "vue";
import Vuex from "vuex";
import { ReverciDataI } from "@/store/modules/reverci";

Vue.use(Vuex);

export interface RootStateI {
  reverci: ReverciDataI;
}

const store = new Vuex.Store<RootStateI>({});

export default store;
