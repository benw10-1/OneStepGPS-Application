<template>
    <div class="device-container">
        <div :class="'device-sidebar ' + (device.online ? device.drive_status : 'offline')"></div>
        <div class="device-inner">
            <div class="device-main-info">
                <div class="device-display-name">
                    {{ device.display_name }}
                </div>
                <div class="device-status">
                    {{ (device.online ? driveStates[device.drive_status] : "No signal") + ' ' +
                            formatTime(device.drive_status_begin_time)
                    }}
                </div>
            </div>
            <div class="device-right">
                <div class="device-speed" v-if="device.online">
                    <span>{{ device.speed + " mph" }}</span>
                </div>
                <div class="buttons">
                    <span>a</span>
                    <div>b</div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
// import { CustomInput } from './CustomInput.vue'
import { dateFormatter } from '../helpers'

export default {
    name: "DeviceComponent",
    components: {
        // CustomInput,
    },
    props: {
        device: {
            type: Object,
            required: true,
        }
    },
    data() {
        return {
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
        }
    },
    methods: {
        formatTime(time) {
            const date = new Date() - new Date(time)

            return dateFormatter(date.valueOf() / 1000)
        },
    }
}
</script>

<style>
.device-container {
    flex: 1;
    height: auto;
    display: flex;
    position: relative;
    align-items: center;
    flex-direction: row;
    box-sizing: border-box;
    padding: 10px 14px;
    border-bottom: 1px solid rgba(0, 0, 0, .12);
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
}
.device-left {
    display: flex;
    align-items: center;
    justify-content: center;
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
</style>