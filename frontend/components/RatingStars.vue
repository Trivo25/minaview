<template>
  <div class="star-wrapper">
    <template v-for="(star, k) in stars">
      <v-icon :key="k" :color="unrated ? 'grey' : 'yellow'">
        {{star}}
      </v-icon>
    </template>
    <span class="average">{{ Math.ceil((this.upStars/(this.upStars + this.downStars) * 5) /0.5)*0.5 }}</span>
    <span class="total-votes">({{ upStars + downStars }})</span>
  </div>
</template>

<script>

export default {
  name: "RatingStars",
  props: ["up", "down"],
  data() {
    return {
      upStars: 0,
      downStars: 0,
      stars: [],
      unrated: false
    }
  },
  mounted() {

    this.upStars = this.$props.up
    this.downStars = this.$props.down

    if(this.upStars == null || this.downStars == null) {
      this.upStars = 0
      this.downStars = 0
    }

    if(this.upStars == 0 && this.downStars == 0) {
      this.unrated = true
      this.stars = [
        "mdi-star",
        "mdi-star",
        "mdi-star",
        "mdi-star",
        "mdi-star"
      ]
      return
    }

    let stars = []
    let k = Math.ceil((this.upStars/(this.upStars + this.downStars) * 5) /0.5)*0.5
    console.log(k)
    switch (k) {
      case 1:
        stars.push("mdi-star")
        break;
      case 1.5:
        stars.push("mdi-star")
        stars.push("mdi-star-half")
        break;
      case 2:
        stars.push("mdi-star")
        stars.push("mdi-star")
        break;
      case 2.5:
        stars.push("mdi-star")
        stars.push("mdi-star")
        stars.push("mdi-star-half")
        break; 
      case 3:
        stars.push("mdi-star")
        stars.push("mdi-star")
        stars.push("mdi-star")
        break; 
      case 3.5:
        stars.push("mdi-star")
        stars.push("mdi-star")
        stars.push("mdi-star")
        stars.push("mdi-star-half")
        break; 
      case 4:
        stars.push("mdi-star")
        stars.push("mdi-star")
        stars.push("mdi-star")
        stars.push("mdi-star")
        break; 
      case 4.5:
        stars.push("mdi-star")
        stars.push("mdi-star")
        stars.push("mdi-star")
        stars.push("mdi-star")
        stars.push("mdi-star-half")
        break; 
      case 5:
        stars.push("mdi-star")
        stars.push("mdi-star")
        stars.push("mdi-star")
        stars.push("mdi-star")
        stars.push("mdi-star")
        break; 
      default:
        break;
    }
    this.stars = stars
  },
}
</script>

<style scoped>

.star-wrapper {
  /* border: solid 1px rgb(77, 77, 77); */
  margin: 0;
  padding: 0;
  width: auto;
}

.total-votes {
  color: grey;
  font-size: 0.9rem;
}

.average {
  color: rgb(187, 187, 187);
}

</style>