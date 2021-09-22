<template>
  <div class="comment-wrapper">
    <div class="left">
      <div class="avatar">
        <Avatar :hash="new Date().toLocaleString()"/>
      </div>
      <div class="date">
        <span>{{ new Date().toLocaleString() }}</span>
      </div>
    </div>
    <div class="right">
      <div class="comment-box">
        <p class="comment">
          Awesome Protocol! Just checked it out and I am amazed.. Can't wait to get started doing something awesome with it!
        </p>
      </div>
    </div>
  </div>
</template>

<script>
import Avatar from "./Avatar.vue"

export default {
  name: "Comment",
  components: {
    Avatar
  },
  data() {
    return {
    }
  },
}
</script>

<style scoped>

.comment-wrapper {
  background-color: rgb(41, 43, 59);
  border-radius: 15px;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}

.right {
  width: 70%;
  /* border: solid 1px white; */
}

.comment-box {
  border: solid 1px rgb(90, 90, 90);
  border-radius: 15px;
  margin: 5px;
  text-align: left;
  padding: 5px;
}

.left {
  width: 30%;
  justify-content: center;
  align-items: center;
  text-align: center;
}

.avatar {
  position: relative;
  left: calc(50% - 32px);
  justify-content: center;
  align-items: center;
  text-align: center;
}

.comment {
  font-weight: 400;
  font-style: italic;
  color: rgb(190, 190, 190);
}


</style>