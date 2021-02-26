<template>
  <input type="text" v-model="value1" />
  <span>{{ operator }}</span>
  <input type="text" v-model="value2" />
  =
  <input type="text" v-model="calcResult" disabled />
  <hr />
  <button @click="add()" id="addButton">Add</button>
</template>

<script>
export default {
  inject: ["$go"],
  props: {
    msg: String,
  },
  data() {
    return {
      operator: "?",
      value1: "",
      value2: "",
      calcResult: "",
    };
  },
  mounted() {
    console.log(this.$go.satu);
  },
  components: {},
  methods: {
    add() {
      let values = [this.value1, this.value2];
      this.$go.add(...values).then((res) => {
        this.calcResult = res;
      });
      this.operator = "+";
    },
  },
};
</script>