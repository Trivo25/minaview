<template>
  <div class="card-wrapper">
    <div class="card">
      <div class="card-header">
        <img class="item-logo" :src="service.ServiceLogo"/>
      </div>
      <h1 class="project-title">{{ service.ServiceName }}</h1>
      <div class="card-content">
        <p>{{ service.ServiceDescription }}</p>
      </div>
      <div class="card-chips">
        <v-chip
          v-for="(tag, t) in tags"
          :key="t"
          class="tag-chip"
          outlined
        >
          <span>{{tag.CategoryTitle.toUpperCase()}}</span>
        </v-chip>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "ItemCard",
  props: ["service", "categories"],
  data() {
    return {
      tags: []
    }
  },
  mounted() {
    this.$props.categories.forEach(category => {
      this.$props.service.CategoryKeys.forEach(key => {
        if(category.CategoryKey == key) {
          this.tags.push(category)
        }
      })
    })
  },

}
</script>

<style scoped>

.card-wrapper {
  border-radius: 10px;
  text-align: center;
  max-height: auto;
}

.item-logo {
  max-height: 150px;
  width: auto;
  max-width: 100%;
  height: auto;
  padding: 15px;
}

.project-title {
  margin-top: 15px;
  font-weight: 600;
  color: grey;
}

.card { 
  background-color: #171717;
  border-radius: 15px;
  padding: 15px;
  width: auto;
  max-height: 100%;
  cursor: pointer;
  animation: hoverCardIn .5s forwards;
}

.card-chips {
  margin-top: 15px;
}

.card:hover {
  animation: hoverCardOut .5s forwards;
  background-color: #202020;
}

@keyframes hoverCardOut {
  0% {
    transform: scale(1.0)
  }
  100% {
    transform: scale(1.03)
  }
}

@keyframes hoverCardIn {
  0% {
    transform: scale(1.03)
  }
  100% {
    transform: scale(1.0)
  }
}

.tag-chip {
  font-weight: 500;
  font-size: 12px;
  color: rgba(255, 115, 0, 0.4);
  border-color: rgb(85, 85, 85);
  margin: 2px;
}

.card-content p {
  font-family: "Roboto"
}
</style>