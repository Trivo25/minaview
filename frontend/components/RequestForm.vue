<template>
  <div class="request-form">
    <div
      height="70%"
    >


      <v-row>
        <v-col cols=6>
          <v-text-field
            placeholder="Service Name"
            :rules="[(v => (v || '' ).length <= 50 || 'Description can not be greater than 50 characters'), v => (v || '' ).length >= 4 || 'Description can not be less than 4 characters']"
          />
        </v-col>
        <v-col cols=6>
          <v-text-field
            placeholder="Website, GitHub, .."
            :rules="[(v => (v || '' ).length <= 150 || 'Description can not be greater than 150 characters'), v => (v || '' ).length >= 4 || 'Description can not be less than 4 characters']"
          />
        </v-col>
      </v-row>
      <v-row>
        <v-col cols=6>
          <v-text-field
            placeholder="Creator (Group, Person, ..)"
            :rules="[(v => (v || '' ).length <= 50 || 'Description can not be greater than 50 characters'), v => (v || '' ).length >= 4 || 'Description can not be less than 4 characters']"
          />
        </v-col>
        <v-col cols=6>
          <v-textarea
            placeholder="Description of what your Service does!"
            :rules="[(v => (v || '' ).length <= 250 || 'Description can not be greater than 250 characters'), v => (v || '' ).length >= 80 || 'Description can not be less than 80 characters']"
          />
        </v-col>
      </v-row>
      <v-row>
        <v-col cols=6>
          <v-text-field
            placeholder="Link to SVG or PNG Logo"
            :rules="[(v => (v || '' ).length <= 150 || 'Description can not be greater than 150 characters'), v => (v || '' ).length >= 4 || 'Description can not be less than 4 characters']"
          />
        </v-col>
        <v-col cols=6>
          <v-text-field
            placeholder="Contact (E-Mail, Discord, Slack, ..)"
            :rules="[(v => (v || '' ).length <= 50 || 'Description can not be greater than 50 characters'), v => (v || '' ).length >= 4 || 'Description can not be less than 4 characters']"
          />
        </v-col>
      </v-row>
      <v-row>
        <v-col cols=6>
          <h1 class="cat">Available Categories</h1>
            <v-chip
              exact
              label
              class="chip-item"
            >
              <span class="category-title">{{ 'Wallets'.toUpperCase() }}</span>
            </v-chip>
        </v-col>
        <v-col cols=6>
          <h1 class="cat">Selected Categories</h1>
          <v-chip>test</v-chip>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols=6>
          <v-text-field
            placeholder="No category fits your need? Request a new one!"
          />
        </v-col>
      </v-row>

      <v-row>
        <v-col>
          <v-btn
          outlined
          plain
          class="service-button"
          >
            Wooooooooooosh!
          </v-btn>
        </v-col>
      </v-row>
    </div>
  </div>
</template>

<script>
export default {
  name: "RequestForm",
  data() {
    return {

    }
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

.v-chip {
  background-color: rgb(39, 39, 39) !important;
  padding: 10px !important;
  margin: 5px !important;
  width: auto !important;
  margin-right: 5px !important;
  color:  rgb(39, 39, 39) !important;
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

</style>