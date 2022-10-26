import { createApp, ref, computed, onMounted } from "vue/dist/vue.esm-bundler";
import VueMultiSelect from "@vueform/multiselect";
import "@vueform/multiselect/themes/default.css";
createApp({
  components: {
    "multi-select": VueMultiSelect,
  },
  setup() {
    const variants = ref([]);
    const variantsPermutation = computed(() => {
      if (variants.value.length === 0) return [];
      const baseVariant = variants.value[0]; //root nodes
      const result = [];
      
      const permute = (variantItems, depth = 1, prev = []) => {
        for (let i = 0; i < variantItems.length; i++) {
          if (depth < variants.value.length) {
            permute(variants.value[depth].items, depth + 1, [
              ...prev,
              variantItems[i],
            ]);
          } else {
            result.push([...prev, variantItems[i]].toString());
          }
        }
      };
      permute(baseVariant.items);
      return result;
    });
    const createVariant = () => {
      variants.value.push({
        name: `Variant ${variants.value.length + 1} `,
        items: [],
      });
    };
    const submit = () => {
      console.log(variants.value);
    };

    return {
      variants,
      createVariant,
      submit,
      variantsPermutation,
    };
  },
}).mount("#root");
