<template>
    <div class="icon-container">
        <div class="icon-button" @mousedown="onMouseDown" @mouseup="onMouseUp" v-on:mouseenter="onMouseEnter"
            v-on:mouseleave="onMouseLeave" @click="e => { e.preventDefault(); e.stopPropagation() }">
            <span class="material-icons-outlined">
                {{ icon }}
            </span>
        </div>
    </div>
</template>

<script>

export default {
    name: "IconButton",
    props: {
        icon: {
            type: String,
            required: true,
        },
        size: {
            type: Number,
            required: true,
        },
        onClick: {
            type: Function,
            required: true,
        },
        disabled: {
            type: Boolean,
            default: false,
        },
    },
    data() {
        return {
            colors: {
                hover: "rgb(25, 118, 210)",
                normal: "rgba(0, 0, 0, .66)",
                disabled: "rgba(0, 0, 0, .33)",
                hoverBG: "rgba(0, 0, 0, .1)",
            },
            hover: false,
            clicked: false,
        }
    },
    methods: {
        onMouseEnter() {
            this.hover = true
        },
        onMouseLeave() {
            this.hover = false
        },
        onMouseDown() {
            this.clicked = true
        },
        onMouseUp() {
            this.clicked = false
            if (!this.disabled) this.onClick()
        },
    }
}
</script>

<style>
.icon-container {
    width: 30px;
    height: 30px;
    display: flex;
    justify-content: center;
    align-items: center;
    cursor: v-bind("disabled ? 'default' : 'pointer'");
    border-radius: 100%;
    background-color: v-bind("(hover && !clicked && !disabled) ? colors.hoverBG : 'transparent'");
    transition: v-bind("clicked ? 'none' : 'all .2s ease-in'");
}

.icon-button {
    width: 100%;
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    border-radius: 100%;
    background-color: transparent;
    color: v-bind("disabled ? colors.disabled : (hover && !clicked) ? colors.hover : colors.normal");
    transition: v-bind("clicked ? 'none' : 'all .2s ease-in'");
    user-select: none;
}
</style>