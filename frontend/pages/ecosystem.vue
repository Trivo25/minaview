<template>
  <div>
    <div class="sticky-wrapper">
      <div class="search-form">
        <v-row>
          <v-text-field
            v-model="nameSearchMatcher"
            placeholder="Type to search.."
          >
          
          </v-text-field>
          <!-- <span class="sort-option">GitHub Stars
            <v-icon v-if="sortGithubStars" @click="sortGithubStars = !sortGithubStars" color="#a5a5a5">mdi-arrow-up</v-icon>
            <v-icon v-if="!sortGithubStars" @click="sortGithubStars = !sortGithubStars" color="#a5a5a5">mdi-arrow-down</v-icon>
          </span>
          <span class="sort-spacer">|</span>
          <span class="sort-option">Likes
            <v-icon color="#a5a5a5" v-if="sortLikes" @click="sortLikes = !sortLikes">mdi-arrow-up</v-icon>
            <v-icon v-if="!sortLikes" @click="sortLikes = !sortLikes" color="#a5a5a5">mdi-arrow-down</v-icon>
          </span>
          <span class="sort-spacer">|</span>
          <span class="sort-option">Downloads
            <v-icon color="#a5a5a5" v-if="sortDownloads" @click="sortDownloads = !sortDownloads">mdi-arrow-up</v-icon>
            <v-icon v-if="!sortDownloads" @click="sortDownloads = !sortDownloads" color="#a5a5a5">mdi-arrow-down</v-icon>
          </span> -->
        </v-row>
      </div>      
    </div>
    <!-- {{parseParams()}} -->
    <div class="list-wrapper">
      <ul :style="gridStyle" class="card-list">
        <!-- TODO: fix below error -->
        <template v-for="(item, i) in taggedItems">
          <li :key="i" class="card-item" v-show="searchTitle(item)"> 
            <ItemCard :project="item" />
          </li>
        </template>
        <div class="not-found" v-show="noMatches()">
          <h1>This service seems to be from another planet..</h1>
          <img src="../assets/not_from_this_world.svg"/>
        </div>
      </ul>
    </div>
  </div>
</template>

<script>
import ItemCard from "../components/ItemCard.vue"

export default {
  name: "Ecosystem",
  props: [],
  data () {
    return {
      sortDownloads: false,
      sortLikes: false,
      sortGithubStars: false,
      nameSearchMatcher: "",
      items: [
        {
          title: "Clorio",
          logo: "https://clor.io/assets/ClorioLogo.svg",
          description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Possimus atque placeat et dicta, odio maiores error modi blanditiis, qui sequi asperiores magnam dolorem nihil eveniet quod consequatur, illum odit eaque.", 
          tags: ["Wallet", "Staking Pool"],
          keys: ["wallets", "staking-pools"],
          link: ""
        },
        {
          title: "MINA",
          logo: "https://minaprotocol.com/wp-content/themes/minaprotocol/img/mina-wordmark-light.svg",
          description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Possimus atque placeat et dicta, odio maiores error modi blanditiis, qui sequi asperiores magnam dolorem nihil eveniet quod consequatur, illum odit eaque.", 
          tags: ["Explorer", "Staking Pool", "Monitoring and Dashboard"],
          keys: ["explorers", "staking-pools", "monitoring-and-dashboards"],
          link: ""
        },
        {
          title: "Auro",
          logo: "https://www.aurowallet.com/wp-content/uploads/2021/04/icon-128.png",
          description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Possimus atque placeat et dicta, odio maiores error modi blanditiis, qui sequi asperiores magnam dolorem nihil eveniet quod consequatur, illum odit eaque.", 
          tags: ["Wallet", "Staking Pool"],
          keys: ["wallets", "staking-pools"],
          link: ""
        },
        {
          title: "MinaExplorer",
          logo: "https://minaexplorer.com/static/global_assets/images/mina-explorer-image.png",
          description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Possimus atque placeat et dicta, odio maiores error modi blanditiis, qui sequi asperiores magnam dolorem nihil eveniet quod consequatur, illum odit eaque.", 
          tags: ["Explorer", "Staking Pool"],
          keys: ["explorers", "staking-pools"],
          link: ""
        },
        {
          title: "TowerStake",
          logo: "https://towerstake.com/wp-content/uploads/2021/04/TowerStake_logo_blanc_WEB.png",
          description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Possimus atque placeat et dicta, odio maiores error modi blanditiis, qui sequi asperiores magnam dolorem nihil eveniet quod consequatur, illum odit eaque.", 
          tags: ["Explorer", "Staking Pool", "Monitoring and Dashboard"],
          keys: ["explorers", "staking-pools", "monitoring-and-dashboards"],
          link: ""
        },
        {
          title: "Piconbello",
          logo: "https://mina.piconbello.com/static/logo-f4ceae98580bca2536ac9e6749f5c3c5.png",
          description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Possimus atque placeat et dicta, odio maiores error modi blanditiis, qui sequi asperiores magnam dolorem nihil eveniet quod consequatur, illum odit eaque.", 
          tags: ["Staking Pool"],
          keys: ["staking-pools"],
          link: ""
        },
        {
          title: "MinaExplorer",
          logo: "https://minaexplorer.com/static/global_assets/images/mina-explorer-image.png",
          description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Possimus atque placeat et dicta, odio maiores error modi blanditiis, qui sequi asperiores magnam dolorem nihil eveniet quod consequatur, illum odit eaque.", 
          tags: ["Explorer", "Staking Pool"],
          keys: ["explorers", "staking-pools"],
          link: ""
        },
      ],
      numberOfColumns: 3,
      noResultsByTag: false
    }
  },
  methods: {
    addCard() {
      this.cards.push('new-card')
    },
    noMatches() {
      return !this.items.find(x => x.title.toLowerCase().includes(this.nameSearchMatcher.toLowerCase())) || this.noResultsByTag
    },
    searchTitle(item) {
      return item.title.toLowerCase().includes(this.nameSearchMatcher.toLowerCase())
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
    taggedItems() {
      //return this.parseParams()[0] == ""  ? this.items : this.items.filter(x => JSON.stringify(x.keys) == JSON.stringify(this.parseParams()))
      let arr = this.parseParams()[0] == "" ? this.items : this.items.filter(x => this.parseParams().every(r => x.keys.includes(r)))
      if(arr.length == 0) this.noResultsByTag = true
      else this.noResultsByTag = false
      return arr
    }
  },
}
</script>


<style >

.not-found {
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

.sticky-wrapper {
  width: 100%;
  height: 50px;
  /* border: solid 1px red; */
  position: sticky;
  backdrop-filter: blur(10px);
  top: 60px;
  z-index: 5;
  /* centering */
  /* margin-right: 0px !important;
  transform: translate(-50%, -50%);
  -ms-transform: translate(-50%, -50%);
  -webkit-transform: translate(-50%, -50%);   */
}

.search-form {
  position: relative;
  width: 100%;
  /* left: 50%; */
}

.card-list {
  height: auto;
  position: relative;
  display: grid;
  margin-top: 50px;
  grid-auto-flow: dense;
  grid-gap: 1em;
  grid-template-columns: repeat(3, auto);
  grid-template-rows: repeat(3, 1fr);
}

ul {
  list-style-type: none;
}
</style>