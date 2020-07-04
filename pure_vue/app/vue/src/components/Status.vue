<template>
  <div class="status">
    <table align="center">
      <tr>
        <td>
          Width:
        </td>
        <td>
          Black:
        </td>
        <td>
          White:
        </td>
        <td>
          <button @click="reset">Reset</button>
        </td>
      </tr>
      <tr>
        <td>
          <select v-model="width">
            <!-- 4 <= x <= 12 -->
            <option v-for="x in 5" :value="2 + 2 * x" :key="'width: ' + x">
              {{ 2 + 2 * x }}
            </option>
          </select>
        </td>
        <td>
          <select v-model="black">
            <option v-for="solver in solvers" :value="solver" :key="solver">{{
              solver
            }}</option>
          </select>
        </td>
        <td>
          <select v-model="white">
            <option v-for="solver in solvers" :value="solver" :key="solver">{{
              solver
            }}</option>
          </select>
        </td>
      </tr>
    </table>
  </div>
</template>

<script lang="ts">
import { Vue, Component, Watch } from "vue-property-decorator";
import { gameDataModule } from "../store/modules/gameData";

@Component({})
export default class Status extends Vue {
  width = 8;
  black = "manual";
  white = "manual";

  get solvers(): string[] {
    return gameDataModule.solver;
  }

  // 盤面の初期化・リサイズ
  reset() {
    gameDataModule.reset();
  }

  created() {
    gameDataModule.getSolver();
    gameDataModule.reset();
  }

  @Watch("width")
  setWidth() {
    gameDataModule.setWidth(this.width);
  }

  @Watch("black")
  setBlack() {
    gameDataModule.setBlack(this.black);
  }

  @Watch("white")
  setWhite() {
    gameDataModule.setWhite(this.white);
  }
}
</script>

<style scoped>
table {
  width: 95%;
  table-layout: fixed;
}
</style>
