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
        <div :class="'device-sidebar ' + (device.online ? device.drive_status : 'offline')"></div>
        <div class="device-inner">
          <div class="device-main-info">
            <div class="device-display-name"> 
              {{ device.display_name }}
            </div>
            <div class="device-status">
              {{ (device.online ? driveStates[device.drive_status] : "No signal") + ' ' + formatTime(device.drive_status_begin_time) }}
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { Holder, dateFormatter } from '../helpers'
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
      // all devices
      devices: [],
      search: '',
      // filtered devices to display
      display: [],
      colors: {
        off: "rgba(255, 0, 0, .67)",
        driving: "rgba(0, 255, 0, .67)",
        idle: "rgba(135, 206, 250, .67)",
        offline: "rgba(218, 223, 225, .8)",
      },
      driveStates: {
        off: "Stopped",
        driving: "Driving",
        idle: "Idle",
      },
      firstLoad: true,
    }
  },
  methods: {
    updateDevices(data) {
      if (!data) return
      // create copy of array to always trigger rerender
      this.devices = data instanceof Array ? [...data] : { ...data }
      this.display = this.filtered()
      if (this.firstLoad) {
        this.firstLoad = false
        this.$emit("loaded")
      }
    },
    setSearch(search) {
      this.search = search
      this.display = this.filtered()
    },
    // simple JS sort function to sort devices by display name
    filtered() {
      return this.devices.filter(device => {
        return device.display_name.toLowerCase().includes(this.search.toLowerCase().trim())
      })
    },
    formatTime(time) {
      const date = new Date() - new Date(time)

      return dateFormatter(date.valueOf() / 1000)
    },
  }
}
</script>

<style scoped>
.sidebar {
  width: 325px;
  height: 100%;
  background: white;
  overflow-x: hidden;
  padding: 0;
  box-shadow: 0 3px 1px -2px rgb(0 0 0 / 20%), 0 2px 2px 0 rgb(0 0 0 / 14%), 0 1px 5px 0 rgb(0 0 0 / 12%);
  outline: none;
}

.search-bar {
  width: 100%;
  height: auto;
  padding: 10px;
  display: flex;
  align-items: center;
  margin: 0;
  box-sizing: border-box;
  border-bottom: 1px solid rgba(0,0,0,.12);
}

.device-container {
  flex: 1;
  height: auto;
  display: flex;
  position: relative;
  align-items: center;
  flex-direction: row;
  box-sizing: border-box;
  padding: 10px 14px;
  border-bottom: 1px solid rgba(0,0,0,.12);
}

.device-sidebar {
  width: 6.5px;
  height: calc(100% - 20px);
  position: absolute;
  left: 0;
  z-index: 1;
  box-sizing: border-box;
}
.off {
  background: v-bind("colors.off");
}
.driving {
  background: v-bind("colors.driving");
}
.idle {
  background: v-bind("colors.idle");
}
.offline {
  background: v-bind("colors.offline");
}

.device-inner {
  height: 50px;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.device-display-name {
  font-size: 14px;
  font-weight: 500;
  color: rgba(0,0,0,.87);
}
.device-status {
  font-size: 12px;
  font-weight: 400;
  color: rgba(0,0,0,.54);
}
</style>
