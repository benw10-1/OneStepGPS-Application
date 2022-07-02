<template>
  <div class="container">
    <SideBar />
  </div>
</template>

<script>
import SideBar from './components/SideBar.vue'

export default {
  name: 'App',
  components: {
    SideBar
  }
}
</script>

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
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-rendering: optimizeLegibility;
  width: 100vw;
  height: 100vh;
}
.container {
  display: flex;
  flex-direction: row;
  height: 100%;
  width: 100%;
}
</style>
