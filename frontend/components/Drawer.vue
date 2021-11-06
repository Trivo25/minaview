<template>
  <div class="drawer">
    <span class="title">{{title}}</span>
    <v-btn
      icon
      @click="show = !show"
    >
      <v-icon>{{ show ? 'mdi-chevron-up' : 'mdi-chevron-down' }}</v-icon>
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
  methods: {}
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

}

.content-wrapper {
  border-left: 2px solid rgb(98, 98, 98);
  height: 100%;
  margin-left: 35px;
}

 .ref-wrapper {
  margin-left: 45px;
 }

 a:link {
  color: rgb(255, 255, 255);
}

a:visited {
  color: white;
}

a:hover {
  color: white;
}

a:active {
  color: white;
}
</style>