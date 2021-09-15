<template>
  <v-app dark>
    <v-app-bar :clipped-left="clipped" fixed app>
      <v-btn to="/ecosystem" class="back-button">
        <v-icon>mdi-menu-left</v-icon>
        BACK
      </v-btn>
      <!-- <img @click="goTo('/')" class="mina-logo" src="../assets/mina_logo_large.svg"/> -->
      <!-- <h1 class="title">{{ title }}</h1> -->
      <v-spacer />
      <v-btn
        outlined
        plain
        to="/request"
        class="add-service-button"
      >
        <v-icon>mdi-plus</v-icon>
        Add your Service
      </v-btn>
    </v-app-bar>
    <v-main>
      <v-container>
        <Nuxt />
      </v-container>
    </v-main>
  </v-app>
</template>

<script>

export default {
  name: "project-page",
  layout: "project-page",
  props: [],
  data () {
    return {

    }
  },
  methods: {
  },
  computed: {
  },
  mounted() {
    
  },
}
</script>

<style scoped>

* {
  font-family: "Roboto";
}  

.v-app-bar, .v-navigation-drawer, .v-list, .v-list-item, .v-footer {
  box-shadow: none !important;
  border-bottom: none !important;
  background: #00000000 !important;
  border: none !important;
  box-shadow: none !important;
}

.v-footer {
  bottom: 0px !important;
}

.theme--dark.v-application {
  background-image: url("data:image/svg+xml,%3csvg xmlns='http://www.w3.org/2000/svg' width='auto' height='auto' version='1.1' xmlns:xlink='http://www.w3.org/1999/xlink' xmlns:svgjs='http://svgjs.com/svgjs'  preserveAspectRatio='none' viewBox='0 0 1440 560'%3e%3cg mask='url(%26quot%3b%23SvgjsMask1036%26quot%3b)' fill='none'%3e%3crect width='1440' height='560' x='0' y='0' fill='url(%23SvgjsLinearGradient1037)'%3e%3c/rect%3e%3cpath d='M1440 0L1316.9 0L1440 212.61z' fill='rgba(255%2c 255%2c 255%2c .1)'%3e%3c/path%3e%3cpath d='M1316.9 0L1440 212.61L1440 232.43L1189.3700000000001 0z' fill='rgba(255%2c 255%2c 255%2c .075)'%3e%3c/path%3e%3cpath d='M1189.37 0L1440 232.43L1440 377.48L450.03999999999985 0z' fill='rgba(255%2c 255%2c 255%2c .05)'%3e%3c/path%3e%3cpath d='M450.03999999999996 0L1440 377.48L1440 467.53000000000003L358.01 0z' fill='rgba(255%2c 255%2c 255%2c .025)'%3e%3c/path%3e%3cpath d='M0 560L480 560L0 515.35z' fill='rgba(0%2c 0%2c 0%2c .1)'%3e%3c/path%3e%3cpath d='M0 515.35L480 560L696.25 560L0 294.96000000000004z' fill='rgba(0%2c 0%2c 0%2c .075)'%3e%3c/path%3e%3cpath d='M0 294.96000000000004L696.25 560L1031.63 560L0 218.33000000000004z' fill='rgba(0%2c 0%2c 0%2c .05)'%3e%3c/path%3e%3cpath d='M0 218.33000000000004L1031.63 560L1273.52 560L0 155.83000000000004z' fill='rgba(0%2c 0%2c 0%2c .025)'%3e%3c/path%3e%3c/g%3e%3cdefs%3e%3cmask id='SvgjsMask1036'%3e%3crect width='1440' height='560' fill='white'%3e%3c/rect%3e%3c/mask%3e%3clinearGradient x1='15.28%25' y1='-39.29%25' x2='84.72%25' y2='139.29%25' gradientUnits='userSpaceOnUse' id='SvgjsLinearGradient1037'%3e%3cstop stop-color='%230e2a47' offset='0'%3e%3c/stop%3e%3cstop stop-color='rgba(44%2c 83%2c 100%2c 1)' offset='0.48'%3e%3c/stop%3e%3cstop stop-color='rgba(16%2c 33%2c 40%2c 1)' offset='1'%3e%3c/stop%3e%3c/linearGradient%3e%3c/defs%3e%3c/svg%3e");
  background-size: 100% 100%;
}

.theme--light.v-app-bar, .theme--light.v-application {
  background: rgba(255, 255, 255, 0) !important;
}

.theme--dark.v-app-bar {
  background-color: rgba(0, 0, 0, 0);
}

.v-footer {
  backdrop-filter: blur(10px);
}

.title {
  font-weight: 100;
  color: grey;
}

.theme--light .title {
  font-weight: 300;
  color: black;
}

.mina-logo {
  height: 100%;
  cursor: pointer;
}

.add-service-button {
  background: rgb(91,0,255);
  background: linear-gradient(148deg, rgba(91,0,255,0.44629521730567223) 0%, rgb(135, 1, 218) 50%, rgb(187, 79, 255) 100%);
  color: rgb(255, 255, 255);
  padding: 10px;
  border: 1px solid black;
  border-radius: 3px;
  font-weight: 500;
}

.back-button {
  /* background: rgb(91,0,255); */
  border: rgb(91,0,255) solid 2px !important;
  background: none !important;
  color: rgb(255, 255, 255);
  padding: 10px;
  border: 1px solid black;
  border-radius: 3px;
  font-weight: 500;
}


.v-footer {
  text-align: center;
  justify-content: center;
}
.v-footer span {
  font-weight: 100;
  color: grey;
}

.made-by {
  font-weight: 100;
  color: grey;
  text-decoration: dashed !important;
}

.made-by:active, .made-by:visited, .made-by:link {
  text-decoration: dashed !important;
}


</style>