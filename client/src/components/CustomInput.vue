<template>
    <div class="custom-input">
        <div class="label-container" v-if="label">
            <span>{{ required ? label + " *" : label }}</span>
        </div>
        <div class="input-relative">
            <div class="input-container" v-on:mouseenter="entered_" v-on:mouseleave="left_" v-on:click="() => {this.$refs.input.focus()}">
            <input v-bind:type="type" v-bind:value="onChange ? value : value_"
                v-on:input="(e) => { onChange_(e.currentTarget.value) }" v-bind:placeholder="placeholder"
                v-bind:disabled="disabled" v-bind:autocorrect="autocomplete" v-bind:autocapitalize="autocomplete"
                v-bind:spellcheck="autocomplete" v-bind:autocomplete="autocomplete" ref="input" v-on:blur="unfocus_"
                v-on:focus="focus_" />
            <div class="error-message" v-if="error_">
                <span>{{ error_ }}</span>
            </div>
        </div>
        </div>
    </div>
</template>

<script>
export default {
    name: 'CustomInput',
    // generic input props + onChange for controlled input
    props: {
        value: {
            type: String,
            default: '',
        },
        placeholder: {
            type: String,
            default: '',
        },
        type: {
            type: String,
            default: 'text',
        },
        disabled: {
            type: Boolean,
            default: false,
        },
        autocomplete: {
            type: String,
            default: 'off',
        },
        error: {
            type: String,
            default: '',
        },
        label: {
            type: String,
            default: '',
        },
        onChange: {
            type: Function,
            default: undefined
        },
        errorColor: {
            type: String,
            default: '#ff0000',
        },
        required: {
            type: Boolean,
            default: false,
        },
        margin: {
            type: String,
            default: '8px 0 16px 0',
        },
        width: {
            type: String,
            default: '100%',
        },
        height: {
            type: String,
            default: '40px',
        },
        labelColor: {
            type: String,
            default: '#000',
        },
        labelFontSize: {
            type: String,
            default: '14px',
        },
        labelFontWeight: {
            type: String,
            default: 'normal',
        },
        labelFocusColor: {
            type: String,
            default: 'rgb(25, 118, 210)',
        },
        labelMargin: {
            type: String,
            default: '0 0 8px 0',
        },
        inputColor: {
            type: String,
            default: '#000',
        },
        inputFontSize: {
            type: String,
            default: '14px',
        },
        inputFontWeight: {
            type: String,
            default: 'normal',
        },
        borderColor: {
            type: String,
            default: 'rgba(0, 0, 0, 0.42)',
        },
        borderHoverColor: {
            type: String,
            default: 'rgba(0, 0, 0, 0.78)',
        },
        borderRadius: {
            type: String,
            default: '4px',
        },
        borderFocusColor: {
            type: String,
            default: 'rgb(25, 118, 210)',
        },
        borderFocusWidth: {
            type: String,
            default: '2px',
        },
        borderHoverWidth: {
            type: String,
            default: '1px',
        },
        borderWidth: {
            type: String,
            default: '1px',
        },
    },
    methods: {
        // controlled input
        onChange_(data) {
            this.value_ = data
            if (this.required && !data) {
                this.error_ = 'Required'
            } else {
                this.error_ = ''
            }
            if (this.onChange) {
                this.onChange(data)
            }
            this.$forceUpdate()
        },
        // focus triggered on blur and click
        focus_() {
            this.focused = true
        },
        unfocus_() {
            this.focused = false
        },
        // trigger rerender when mouse enters or leaves visible div
        entered_() {
            this.hovered = true
        },
        left_() {
            this.hovered = false
        }
    },
    data() {
        return {
            focused: false,
            hovered: false,
            value_: this.onChange ? this.value : '',
            error_: this.error,
        }
    },
}
</script>

<style scoped>
.custom-input {
    width: v-bind("flexItem ? '100%' : width");
    flex: v-bind("flexItem ? '1' : 'unset'");
    height: auto;
    display: flex;
    flex-direction: column;
    margin: v-bind(margin);
}

.label-container {
    margin: 0;
    padding: 0;
    font-size: 14px;
    font-weight: normal;
    color: v-bind("error_ ? errorColor : (focused ? labelFocusColor : labelColor)");
    margin: v-bind(labelMargin);
    font-size: v-bind(labelFontSize);
    font-weight: v-bind(labelFontWeight);
}
.input-relative {
    height: calc(v-bind(height));
    position: relative;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    box-sizing: border-box;
}
.input-container {
    display: flex;
    height: 100%;
    width: 100%;
    border: v-bind("(focused ? borderFocusWidth : hovered ? borderHoverWidth : borderWidth) + ' solid ' + (error_ ? errorColor : focused ? borderFocusColor : hovered ? borderHoverColor : borderColor)");
    border-radius: v-bind(borderRadius);
    background-color: v-bind("focused ? '#fff' : '#fff'");
    padding: 0 .65em;
    box-sizing: border-box;
    margin: 0;
    cursor: text;
}

.input-container input {
    position: absolute;
    outline: none;
    border: none;
    flex: 1;
    height: 100%;
    font-size: v-bind(inputFontSize);
    font-weight: v-bind(inputFontWeight);
    color: v-bind(inputColor);
    background-color: transparent;
    left: .65em;
    top: 50%;
    transform: translateY(-50%);
    margin: 0;
}

.error-message {
    position: absolute;
    bottom: 0;
    left: 0;
    font-size: 14px;
    font-weight: normal;
    color: v-bind(errorColor);
    margin: 0;
    padding: 0;
    transform: translate(0, 120%);
}
</style>
