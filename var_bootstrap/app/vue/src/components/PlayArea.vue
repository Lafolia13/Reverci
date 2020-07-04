<template>
  <div class="playArea">
    <table class="table mx-auto mb-5" ref="table-reversi">
      <tr v-for="h in width" :key="'h:' + h">
        <td
          class="bg-success border border-dark align-middle p-0"
          height="0px"
          v-for="w in width"
          :key="'w:' + w"
          @click="putStone(h - 1, w - 1)"
        >
          <div class="text-center">
            <Stone :color="stoneColor(h - 1, w - 1)" />
          </div>
        </td>
      </tr>
    </table>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
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
    if (reversiDataModule.putOK == false) return;
    const stonePos: PointI = {
      h: h,
      w: w
    };
    reversiDataModule.send(stonePos);
  }
}
</script>
