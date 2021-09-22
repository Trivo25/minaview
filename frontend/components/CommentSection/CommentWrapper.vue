<template>
  <div class="comment-wrapper">
    <div class="comment-list">
      <Comment />
      <div class="seperator"></div>
      <Comment />
      <div class="seperator"></div>
      <Comment />
      <div class="seperator"></div>
      <Comment />
    </div>
    <div class="form-section">

    </div>
  </div>
</template>

<script>
import Comment from "./Comment.vue"
export default {
  name: "CommentWrapper",
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
  padding: 5px;
}

.seperator {
  width: 100%;
  margin-top: 20px;
  margin-bottom: 20px;
  border-bottom: solid 1px rgba(160, 160, 160, 0.4);
  border-radius: 50px;
}
</style>