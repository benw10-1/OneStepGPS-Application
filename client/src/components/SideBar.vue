<template>
  <div class="sidebar">
    <div class="sorting-container">
      <div class="search-bar">
        <CustomInput placeholder="Search..." margin="0" height="40px" flexItem :onChange="setSearch" :value="search" />
      </div>
      <div class="">

      </div>
    </div>
    <div class="devices-container">
      <div v-for="device of display" class="device-container" v-bind:key="device?.device_id">
        <div class="data">{{ device?.display_name }}</div>
      </div>
    </div>
  </div>
</template>

<script>
import { Holder } from '../helpers'
import CustomInput from './CustomInput.vue'

export default {
  name: 'SideBar',
  components: {
    CustomInput
  },
  mounted() {
    Holder.onUpdate(this.updateDevices)
  },
  unmounted() {
    Holder.removeUpdate(this.updateDevices)
  },
  data() {
    return {
      devices: [],
      search: '',
      display: []
    }
  },
  methods: {
    updateDevices(data) {
      this.devices = data instanceof Array ? [...data] : { ...data }
      this.display = this.filtered()
    },
    setSearch(search) {
      this.search = search
      this.display = this.filtered()
    },
    filtered() {
      return this.devices.filter(device => {
        return device.display_name.toLowerCase().includes(this.search.toLowerCase())
      })
    }
  }
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

.search-bar {
  width: 100%;
  height: auto;
  background: #f5f5f5;
  padding: 10px;
  display: flex;
  align-items: center;
  margin: 0;
  box-sizing: border-box;
  border-bottom: 1px solid #e5e5e5;
}

.device-container {
  flex: 1;
  height: auto;
  display: flex;
  position: relative;
  align-items: center;
  flex-direction: row;
  box-sizing: border-box;
  padding: 10px;
}
</style>
