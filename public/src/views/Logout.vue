<template>
  <p>Welcome home!</p>
</template>


<script>
export default {
  created() {
    this.$store.dispatch("destroyToken")
     .then(() => {
          this.$router.push({ path: "/" });
        })
  },
};
</script>