<template>
  <div  style="left: 500px;">{{getServiceHash()}}</div>
</template>

<script>

export default {
  name: "project",
  layout: "project-page",
  props: [],
  components: {
  },
  data () {
    return {
    }
  },
  methods: {
    getServiceHash() {
      return this.$route.query.prj
    }
  },
  computed: {
  },
  async mounted() {
    try {
        let res = await this.$axios.post("/getServiceByHash", {
          Hash: this.getServiceHash(),
        })
        if(res.data.Error != undefined) {
          this.error.error = {
            Error: res.data.Error,
            Code: ""
          }
          this.error.show = true
          this.error.type = "error"
        } else {
          this.error.error = {
            Error: "Your request has been sent!",
            Code: ""
          }
          this.error.show = true
          this.error.type = "success"
        }
    } catch (error) {
      
    }
  }
}
</script>

<style scoped>

</style>