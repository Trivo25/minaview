<template>
  <div>
    <span>Test</span>
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
            I'm a thing. But, like most politicians, he promised more than he could deliver. You won't have time for sleeping, soldier, not with all the bed making you'll be doing. Then we'll go with that data file! Hey, you add a one and two zeros to that or we walk! You're going to do his laundry? I've got to find a way to escape.
          </v-card-text>
        </div>
      </v-expand-transition>
  </div>
</template>

<script>
export default {
  name: "Drawer",
  props: ["title", "content"],
  data() {
    return {
     show: false 
    }
  },
  methods: {}
}
</script>