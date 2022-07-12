<template>
    <div class="dashboard-page">
        <div v-if="!actualLoaded" class="loading-panel">
            <div class="banner"></div>
            <div class="loader" style="margin-top: 30px"></div>
        </div>
        <SideBar @loaded="() => {setSideLoad(true)}" />
        <MapComponent @map-loaded="() => {setMapLoad(true)}" />
    </div>
</template>

<script setup>
    document.title = "Dashboard"
</script>

<script>
import SideBar from '../components/SideBar.vue'
import MapComponent from '../components/MapComponent.vue'
import "../assets/one-step.png"

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
        }
    },
    methods: {
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
    }
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
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}
</style>
