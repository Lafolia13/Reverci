<template>
  <div class="home">
    <Status />
    <Reversi />
  </div>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator";
import Reversi from "@/components/Reversi.vue";
import Status from "@/components/Status.vue";

@Component({
  components: {
    Reversi,
    Status
  }
})
export default class Home extends Vue {}
</script>
