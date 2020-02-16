<template>
  <span v-html="stone"></span>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
@Component
export default class Stone extends Vue {
  @Prop() private type!: number;

  get stone(): string {
    if (this.type === 0) {
      return "&nbsp;";
    } else if (this.type === 1) {
      return "&#x26AB";
    } else if (this.type === 2) {
      return "&#x26AA";
    } else {
      return "other";
    }
  }
}
</script>
