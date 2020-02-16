<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <h1>{{ getStatus }}</h1>
    <table align="center">
      <tr v-for="h in 8" :key="'h:' + h">
        <td v-for="w in 8" :key="'w:' + w" @click="put(h - 1, w - 1)">
          <Stone :type="getStone(h - 1, w - 1)" />
        </td>
      </tr>
    </table>

    <button @click="sendData">send</button>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import Stone from "@/components/Stone.vue";
import { reverciModule, PositionI } from "../store/modules/reverci";

@Component({
  components: {
    Stone
  }
})
export default class HelloWorld extends Vue {
  @Prop() private msg!: string;

  get getStatus(): string {
    return reverciModule.nowStatus;
  }

  get getStone() {
    return (h: number, w: number) => {
      const pos: PositionI = {
        h: h,
        w: w
      };
      return reverciModule.nowStone(pos);
    };
  }

  sendData() {
    reverciModule.send();
  }

  put(h: number, w: number) {
    const stonePos: PositionI = {
      h: h,
      w: w
    };
    reverciModule.send(stonePos);
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
table {
  border-collapse: collapse;
}
td {
  border: 1px solid #999;
  background: #f0f0f0;
  width: 30px;
  height: 30px;
  text-align: center;
  vertical-align: middle;
}
</style>
