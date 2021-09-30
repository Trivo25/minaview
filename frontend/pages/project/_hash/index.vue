<template>
  <div class="project-wrapper">
    <div v-if="isLoading">
      <Loader/>
    </div>
    <div v-else>
      <div class="header">
        <div class="logo">
          <img :src="service.ServiceLogo" alt="Project Logo">
        </div>
        <div class="title">
          <div>
            <h2>{{ service.ServiceName }}</h2>
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
                >mdi-shield-lock</v-icon>
                </template>
                <span>This project has been audited by a security firm.</span>
              </v-tooltip>
            </template>
          </div>
          <RatingStars :up="service.Up" :down="service.Down" />
          <p>{{ service.ServiceDescription }}</p>
        </div>
        <div class="card-chips">
          <v-chip
            v-for="(tag, t) in tags"
            :key="t"
            class="tag-chip"
            outlined
          >
            <span>{{ tag.CategoryTitle.toUpperCase()}}</span>
          </v-chip>
        </div>
        <div>
          <v-btn
            class="open-website-button"
            :href="service.ServiceWebsite"
            target="_blank"
            large
          >
            Visit Website
            <v-icon>mdi-menu-right</v-icon>
          </v-btn>
        </div>
      </div>
      <div class="seperator"></div>
      <div class="content">
        <div class="left">
          <h2>Social Media</h2>
          <div class="socials">
            <ul>
              <li v-if="service.Github != ''">
                <v-icon large class="social" color="#4078c0">mdi-github</v-icon>
                <a :href="service.Github">{{ service.Github }}</a>
              </li>
              <li v-if="service.Telegram != ''">
                <v-icon large class="social" color="#26A5E4">mdi-send</v-icon>
                <a :href="service.Telegram">{{ service.Telegram }}</a>
              </li>
              <li v-if="service.Reddit != ''">
                <v-icon large class="social" color="#FF4500">mdi-reddit</v-icon>
                <a :href="service.Reddit">{{ service.Reddit }}</a>
              </li>
              <li v-if="service.Discord != ''">
                <v-icon large class="social" color="#5865F2">mdi-discord</v-icon>
                <a :href="service.Discord">{{ service.Discord }}</a>
              </li>
              <li v-if="service.Slack != ''" >
                <v-icon large class="social" color="#4A154B">mdi-slack</v-icon>
                <a :href="service.Slack">{{ service.Slack }}</a>
              </li>
              <li v-if="service.Twitter != ''">
                <v-icon large class="social" color="#1DA1F2">mdi-twitter</v-icon>
                <a :href="service.Twitter">{{ service.Twitter }}</a>
              </li>
            </ul>
          </div>
          <Tweet
            class="tweet"
            id="1440369345817837574"
            :options="{ theme: 'light' }"
            error-message="This tweet could not be loaded"
          >
          </Tweet>
        </div>
        <div class="right">
          <h2>What users think</h2>
          <RatedBar :hash="getServiceHash()" :up="service.Up" :down="service.Down" />
          <div class="seperator"></div>
          <CommentWrapper :comments="service.Comments" :service_hash="service.ServiceHash" />
        </div>
      </div>
    </div>
    <Error @closeNotification="error.show = false" v-if="error.show" :error="error.error" :type="error.type" />
  </div>
</template>

<script>
import Error from "../../../components/Error.vue"
import { Tweet } from 'vue-tweet-embed'
import Loader from "../../../components/Loader.vue"
import RatedBar from "../../../components/RatedBar.vue"
import CommentWrapper from "../../../components/CommentSection/CommentWrapper.vue"
import RatingStars from "../../../components/RatingStars.vue"

export default {
  name: "project",
  layout: "project-page",
  props: [],
  components: {
    Error,
    Tweet,
    Loader,
    RatedBar,
    CommentWrapper,
    RatingStars
  },
  data () {
    return {
      service: new Object,
      isLoading: false,
      tags: [],
      error: {
        show: false,
        error: null,
        type: "error"
      }
    }
  },
  methods: {
    getServiceHash() {
      return this.$route.params.hash
    },
  },
  computed: {
  },
  created() {
    this.isLoading = true 
  },
  async mounted() {
    this.isLoading = true

    try {

      let res = await this.$axios.post("/getServiceByHash", {
        Hash: this.getServiceHash(),
      })
      this.service = res.data.Data

      res = await this.$axios.post("/hitService", this.service)

      if(this.$store.state.categories == null || this.$store.state.categories == undefined || this.$store.state.categories.length == 0) {
        let categories = (await this.$axios.post("/getCategories")).data.Data
        categories.forEach(category => {
        this.service.CategoryKeys.forEach(key => {
          if(category.CategoryKey == key) {
            this.tags.push(category)
          }
        })})
      } else {
        this.$store.state.categories.forEach(category => {
        this.service.CategoryKeys.forEach(key => {
          if(category.CategoryKey == key) {
            this.tags.push(category)
          }
        })
      })
      }

      

    } catch (error) {
      console.log(error)
    }
    this.isLoading = false
  },
}
</script>

<style scoped>

@import url('https://fonts.googleapis.com/css2?family=Open+Sans:ital,wght@0,300;0,400;0,600;0,700;0,800;1,300;1,400;1,600;1,700;1,800&display=swap');

* {
  font-family: "Open Sans"
}


.open-website-button {
  /* background: rgb(91,0,255); */
  border: rgb(91,0,255) solid 2px !important;
  background: none !important;
  color: rgb(255, 255, 255);
  padding: 10px;
  border: 1px solid black;
  border-radius: 3px;
  font-weight: 500;
  margin-bottom: 10px;
}

.tweet-not-found {
  /* color: red !important; */
  font-size: 50px;
}

.project-wrapper {
  position: relative;
}

.header {
  display: inline;
  justify-content: center;
  text-align: center;
  width: 100%;
  height: auto;
  align-items: center;

}

.content {
  display: flex;
  width: 100%;
  height: 100%;
}

.left, .right {
  width: 50%;
  text-align: center;
  /* border: solid 1px red; */
}

.left h2, .right h2 {
  font-weight: 400;
  color: rgb(146, 146, 146);
  text-shadow: 2px 2px 2px rgb(42, 42, 42);
  margin-bottom: 10px;
}

.header .logo {
  height: 50%;
}

.header .logo img {
  max-height: 200px;
}

.header .title {
  width: 100%;
  height: 50%;
}

.header .title h2 {
  font-weight: 400;
  color: rgb(255, 255, 255);
  text-shadow: 2px 2px 2px black;
}

.header .title p {
  margin-top: 10px;
  font-family: "Open Sans";
  color: rgb(255, 255, 255);
  text-shadow: 2px 1px 2px black;
}

.socials {
  align-items: center;
  margin: 8px;
  height: auto;
  text-align: center;
}

.socials ul {
  list-style: none;
}

.socials ul a {
  color: grey;
  font-weight: 300;
}

.socials ul .v-icon {
  margin: 3px;
}

.social {
  animation: scaleOut 0.5s forwards;
}

.social:hover {
  cursor: pointer;
  animation: scaleIn 0.5s forwards;
}

.tweet {
  /* width: 50%; */
  margin-left: 20%;
  margin-right: 20%;
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

.card-chips {
  margin-top: 15px;
  margin-bottom: 15px;
}

.seperator {
  width: 100%;
  margin-top: 20px;
  margin-bottom: 20px;
  border-bottom: solid 1px rgba(160, 160, 160, 0.4);
  border-radius: 50px;
}

</style>