<template>
    <div class="address-container">
        <div v-if="loading" class="loader"></div>
        <div v-else class="address-text">
            <p>{{ address }}</p>
            <p>{{ latlng[0].toFixed(6) }}, {{ latlng[1].toFixed(6) }}</p>
        </div>
    </div>
</template>

<script>
import { locationCache } from '@/helpers'

export default {
    name: "AddressComponent",
    props: {
        latlng: {
            type: Array,
            required: true,
        },
    },
    data() {
        return {
            loading: true,
            address: '',
        }
    },
    watch: {
        latlng: {
            immediate: true,
            handler(latlng) {
                this.getAddress(latlng)
            },
        },
    },
    methods: {
        getAddress(latlng=this.latlng) {
            this.loading = true
            locationCache.requestLocation(latlng).then(address => {
                this.address = address
                this.loading = false
            })
        },
    },
    mounted() {
        this.getAddress()
    },
}
</script>

<style>
.address-container {
    display: flex;
    justify-content: v-bind("loading ? 'center' : 'unset'");
    align-items: v-bind("loading ? 'center' : 'unset'");
    width: 100%;
    height: 100%;
}
.address-text {
    font-size: inherit;
    font-weight: inherit;
    color: inherit;
}
</style>