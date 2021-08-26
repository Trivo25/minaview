<template>
  <div class="card-wrapper">
    <div class="card">
      <div class="card-header">
        <img class="item-logo" :src="project.logo"/>
      </div>
      <h1 class="project-title">{{ project.title }}</h1>
      <div class="card-content">
        {{ project.description }}
      </div>
      <div class="card-chips">
        <v-chip
          v-for="(tag, t) in project.tags"
          :key="t"
          class="tag-chip"
          outlined
        >
          <span>{{tag.toUpperCase()}}</span>
        </v-chip>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "ItemCard",
  props: ["project"],
  data() {
    return {

    }
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
  font-weight: 500;
  color: grey;
}

.card { 
  background-color: #171717;
  border-radius: 15px;
  padding: 15px;
  width: auto;
  max-height: 100%;
  cursor: pointer;
  animation: hoverCardIn 1s forwards;
}

.card-chips {
  margin-top: 15px;
}

.card:hover {
  animation: hoverCardOut 1s forwards;
}

@keyframes hoverCardOut {
  0% {
    transform: scale(1.0)
  }
  100% {
    transform: scale(1.05)
  }
}

@keyframes hoverCardIn {
  0% {
    transform: scale(1.05)
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

</style>