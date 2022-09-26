const Webpack = require('webpack')
const path = require('path')
const glob = require('glob')
const { VueLoaderPlugin } = require('vue-loader')

const getEntries = (pattern="", exactFolder)=>{
    const entries = {}
    glob.sync(pattern,{
        // ignore:ignorePatterns
    }).forEach((file)=>{
      console.log(file)
        const entryKey = file.replace(exactFolder,"")
       entries[entryKey] = path.resolve(__dirname, file)
    })
    return entries
}

module.exports = {
    mode: "development",
    watch:true,
    entry:getEntries("assets/js/**/*.js", "assets/js/"),
    resolve: {
      extensions: ['.vue','.js'],
      alias: {
        '@components': path.resolve(__dirname, "templates/components" ),
      }
    },
    output:{
        path: path.resolve(__dirname,'dist/js/'),
        filename: "[name]"
    },
    module: {
        rules: [
          {
            test: /\.vue$/,
            loader: 'vue-loader',
          }
        ]
      },
    plugins : [
        new Webpack.DefinePlugin({ __VUE_OPTIONS_API__: true, __VUE_PROD_DEVTOOLS__: true }), // to remove warn in browser console: runtime-core.esm-bundler.js:3607 Feature flags __VUE_OPTIONS_API__, __VUE_PROD_DEVTOOLS__ are not explicitly defined...
        new VueLoaderPlugin()
    ],
    
}