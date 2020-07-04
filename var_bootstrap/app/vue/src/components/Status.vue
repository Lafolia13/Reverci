<template>
  <div class="status">
    <div class="container pt-4">
      <div class="form-row">
        <div class="form-group col-6">
          <label class="col-5 col-form-label" for="black">Black:</label>
          <select class="custom-select col-7" id="black" v-model="black">
            <option v-for="solver in solvers" :value="solver" :key="solver">
              {{ solver }}
            </option>
          </select>
        </div>
        <div class="form-group col-6">
          <label class="col-5 col-form-label" for="white">White:</label>
          <select class="custom-select col-7" id="white" v-model="white">
            <option v-for="solver in solvers" :value="solver" :key="solver">
              {{ solver }}
            </option>
          </select>
        </div>
      </div>

      <div class="form-row">
        <div class="form-group col-6">
          <label class="col-5 col-form-label" for="width">Width:</label>
          <select class="custom-select col-7" id="width" v-model="width">
            <option v-for="x in 5" :value="2 + 2 * x" :key="'width: ' + x">
              {{ 2 + 2 * x }}
            </option>
          </select>
        </div>

        <div class="col">
          <div class="text-center">
            <button class="btn btn-light border" @click="reset">
              Reset
            </button>
          </div>
        </div>
      </div>
    </div>
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
