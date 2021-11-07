<template>
  <div class="faq">
    <h1 class="faq-title">Frequently Asked Questions</h1>
    <div class="search-form-wrapper">
      <v-text-field
        placeholder="Searching for .."
        class="search-form"
        v-model="filter"
      >

      </v-text-field>
    </div>
    <div class="content">

      <template v-for="(subsection, s) in filteredItems">
        <div :key="s">
          <h2 class="sub-section">{{(s + 1) + ' - ' + subsection.title}}</h2> 
          <v-divider></v-divider>
          <template v-for="(question, i) in subsection.items">
            <div :key="i">
              <Drawer  :title="(s + 1) + '.' +(i + 1) + ' - ' +  question.title" :content="question.content" :references="question.references" />
            </div>
          </template>
        </div>
      </template>

    </div>
  </div>
</template>

<script>
import Drawer from "../components/Drawer.vue"

export default {
  name: "faq",
  layout: "faq-layout",
  props: [],
  components: {
    Drawer
  },
  data() {
    return {
      items: [],
      filter: "",
    }
  },
  async mounted() {
    const response = await fetch('/faq.json');
    let json = await response.json()
    this.items = json
  },
  computed: {
    filteredItems() {
      if(this.filter == "") {
        return this.items
      }
      
      // let filteredSections = new Array()

      // this.items.forEach(section => {
      //   if (section.title.toLowerCase().includes(this.filter.toLowerCase()))
      //     filteredSections.push(section)
      // })


      // return filteredSections


      return this.items.filter(item => item.items.some(s => s.title.toLowerCase().includes(this.filter.toLowerCase())));

    }  
  }
}
</script>


<style scoped>
* {
  -webkit-text-stroke: 0.5px rgb(0, 0, 0);
}
.faq {
  justify-content: center;
  text-align: center;
  width: 100%;
  height: 600px !important;
}

.content {
  width: 100%;
  justify-content: left;
  text-align: left;
  float: left;
}

.sub-section {
  margin-left: 15px;
  font-weight: 400;
}

.faq-title {
  font-size: 1.5rem;
  font-weight: 400;
  color: rgb(255, 255, 255);

}

.search-form-wrapper {
  justify-content: center;
  align-items: center;
}

.search-form {
  width: 400px;

}

</style>