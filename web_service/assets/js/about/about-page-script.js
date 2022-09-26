import {createApp} from 'vue/dist/vue.esm-bundler.js'
import ButtonWithCounter from '@components/ButtonWithCounter'


createApp({
    components:{
        "button-with-counter": ButtonWithCounter
    }
}).mount("#root")