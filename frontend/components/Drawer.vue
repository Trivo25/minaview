<template>
  <div class="drawer">
    <!-- {{ selected }} -->
    <span :id="parseAnker(title)" class="title">{{id + title}}</span>
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
              <span v-if="references.length != 0">References</span><br>
              <template v-for="(ref, r) in references">
                <div :key="r">
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
  props: ["title", "content", "references", "selected", "id"],
  data() {
    return {
     show: false 
    }
  },
  methods: {
    parseAnker(title) {
      return title.replace(/ /g,"_")
    }
  },
  mounted() {
    if(this.$props.selected == this.parseAnker(this.$props.title)) {
      this.show = true
    }
  },
}
</script>

<style scoped>

.drawer {
  margin-top: 5px;
}

.title {
  margin-left: 35px;
  -webkit-text-stroke: 0.5px rgb(0, 0, 0);
}

.content {
  margin-left: 10px;
  font-size: 1.1rem;
  -webkit-text-stroke: 0.5px rgb(0, 0, 0);
}

.content-wrapper {
  border-left: 1px solid rgb(100, 100, 100);
  height: 100%;
  background-color: rgba(29, 29, 29, 1);
  -webkit-text-stroke: none;
  margin-left: 35px;
  padding: 5px;
}

.ref-wrapper {
  margin-left: 45px;
  -webkit-text-stroke: 0.5px rgb(0, 0, 0);
}

a, a:link, a:visited, a:hover, a:active {
  color: rgb(184, 184, 184);
  text-decoration: none;
  font-weight: 200;
}

</style>