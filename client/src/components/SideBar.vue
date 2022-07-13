<template>
  <div class="sidebar">
    <div class="sorting-container">
      <div class="search-bar">
        <CustomInput placeholder="Search by name..." margin="0" height="40px" flexItem :onChange="setSearch"
          :value="search">
          <span class="material-icons-outlined">
            search
          </span>
        </CustomInput>
      </div>
      <div class="">

      </div>
    </div>
    <div class="devices-container">
      <DeviceComponent v-for="device of display" :device="device" v-bind:key="device.device_id" @edit="editDevice"
        @locate="locateDevice" />
    </div>
  </div>
</template>

<script>
import { Holder } from '../helpers'
import CustomInput from './CustomInput.vue'
import DeviceComponent from './Device.vue'

export default {
  name: 'SideBar',
  components: {
    CustomInput,
    DeviceComponent
  },
  mounted() {
    Holder.onUpdate(this.updateDevices)
  },
  unmounted() {
    Holder.removeUpdate(this.updateDevices)
  },
  props: {
    setMapCenter: {
      type: Function,
      default: null
    },
    select: {
      type: Function,
      default: null
    },
  },
  data() {
    return {
      // all devices
      devices: [],
      search: '',
      // filtered devices to display
      display: [],
      firstLoad: true,
      editingDevice: false,
    }
  },
  methods: {
    editDevice(device) {
      this.editingDevice = device
    },
    stopEditing() {
      this.editingDevice = false
    },
    locateDevice(device) {
      if (this.select) this.select(device)
    },
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
  border-bottom: 1px solid rgba(0, 0, 0, .12);
}
</style>
