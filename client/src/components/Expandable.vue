<template>
    <div class="expandable-container">
        <div class="expandable-header" @mouseenter="onMouseEnter" @mouseleave="onMouseLeave" @mousedown="onMouseDown" @mouseup="onMouseUp" >
            <slot name="header"></slot>
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
            colors: {
                hovered: "rgba(155, 155, 155, 0.1)",
                clicked: "rgba(155, 155, 155, 0.2)",
            },
            maxHeight: 0,
        }
    },
    mounted() {
        const margin = parseInt(window.getComputedStyle(this.$refs.content).marginTop) + parseInt(window.getComputedStyle(this.$refs.content).marginBottom);
        this.maxHeight = this.$refs.content.offsetHeight + margin + "px";
        console.log(this.maxHeight);
    },
    methods: {
        onMouseEnter() {
            this.hovered = true;
        },
        onClick(e) {
            if (e.target === this.$refs.content) {
                this.onChange(!this.expanded);
            }
        },
        onMouseLeave() {
            this.hovered = false;
        },
    },
}
</script>

<style>
.expandable-container {
    width: 100%;
    box-sizing: border-box;
}
.expandable-header {
    width: 100%;
    box-sizing: border-box;
    cursor: pointer;
    user-select: none;
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    background-color: v-bind("hovered ? colors.hovered : 'transparent'");
}
.expandable-content-anchor {
    width: 100%;
    box-sizing: border-box;
    position: relative;
    overflow: hidden;
    transition: v-bind("snap ? '' : 'max-height 0.3s ease-in'");
    max-height: v-bind("expanded ? maxHeight : 0");
}
.expandable-content {
    width: 100%;
    box-sizing: border-box;
    position: absolute;
    bottom: 0;
}
</style>