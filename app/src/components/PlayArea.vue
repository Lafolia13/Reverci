<template>
  <div class="playArea">
    <table align="center">
      <tr v-for="h in 8" :key="'h:' + h">
        <td v-for="w in 8" :key="'w:' + w" @click="putStone(h - 1, w - 1)">
          <Stone :color="stoneColor(h - 1, w - 1)" />
        </td>
      </tr>
    </table>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import Stone from "@/components/Stone.vue";
import { reversiModule, PositionI } from "../store/modules/reversi";

@Component({
  components: {
    Stone
  }
})
export default class PlayArea extends Vue {
  get stoneColor() {
    return (h: number, w: number) => {
      const stonePos: PositionI = {
        h: h,
        w: w
      };
      return reversiModule.getStone(stonePos);
    };
  }

  putStone(h: number, w: number) {
    const stonePos: PositionI = {
      h: h,
      w: w
    };
    reversiModule.send(stonePos);
  }
}
</script>

<style scoped>
table {
  border-collapse: collapse;
}
td {
  border: 3px solid #000000;
  background: #42b983;
  width: 50px;
  height: 50px;
  font-size: 30px;
  text-align: center;
  vertical-align: middle;
}
</style>
