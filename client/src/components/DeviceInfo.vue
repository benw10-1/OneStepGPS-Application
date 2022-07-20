<template>
    <div class="device-content">
        <div class="device-info">
            <div class="device-info-header">
                <span class="device-info-header-text">Location</span>
            </div>
            <div class="device-info-content">
                <div class="device-info-content-text">
                    <AddressVue :latlng="[device.lat, device.lng]" />
                </div>
            </div>
        </div>
        <div class="device-info" style="margin: 14px 0 0">
            <div class="device-info-content">
                <div class="device-info-row">
                    <div class="info-block">
                        <div class="device-info-header-text">Altitude</div>
                        <span>{{ device.altitude + " ft" }}</span>
                    </div>
                    <div class="info-block" v-if="device.odometer">
                        <div class="device-info-header-text">Odometer</div>
                        <span>{{ device.odometer + " mi" }}</span>
                    </div>
                    <div class="info-block">
                        <div class="device-info-content-text">Voltage</div>
                        <span>{{ device.voltage + " V" }}</span>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import AddressVue from './Address.vue'

export default {
    name: "DeviceInfo",
    props: {
        device: {
            type: Object,
            required: true
        }
    },
    components: {
        AddressVue
    }
}
</script>