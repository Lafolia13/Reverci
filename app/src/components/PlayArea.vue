<template>
  <div class="playArea">
    <table align="center">
      <tr v-for="h in width" :key="'h:' + h">
        <td v-for="w in width" :key="'w:' + w" @click="putStone(h - 1, w - 1)">
          <Stone :color="stoneColor(h - 1, w - 1)" />
        </td>
      </tr>
    </table>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from "vue-property-decorator";
import Stone from "@/components/Stone.vue";
import { reversiDataModule, PointI } from "../store/modules/reversiData";

@Component({
  components: {
    Stone
  }
})
export default class PlayArea extends Vue {
  get width(): number {
    return reversiDataModule.field.length;
  }

  get stoneColor() {
    return (h: number, w: number) => {
      const stonePos: PointI = {
        h: h,
        w: w
      };
      return reversiDataModule.getStone(stonePos);
    };
  }

  get status(): string {
    return reversiDataModule.status;
  }

  putStone(h: number, w: number) {
    const stonePos: PointI = {
      h: h,
      w: w
    };
    reversiDataModule.send(stonePos);
  }
}
</script>

<style scoped>
div {
  overflow: auto;
}
table {
  border-collapse: collapse;
  table-layout: fixed;
}
td {
  border: 3px solid #000000;
  background: #42b983;
  width: 50px;
  height: width;
  font-size: 30px;
  text-align: center;
  vertical-align: middle;
}
</style>
