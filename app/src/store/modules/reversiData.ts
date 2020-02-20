import {
  Module,
  VuexModule,
  Mutation,
  Action,
  getModule
} from "vuex-module-decorators";
import axios from "axios";
import store from "../index";

export interface PointI {
  h: number;
  w: number;
}

export interface ScoresI {
  black: number;
  white: number;
}

export interface ReversiDataI {
  turn: number;
  score: ScoresI;
  field: number[][];
  request: PointI;
  status: string;
}

@Module({ dynamic: true, store: store, name: "reversiData", namespaced: true })
class ReversiData extends VuexModule {
  turn = 0;
  score: ScoresI = {
    black: 2,
    white: 2
  };
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
  request: PointI = {
    h: 0,
    w: 0
  };
  status = "none";

  @Mutation
  update(next: ReversiDataI) {
    this.turn = next.turn;
    this.score = next.score;
    this.field = next.field;
    this.status = next.status;
  }

  @Action({})
  async updateAction(updateData: ReversiDataI) {
    try {
      await this.update(updateData);
    } catch (err) {
      console.log(err);
    }

    const sleep = (msec: number) =>
      new Promise(resolve => setTimeout(resolve, msec));

    if (this.status === "skip") {
      await sleep(2000);
      this.skip();
    }
  }

  @Action({})
  async send(stonePos: PointI) {
    try {
      const sendIP = "http://192.168.1.22:8888/userRequest";
      const sendData: ReversiDataI = {
        turn: this.turn,
        score: this.score,
        field: this.field,
        request: stonePos,
        status: this.status
      };
      const res = await axios.post(sendIP, JSON.stringify(sendData));
      const retData: ReversiDataI = {
        turn: res.data.turn,
        score: res.data.score,
        field: res.data.field,
        request: res.data.request,
        status: res.data.status
      };
      this.updateAction(retData);
    } catch (err) {
      console.log(err);
    }
  }

  @Action({})
  async skip() {
    try {
      const sendIP = "http://192.168.1.22:8888/skip";
      const sendData: ReversiDataI = {
        turn: this.turn,
        score: this.score,
        field: this.field,
        request: { h: -1, w: -1 },
        status: this.status
      };
      const res = await axios.post(sendIP, JSON.stringify(sendData));
      const retData: ReversiDataI = {
        turn: res.data.turn,
        score: res.data.score,
        field: res.data.field,
        request: res.data.request,
        status: res.data.status
      };
      this.updateAction(retData);
    } catch (err) {
      console.log(err);
    }
  }

  get getStone() {
    return (pos: PointI) => {
      return this.field[pos.h][pos.w];
    };
  }
}

export const reversiDataModule = getModule(ReversiData);
