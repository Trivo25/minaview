<template>
  <div class="comment-wrapper">
    <div class="form-section">
      <h2>What is on your mind?</h2>
      <v-textarea
        class="text-field"
        placeholder="Share your experience!"
        rows="3"
        :rules="[(v => (v || '' ).length <= 250 || 'Description can not be greater than 250 characters'), v => (v || '' ).length >= 10 || 'Description can not be less than 10 characters']"
      />
      <v-btn
        class="submit-button"
      >SUBMIT 
      <v-icon>mdi-send</v-icon>
      </v-btn>
    </div>
    <div class="comment-list">
      <Comment />
      <div class="seperator"></div>
      <Comment />
      <div class="seperator"></div>
      <Comment />
      <div class="seperator"></div>
      <Comment />
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

}

.seperator {
  width: 100%;
  margin-top: 20px;
  margin-bottom: 20px;
  border-bottom: solid 1px rgba(160, 160, 160, 0.4);
  border-radius: 50px;
  
}

.comment-list, .form-section {
  position: relative;
  width: 100%;
  margin-top: 20px;
  text-align: left;
  /* background-color: rgb(41, 43, 59); */
  background-color: rgb(17, 44, 56);
  border-radius: 15px;
  justify-content: center;
  align-items: center;
  text-align: center;
  padding: 5px;
  border: solid 1px rgba(112, 112, 112, 0.5);
}
.form-section h2 {
  font-weight: 200;
  color: rgb(146, 146, 146);
  text-shadow: 2px 2px 2px rgb(42, 42, 42);
}

.form-section {
  justify-content: center;
  align-items: center;
  text-align: center;
}

.text-field {
  padding: 15px;
}

.submit-button {
  border: rgb(91,0,255) solid 2px !important;
  background: none !important;
  color: rgb(255, 255, 255);
  padding: 10px;
  border: 1px solid black;
  border-radius: 3px;
  font-weight: 500;
  margin-bottom: 15px;
}
</style>