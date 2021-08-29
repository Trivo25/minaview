<template>
  <div class="request-form">
    <div
      height="70%"
    >

    <h1 style="color: red; font-weight: 300;">For changes or new services please follow <a href="https://github.com/Trivo25/mina-view/issues/6">this link </a></h1>
      <v-row>
        <v-col cols=6>
          <v-text-field
            class="text-field"
            v-model="request.ServiceName"
            placeholder="Service Name"
            :rules="[(v => (v || '' ).length <= 50 || 'Description can not be greater than 50 characters'), v => (v || '' ).length >= 4 || 'Description can not be less than 4 characters']"
          />
        </v-col>
        <v-col cols=6>
          <v-text-field
            class="text-field"
            v-model="request.ServiceWebsite"
            placeholder="Website, GitHub, .."
            :rules="[(v => (v || '' ).length <= 150 || 'Description can not be greater than 150 characters'), v => (v || '' ).length >= 4 || 'Description can not be less than 4 characters']"
          />
        </v-col>
      </v-row>
      <v-row>
        <v-col cols=6>
          <v-text-field
            class="text-field"
            v-model="request.ServiceCreator"
            placeholder="Creator (Group, Person, ..)"
            :rules="[(v => (v || '' ).length <= 50 || 'Description can not be greater than 50 characters'), v => (v || '' ).length >= 4 || 'Description can not be less than 4 characters']"
          />
        </v-col>
        <v-col cols=6>
          <v-textarea
            class="text-field"

            v-model="request.ServiceDescription"
            placeholder="Description of what your Service does!"
            :rules="[(v => (v || '' ).length <= 250 || 'Description can not be greater than 250 characters'), v => (v || '' ).length >= 80 || 'Description can not be less than 80 characters']"
          />
        </v-col>
      </v-row>
      <v-row>
        <v-col cols=6>
          <v-text-field
            class="text-field"
            v-model="request.ServiceLogo"
            placeholder="Link to SVG or PNG Logo"
            :rules="[(v => (v || '' ).length <= 150 || 'Description can not be greater than 150 characters'), v => (v || '' ).length >= 4 || 'Description can not be less than 4 characters']"
          />
        </v-col>
        <!-- discord, telegram, email, slack, reddit, github -->
        <v-col cols=3>
          <v-text-field
            class="text-field"
            v-model="request.Twitter"
            placeholder="Twitter - e.g MinaProtocol"
            :rules="[(v => (v || '' ).length <= 50 || 'Description can not be greater than 50 characters'), v => (v || '' ).length >= 4 || 'Description can not be less than 4 characters']"
          />
        </v-col>
        <v-col cols=3>
          <v-text-field
            class="text-field"
            v-model="request.Discord"
            placeholder="Discord - e.g discord.gg/myCoolInvite"
            :rules="[(v => (v || '' ).length <= 50 || 'Description can not be greater than 50 characters'), v => (v || '' ).length >= 4 || 'Description can not be less than 4 characters']"
          />
        </v-col>
        <v-col cols=3>
          <v-text-field
            class="text-field"
            v-model="request.Telegram"
            placeholder="Telegram - e.g MyGroupName"
            :rules="[(v => (v || '' ).length <= 50 || 'Description can not be greater than 50 characters'), v => (v || '' ).length >= 4 || 'Description can not be less than 4 characters']"
          />
        </v-col>
        <v-col cols=3>
          <v-text-field
            class="text-field"
            v-model="request.Email"
            placeholder="Email - e.g funnyMail@smile.com"
            :rules="[(v => (v || '' ).length <= 50 || 'Description can not be greater than 50 characters'), v => (v || '' ).length >= 4 || 'Description can not be less than 4 characters']"
          />
        </v-col>
        <v-col cols=3>
          <v-text-field
            class="text-field"
            v-model="request.Slack"
            placeholder="Slack"
            :rules="[(v => (v || '' ).length <= 50 || 'Description can not be greater than 50 characters'), v => (v || '' ).length >= 4 || 'Description can not be less than 4 characters']"
          />
        </v-col>
        <v-col cols=3>
          <v-text-field
            class="text-field"
            v-model="request.Reddit"
            placeholder="Reddit - e.g /r/minaprotocol or /u/Trivo_"
            :rules="[(v => (v || '' ).length <= 50 || 'Description can not be greater than 50 characters'), v => (v || '' ).length >= 4 || 'Description can not be less than 4 characters']"
          />
        </v-col>
        <v-col cols=3>
          <v-text-field
            class="text-field"
            v-model="request.Github"
            placeholder="Github - e.g Trivo25"
            :rules="[(v => (v || '' ).length <= 50 || 'Description can not be greater than 50 characters'), v => (v || '' ).length >= 4 || 'Description can not be less than 4 characters']"
          />
        </v-col>
      </v-row>
      <v-row>
        <v-col cols=6>
          <h1 class="cat">Available Categories</h1>
            <div class="card-chips">
              <v-chip
                v-for="(tag, t) in availableCategories"
                :key="t"
                class="tag-chip"
                outlined
                close
                close-icon="mdi-plus"
                @click:close="addSelected(tag)"
              >
                <span>{{tag.CategoryTitle.toUpperCase()}}</span>
              </v-chip>
            </div>  
        </v-col>
        <v-col cols=6>
          <h1 class="cat">Selected Categories</h1>
            <div class="card-chips">
              <v-chip
                v-for="(tag, t) in selectedCategories"
                :key="t"
                class="tag-chip"
                outlined
                close
                close-icon="mdi-close-outline"
                @click:close="removeSelected(tag)"
              >
                <span>{{tag.CategoryTitle.toUpperCase()}}</span>
              </v-chip>
            </div>  
        </v-col>
      </v-row>
      <v-row>
        <v-col cols=6>
          <v-text-field
            class="text-field"
            v-model="request.NewCategory"
            placeholder="No category fits your need? Request a new one!"
          />
        </v-col>
      </v-row>

      <v-row>
        <v-col>
          <v-btn
          outlined
          plain
          @click="submit()"
          class="service-button"
          >
            Wooooooooooosh!
          </v-btn>
        </v-col>
      </v-row>
    </div>
    <Error @closeNotification="error.show = false" v-if="error.show" :error="error.error" :type="error.type" />
  </div>
