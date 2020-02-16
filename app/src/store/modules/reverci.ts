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

export interface ReverciDataI {
  turn: number;
  field: number[][];
  request: {
    h: number;
    w: number;
  };
  status: string;
}

@Module({ dynamic: true, store: store, name: "reverci", namespaced: true })
class ReverciData extends VuexModule {
  turn = 0;
  field: number[][] = [
    [0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 1, 2, 0, 0, 0],
    [0, 0, 0, 2, 1, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0]
  ];
  request: PositionI = {
    h: 0,
    w: 0
  };
  status = "none";

  @Mutation
  update(next: ReverciDataI) {
    this.turn = next.turn;
    this.field = next.field;
    this.request = next.request;
    this.status = next.status;
  }

  @Action({})
  async send(stonePos: PositionI) {
    try {
      const sendData: ReverciDataI = {
        turn: this.turn,
        field: this.field,
        request: stonePos,
        status: this.status
      };
      const res = await axios.post(
        "http://localhost:8888",
        JSON.stringify(sendData)
      );
      const retData: ReverciDataI = {
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

  get nowTurn(): number {
    return this.turn;
  }

  get nowStatus(): string {
    return this.status;
  }

  get nowStone() {
    return (pos: PositionI) => {
      return this.field[pos.h][pos.w];
    };
  }
}

export const reverciModule = getModule(ReverciData);
