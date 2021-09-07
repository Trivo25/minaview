<template>
  <div @click="handleClick()" class="card-wrapper">
    <div class="card">
      <div class="stats">
        <div class="left">

          <template
            v-if="service.Verified"
          >
            <v-tooltip bottom>
              <template v-slot:activator="{ on, attrs }">
                <v-icon
                color="green"
                v-bind="attrs"
                v-on="on"
                >mdi-shield-check</v-icon>
                </template>
              <span>This project has been verified to be a legitimate service.</span>
            </v-tooltip>
          </template>
          <template
            v-if="service.Audited"
          >
            <v-tooltip bottom>
              <template v-slot:activator="{ on, attrs }">
                <v-icon
                color="green"
                v-bind="attrs"
                v-on="on"
                @click="goTo(service.Audit)"
              >mdi-shield-lock</v-icon>
              </template>
              <span>This project has been audited by a security firm. Click for more information.</span>
            </v-tooltip>
          </template>
        </div>
        <div v-if="service.Github != null && service.Github != ''" class="right"> 
          <v-icon>mdi-star</v-icon>
          <span>{{service.GithubStars}}</span>

          <v-icon>mdi-source-branch</v-icon>
          <span>{{service.GithubForks}}</span>
        </div>
        
      </div>
      <div class="card-header">
        <img v-if="service.ServiceLogo != ''" class="item-logo" :src="service.ServiceLogo"/>
        <!-- <img v-else class="item-logo" src="../assets/placeholder.svg"/> -->
        <!-- might add a placeholder later -->
        
      </div>
      <h1 class="project-title">{{ service.ServiceName }}</h1>
       <div class="socials">
        <v-btn @click="goTo(service.Github)" v-if="service.Github != ''" icon><v-icon class="social" color="#4078c0">mdi-github</v-icon></v-btn>
        <v-btn @click="goTo(service.Telegram)" v-if="service.Telegram != ''" icon><v-icon class="social" color="#26A5E4">mdi-telegram</v-icon></v-btn>
        <v-btn @click="goTo(service.Reddit)" v-if="service.Reddit != ''" icon><v-icon class="social" color="#FF4500">mdi-reddit</v-icon></v-btn>
        <v-btn @click="goTo(service.Discord)" v-if="service.Discord != ''" icon><v-icon class="social" color="#5865F2">mdi-discord</v-icon></v-btn>
        <v-btn @click="goTo(service.Slack)" v-if="service.Slack != ''" icon><v-icon class="social" color="#4A154B">mdi-slack</v-icon></v-btn>
        <v-btn @click="goTo(service.Twitter)" v-if="service.Twitter != ''" icon><v-icon class="social" color="#1DA1F2">mdi-twitter</v-icon></v-btn>
      </div>
      <div class="card-content">
        <p>{{ service.ServiceDescription }}</p>
      </div>
      <div class="card-chips">
        <v-chip
          v-for="(tag, t) in tags"
          :key="t"
          class="tag-chip"
          outlined
        >
          <span>{{tag.CategoryTitle.toUpperCase()}}</span>
        </v-chip>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "ItemCard",
  props: ["service", "categories"],
  data() {
    return {
      tags: []
    }
  },
  methods: {
    async handleClick() {
      let res = await this.$axios.post("/hitService", this.$props.service)
      window.open(this.$props.service.ServiceWebsite, '_blank')
    },
    goTo(url) {
      window.open(url, '_blank')
    }
  },
  mounted() {
    // this.$props.categories.forEach(category => {
    this.$store.state.categories.forEach(category => {
      this.$props.service.CategoryKeys.forEach(key => {
        if(category.CategoryKey == key) {
          this.tags.push(category)
          // console.log(category)
        }
      })
    })
  },

}
</script>

<style scoped>

.stats {
  float: right;
  align-content: right;
  text-align: right;
  justify-content: right;
  margin-left: 5px;
  width: 100%;
}

