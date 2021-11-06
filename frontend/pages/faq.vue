<template>
  <div class="faq">
    <h1 class="faq-title">Frequently Asked Questions</h1>
    <Drawer/>
  </div>
</template>

<script>
import Drawer from "../components/Drawer.vue"

export default {
  name: "faq",
  layout: "faq-layout",
  props: [],
  components: {
    Drawer
  },
}
</script>


<style scoped>

.faq {
  justify-content: center;
  text-align: center;
  width: 100%;
  height: 600px !important;
}

.faq-title {
  font-size: 1.5rem;
  font-weight: 400;
  color: grey;

}

</style>