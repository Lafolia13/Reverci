<template>
  <div class="home">
    <PlayArea />
  </div>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator";
import PlayArea from "@/components/PlayArea.vue";

@Component({
  components: {
    PlayArea
  }
})
export default class Home extends Vue {}
</script>
