<template>
  <div class="avatar-wrapper">
    <div class="avatar" v-html="avatar"></div>
  </div>
</template>

<script>
import { createAvatar } from '@dicebear/avatars';
import * as style from '@dicebear/avatars-identicon-sprites';

export default {
  name: "Avatar",
  props: ["hash"],
  data() {
    return {
      avatar: null
    }
  },
  mounted() {
    let svg = createAvatar(style, {
      seed: this.$props.hash,
    });
    this.avatar = svg
    console.log(svg)
  }
}
</script>

<style scoped>

.avatar-wrapper {
  width: 64px;
  height: 64px;
}

.avatar {
  background-color: white;
  margin: 4px;
  padding: 4px;
  border-radius: 5px;
}

</style>