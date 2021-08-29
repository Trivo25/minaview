<template>
  <div class="trending-card">
    
  </div>
</template>

<script>
export default {
  name: "TrendingCard"
}
</script>

<style scoped>
.trending-card {
  position: absolute;
  display: inline;
  z-index: 15;
  left: calc(50% - 175px);
  top: 10px;
  width: 500px;
  height: 200px;
  border-radius: 5px;
  border: 1px solid rgba(6, 170, 241, 0.46);
  background-color: rgba(23, 23, 23, 0.3);
  box-shadow: 0px 0px 2px #0396ff;
  color: #0396ff;
  text-shadow: 2px 1px #00040a;
  font-weight: 400;
  transition:0.5s;
  align-content: center;
  margin-bottom: 0;
  justify-items: center;
  cursor: default;
}
</style>