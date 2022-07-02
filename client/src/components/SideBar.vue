<template>
  <div class="sidebar" @click="apiStore.refresh">
    {{ apiStore.data }}
  </div>
</template>

<script>
import { apiStore } from '../stores'

export default {
  name: 'SideBar',
  props: {
    
  },
  data() {
    return {
      apiStore
    }
  },
}

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.sidebar {
  width: 400px;
  height: 100%;
  background: #f5f5f5;
  overflow-x: hidden;
  padding: 0;
}
</style>
