import {
  Module,
  VuexModule,
  Mutation,
  Action,
  getModule
} from "vuex-module-decorators";
import store from "../index";
import { ReversiDataI, reversiDataModule } from "@/store/modules/reversiData";
import axios from "axios";

// export interface SolversI {
//   solver: string[];
// }

export interface GameDataI {
  width: number;
  black: string;
  white: string;
}

@Module({ dynamic: true, store: store, name: "gameData", namespaced: true })
class GameData extends VuexModule {
  sendIP = "http://192.168.1.22:8888/";
  width = 8;
  black = "manual";
  white = "manual";
  solver: string[] = ["manual"];

  @Mutation
  updateSolvers(solver: string[]) {
    this.solver = solver;
  }

  @Mutation
  updateWidth(width: number) {
    this.width = width;
  }

  @Mutation
  updateBlack(solver: string) {
    this.black = solver;
  }

  @Mutation
  updateWhite(solver: string) {
    this.white = solver;
  }

  @Action({})
  setWidth(width: number) {
    this.updateWidth(width);
  }

  @Action({})
  setBlack(solver: string) {
    this.updateBlack(solver);
  }

  @Action({})
  setWhite(solver: string) {
    this.updateWhite(solver);
  }

  @Action({})
  async reset() {
    const retData: ReversiDataI = {
      turn: 0,
      score: { black: 2, white: 2 },
      field: Array<number[]>(this.width),
      request: { h: 0, w: 0 },
      status: "start"
    };
    for (let h = 0; h < this.width; h++) {
      retData.field[h] = Array<number>(this.width);
      for (let w = 0; w < this.width; w++) {
        retData.field[h][w] = 0;
      }
    }
    retData.field[this.width / 2 - 1][this.width / 2 - 1] = 1;
    retData.field[this.width / 2][this.width / 2] = 1;
    retData.field[this.width / 2 - 1][this.width / 2] = 2;
    retData.field[this.width / 2][this.width / 2 - 1] = 2;
    retData.field[this.width / 2 - 2][this.width / 2] = 3;
    retData.field[this.width / 2 - 1][this.width / 2 + 1] = 3;
    retData.field[this.width / 2][this.width / 2 - 2] = 3;
    retData.field[this.width / 2 + 1][this.width / 2 - 1] = 3;

    await reversiDataModule.set(retData);
  }

  @Action({})
  async getSolver() {
    try {
      const res = await axios.post(this.sendIP + "solverRequest");
      this.updateSolvers(res.data.solver);
    } catch (err) {
      console.log(err);
    }
  }
}

export const gameDataModule = getModule(GameData);
