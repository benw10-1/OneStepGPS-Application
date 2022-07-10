<template>
    <div class="dashboard-page">
        <SideBar />
    </div>
</template>

<script setup>
    document.title = "Dashboard"
</script>

<script>
import SideBar from '../components/SideBar.vue'

export default {
    name: 'DashBoard',
    components: {
        SideBar
    },
    data() {

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
