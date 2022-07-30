const path  = require('path')
const Webpack = require('webpack')
module.exports = {
    mode: "production",
    watch:true,
    entry: {
        home:[
            path.resolve(__dirname,'web/static/js/home/homepage.js')
        ],
        product:[
            path.resolve(__dirname,'web/static/js/product/index.js')
        ],
    },
    output:{
        path: path.resolve(__dirname,'web/static/js/build'),
        filename: "[name].js"
    },
    plugins : [
        new Webpack.DefinePlugin({ __VUE_OPTIONS_API__: true, __VUE_PROD_DEVTOOLS__: true }), // to remove warn in browser console: runtime-core.esm-bundler.js:3607 Feature flags __VUE_OPTIONS_API__, __VUE_PROD_DEVTOOLS__ are not explicitly defined...
      
    ],
}