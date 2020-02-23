import {
  Module,
  VuexModule,
  Mutation,
  Action,
  getModule
} from "vuex-module-decorators";
import axios from "axios";
import store from "../index";
import { gameDataModule } from "@/store/modules/gameData";

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
  sendIP = "http://192.168.1.22:8888/";
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
  putOK = false;

  @Mutation
  update(next: ReversiDataI) {
    this.turn = next.turn;
    this.score = next.score;
    this.field = next.field;
    this.status = next.status;
  }

  @Mutation
  updateStatus(status: string) {
    this.status = status;
  }

  @Mutation
  updatePutOK(status: boolean) {
    this.putOK = status;
  }

  @Action({})
  async set(updateData: ReversiDataI) {
    try {
      await this.update(updateData);
    } catch (err) {
      console.log(err);
    }

    const sleep = (msec: number) =>
      new Promise(resolve => setTimeout(resolve, msec));

    if (this.status === "skip") {
      await sleep(1500);
      this.skip();
    } else if (this.status !== "finish") {
      if (this.turn % 2 == 0 && gameDataModule.black !== "manual") {
        this.updatePutOK(false);
        await this.cpu(gameDataModule.black);
      } else if (this.turn % 2 == 1 && gameDataModule.white !== "manual") {
        this.updatePutOK(false);
        await this.cpu(gameDataModule.white);
      } else {
        this.updatePutOK(true);
      }
    }
  }

  @Action({})
  async cpu(solver: string) {
    this.updateStatus("cpu");
    try {
      const sendData: ReversiDataI = {
        turn: this.turn,
        score: this.score,
        field: this.field,
        request: { h: 0, w: 0 },
        status: solver
      };
      const res = await axios.post(
        this.sendIP + "cpuRequest",
        JSON.stringify(sendData)
      );
      const retData: ReversiDataI = {
        turn: res.data.turn,
        score: res.data.score,
        field: res.data.field,
        request: res.data.request,
        status: res.data.status
      };
      this.set(retData);
    } catch (err) {
      console.log(err);
    }
  }

  @Action({})
  async send(stonePos: PointI) {
    try {
      const sendData: ReversiDataI = {
        turn: this.turn,
        score: this.score,
        field: this.field,
        request: stonePos,
        status: this.status
      };
      const res = await axios.post(
        this.sendIP + "userRequest",
        JSON.stringify(sendData)
      );
      const retData: ReversiDataI = {
        turn: res.data.turn,
        score: res.data.score,
        field: res.data.field,
        request: res.data.request,
        status: res.data.status
      };
      this.set(retData);
    } catch (err) {
      console.log(err);
    }
  }

  @Action({})
  async skip() {
    try {
      const sendData: ReversiDataI = {
        turn: this.turn,
        score: this.score,
        field: this.field,
        request: { h: -1, w: -1 },
        status: this.status
      };
      const res = await axios.post(
        this.sendIP + "skip",
        JSON.stringify(sendData)
      );
      const retData: ReversiDataI = {
        turn: res.data.turn,
        score: res.data.score,
        field: res.data.field,
        request: res.data.request,
        status: res.data.status
      };
      this.set(retData);
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
