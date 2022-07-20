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
      <SortBar />
    </div>
    <div class="devices-container">
      <DeviceComponent v-for="device of display" :device="device" v-bind:key="device.device_id" @edit="editDevice"
        @locate="locateDevice" @drag="() => {console.log('dragged')}" />
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
          <div class="icon-selector">
            <div class="icon-selector-header">Icon</div>
            <div class="icon-selector-body">
              <div class="icon-selector-icon" v-for="key in Object.keys(svgMap)" :key="key" @click="() => {updateIcon(key)}" :class="{selected: icon === key}">
                <SVGIcon :icon="key" />
              </div>
            </div>
          </div>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<script>
import { Holder, PreferenceHolder } from '../helpers'
import SVGIcon from './SVGIcon.vue'
import CustomInput from './CustomInput.vue'
import DeviceComponent from './Device.vue'
import FilterBar from './FilterBar.vue'
import SortBar from './SortBar.vue'
import fileMap from '../assets/imageMap.json'

export default {
  name: 'SideBar',
  components: {
    CustomInput,
    DeviceComponent,
    FilterBar,
    SortBar,
    SVGIcon,
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
    // map controller
    setMapCenter: {
      type: Function,
      default: null
    },
    // selected device controller
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
      // modal state
      displayName: '',
      icon: 'car',
      svgMap: fileMap,
    }
  },
  methods: {
    editDevice(device) {
      this.editingDevice = device
      // get cached display name and icon
      const pref = PreferenceHolder.get().deviceSettings?.[device.device_id] ?? {}
      this.displayName = pref.displayName || device.display_name
      this.icon = pref.icon || 'car'
      window.addEventListener("click", this.stopEditing)
    },
    setDeviceName(name) {
      this.displayName = name
    },
    updateDisplayname() {
      if (!this.editingDevice || !this.displayName) return false
      const prev = PreferenceHolder.get().deviceSettings ?? {}
      // overwrite display name
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
    updateIcon(icon) {
      if (!this.editingDevice || !this.icon) return false
      const prev = PreferenceHolder.get().deviceSettings ?? {}
      this.icon = icon
      // overwrite icon
      PreferenceHolder.set({
        deviceSettings: {
          ...prev,
          [this.editingDevice.device_id]: {
            ...prev?.[this.editingDevice.device_id] ?? {},
            icon: this.icon
          }
        }
      })
    },
    stopEditing(e) {
      // ignores close button
      if (!e.target.classList?.contains("close")) {
        for (const x of e.composedPath()) {
          if (x.classList?.contains("modal-content")) {
            return
          }
        }
      }
      // update state on modal close
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
      // prefrences get
      const prefs = PreferenceHolder.get()?.filter
      const sort = PreferenceHolder.get()?.sort === "none" ? null : PreferenceHolder.get()?.sort
      // filter by filter tags
      const devSettings = PreferenceHolder.get().deviceSettings ?? {}
      const filtered = this.devices.filter(device => {
        const disp = devSettings[device.device_id]?.displayName ?? device.display_name
        return disp.toLowerCase().includes(this.search.toLowerCase().trim()) && (!prefs || prefs.disabled || prefs.tags.includes(device.online ? device.drive_status : 'offline'))
      })

      if (sort) {
        if (sort === "asc") {
          filtered.sort((a, b) => {
            const namea = devSettings[a.device_id]?.displayName ?? a.display_name.toLowerCase()
            const nameb = devSettings[b.device_id]?.displayName ?? b.display_name.toLowerCase()
            return namea.localeCompare(nameb)
          })
        }
        else {
          filtered.sort((a, b) => {
            const namea = devSettings[a.device_id]?.displayName ?? a.display_name.toLowerCase()
            const nameb = devSettings[b.device_id]?.displayName ?? b.display_name.toLowerCase()
            return nameb.localeCompare(namea)
          })
        }
      }

      return filtered
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

.icon-selector {
  width: 100%;
  height: auto;
  display: flex;
  flex-direction: column;
  margin-top: 10px;
}
.icon-selector-body {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  border-radius: 4px;
  border: 1px solid rgba(0, 0, 0, .4);
  flex-wrap: wrap;
  width: 100%;
}
.icon-selector-header {
  font-size: 14px;
  font-weight: normal;
  margin-bottom: 18px;
}

.icon-selector-icon {
  width: 50px;
  height: 50px;
  margin: 5px;
  border-radius: 5px;
  display: flex;
  justify-content: center;
  align-items: center;
  cursor: pointer;
  transition: all .2s ease-in;
  user-select: none;
}
.icon-selector-icon:hover, .icon-selector-icon.selected {
  background: rgba(111, 111, 111, .2);
}
</style>
