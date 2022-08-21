
const Webpack = require('webpack')
const path = require('path')
const entries = require('./js-walker')

module.exports = {
    mode: "development",
    watch:true,
    entry:entries,
    output:{
        path: path.resolve(__dirname,'web/static/js/dist'),
        filename: "[name].js"
    },
    plugins : [
        new Webpack.DefinePlugin({ __VUE_OPTIONS_API__: true, __VUE_PROD_DEVTOOLS__: true }), // to remove warn in browser console: runtime-core.esm-bundler.js:3607 Feature flags __VUE_OPTIONS_API__, __VUE_PROD_DEVTOOLS__ are not explicitly defined...
    ],
}