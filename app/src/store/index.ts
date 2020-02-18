import Vue from "vue";
import Vuex from "vuex";
import { ReversiDataI } from "@/store/modules/reversi";

Vue.use(Vuex);

export interface RootStateI {
  reverci: ReversiDataI;
}

const store = new Vuex.Store<RootStateI>({});

export default store;
