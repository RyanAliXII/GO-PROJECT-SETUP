import {createApp} from 'vue/dist/vue.esm-bundler'
import Layout from '@components/Layout'
//call this to initialize 
createApp({
        components:{
            "layout": Layout
        },
}).mount("#root")

