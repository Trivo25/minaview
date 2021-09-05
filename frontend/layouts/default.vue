<template>
  <v-app dark>
    <NotificationCard @closeNotification="showInformation = false" v-if="showInformation" />
    <v-navigation-drawer
      v-model="drawer"
      :clipped="clipped"
      fixed
      app
    >
    <TagNav :isLoading="isLoading" :categories="categories"/>
    <v-spacer></v-spacer>
    <div class="socials">
      <span>Made by Trivo on <a href="https://www.github.com/trivo25/mina-view" class="mdi mdi-github"></a></span>
      <!-- <span>Please leave feedback your <a href="https://www.github.com/trivo25/mina-view/issues" class="mdi mdi-github">here</a>!</span> -->
    </div>
    </v-navigation-drawer>
    <v-app-bar :clipped-left="clipped" fixed app>
      <v-app-bar-nav-icon color="#6a00ff" @click.stop="drawer = !drawer" />
      <img @click="goTo('/')" class="mina-logo" src="../assets/mina_logo_large.svg"/>
      <h1 class="title">{{ title }}</h1>
      <v-spacer />
      <!-- <v-btn
        icon
        @click="$vuetify.theme.dark = !$vuetify.theme.dark"
      >
        <v-icon
          :color=" $vuetify.theme.dark ? 'yellow' : 'black'"
        >
          mdi-theme-light-dark
        </v-icon>
      </v-btn> -->
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
    <!-- <v-footer fixed=false app>
      <span>&copy; {{ new Date().getFullYear() }} - MINAwatch is a community developed project and is not affiliated to Mina Foundation or O(1)Labs. </span>
      <span><a class="made-by" target="_blank" href="https://github.com/Trivo25">Made by Trivo on GitHub </a> <v-icon>mdi-github</v-icon></span>
    </v-footer> -->
    <Error @closeNotification="error.show = false" v-if="error.show" :error="error.error" :type="error.type" />
  </v-app>
</template>

<script>
import TagNav from "../components/TagNav.vue"
import NotificationCard from "../components/NotificationCard.vue"
import Error from "../components/Error.vue"

export default {
  name: "default",
  components: {
    TagNav,
    NotificationCard,
    Error
  },
  data() {
    return {
      clipped: true,
      drawer: true,
      fixed: false,
      categories: [],
      services: [],
      miniVariant: false,
      title: 'View',
      showInformation: false,
      isLoading: false,
      error: {
        show: false,
        error: null,
        type: "error"
      }
    }
  },
  methods: {
    goTo(url) {
      window.location.href = url
    }
  },
  async mounted() {
    this.isLoading = true

    await this.$store.dispatch("getCategories")
      .then(res => {
        this.categories = res
      }, error => {
        this.error.error = error
        this.error.type == "error"
        this.error.show = true
        
      })


    this.isLoading = false

    setTimeout(() => {
      this.showInformation = true
    }, 10000)
  },
}
</script>

<style scoped>
* {
  font-family: "Roboto";
}  

.socials {
  height: 50px;
  width: 100%;
  position: absolute;
  bottom: 0 !important;
  text-align: center;
  align-items: center;
  vertical-align: middle;
  line-height: 50px;
}

.socials span {
  color: grey;
  font-weight: 300;
  text-align: center;
  align-items: center;
  vertical-align: middle;
  line-height: 50px;
}

.socials a:active {
  text-decoration: none;
}

.socials a:link {
  text-decoration: none;
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
  /* background: linear-gradient(
    to right,
    rgb(16, 33, 40),
    rgb(32, 58, 67),
    rgb(44, 83, 100)
  ) !important;background-color: #0F2027; */
/* background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='221' height='221' viewBox='0 0 800 800'%3E%3Cg fill='none' stroke='%23203A43' stroke-width='1'%3E%3Cpath d='M769 229L1037 260.9M927 880L731 737 520 660 309 538 40 599 295 764 126.5 879.5 40 599-197 493 102 382-31 229 126.5 79.5-69-63'/%3E%3Cpath d='M-31 229L237 261 390 382 603 493 308.5 537.5 101.5 381.5M370 905L295 764'/%3E%3Cpath d='M520 660L578 842 731 737 840 599 603 493 520 660 295 764 309 538 390 382 539 269 769 229 577.5 41.5 370 105 295 -36 126.5 79.5 237 261 102 382 40 599 -69 737 127 880'/%3E%3Cpath d='M520-140L578.5 42.5 731-63M603 493L539 269 237 261 370 105M902 382L539 269M390 382L102 382'/%3E%3Cpath d='M-222 42L126.5 79.5 370 105 539 269 577.5 41.5 927 80 769 229 902 382 603 493 731 737M295-36L577.5 41.5M578 842L295 764M40-201L127 80M102 382L-261 269'/%3E%3C/g%3E%3Cg fill='%232C5364'%3E%3Ccircle cx='769' cy='229' r='6'/%3E%3Ccircle cx='539' cy='269' r='6'/%3E%3Ccircle cx='603' cy='493' r='6'/%3E%3Ccircle cx='731' cy='737' r='6'/%3E%3Ccircle cx='520' cy='660' r='6'/%3E%3Ccircle cx='309' cy='538' r='6'/%3E%3Ccircle cx='295' cy='764' r='6'/%3E%3Ccircle cx='40' cy='599' r='6'/%3E%3Ccircle cx='102' cy='382' r='6'/%3E%3Ccircle cx='127' cy='80' r='6'/%3E%3Ccircle cx='370' cy='105' r='6'/%3E%3Ccircle cx='578' cy='42' r='6'/%3E%3Ccircle cx='237' cy='261' r='6'/%3E%3Ccircle cx='390' cy='382' r='6'/%3E%3C/g%3E%3C/svg%3E"); */

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
