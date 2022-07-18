<template>
  <div class="container">
    <Transition>
      <DashBoard v-if="loggedIn" />
      <LogIn v-else />
    </Transition>
  </div>
</template>

<script>
import DashBoard from './pages/DashBoard.vue'
import LogIn from './pages/LogIn.vue'
import { Auth } from './helpers'
import "material-icons/iconfont/filled.css"
import "material-icons/iconfont/outlined.css"

export default {
  name: 'App',
  components: {
    DashBoard,
    LogIn
  },
  data() {
    return {
      loggedIn: Auth.loggedIn() && Auth.getProfile()?.APIKey
    }
  },
  mounted() {
    Auth.onLogin((data) => {
      this.loggedIn = !!data?.APIKey
    })
    Auth.onLogout(() => {
      this.loggedIn = false
    })
  },
}
</script>
<style>
  @import url('https://fonts.googleapis.com/css2?family=Roboto:ital,wght@0,100;0,300;0,400;0,500;0,700;0,900;1,100;1,300;1,400;1,500;1,700;1,900&display=swap');
</style>
<style>
html {
  font-size: 16px;
  font-family: 'Roboto', sans-serif;
  margin: 0;
  padding: 0;
}

body {
  margin: 0;
  padding: 0;
}

#app {
  font-family: "Roboto", Helvetica, Arial, sans-serif;
  /* -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-rendering: optimizeLegibility; */
  width: 100vw;
  height: 100vh;
  overflow: hidden;
}

.container {
  height: 100%;
  width: 100%;
  background-color: #f5f5f5;
}

.v-enter-active,
.v-leave-active {
  transition: opacity 0.2s ease-in;
}

.v-enter-from,
.v-leave-to {
  opacity: 0;
}

p {
  font-size: 1.08rem;
  font-weight: 300;
  /* line-height: 1.5; */
  margin: 0;
  padding: 0;
}
</style>
