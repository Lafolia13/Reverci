<template>
  <img width="100%" height="100%" :src="stone" />
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
@Component
export default class Stone extends Vue {
  @Prop() private color!: number;

  get stone(): any {
    if (this.color === 0) {
      return require("../assets/images/none.png");
    } else if (this.color === 1) {
      return require("@/assets/images/black.png");
    } else if (this.color === 2) {
      return require("@/assets/images/white.png");
    } else {
      return require("@/assets/images/able.png");
    }
  }
}
</script>

<style scoped>
/* 文字を選択できないようにする */
img {
  user-select: none;
  -moz-user-select: none;
  -webkit-user-select: none;
  -ms-user-select: none;
}
</style>
