import {createApp} from 'vue/dist/vue.esm-bundler.js'


 createApp({
    compilerOptions:{
        delimiters:["{%", "%}"]
    },
    data(){
        return {
            count:0,
        }
    },
    mounted(){
        console.log("MOUNTED")
    },
    methods:{
        incrementCount(){
            this.count += 1;
        }
    }   


}).mount("#root")