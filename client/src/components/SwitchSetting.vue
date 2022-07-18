<template>
    <div class="switch-setting">
        <div class="option-label">
            {{ text }}
        </div>
        <button class="options-container" @mouseenter="enter_" @mouseleave="exit_" @mousedown="onMouseDown" @mouseup="onMouseUp">
            <div v-for="(item, index) of switchArray" class="switch-option" :style='`right: ${100 * (index - switchIndex)}%`' :key="index" @click="next">
                <span :class="isIcon ? 'material-icons-outlined' : 'underlined'">{{ item }}</span>
            </div>
        </button>
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
            default: 0,
        },
        isIcon: {
            type: Boolean,
            default: false,
        },
    },
    data() {
        return {
            switchIndex: this.initial,
            colors: {
                selected: "rgba(0, 0, 255, .67)",
                unselected: "rgba(0, 0, 0, .67)",
            },
            hovered: false,
            clicked: false,
        }
    },
    methods: {
        next() {
            this.switchIndex = (this.switchIndex + 1) % this.switchArray.length
            this.onChange(this.switchArray[this.switchIndex])
        },
        enter_() {
            this.hovered = true
        },
        exit_() {
            this.hovered = false
        },
        onMouseDown() {
            this.clicked = true
        },
        onMouseUp() {
            this.clicked = false
        },
    },
}
</script>

<style>
.underlined {
    text-decoration: underline;
}
.switch-setting {
    flex: 1;
    height: 50px;
    display: flex;
    align-items: center;
}
.option-label {
    flex: 1;
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 17px;
    font-weight: bold;
    color: v-bind("hovered ? colors.selected : colors.unselected");
    transition: all .2s ease-in-out;
}
.options-container {
    position: relative;
    flex: 1;
    height: 100%;
    overflow: hidden;
    background-color: v-bind("(hovered && !clicked) ? 'rgba(155, 155, 155, .47)' : 'transparent'");
    border-radius: 5px;
    transition: v-bind("clicked ? 'none' : 'all .2s ease-in-out'");
    cursor: pointer;
    user-select: none;
    outline: none;
    border: none;
}
/* .options-container:focus {
    outline: none;
    background-color: #e0e0e0;
} */
.switch-option {
    position: absolute;
    width: 100%;
    height: 100%;
    top: 0;
    display: grid;
    place-items: center;
    color: v-bind("hovered ? colors.selected : colors.unselected");
    transition: all .2s ease-in;
}
</style>
