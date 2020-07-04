<template>
  <div class="container">
    <div class="row">
      <div class="col-6">
        <div class="text-center">
          Turn {{ turn }} :
          <Stone width="16px" height="16px" :color="color" />
        </div>
      </div>
      <div class="col-6">
        <div class="text-center">Status: {{ status }}</div>
      </div>
    </div>

    <div class="row">
      <div class="col-6">
        <div class="text-center">
          <Stone width="16px" height="16px" :color="1" />: {{ blackScore }}
        </div>
      </div>
      <div class="col-6">
        <div class="text-center">
          <Stone width="16px" height="16px" :color="2" />: {{ whiteScore }}
        </div>
      </div>
    </div>

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
