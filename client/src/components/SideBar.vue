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
      <FilterBar />
    </div>
    <div class="devices-container">
      <DeviceComponent v-for="device of display" :device="device" v-bind:key="device.device_id" @edit="editDevice"
        @locate="locateDevice" />
    </div>
  </div>
  <Teleport to="body">
    <Transition>
      <div class="modal-container" v-if="editingDevice">
        <div class="modal-content">
          <div class="close material-icons-outlined">close</div>
          <div class="modal-header">Edit Device</div>
          <CustomInput :value="displayName" :label="'Display Name'" :onChange="setDeviceName" :required="true"
            :width="'100%'" :height="'40px'" @blur="updateDisplayname" />
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<script>
import { Holder, PreferenceHolder } from '../helpers'
import CustomInput from './CustomInput.vue'
import DeviceComponent from './Device.vue'
import FilterBar from './FilterBar.vue'

export default {
  name: 'SideBar',
  components: {
    CustomInput,
    DeviceComponent,
    FilterBar
  },
  mounted() {
    Holder.onUpdate(this.updateDevices)
    PreferenceHolder.onUpdate(this.refreshDevices)
  },
  unmounted() {
    Holder.removeUpdate(this.updateDevices)
    PreferenceHolder.removeUpdate(this.refreshDevices)
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
      displayName: '',
    }
  },
  methods: {
    editDevice(device) {
      this.editingDevice = device
      this.displayName = PreferenceHolder.get().deviceSettings?.[device.device_id]?.displayName || device.display_name
      window.addEventListener("click", this.stopEditing)
    },
    setDeviceName(name) {
      this.displayName = name
    },
    updateDisplayname() {
      if (!this.editingDevice || !this.displayName) return false
      const prev = PreferenceHolder.get().deviceSettings ?? {}
      PreferenceHolder.set({
        deviceSettings: {
          ...prev,
          [this.editingDevice.device_id]: {
            ...prev?.[this.editingDevice.device_id] ?? {},
            displayName: this.displayName.trim()
          }
        }
      })
    },
    stopEditing(e) {
      if (!e.target.classList?.contains("close")) {
        for (const x of e.composedPath()) {
          if (x.classList?.contains("modal-content")) {
            return
          }
        }
      }
      this.updateDisplayname()
      this.editingDevice = false
      window.removeEventListener("click", this.stopEditing)
    },
    locateDevice(device) {
      if (this.select) this.select(device)
    },
    refreshDevices() {
      this.updateDevices(this.devices)
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
      const prefs = PreferenceHolder.get()?.filter
      return this.devices.filter(device => {
        return device.display_name.toLowerCase().includes(this.search.toLowerCase().trim()) && (!prefs || prefs.disabled || prefs.tags.includes(device.online ? device.drive_status : 'offline'))
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

.modal-container {
  width: 100%;
  height: 100%;
  position: absolute;
  top: 0;
  left: 0;
  z-index: 10001;
  background: rgba(0, 0, 0, .5);
  display: flex;
  justify-content: center;
  align-items: center;
}

.modal-content {
  background: white;
  width: 400px;
  height: auto;
  border-radius: 5px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  box-shadow: 0 3px 1px -2px rgb(0 0 0 / 20%), 0 2px 2px 0 rgb(0 0 0 / 14%), 0 1px 5px 0 rgb(0 0 0 / 12%);
  align-items: center;
  padding: 20px;
  position: relative;
}

.close {
  position: absolute;
  top: 10px;
  right: 10px;
  cursor: pointer;
  color: rgba(0, 0, 0, .5);
  transition: all .2s ease-in;
  user-select: none;
}

.close:hover {
  color: rgb(25, 118, 210);
}
.modal-header {
  font-size: 20px;
  font-weight: bold;
  margin-bottom: 20px;
}
</style>
