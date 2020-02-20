<template>
  <div class="home">
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
            <option value="4">4</option>
            <option value="6">6</option>
            <option value="8">8</option>
            <option value="10">10</option>
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
        <td>
          <button @click="getSolver">Get solver</button>
        </td>
      </tr>
    </table>

    <Reversi />
  </div>
</template>

<script lang="ts">
import { Vue, Component, Watch } from "vue-property-decorator";
import Reversi from "@/components/Reversi.vue";
import { gameDataModule } from "../store/modules/gameData";

@Component({
  components: {
    Reversi
  }
})
export default class Home extends Vue {
  width = 8;
  black = "manual";
  white = "manual";

  get solvers(): string[] {
    return gameDataModule.solver;
  }

  reset() {
    gameDataModule.reset();
  }

  getSolver() {
    gameDataModule.getSolver();
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
  width: 400px;
}
/* td {
  width: calc(100% / 3);
} */
</style>
