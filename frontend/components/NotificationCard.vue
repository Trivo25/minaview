<template>
   <div class="notification-card">
      <div class="top">
        <v-icon @click="$emit('closeNotification')" class="close-icon">mdi-close</v-icon>
      </div>
     <div class="content">
       <div class="left">
          <a style="font-size: 2rem; cursor: default;">&#128526;</a>
       </div>

       <div class="right">
         <span>
            Like what you are seeing?
            Please consider leaving a Star on <a href="https://github.com/Trivo25/mina-view" target="_blank">GitHub!</a>
         </span>
       </div>
     </div>
  </div>
</template>

<script>
export default {
  name:"NotificationCard",
  props: ["type"]
}
</script>

<style scoped>
.notification-card {
  position: absolute;
  display: inline;
  z-index: 15;
  left: calc(50% - 175px);
  top: 10px;
  width: 350px;
  height: 80px;
  border-radius: 5px;
  border: 1px solid rgba(6, 44, 241, 0.46);
  background-color: rgba(7, 73, 149, 0.12156862745098039);
  box-shadow: 0px 0px 2px #0396ff;
  color: #0396ff;
  text-shadow: 2px 1px #00040a;
  font-weight: 400;
  transition:0.5s;
  align-content: center;
  margin-bottom: 0;
  justify-items: center;
  animation: floatIn 1s forwards;
  cursor: default;
}



.theme--light .notification-card {
  text-shadow: none;
}

@keyframes floatIn {
  0% {
    top: -50px;
  }
  100% {
    top: 10px;
  }
}
.content {
  position: absolute;
  display: flex;
  align-items: center;
  text-align: center;
  width: 100%;
}

.left {
  width: 20%;
  position: absolute;
  left: 0;
  margin-top: 25px;
}

.right {
  width: 80%;
  position: absolute;
  right: 0;
  margin-top: 25px;
}

.top {
  width: 100%;
  align-content: right;
  text-align: right;
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