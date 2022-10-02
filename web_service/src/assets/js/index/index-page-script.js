import {createApp} from 'vue/dist/vue.esm-bundler.js'
import ButtonWithCounter from '@components/ButtonWithCounter'
import Multiselect from '@vueform/multiselect'
import "@vueform/multiselect/themes/default.css"

createApp({
    components:{
       "button-with-counter":ButtonWithCounter,
        "multi-select":Multiselect
    
    },
    data(){
        return {
            selected: null,
            options: [
                { value: 'batman', label: 'Batman' },
                { value: 'robin', label: 'Robin' },
                { value: 'joker', label: 'Joker' },
            ]
        }
    },
    mounted(){
        console.log("MOUNTED 2")
    },
    compilerOptions:{
        delimiters:["{%", "%}"],
    },
    methods:{
    }

}).mount("#root")


