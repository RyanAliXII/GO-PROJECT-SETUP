import {createApp} from 'vue/dist/vue.esm-bundler'
import LoginLayout from '@components/LoginLayout'
createApp({
    components:{
        "login-layout": LoginLayout
    },
    data(){
        return{
            form:{
                email:"ryanali456@gmail.com",
                password: "testpass"
            }
        }
    },
    methods:{
        submit(){
                console.log(this.form)
        }
    }
}).mount("#root")