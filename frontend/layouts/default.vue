<template>
  <v-app dark>
    <NotificationCard @closeNotification="showInformation = false" v-if="showInformation" />
    <v-navigation-drawer
      v-model="drawer"
      :clipped="clipped"
      fixed
      app
    >
    <TagNav/>
    </v-navigation-drawer>
    <v-app-bar :clipped-left="clipped" fixed app>
      <v-app-bar-nav-icon color="#6a00ff" @click.stop="drawer = !drawer" />
      <img class="mina-logo" src="../assets/mina_logo_large.svg"/>
      <h1 class="title">{{ title }}</h1>
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
    <v-footer fixed=true app>
      <!-- <span>&copy; {{ new Date().getFullYear() }} - MINAwatch is a community developed project and is not affiliated to Mina Foundation or O(1)Labs. </span>  -->
      <span><a class="made-by" target="_blank" href="https://github.com/Trivo25">Made by Trivo on GitHub </a> <v-icon>mdi-github</v-icon></span>
    </v-footer>
  </v-app>
</template>

<script>
import TagNav from "../components/TagNav.vue"
import NotificationCard from "../components/NotificationCard.vue"

export default {
  name: "default",
  components: {
    TagNav,
    NotificationCard
  },
  data() {
    return {
      clipped: true,
      drawer: true,
      fixed: false,
      items: [
        {
          icon: 'mdi-apps',
          iconColor: "#5442f5",
          title: 'All',
          to: '/',
          total: 5
        },
        {
          icon: 'mdi-chart-bubble',
          iconColor: "#7542f5",
          title: 'Official Resources',
          to: '/official-resources',
          total: 5
        },
        {
          icon: 'mdi-chart-bubble',
          iconColor: "#8d42f5",
          title: 'Wallets',
          to: '/wallets',
          total: 5
        },
        {
          icon: 'mdi-chart-bubble',
          iconColor: "#a442f5",
          title: 'Explorers',
          to: '/explorers',
          total: 5
        },
        {
          icon: 'mdi-chart-bubble',
          iconColor: "#b642f5",
          title: 'Tools',
          to: '/tools',
          total: 5
        },
        {
          icon: 'mdi-chart-bubble',
          iconColor: "#8d42f5",
          title: 'Staking Pools',
          to: '/staking-pools',
          total: 5
        },
        {
          icon: 'mdi-chart-bubble',
          iconColor: "#d442f5",
          title: 'Monitoring and Dashboards',
          to: '/monitoring-und-dashboards',
          total: 5
        },
        {
          icon: 'mdi-chart-bubble',
          iconColor: "#f542ef",
          title: 'Ledger Apps',
          to: '/ledger',
          total: 5
        },
        {
          icon: 'mdi-chart-bubble',
          iconColor: "#f5426f",
          title: 'News, Resources and Articles',
          to: '/news-articles',
          total: 5
        },
        {
          icon: 'mdi-chart-bubble',
          iconColor: "#f54242",
          title: 'International Communities',
          to: '/communities',
          total: 5
        },
      ],
      miniVariant: false,
      title: 'View',
      showInformation: false
    }
  },
  async mounted() {
    // await this.$store.dispatch("getCategories")
    // await this.$store.dispatch("getServices")

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

.theme--dark.v-app-bar {
  background: linear-gradient(
    to right,
    rgb(15, 32, 39),
    rgb(32, 58, 67),
    rgb(44, 83, 100)
  ) !important;
}

.theme--light.v-app-bar {
  background: rgb(5,0,255);
  background: linear-gradient(148deg, rgba(5,0,255,0.44629521730567223) 0%, rgba(0,167,255,1) 50%, rgba(202,255,234,0.7708918539325843) 100%);
}

.theme--dark.v-application {
  background: #0f2027;
  background: linear-gradient(
    to right,
    rgb(15, 32, 39),
    rgb(32, 58, 67),
    rgb(44, 83, 100)
  ) !important; 
}

.v-footer {
  backdrop-filter: blur(10px);
}


.title {
  font-weight: 100;
  color: grey;
}

.mina-logo {
  height: 100%;
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
