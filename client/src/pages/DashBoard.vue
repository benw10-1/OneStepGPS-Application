<template>
    <div class="dashboard-page">
        <Transition>
            <div v-if="!actualLoaded" class="loading-panel">
                <div class="banner"></div>
                <div class="loader" style="margin-top: 30px"></div>
            </div>
        </Transition>
        <SideBar @loaded="() => { setSideLoad(true) }" :set-map-center="mapBridge.moveCenter"
            :select="mapBridge.select" />
        <MapComponent @map-loaded="(obj) => { setMapLoad(true); setMapBridge(obj) }" />
    </div>
</template>

<script setup>
    document.title = "Dashboard"
</script>

<script>
import SideBar from '../components/SideBar.vue'
import MapComponent from '../components/MapComponent.vue'
import "../assets/one-step.png"
import { settingsStore, deviceStore } from '@/helpers';
import { mapActions } from 'pinia';

export default {
    name: 'DashBoard',
    components: {
        SideBar,
        MapComponent
    },
    data() {
        return {
            sideLoaded: false,
            mapLoaded: false,
            actualLoaded: false,
            prefLoad: false,
            mapBridge: {},
        }
    },
    methods: {
        ...mapActions(deviceStore, ['startInterval', 'stopInterval', 'refreshDevices']),
        ...mapActions(settingsStore, ['refreshSettings']),
        // loading triggers + handlers
        setSideLoad(loaded) {
            this.sideLoaded = loaded
            if (this.sideLoaded && this.mapLoaded) {
                setTimeout(() => {
                    this.actualLoaded = true
                }, 500)
            }
        },
        setMapLoad(loaded) {
            this.mapLoaded = loaded
            if (this.sideLoaded && this.mapLoaded) {
                setTimeout(() => {
                    this.actualLoaded = true
                }, 500)
            }
        },
        setMapBridge(bridge) {
            this.mapBridge = bridge
        },
    },
    mounted() {
        this.refreshSettings().then(() => {
            this.prefLoad = true
        })
        this.refreshDevices()
        this.startInterval()
    },
    unmounted() {
        this.stopInterval()
    },
}
</script>

<style>
.dashboard-page {
    width: 100%;
    height: 100%;
    display: flex;
}

.loading-panel {
    position: absolute;
    top: 0;
    left: 0;
    height: 100%;
    width: 100%;
    z-index: 1111;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    background-color: #f5f5f5;
}

.banner {
    width: 100%;
    height: 50%;
    background: url("../assets/one-step.png") no-repeat center;
}

.loader {
    border: 10px solid #f5f5f5;
    border-top: 10px solid rgb(25, 118, 210);
    border-radius: 50%;
    width: 50px;
    height: 50px;
    animation: spin 2s linear infinite;
}

@keyframes spin {
    0% {
        transform: rotate(0deg);
    }

    100% {
        transform: rotate(360deg);
    }
}
</style>
