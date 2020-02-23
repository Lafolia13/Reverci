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
            <option v-for="x in 9" :value="2 + 2 * x" :key="'width: ' + x">
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
