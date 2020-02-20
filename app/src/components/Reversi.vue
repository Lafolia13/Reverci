<template>
  <div class="reversi">
    <p align="center">Turn {{ turn }}: <Stone :color="color" /></p>
    <p align="center">
      <span> <Stone :color="1" />: {{ blackScore }}</span>
      <span><Stone :color="2" />: {{ whiteScore }}</span>
    </p>
    <p align="center">{{ status }}</p>
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
span {
  margin-right: 10px;
}
</style>
