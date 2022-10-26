import {createApp, onMounted, ref} from 'vue/dist/vue.esm-bundler'
import LoginLayout from '@components/LoginLayout'
import axios from 'axios'
import {LoginSchema} from './yup-schema'
import UseForm from '@hooks/useForm'

createApp({
    components:{
        "login-layout": LoginLayout
    },

    setup(){
        const {form, changeFormValue, errors} = UseForm({initialData:{email:" ", password:""}, yupSchema:LoginSchema})
        const submit = async()=>{
                console.log(errors.value)
         }
         const handleFormInputs = (event)=>{
            const input = event.target
            changeFormValue(input.name, input.value)
         }
         onMounted(()=>{
            console.log("MOUNTED")
        console.log(form.value)
         })
        return {
            submit,
            handleFormInputs,
            errors
        }
    }
}).mount("#root")