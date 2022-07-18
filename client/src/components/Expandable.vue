<template>
    <div class="expandable-container">
        <div class="expandable-header" @mouseenter="onMouseEnter" @mouseleave="onMouseLeave" @click="onClick" >
            <slot name="header"></slot>
            <div v-if="expanded" class="bottom-shadow"></div>
        </div>
        <div class="expandable-content-anchor">
            <div class="expandable-content" ref="content">
                <slot></slot>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: "ExpandableComponent",
    props: {
        expanded: {
            type: Boolean,
            default: false,
        },
        onChange: {
            type: Function,
            required: true,
        },
        snap: {
            type: Boolean,
            default: false,
        },
    },
    data() {
        return {
            hovered: false,
            clicked: false,
            snap_: true,
            colors: {
                hovered: "rgba(155, 155, 155, 0.1)",
                clicked: "rgba(155, 155, 155, 0.2)",
            },
            maxHeight: 0,
        }
    },
    mounted() {
        setTimeout(this.updateHeight, 300)
        this.$refs.content.addEventListener("resize", this.updateHeight)
    },
    methods: {
        onMouseEnter() {
            this.hovered = true;
        },
        onClick() {
            this.onChange(!this.expanded)
        },
        onMouseLeave() {
            this.hovered = false;
        },
        updateHeight() {
            const margin = parseInt(window.getComputedStyle(this.$refs.content).marginTop) + parseInt(window.getComputedStyle(this.$refs.content).marginBottom);
            this.maxHeight = this.$refs.content.offsetHeight + margin;
            this.snap_ = this.snap;
        },
    },
}
</script>

<style>
.expandable-container {
    width: 100%;
    box-sizing: border-box;
    border-bottom: 1px solid #e0e0e0;
}
.expandable-header {
    position: relative;
    width: 100%;
    box-sizing: border-box;
    cursor: pointer;
    user-select: none;
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    background-color: v-bind("hovered ? colors.hovered : 'transparent'");
    transition: v-bind("clicked ? 'none' : 'all .2s ease-in-out'");
    z-index: 1;
    box-shadow: v-bind("expanded ? '0 5px 2px -2px rgba(0, 0, 0, .44)' : 'none'");
}
.expandable-content-anchor {
    width: 100%;
    box-sizing: border-box;
    position: relative;
    overflow: hidden;
    transition: v-bind("snap_ ? '' : 'height 0.3s ease-in-out'");
    height: v-bind("(expanded ? maxHeight : 0) + 'px'");
}
.expandable-content {
    width: 100%;
    box-sizing: border-box;
    position: absolute;
    bottom: 0;
}
.expandable-header::after {
    position: absolute;
    width: 100%;
    bottom: 1px;
    z-index: 0;
    left: 0;
    background-color: #e0e0e0;
    box-shadow: 0px 0px 8px 2px #000000;
}
</style>