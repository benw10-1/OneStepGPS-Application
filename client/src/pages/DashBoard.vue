<template>
    <div class="dashboard-page">
        <SideBar />
        <MapComponent />
    </div>
</template>

<script setup>
    document.title = "Dashboard"
</script>

<script>
import SideBar from '../components/SideBar.vue'
import MapComponent from '../components/MapComponent.vue'

export default {
    name: 'DashBoard',
    components: {
        SideBar,
        MapComponent
    },
}
</script>

<style>
.dashboard-page {
    width: 100%;
    height: 100%;
    display: flex;
}
</style>
