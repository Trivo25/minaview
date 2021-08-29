<template>
  <div class="nav-tags-wrapper">
    <v-text-field
      class="category-filter-field"
      v-model="searchCategoryFilter"
      placeholder="e.g. Wallets"
    />
    <v-row>
      <v-chip
        router
        exact
        label
        class="chip-item"
        :class="{'chip-active' : allChip.isActive }"
        @click="clearSelection()"
      >
        <span class="category-title">{{ allChip.CategoryTitle.toUpperCase() }}</span>
        <span class="category-total">({{ allChip.ServiceCount ? allChip.ServiceCount : 0 }})</span>
      </v-chip>
    </v-row>
    <v-row
      v-for="(item, i) in categories"
      :key="i"
    >
      <v-chip
        :to="routeParams"
        router
        exact
        label
        v-show="filterCategories(item)"
        class="chip-item"
        :class="{'chip-active' : item.isActive }"
        @click="changeSelection(item)"
      >
        <span class="category-title">{{ item.CategoryTitle.toUpperCase() }}</span>
        <span class="category-total">({{ item.ServiceCount ? item.ServiceCount : 0 }})</span>
      </v-chip>
    </v-row>
  </div>
</template>

<script>
export default {
  name: "TagNav",
  props: [],
  data() {
    return {
      searchCategoryFilter: "",
      targetRoute: "/ecosystem?categories=",
      parsedTargetPath: "",
      params: [],
      allChip: {
          CategoryTitle: 'All',
          to: '/ecosystem',
          ServiceCount: 0,
          isActive: true
        },
      categories: [
      ],
    }
  },
  methods: {
    filterCategories(item) {
      return item.CategoryTitle.toLowerCase().includes((this.searchCategoryFilter).toLowerCase())
    },
    changeSelection(item) {
      this.params.includes(item.CategoryKey) && item.isActive ?
        this.params.splice(this.params.indexOf(item.CategoryKey), 1) :
        this.params.push(item.CategoryKey); item.isActive = !item.isActive;

      this.allChip.isActive = false;
      this.$router.push({
        path: this.parsedTargetPath
      })
    },
    clearSelection() {
      this.allChip.isActive = true
      this.params = new Array()
      this.categories.forEach(item => {
        item.isActive = false
      })
    },
  },
  computed: {
    routeParams() {
      let paramString = ""
      this.params.forEach((param, i) => {
        paramString += (i == 0 ? '' :  ",") + param
      })
      let parsedTargetPath = `${this.targetRoute}${paramString}`
      this.$router.push({
        path: parsedTargetPath
      })
      return parsedTargetPath
    }
  },
  async mounted() {
    await this.$store.dispatch("getCategories")
    this.categories = await this.$store.state.categories
    this.categories.forEach(cat => {
      cat.isActive = false
      this.allChip.ServiceCount += cat.ServiceCount
    })
  },
}
</script>

<style scoped>
  

.nav-tags-wrapper {
  margin-left: 15px;
  margin-right: 15px;
}

.chip-item {
  background-color: rgb(39, 39, 39) !important;
  padding: 10px !important;
  margin: 5px !important;
  width: auto !important;
  margin-right: 5px !important;
  color:  rgb(39, 39, 39) !important;
  width: auto !important;
  margin-right: 5px !important;
  font-size: 0.8rem !important;
}

.chip-item:hover, .v-chip:hover {
  background-color: rgb(106, 0, 255) !important;
  color: none !important;
}

.v-chip--active {
  background-color: rgb(39, 39, 39) !important;
  color:  rgb(39, 39, 39) !important;
}

.chip-active {
  background-color: rgb(106, 0, 255) !important;
  color: white;
}

.category-filter-field {
  margin-right: 20px;
  margin-left: 20px;
}

.category-title {
  color: rgb(255, 255, 255);
  font-weight: 400;
}

.category-total {
  color: rgb(146, 146, 146);
  font-weight: 300;
  margin-left: 2px;
}

</style>