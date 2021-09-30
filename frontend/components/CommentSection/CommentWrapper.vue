<template>
  <div class="comment-wrapper">
    <div class="form-section">
      <h2>What is on your mind?</h2>
      <v-row>
        <v-col
          cols="4"
        >
          <v-text-field
            placeholder="Satoshi Nakamoto"
            label="Name"
            v-model="username"
          ></v-text-field>
        </v-col>
      </v-row>
      <v-textarea
        class="text-field"
        placeholder="Share your experience!"
        rows="3"
        v-model="comment"
        :rules="[(v => (v || '' ).length <= 250 || 'Description can not be greater than 250 characters'), v => (v || '' ).length >= 10 || 'Description can not be less than 10 characters']"
      />
      <span v-if="profane" class="profane-language-hint">Oooopsie wooopsie.. Please double check what you wrote - no profane language!</span>
      <span v-if="sent" class="message-sent-hint">Your comment has been sent and will be reviewed for spam and insults before posting.</span>
      <v-btn
        class="submit-button"
        @click="sendComment()"
      >SUBMIT 
      <v-icon>mdi-send</v-icon>
      </v-btn>
    </div>
    <div class="comment-list">
      <template v-for="(comment, c) in comments">
        <div :key="c">
          <Comment :name="comment.username" :message="comment.text" :hash="comment.hash" :time="comment.time"/>
          <div class="seperator"></div>
        </div>
      </template>
    </div>
  <!-- <Error @closeNotification="error.show = false" v-if="error.show" :error="error.error" :type="error.type" /> -->
  </div>
</template>

<script>
import Comment from "./Comment.vue"
import FingerprintJS from '@fingerprintjs/fingerprintjs'

export default {
  name: "CommentWrapper",
  props: ["comments", "service_hash"],
  components: {
    Comment
  },
  data() {
    return {
      username: "",
      comment: "",
      profane: false,
      sent: false,
      error: {
        error: null,
        show: false,
        type:"success"
      }
    }
  },
  methods: {
    check(msg) {
      var Filter = require('bad-words')
      var f = new Filter();
      return f.isProfane(msg)
    },
    async sendComment() {
      let fingerprint = await this.getFingerprint()
      console.log(fingerprint)
      this.sent = false
      if(this.check(this.username) || this.check(this.comment)) {
        this.profane = true
        return
      } else {
        this.profane = false


        let res = (await this.$axios.post("/comment", {
          service_hash: this.$props.service_hash,
          fingerprint: fingerprint,
          username: this.username,
          text: this.comment,

        }))


        this.sent = true
      }
    },
    async getFingerprint() {
      const fpPromise = FingerprintJS.load()
      const fp = await fpPromise
      const result = await fp.get()

      return result.visitorId
    }
  },
  mounted() {
    
  }
}
</script>

<style scoped>

.comment-wrapper {

}

.profane-language-hint {
  margin-top: 0;
  font-size: 0.8rem;
  color: red;
}
.message-sent-hint {
  margin-top: 0;
  font-size: 0.8rem;
  color: green;
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
  padding: 5px;
  margin-bottom: 0;
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