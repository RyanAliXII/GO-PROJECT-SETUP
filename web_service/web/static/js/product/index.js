import { createApp } from 'vue/dist/vue.esm-bundler';

createApp({
    compilerOptions:{
        delimiters: ["${", "}$"]
    },
    mounted(){
        console.log("mounted");
    },
    data(){
        return{
            message:"HELLO WORLD",
            count:0,
            value:2,
        }
    },
    methods:{
        increment(){
            this.count = this.count + 1
        },
        decrement(){
            this.count = this.count - 1
        }
    }
}).mount("#createProductsPage")