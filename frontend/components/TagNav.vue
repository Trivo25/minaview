<template>
  <div class="tag-nav-wrapper">
    <v-text-field
      class="category-filter-field"
      v-model="searchCategoryFilter"
      placeholder="e.g. Wallets"
    />
    <v-chip
      :to="allChip.to"
      router
      exact
      label
      class="chip-item"
      :class="{'chip-active' : allChip.isActive }"
      @click="clearSelection()"
    >
      <span class="category-title">{{ allChip.title.toUpperCase() }}</span>
      <span class="category-total">({{ allChip.total }})</span>
    </v-chip>
    <v-chip
      v-for="(item, i) in items"
      :key="i"
      :to="routeParams"
      router
      exact
      label
      v-show="filterCategories(item)"
      class="chip-item"
      :class="{'chip-active' : item.isActive }"
      @click="changeSelection(item)"
    >
      <span class="category-title">{{ item.title.toUpperCase() }}</span>
      <span class="category-total">({{ item.total }})</span>
    </v-chip>
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
          icon: 'mdi-apps',
          iconColor: "#5442f5",
          title: 'All',
          to: '/ecosystem',
          total: 5,
          isActive: true
        },
      items: [
        {
          icon: 'mdi-chart-bubble',
          iconColor: "#7542f5",
          title: 'Official Resources',
          param: 'official-resources',
          total: 5,
          isActive: false
        },
        {
          icon: 'mdi-chart-bubble',
          iconColor: "#8d42f5",
          title: 'Wallets',
          param: 'wallets',
          total: 5,
          isActive: false
        },
        {
          icon: 'mdi-chart-bubble',
          iconColor: "#a442f5",
          title: 'Explorers',
          param: 'explorers',
          total: 5,
          isActive: false
        },
        {
          icon: 'mdi-chart-bubble',
          iconColor: "#b642f5",
          title: 'Tools',
          param: 'tools',
          total: 5,
          isActive: false
        },
        {
          icon: 'mdi-chart-bubble',
          iconColor: "#8d42f5",
          title: 'Staking Pools',
          param: 'staking-pools',
          total: 5,
          isActive: false
        },
        {
          icon: 'mdi-chart-bubble',
          iconColor: "#d442f5",
          title: 'Monitoring and Dashboards',
          param: 'monitoring-und-dashboards',
          total: 5,
          isActive: false
        },
        {
          icon: 'mdi-chart-bubble',
          iconColor: "#f542ef",
          title: 'Ledger Apps',
          param: 'ledger',
          total: 5,
          isActive: false
        },
        {
          icon: 'mdi-chart-bubble',
          iconColor: "#f5426f",
          title: 'News, Resources and Articles',
          param: 'news-articles',
          total: 5,
          isActive: false
        },
        {
          icon: 'mdi-chart-bubble',
          iconColor: "#f54242",
          title: 'International Communities',
          param: 'communities',
          total: 5,
          isActive: false
        },
      ],
    }
  },
  methods: {
    filterCategories(item) {
      return ((item.title).toLowerCase()).includes((this.searchCategoryFilter).toLowerCase())
    },
    changeSelection(item) {
      this.params.includes(item.param) && item.isActive ?
        this.params.splice(this.params.indexOf(item.param), 1) :
        this.params.push(item.param); item.isActive = !item.isActive;

      this.allChip.isActive = false;
      // this.$router.push({
      //   path: this.parsedTargetPath
      // })
    },
    clearSelection() {
      this.allChip.isActive = true
      this.params = new Array()
      this.items.forEach(item => {
        item.isActive = false
      })
    }
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
  }
}
</script>

<style scoped>
  

.tag-nav-wrapper {

}

.chip-item {
  background-color: rgb(39, 39, 39) !important;
  padding: 10px !important;
  margin: 5px !important;
  width: auto !important;
  margin-right: 5px !important;
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