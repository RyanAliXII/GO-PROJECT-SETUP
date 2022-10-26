import { ref } from "vue/dist/vue.esm-bundler"
import {ValidationError} from "yup"
const useForm = ({ initialData, yupSchema, onlyValidateActiveInput}) => {

    const form = ref(initialData)
    const errors = ref({})
    const currentActiveInput = ref("")
    const changeFormValue = async (key, value) => {
        form.value[key] = value
        validateForm()
    }

    const validateForm = async () => {
        errors.value = {} // clear errors before re-eval
        try {
            await yupSchema.validate(form.value, { abortEarly: false })
        }
        catch (error) {
            if (error instanceof ValidationError) {
                setErrors(error.inner)
            }
        }
    }
    const validateField = async ()=>{


        
    }
    const setErrors = (yupErrors)=>{
        yupErrors.forEach((error)=>{
            errors.value[error.path] = error.message
        })
    }
    return {
        form,
        changeFormValue,
        errors,
        currentActiveInput
    }
}

export default useForm