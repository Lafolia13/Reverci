<template>
  <div class="home">
    <Reversi />
  </div>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator";
import Reversi from "@/components/Reversi.vue";

@Component({
  components: {
    Reversi
  }
})
export default class Home extends Vue {}
</script>
