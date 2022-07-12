<template>
    <div class="switch-setting">
        <div class="option-label">
            {{ text }}
        </div>
        <div class="options-container" @mouseenter="enter_" @mouseleave="exit_">
            <div v-for="(item, index) of switchArray" class="switch-option" :style="`left: ${100 * (index - switchIndex)}%`" :key="index">
                <div class="switch-content">
                    {{ item }}
                </div>
            </div>
        </div>
    </div>
</template>

<script>

export default {
    name: "SwitchSetting",
    props: {
        switchArray: {
            type: Array,
            required: true,
        },
        onChange: {
            type: Function,
            required: true,
        },
        text: {
            type: String,
            required: true,
        },
        initial: {
            type: Number,
            required: true,
        },
    },
    data() {
        return {
            switchIndex: this.initial,
            colors: {
                selected: "rgba(0, 0, 255, .67)",
                unselected: "rgba(255, 255, 255, .67)",
            },
            hovered: false,
        }
    },
    methods: {
        next() {
            this.index = (this.index + 1) % this.switchArray.length
            this.onChange(this.switchArray[this.index])
        },
        enter_() {
            this.hovered = true
        },
        exit_() {
            this.hovered = false
        },
    },
}
</script>

<style>
.switch-setting {
    flex: 1;
    height: 50px;
    display: flex;
    align-items: center;
}
.option-label {
    flex: 1;
    height: 100%;
    display: grid;
    place-items: center;
    font-size: 20px;
    font-weight: bold;
    color: v-bind("hovered ? colors.selected : colors.unselected");
    transition: all .2s ease-in;
}
.options-container {
    flex: 1;
    height: 100%;
    display: flex;
    overflow: hidden;
}
.switch-option {
    position: relative;
    flex: 1;
    height: 100%;
    background-color: v-bind("hovered ? 'rgba(155, 155, 155, .67)' : 'transparent'");
    transition: all .2s ease-in;
}
.switch-content {
    width: 100%;
    height: 100%;
    color: v-bind("hovered ? colors.selected : colors.unselected");
}
</style>
