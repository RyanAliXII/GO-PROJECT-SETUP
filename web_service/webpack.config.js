const Webpack = require('webpack')
const path = require('path')
const glob = require('glob')

const getEntries = (pattern="", ignorePatterns=[])=>{
    const entries = {}
    glob.sync(pattern,{
        ignore:ignorePatterns
    }).forEach((file)=>{
        const entryKey = file.replace("public/js/","")
       entries[entryKey] = path.resolve(__dirname, file)
    })
    return entries
}


module.exports = {
    mode: "development",
    watch:true,
    entry:getEntries("public/js/**/*.js",["public/js/dist/**/*.js"]),
    output:{
        path: path.resolve(__dirname,'public/js/dist'),
        filename: "[name]"
    },
    plugins : [
        new Webpack.DefinePlugin({ __VUE_OPTIONS_API__: true, __VUE_PROD_DEVTOOLS__: true }), // to remove warn in browser console: runtime-core.esm-bundler.js:3607 Feature flags __VUE_OPTIONS_API__, __VUE_PROD_DEVTOOLS__ are not explicitly defined...
    ],
}