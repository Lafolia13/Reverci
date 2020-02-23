<template>
  <div class="reversi">
    <table align="center">
      <tr>
        <td>Turn {{ turn }}: <Stone :color="color" /></td>
      </tr>
      <tr>
        <td><Stone :color="1" />: {{ blackScore }}</td>
        <td><Stone :color="2" />: {{ whiteScore }}</td>
        <td>
          <span> Status: {{ status }} </span>
        </td>
      </tr>
    </table>

    <PlayArea />
  </div>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator";
import PlayArea from "@/components/PlayArea.vue";
import Stone from "@/components/Stone.vue";
import { reversiDataModule } from "../store/modules/reversiData";

@Component({
  components: {
    PlayArea,
    Stone
  }
})
export default class Reversi extends Vue {
  get turn(): number {
    return reversiDataModule.turn;
  }
  get color(): number {
    return (reversiDataModule.turn % 2) + 1;
  }
  get blackScore(): number {
    return reversiDataModule.score.black;
  }
  get whiteScore(): number {
    return reversiDataModule.score.white;
  }
  get status(): string {
    return reversiDataModule.status;
  }
}
</script>

<style scoped>
table {
  width: 90%;
}
td {
  width: calc(90% / 3);
}
</style>
