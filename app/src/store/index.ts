import Vue from "vue";
import Vuex from "vuex";
import { ReversiDataI } from "../store/modules/reversiData";
import { SolversI, GameDataI } from "../store/modules/gameData";

Vue.use(Vuex);

export interface RootStateI {
  reverci: ReversiDataI;
  solver: SolversI;
  gameData: GameDataI;
}

const store = new Vuex.Store<RootStateI>({});

export default store;
