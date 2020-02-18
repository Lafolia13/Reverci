<template>
  <span v-html="stone"></span>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
@Component
export default class Stone extends Vue {
  @Prop() private color!: number;

  get stone(): string {
    if (this.color === 0) {
      return "&nbsp;"; // none
    } else if (this.color === 1) {
      return "&#x26AB"; // black
    } else if (this.color === 2) {
      return "&#x26AA"; // white
    } else {
      return "&#x1f914;"; // :thinking_face:
    }
  }
}
</script>

<style scoped>
span {
  user-select: none;
  -moz-user-select: none;
  -webkit-user-select: none;
  -ms-user-select: none;
}
</style>
