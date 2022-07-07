<template>
  <div class="sidebar">
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
    this.funcIndex = Holder.onUpdate((data) => {
      this.devices = data
    })
  },
  unmounted() {
    Holder.removeUpdate(this.funcIndex)
  },
  data() {
    return {
      funcIndex: 0,
      devices: [],
      search: '',
    }
  },
}
</script>

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
