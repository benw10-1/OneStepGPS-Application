<template>
  <div class="sidebar" @click="Holder.store.log">
    <div class="sorting-container">
      <div class="search-bar">
        <!-- <input type="text" placeholder="Search" v-model="search" @keyup.enter="search"> -->
      </div>
      <div class="">

      </div>
    </div>
    <div v-for="device in Holder.store.data" class="device-container" v-bind:key="device.device_id">
      <div class=""></div>
    </div>
  </div>
</template>

<script>
import { Holder } from '../helpers'

export default {
  name: 'SideBar',
  props: {

  },
  mounted() {
    Holder.onUpdate(this.$forceUpdate)
  },
  unmounted() {
    Holder.removeUpdate(this.$forceUpdate)
  },
  data() {
    return {
      Holder
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

.device-container {
  flex: 1;
  height: 42px;
  display: flex;
  align-items: center;
  flex-direction: row;
}

</style>
