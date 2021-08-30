<template>
  <div @click="handleClick()" class="card-wrapper">
    <div class="card">
      <div class="card-header">
        <img v-if="service.ServiceLogo != ''" class="item-logo" :src="service.ServiceLogo"/>
        <img v-else class="item-logo" src="../assets/placeholder.svg"/>
        
      </div>
      <h1 class="project-title">{{ service.ServiceName }}</h1>
       <div class="socials">
        <v-btn @click="goTo(service.Github)" v-if="service.Github != ''" icon><v-icon class="social" color="#F05032">mdi-git</v-icon></v-btn>
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
      window.open(this.$props.service.ServiceWebsite, '_blank').focus();
    },
    goTo(url) {
      window.open(url, '_blank').focus();
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
  width: 100%;
  text-align: center;
  max-height: auto;
  max-width: 100%;
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