.left {
  float: left;
  align-content: left;
  text-align: left;
  justify-content: left;
  width: 50%;
}

.left .v-icon {
  /* -webkit-animation: glow 1s ease-in-out infinite alternate;
  -moz-animation: glow 1s ease-in-out infinite alternate;
  animation: glow 1s ease-in-out infinite alternate; */
  /* text-shadow: 0 0 2px rgba(255, 255, 255, 0.207), 0 0 8px rgba(255, 255, 255, 0.234), 0 0 6px #115f001e, 0 0 8px #115f001e, 0 0 9px #115f001e, 0 0 10px #115f001e, 0 0 12px #115f001e; */
  animation: hoverIconOut 0.3s forwards;;
}

.left .v-icon:hover {
  animation: hoverIcon 1s forwards;;
}
@keyframes hoverIcon {
  0% {
    transform: scale(1.0);
  }
  100% {
    transform: scale(1.2);
  }
}


@keyframes hoverIconOut {
  0% {
    transform: scale(1.2);
  }
  100% {
    transform: scale(1.0);
  }
}



.right {
  float: right;
  align-content: right;
  text-align: right;
  justify-content: right;
  width: 50%;
}

.stats .v-icon {
  color: rgb(100, 100, 100);
}

.stats span {
  color: grey;
}

.socials {
  align-items: center;
  margin: 5px;

}

.socials .v-icon {
  margin: 5px;
}

.social {
  animation: scaleOut 0.5s forwards;
}

.social:hover {
  cursor: pointer;
  animation: scaleIn 0.5s forwards;
}

@keyframes scaleIn {
  0% {
    transform: scale(1.0);
  }
  100% {
    transform: scale(1.2);
  }
}

@keyframes scaleOut {
  0% {
    transform: scale(1.2);
  }
  100% {
    transform: scale(1.0);
  }
}

.card-wrapper {
  border-radius: 10px;
  text-align: center;
  max-height: auto;
  max-width: auto;
}

.item-logo {
  display: inline;
  height: auto;
  width: auto;
  max-height: 150px;
  max-width: 300px;
  padding: 15px;
}

.project-title {
  margin-top: 15px;
  font-weight: 400;
  font-family: "Roboto";
  color: rgb(255, 255, 255);
  text-shadow: 2px 2px 2px black;
}




.theme--light .project-title, .theme--light .card-content p {
  font-weight: 500;
  color: black;
  text-shadow: none;
}

.card { 
    /* background: linear-gradient(
    to right,
    rgb(15, 32, 39),
    rgb(32, 58, 67),
    rgb(44, 83, 100)
  ) !important; */
    margin-top: 5px;
  background-color: rgb(17, 44, 56);
  border-radius: 15px;
  border: solid 1px rgb(80, 80, 80);
  padding: 15px;
  width: auto;
  max-height: 100%;
  cursor: pointer;
  animation: hoverCardIn .5s forwards;
}

.theme--light .card {
  background-color: rgb(238, 238, 238);
  border: 2px black solid;
}

.card-chips {
  margin-top: 15px;
}

.card:hover {
  animation: hoverCardOut .5s forwards;
  background-color: rgb(16, 56, 73);
}

.theme--light .card:hover {
  background-color: rgb(223, 223, 223);
}

@keyframes hoverCardOut {
  0% {
    transform: scale(1.0)
  }
  100% {
    transform: scale(1.03)
  }
}

@keyframes hoverCardIn {
  0% {
    transform: scale(1.03)
  }
  100% {
    transform: scale(1.0)
  }
}

.tag-chip {
  font-weight: 500;
  font-size: 12px;
  color: rgba(255, 115, 0, 0.696);
  border-color: rgb(85, 85, 85);
  margin: 2px;
}


.theme--light .tag-chip {
  color: rgba(255, 115, 0, 1);
}

.card-content p {
  font-family: "Roboto";
  font-size: 1.1rem;
  text-shadow: 1px 1px 1px black;
}
</style>