</template>

<script>
import Error from "./Error.vue"


export default {
  name: "RequestForm",
  components: {
    Error
  },
  data() {
    return {
      request: {},
      availableCategories: [],
      selectedCategories: [],
      error: {
        error: null,
        show: false,
        type:"success"
      }
    }
  },
  methods: {
    removeSelected(tag) {
      // NOTE: weird workaround but the normal methods didnt work for some reason, conflicting with vuex ?
      let temp = []
      this.selectedCategories.forEach(el => {
        if(el.CategoryID != tag.CategoryID) temp.push(el)
      })
      this.selectedCategories = temp
      this.availableCategories.push(tag)
    },
    addSelected(tag) {
      // NOTE: weird workaround but the normal methods didnt work for some reason, conflicting with vuex ?
      let temp = []
      this.availableCategories.forEach(el => {
        if(el.CategoryID != tag.CategoryID) temp.push(el)
      })
      this.availableCategories = temp
      this.selectedCategories.push(tag)
    },
    async submit() {

      this.request.Categories = this.selectedCategories


      let res = await this.$axios.post("http://localhost:8000/requestService", this.request)


      this.error.error = {
        Error: "Success! Your request has been submitted",
        Code: ""
      }
      this.error.show = true
      this.error.type = "success"
    }
  },
  async mounted() {
    if(this.$store.state.categories.length == 0) {
      await this.$store.dispatch("getCategories")
        .then(res => {
          this.availableCategories = res
        }, error => {
          this.error.error = error
          this.error.show = true
        })
    }
    this.availableCategories = this.$store.state.categories
  }
}
</script>

<style scoped>

.request-form {

}

.service-button {
  background: rgb(75, 144, 0);
  background: linear-gradient(148deg, rgba(91,0,255,0.44629521730567223) 0%, rgba(156,0,255,1) 50%, rgba(202,255,234,0.7708918539325843) 100%);
  color: rgb(204, 204, 204);
  padding: 10px;
  border: 1px solid black;
  border-radius: 3px;
  font-weight: 500;
}


.category-filter-field {
  margin-right: 20px;
  margin-left: 20px;
}

.category-title {
  color: rgb(255, 255, 255);
  font-weight: 400;
}
.cat {
  color: rgb(177, 177, 177);
  font-weight: 300;
  font-size: 1.5rem;
}

.card-chips {
  margin-top: 15px;
}

.tag-chip {
  font-weight: 500;
  font-size: 12px;
  color: rgba(255, 115, 0, 0.4);
  border-color: rgb(85, 85, 85);
  margin: 2px;
}

.theme--light .tag-chip {
  color: rgba(255, 115, 0, 1);
}

.theme--light .text-field {
  color: black !important;
  font-weight: 500 !important;
}

</style>