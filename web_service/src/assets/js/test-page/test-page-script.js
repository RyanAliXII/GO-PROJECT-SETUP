import {createApp} from 'vue/dist/vue.esm-bundler.js'
import { plugin, defaultConfig } from '@formkit/vue'
import '@formkit/themes/genesis'

createApp({
  
    mounted(){
        console.log("MOUNTED")
    },
    compilerOptions:{
        delimiters:["{%", "%}"],
    },
    methods:{
        submitForm(formData){
            console.log(formData)
        }   
    }
}).use(plugin, defaultConfig).mount("#root")
