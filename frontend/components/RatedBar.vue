<template>
  <div>
    <div class="bar-wrapper" v-if="voted">
      <div class="info-card">
        <h3>You voted: {{ userVote ? '👍' : '👎'}}</h3>
        <p>Check out what other users voted!</p>
      </div>
      <div class="bar-card">
        <div class="bar">
          <div class="green" v-bind:style="{ width: (up/(up + down) * 100) + '%', backgroundColor: 'green' }">
          </div>
          <div class="red" v-bind:style="{ width: (down/(up + down) * 100) + '%', backgroundColor: 'red'  }">
          </div>
        </div>
        <div class="types">
          <div class="up">
            <span class="vote-count">{{ up }} Votes - Great 👍</span>
          </div>
          <div class="down">
            <span class="vote-count">{{ down }} Votes - Bad 👎</span> 
          </div>
        </div>
      </div>
    </div>
    <div class="bar-wrapper" v-else>
      <div class="vote-selection">
        <div class="header">
          <h2>What do you think about this project?</h2>
        </div>
        <div class="vote-buttons">
          <span class="button" @click="vote(true)">👍 Amazing!</span>
          <span class="button" @click="vote(false)">👎 Meh :/</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import FingerprintJS from '@fingerprintjs/fingerprintjs'

export default {
  name: "RatedBar",
  props: ["hash", "up", "down"],
  data() {
    return {
      up: 0,
      down: 0,
      userVote: true,
      hasVoted: false,
      nextVote: null
    }
  },
  computed: {
    voted() {
      if(this.hasVoted) return true
      // if (this.nextVote <= Date.now()) {
      //   return false
      // } else if(this.nextVote == null) {
      //   return false
      // }
      // return true
    }
  },
  methods: {
    async vote(v) {
      let up = 0
      let down = 0
      let fingerprint = await this.getFingerprint()
      if(v) {
        this.userVote = true
        this.up += 1
        up += 1
      } else {
        this.userVote = false
        this.down += 1
        down += 1  
      }
      let n = Date.now() + 8.64e+7
      // this.nextVote = n
      // localStorage.setItem(`nv_${this.$props.hash}`, n)
      this.hasVoted = true

      let res = (await this.$axios.post("/vote", {
        serviceHash: this.$props.hash,
        fingerprint: fingerprint,
        up: up,
        down: down,
        nextVote: n,
        currentTime: Date.now()
      }))

    },
    async getFingerprint() {
      const fpPromise = FingerprintJS.load()
      const fp = await fpPromise
      const result = await fp.get()

      return result.visitorId
    }
  },
  async mounted() {
    // this.nextVote = localStorage.getItem(`nv_${this.$props.hash}`)

    // TODO: get last vote by fingerprint

    this.up = this.$props.up
    this.down = this.$props.down
    let fingerprint = await this.getFingerprint()

    let res = (await this.$axios.post("/hasVoted", {
        fingerprint: fingerprint,
        serviceHash: this.$props.hash
      })).data.Data
    res.HasVoted ? this.hasVoted = true : this.hasVoted = false

  }
}
</script>

<style scoped>
  
.bar-wrapper {
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
  border-radius: 15px;
  width: 100%;
  height: 80px;
  background-color: rgb(17, 44, 56);
  max-width: 700px;
  margin: auto;
  border: solid 1px rgba(112, 112, 112, 0.5);
}

.info-card {
  width: 30%;
  height: 100%;
  align-items: center;
  justify-content: center;
  text-align: center;
}

.bar-card {
  width: 70%;
  height: 100%;
}

.types {
  width: 100%;
  display: flex;
  margin-top: 1px;
}

.up {
  width: 50%;
}

.down {
  width: 50%;
}

.bar {
  display: flex;
  margin: 15px;
  margin-bottom: 0;
  height: 20px;
  border-radius: 15px;
}

.green {
  justify-content: center;
  align-items: center;
  text-align: center;
  margin: 0;
  height: 100%;
  background-color: white;
  border-top-left-radius: 15px;
  border-bottom-left-radius: 15px;
  /* width: 50%; */
}

.red {
  justify-content: center;
  align-items: center;
  text-align: center;
  margin: 0;
  height: 100%;
  background-color: white;
  border-top-right-radius: 15px;
  border-bottom-right-radius: 15px;
  /* width: 50%; */
}

.up span, .down span {
  margin: 0;
  font-size: 0.8rem;
  color: rgb(255, 255, 255);
  font-weight: 600;
}

.vote-selection {

}

.header {
  width: 100%;
  align-items: center;
  justify-content: center;
  text-align: center;
  margin-bottom: 10px;
}

.header h2 {
  margin: 0;
  font-size: 1.1rem;
  color: rgb(255, 255, 255);
  font-weight: 600;
}

.vote-buttons {
  width: 100%;
  align-items: center;
  justify-content: center;
  text-align: center;
}

.button {
  border: solid 1px rgb(98, 98, 98);  
  margin: 15px;
  padding: 6px;
  border-radius: 6px;
  color: rgb(255, 255, 255);
  font-weight: 600;
  animation: buttonHoverOut .5s forwards; 
}

.button:hover {
  cursor: pointer;
  animation: buttonHoverIn .5s forwards; 
}

@keyframes buttonHoverIn {
  0% {
    border: solid 1px rgb(98, 98, 98);  
  }
  100% {
    border: solid 1px rgb(212, 212, 212);  
  }
}

@keyframes buttonHoverOut {
  0% {
    border: solid 1px rgb(212, 212, 212);  
  }
  100% {
    border: solid 1px rgb(98, 98, 98);  
  }
}

</style>