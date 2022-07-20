<template>
    <div class="svg-image-container" @click="onClick">
        <img :src="source" class="svg-image" />
    </div>
</template>

<script>
import fileMap from '../assets/imageMap.json'

export default {
    name: "SVGIcon",
    props: {
        icon: {
            type: String,
            default: 'default'
        }
    },
    data() {
        return {
            source: "/images/" + fileMap[this.icon]
        }
    },
    methods: {
        onClick() {
            this.$emit('click')
        }
    },
    watch: {
        icon(newIcon) {
            this.source = "/images/" + fileMap[newIcon]
        }
    }
}
</script>

<style>
.svg-image-container {
    display: grid;
    place-items: center;
    height: 30px;
    width: 30px;
}
.svg-image {
    width: 25px;
    height: 25px;
    /* object-fit: contain; */
}
</style>