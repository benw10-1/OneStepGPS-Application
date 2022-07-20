<template>
    <ExpandableVue :onChange="onExpand" :expanded="expanded">
        <template v-slot:header>
            <div class="device-container">
                <div style="margin-right: 10px">
                    <SVGIcon :icon="settings.icon ?? 'car'" />
                </div>
                <div class="device-inner">
                    <div class="device-main-info">
                        <div class="device-display-name">
                            {{ settings.displayName }}
                        </div>
                        <div class="device-status">
                            <div :class="'status-text ' + (device.online ? device.drive_status : 'offline')">{{ device.online ? driveStates[device.drive_status] : "No signal" }}</div>
                            <span>{{ time }}</span>
                        </div>
                    </div>
                    <div class="device-right">
                        <div class="device-speed" v-if="device.online">
                            <span>{{ device.speed + " mph" }}</span>
                        </div>
                        <div class="buttons">
                            <IconButton :icon="'settings'" :onClick="() => {
                                $emit('edit', device)
                            }" />
                            <IconButton :icon="settings.visible ? 'visibility' : 'visibility_off'" :onClick="() => {
                                updateSettings({
                                    visible: !settings.visible
                                })
                            }" />
                            <IconButton :icon="'my_location'" :disabled="!settings.visible" :onClick="() => {
                                $emit('locate', device)
                            }" />
                        </div>
                    </div>
                </div>
            </div>
        </template>
        <template v-slot:default>
            <DeviceInfo :device="device" />
        </template>
    </ExpandableVue>
</template>

<script>
import IconButton from './IconButton.vue'
import { dateFormatter, PreferenceHolder } from '../helpers'
import ExpandableVue from './Expandable.vue'
import DeviceInfo from './DeviceInfo.vue'
import SVGIcon from './SVGIcon.vue'

export default {
    name: "DeviceComponent",
    components: {
        IconButton,
        ExpandableVue,
        DeviceInfo,
        SVGIcon,
    },
    props: {
        device: {
            type: Object,
            required: true,
        },
    },
    mounted() {
        PreferenceHolder.onUpdate(this.updateSettings)
    },
    unmounted() {
        PreferenceHolder.removeUpdate(this.updateSettings)
    },
    data() {
        return {
            colors: {
                off: "rgba(255, 0, 0, .67)",
                driving: "rgba(0, 255, 0, .67)",
                idle: "rgba(135, 206, 250, .67)",
                offline: "rgba(218, 223, 225, .8)",
            },
            time: this.formatTime(this.device.drive_status_begin_time),
            driveStates: {
                off: "Stopped",
                driving: "Driving",
                idle: "Idle",
            },
            // preference map
            settings: PreferenceHolder.get()?.deviceSettings?.[this.device.device_id] ?? {
                visible: true,
                displayName: this.device.display_name,
                icon: "car",
            },
            expanded: false
        }
    },
    methods: {
        formatTime(time) {
            const date = new Date() - new Date(time)

            return dateFormatter(date.valueOf() / 1000)
        },
        onExpand(expanded) {
            this.expanded = expanded
        },
        updateSettings(settings) {
            let safe = false
            if (settings.isPrefs__) {
                settings = settings?.deviceSettings?.[this.device.device_id] ?? {}
            }
            else safe = true
            this.settings = {
                ...this.settings,
                ...settings,
            }

            // prevent infinite loop
            if (safe) PreferenceHolder.set({
                deviceSettings: {
                    ...PreferenceHolder.get()?.deviceSettings ?? {},
                    [this.device.device_id]: this.settings,
                },
            }, this.updateSettings)
        },
    },
    watch: {
        device: {
            handler(newDevice) {
                // update time on device update, not on state change
                this.time = this.formatTime(newDevice.drive_status_begin_time)
            },
            immediate: true,
        },
    },
}
</script>

<style>
.device-container {
    flex: 1;
    height: 70px;
    display: flex;
    position: relative;
    align-items: center;
    flex-direction: row;
    box-sizing: border-box;
    padding: 0 7px;
    background-color: v-bind("settings.visible ? 'transparent' : 'rgba(155, 155, 155, .1)'");
    transition: all .2s ease-in;
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
    color: white;
}

.driving {
    background: v-bind("colors.driving");
    color: black;
}

.idle {
    background: v-bind("colors.idle");
    color: black;
}

.offline {
    background: v-bind("colors.offline");
    color: black;
}

.device-inner {
    height: 50px;
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: space-between;
}

.device-display-name {
    font-size: 14px;
    font-weight: 500;
    color: rgba(0, 0, 0, .87);
}

.device-status {
    font-size: 12px;
    font-weight: 400;
    color: rgba(0, 0, 0, .54);
}

.device-speed {
    font-size: 14px;
    font-weight: 400;
    color: rgba(0, 0, 0, .84);
    margin-bottom: 5px;
    text-align: center;
}

.device-left {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 100%;
    flex-direction: column;
}

.device-right {
    display: flex;
    align-items: center;
    justify-content: space-around;
    height: 100%;
    flex-direction: column;
}

.buttons {
    width: 100%;
    height: auto;
    display: flex;
    align-items: center;
    justify-content: flex-end;
}

.device-content {
    width: 100%;
    padding: .5rem .75rem;
    box-sizing: border-box;
    display: flex;
    flex-direction: column;
    justify-content: space-around;
}
.device-info {
    width: 100%
}
.device-info-content-text p {
    margin-top: .4rem;
}
.info-block {
    display: flex;
    flex-direction: column;
    align-items: center;
    width: fit-content;
    margin-right: 22px;
}
.info-block span {
    font-size: 14px;
    font-weight: 400;
    color: rgba(0, 0, 0, .54);
    margin-top: 3px;
}
.device-info-row {
    display: flex;
    flex-direction: row;
    align-items: center;
    width: 100%;
}
.device-status {
    display: flex;
    flex-direction: row;
    align-items: center;
}
.status-text {
    display: grid;
    place-items: center;
    font-size: 14px;
    padding: 4px;
    font-weight: 600;
    margin-right: 5px;
    border-radius: 4px;
}
</style>