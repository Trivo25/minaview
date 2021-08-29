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
    </v-navigation-drawer>
    <v-app-bar :clipped-left="clipped" fixed app>
      <v-app-bar-nav-icon color="#6a00ff" @click.stop="drawer = !drawer" />
      <img class="mina-logo" src="../assets/mina_logo_large.svg"/>
      <h1 class="title">{{ title }}</h1>
      <v-spacer />
      <v-btn
        icon
        @click="$vuetify.theme.dark = !$vuetify.theme.dark"
      >
        <v-icon
          :color=" $vuetify.theme.dark ? 'yellow' : 'black'"
        >
          mdi-theme-light-dark
        </v-icon>
      </v-btn>
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
    <v-footer fixed=true app>
      <!-- <span>&copy; {{ new Date().getFullYear() }} - MINAwatch is a community developed project and is not affiliated to Mina Foundation or O(1)Labs. </span>  -->
      <span><a class="made-by" target="_blank" href="https://github.com/Trivo25">Made by Trivo on GitHub </a> <v-icon>mdi-github</v-icon></span>
    </v-footer>
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
    }, 8000)
  }

}
</script>

<style scoped>
* {
  font-family: "Roboto";
}  

.v-app-bar, .v-navigation-drawer, .v-list, .v-list-item, .v-footer {
  box-shadow: none !important;
  border-bottom: none !important;
  background: #00000000;
  border: none !important;
  box-shadow: none !important;
}

.theme--dark.v-app-bar, .theme--dark.v-application {
  background: linear-gradient(
    to right,
    rgb(15, 32, 39),
    rgb(32, 58, 67),
    rgb(44, 83, 100)
  ) !important;
}

.theme--light.v-app-bar, .theme--light.v-application {
  /* background: linear-gradient(148deg, rgba(5,0,255,0.44629521730567223) 0%, rgba(0,167,255,1) 50%, rgba(202,255,234,0.7708918539325843) 100%) !important; */
  background-color: white !important;
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
  background: linear-gradient(148deg, rgba(91,0,255,0.44629521730567223) 0%, rgba(156,0,255,1) 50%, rgba(153, 255, 214, 0.771) 100%);
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
