<template>
  <div class="project-wrapper">
    <div class="header">
      <div class="logo">
        <img :src="service.ServiceLogo" alt="Project Logo">
      </div>
      <div class="title">
        {{ service.ServiceName }}
      </div>
    </div>
    <div class="socials">
      scials
    </div>
    <div class="content">
      content
    </div>
    <Error @closeNotification="error.show = false" v-if="error.show" :error="error.error" :type="error.type" />
  </div>
</template>

<script>
import Error from "../components/Error.vue"

export default {
  name: "project",
  layout: "project-page",
  props: [],
  components: {
    Error
  },
  data () {
    return {
      service: new Object,
      error: {
        show: false,
        error: null,
        type: "error"
      }
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

        this.service = res.data.Data
    } catch (error) {
      console.log(error)
    }
  }
}
</script>

<style scoped>

</style>