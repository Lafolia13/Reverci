import {
  Module,
  VuexModule,
  Mutation,
  Action,
  getModule
} from "vuex-module-decorators";
import axios from "axios";
import store from "@/store/index";

export interface PositionI {
  h: number;
  w: number;
}

export interface ReversiDataI {
  turn: number;
  field: number[][];
  request: {
    h: number;
    w: number;
  };
  status: string;
}

@Module({ dynamic: true, store: store, name: "reversi", namespaced: true })
class ReversiData extends VuexModule {
  turn = 0;
  field: number[][] = [
    [0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 3, 0, 0, 0],
    [0, 0, 0, 1, 2, 3, 0, 0],
    [0, 0, 3, 2, 1, 0, 0, 0],
    [0, 0, 0, 3, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0]
  ];
  request: PositionI = {
    h: 0,
    w: 0
  };
  status = "none";

  @Mutation
  update(next: ReversiDataI) {
    this.turn = next.turn;
    this.field = next.field;
    this.status = next.status;
  }

  @Action({})
  async send(stonePos: PositionI) {
    try {
      const sendData: ReversiDataI = {
        turn: this.turn,
        field: this.field,
        request: stonePos,
        status: this.status
      };
      const res = await axios.post(
        "http://localhost:8888/userRequest",
        JSON.stringify(sendData)
      );
      const retData: ReversiDataI = {
        turn: res.data.turn,
        field: res.data.field,
        request: res.data.request,
        status: res.data.status
      };
      this.update(retData);
    } catch (err) {
      console.log(err);
    }
  }

  get getStone() {
    return (pos: PositionI) => {
      return this.field[pos.h][pos.w];
    };
  }
}

export const reversiModule = getModule(ReversiData);
