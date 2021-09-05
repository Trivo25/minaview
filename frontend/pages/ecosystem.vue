<template>
  <div>
    <div class="sticky-wrapper">
      <div class="search-form">
        <v-row>
          <v-text-field
            v-model="nameSearchMatcher"
            placeholder="Type to search.."
          >
          <!-- {{$store.services}} -->
          </v-text-field>
          <!-- <span class="sort-option">GitHub Stars
            <v-icon v-if="sortGithubStars" @click="sortGithubStars = !sortGithubStars" color="#a5a5a5">mdi-arrow-up</v-icon>
            <v-icon v-if="!sortGithubStars" @click="sortGithubStars = !sortGithubStars" color="#a5a5a5">mdi-arrow-down</v-icon>
          </span> -->
        </v-row>
      </div>      
    </div>
    <!-- {{parseParams()}} -->
            
    <div class="not-found" v-show="noMatches()">
          <h1>This service seems to be from another planet..</h1>
          <img class="not-found" src="../assets/not_from_this_world.svg"/>
    </div>
    <div v-if="!isLoading" class="card-list">
      <ItemCard v-for="(item, i) in taggedServices" :key="i" class="card-item" v-show="searchTitle(item)" :categories="categories" :service="item" />
    </div>
    <div v-else>
      <Loader/>
    </div>
    <Error @closeNotification="error.show = false" v-if="error.show" :error="error.error" :type="error.type"/>
  </div>
</template>

<script>
import ItemCard from "../components/ItemCard.vue"
import Loader from "../components/Loader.vue"


export default {
  name: "Ecosystem",
  props: [],
  components: {
    ItemCard,
    Loader,
  },
  data () {
    return {
      sortDownloads: false,
      sortLikes: false,
      sortGithubStars: false,
      nameSearchMatcher: "",
      services: [],
      categories: [],
      numberOfColumns: 3,
      noResultsByTag: false,
      isLoading: false,
      error: {
        error: null,
        show: false,
        type:"error"
      }
    }
  },
  methods: {
    addCard() {
      this.cards.push('new-card')
    },
    noMatches() {
      return !this.services.find(x => x.ServiceName.toLowerCase().includes(this.nameSearchMatcher.toLowerCase())) || this.noResultsByTag
    },
    searchTitle(item) {
      return item.ServiceName.toLowerCase().includes(this.nameSearchMatcher.toLowerCase())
    },
    parseParams() {
      let categories = this.$route.query.categories
      if(categories == null || categories == undefined) return [""]
      return categories.split(',');
    },
  },
  computed: {
    gridStyle() {
      return {
        gridTemplateColumns: ""//`repeat(${this.numberOfColumns}, minmax(100px, 1fr))`
      }
    },
    // parseParams() {
    //   return this.$route.query.categories.split(',');
    // },
    taggedServices() {
      //return this.parseParams()[0] == ""  ? this.items : this.items.filter(x => JSON.stringify(x.keys) == JSON.stringify(this.parseParams()))
      let arr = this.parseParams()[0] == "" ? this.services : this.services.filter(x => this.parseParams().every(r => x.CategoryKeys.includes(r)))
      if(arr.length == 0) this.noResultsByTag = true
      else this.noResultsByTag = false
      return arr
    }
  },
  async mounted()  {

    this.isLoading = true

    await this.$store.dispatch("getServices")
      .then(res => {
        this.services = res
      }, error => {
        this.error.error = error
        this.error.show = true
        this.error.type = "error"
      })
    await this.$store.dispatch("getCategories")
      .then(res => {
        this.categories = res
      }, error => {
        this.error.error = error
        this.error.show = true
        this.error.type = "error"
      })
    // this.$store.dispatch("getCategories")
    // this.services = await this.$store.state.services
    // this.categories = await this.$store.state.categories

    this.isLoading = false
  }
}
</script>


<style scoped>


.card-item {
  margin: 0;
  display: grid;
  grid-template-rows: 1fr auto;
  margin-bottom: 10px;
  break-inside: avoid;
  width: 100%;
  z-index: 5;
}

.card-list {
  margin-top: 35px;
  column-count: 3;
  column-gap: 15px;
}

.not-found {
  margin-top: 15px;
  display: flex;
  text-align: center !important;
  justify-content: center;
  align-items: center;
}

.not-found img {
  position: relative;
  max-height: 80%;
  max-width: 50%;
}

.not-found h1 {
  font-weight: 300;
  color: grey;
}

.ecosystem {
  display: block;
}

.sort-option {
  font-weight: 300;
  color: #a5a5a5;
  line-height: 60px;
  margin-left: 15px;
}
.sort-option .v-icon {
  cursor: pointer;
}

.sort-spacer {
  font-weight: 300;
  color: rgb(165, 165, 165);
  line-height: 60px;
  margin-left: 15px;
}

.search-form {
  position: relative;
  width: 100%;

}


ul {
  list-style-type: none;
}

@keyframes flying {
  0% {
    top: 0px;
  }
  50% {
    top: 50px;
  }
  100% {
    top: 0px;
  }
}

@media only screen and (max-width: 1273px) {
  .card-list {
    column-count: 1 !important;
  }
}
</style>