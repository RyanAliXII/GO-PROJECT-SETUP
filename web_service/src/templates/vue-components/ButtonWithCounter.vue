
<template>
<div class="flex flex-col mt-2">
  <span>Count: {{count}}</span>
  <button @click="incrementCount()" class="bg-blue-400 rounded-sm text-white w-24 h-8">{{buttonText}}</button>
</div>
</template>

<script>
export default {
   data(){
       return {
           count:0,
       }
   },
   props:{
        buttonText: {
            type: String,
            default:"Click to increment count"
        }
   },
   methods:{
       incrementCount(){
           this.count += 1;
       }
   }   
   
   
}
</script>