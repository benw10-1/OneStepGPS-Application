<template>
  <div class="container">
    <DashBoard v-if="loggedIn" />
    <LogIn v-else />
  </div>
</template>

<script>
import DashBoard from './pages/DashBoard.vue'
import LogIn from './pages/LogIn.vue'
import { Auth } from './helpers'

export default {
  name: 'App',
  components: {
    DashBoard,
    LogIn
  },
  data() {
    return {
      loggedIn: Auth.loggedIn()
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
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-rendering: optimizeLegibility;
  width: 100vw;
  height: 100vh;
}

.container {
  height: 100%;
  width: 100%;
  background-color: #fafafa;
}
</style>
