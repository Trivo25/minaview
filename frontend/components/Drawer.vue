<template>
  <div class="drawer">
    <span :id="parseAnker(title)" class="title">{{title}}</span>
    <v-btn
      icon
      @click="show = !show"
    >
      <v-icon>{{ show ? 'mdi-chevron-up' : 'mdi-chevron-down' }}</v-icon>
    </v-btn>
    <v-btn
      icon
      :href="'#' + parseAnker(title)"
    >
      <v-icon>mdi-link</v-icon>
    </v-btn>
    
    <v-expand-transition>
        <div v-show="show">
          <v-divider></v-divider>

          <v-card-text>
            <div class="content-wrapper">
              <p class="content" v-html="content"></p>
            </div>
            <div class="ref-wrapper">
              <template v-for="(ref, r) in references">
                <div :key="r">
                  <span>References</span><br>
                  <span class="">[{{ r }}]</span> <a class="ref-link" :href="ref" target="_blank">{{ ref }}</a>
                </div>
              </template>
            </div>
          </v-card-text>
        </div>
      </v-expand-transition>
  </div>
</template>

<script>
export default {
  name: "Drawer",
  props: ["title", "content", "references"],
  data() {
    return {
     show: false 
    }
  },
  methods: {
    parseAnker(title) {
      return title.replace(/ /g,"_")
    }
  }
}
</script>

<style scoped>

.drawer {
  margin-top: 5px;
}

.title {
  margin-left: 35px;
}

.content {
  margin-left: 10px;
  font-size: 1.1rem;
}

.content-wrapper {
  border-left: 2px solid rgb(98, 98, 98);
  height: 100%;
  margin-left: 35px;
}

 .ref-wrapper {
  margin-left: 45px;
 }

a {
  color: rgb(255, 255, 255);
  text-decoration: none;
  font-weight: 200;
}

a:link {
  color: rgb(255, 255, 255);
  text-decoration: none;
  font-weight: 200;
}

a:visited {
  color: white;
  text-decoration: none;
  font-weight: 200;
}

a:hover {
  color: white;
  text-decoration: none;
  font-weight: 200;
}

a:active {
  color: white;
  text-decoration: none;
  font-weight: 200;
}
</style>