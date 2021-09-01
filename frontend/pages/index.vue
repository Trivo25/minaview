<template>
  <div class="index">
    <!-- <TrendingCard /> -->
    <!-- <img src="../assets/mina_logo_large.svg"/>
    <h1>The world's lightest blockchain, powered by participants.</h1>
    <h3>MINAview offers an overview of the fast-growing ecosystem.</h3>
    <v-btn
      class="explore-button bn"
      outlined
      x-large
      to="/ecosystem"
      elevation="24"
    >
      <span>Explore the Ecosystem</span>
    </v-btn> -->
    <PeerMap />
  </div>
</template>

<script>
import TrendingCard from "../components/TrendingCard.vue"
import PeerMap from "../components/PeerMap.vue"

export default {
  name: "index",
  props: [],
  components: {
    PeerMap
  },
  data () {
    return {

    }
  },
  methods: {
  },
  computed: {
  },
}
</script>


<style scoped>
.bn {
  padding: 0.6em 2em;
  border: none;
  outline: none;
  color: rgb(255, 255, 255);
  background: #111;
  cursor: pointer;
  position: relative;
  z-index: 0;
  border-radius: 10px;
}

.bn:before {
  content: "";
  background: linear-gradient(
    45deg,
    #ff0000,
    #ff7300,
    #fffb00,
    #48ff00,
    #00ffd5,
    #002bff,
    #7a00ff,
    #ff00c8,
    #ff0000
  );
  position: absolute;
  top: -2px;
  left: -2px;
  background-size: 400%;
  z-index: -1;
  filter: blur(5px);
  width: calc(100% + 4px);
  height: calc(100% + 4px);
  animation: glowingbn5 20s linear infinite;
  opacity: 0;
  transition: opacity 0.3s ease-in-out;
  border-radius: 10px;
}

@keyframes glowingbn5 {
  0% {
    background-position: 0 0;
  }
  50% {
    background-position: 400% 0;
  }
  100% {
    background-position: 0 0;
  }
}

.bn:active {
  color: #000;
}

.bn:active:after {
  background: transparent;
}

.bn:hover:before {
  opacity: 1;
}

.bn:after {
  z-index: -1;
  content: "";
  position: absolute;
  width: 100%;
  height: 100%;
  background: #191919;
  left: 0;
  top: 0;
  border-radius: 10px;
}

.index {
  justify-content: center;
  text-align: center;
  width: 100%;
  height: 100% !important; 
  

}

.theme--light h1, .theme--light h3 {
  color: black;
}

h1, h3 {
  font-weight: 400;
  color: rgb(214, 214, 214);
}

.explore-button {
  margin-top: 15px;
  justify-content: center;
  align-items: center;
  text-align: center;
  line-height: 50px;
  font-size: 1.5rem;
  border: solid 1px black;
}


</style>