<template>
   <div v-bind:class="[ type == 'error' ? 'e' : 's']">
      <div class="top">
        <v-icon @click="$emit('closeNotification')" class="close-icon">mdi-close</v-icon>
      </div>
     <div class="content">
      <span>{{error.Error + " " +  error.Code}}</span>
     </div>
  </div>
</template>

<script>
export default {
  props: ["error", "type"],
  name:"Error"
}
</script>

<style scoped>
.s, .e {
  position: fixed;
  z-index: 15;
  width: 300px;
  height: 100px;
  right: 20px;
  bottom:10px;
  border-radius: 5px;
  border: 1px solid rgba(241, 6, 6, 0.46);
  background-color: rgba(211, 0, 0, 0.122);
  box-shadow: 0px 0px 2px #0396ff;
  color: #c01212;
  text-shadow: 2px 1px #00040a;
  font-weight: 400;
  transition:0.5s;
  align-content: center;
  justify-items: center;
  animation: floatIn 2s forwards;
  cursor: default; 
}


.s {
  border: 1px solid rgba(147, 241, 6, 0.46);
  background-color: rgba(6, 177, 38, 0.122);
  box-shadow: 0px 0px 2px #0396ff;
  color: #12c01e;
}
/* 
.error {
  border: 1px solid rgba(241, 6, 6, 0.46);
  background-color: rgba(211, 0, 0, 0.122);
  box-shadow: 0px 0px 2px #0396ff;
  color: #c01212;
  text-shadow: 2px 1px #00040a;
}

.success {
  border: 1px solid rgba(241, 6, 6, 0.46);
  background-color: rgba(211, 0, 0, 0.122);
  box-shadow: 0px 0px 2px #0396ff;
  color: #c01212;
  text-shadow: 2px 1px #00040a;
} */

@keyframes floatIn {
  0% {
    bottom: -100px;
  }
  100% {
    bottom:10px;
  }
}
.content {
  position: absolute;
  display: flex;
  align-items: center;
  text-align: center;
  width: 100%;
}

.notification-card:hover{
  background-color: #00000059;
  transition:0.5s;
}

.close-icon {
  z-index: 15;
  color: #ffffff59 !important;
  cursor: pointer;
}



</style>