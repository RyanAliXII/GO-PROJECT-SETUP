const Webpack = require('webpack')
const fsw = require('@nodelib/fs.walk');
const path = require('path')
var CopyWebpackPlugin = require('copy-webpack-plugin');
const { VueLoaderPlugin } = require('vue-loader')

const getEntries = (filesPath) => {
  const files = fsw.walkSync(filesPath)
  const entries = {}
  files.forEach((entry)=>{
      const isDirectory = entry.dirent.isDirectory()
      if(!isDirectory){
          entries[entry.path.replace(filesPath,"")] = path.resolve(__dirname, entry.path)
      }
  })
    return entries
}
module.exports = {
  mode: "production",
  watch:false,
  entry: getEntries("src/assets/js"),
  resolve: {
    extensions: ['.vue', '.js','.css'],
    alias: {
      '@components': path.resolve(__dirname, "src/vue/components"),
      '@hooks': path.resolve(__dirname, "src/vue/hooks"),
    }
  },
  output: {
    path: path.resolve(__dirname, 'dist/js/'),
    filename: "[name]"
  },
  module: {
    rules: [
      {
        test: /\.vue$/,
        loader: 'vue-loader',
      },
      {
        test: /\.css$/i,
        use: ["style-loader", "css-loader"],
      },
    ]
  },
  plugins: [
    new Webpack.DefinePlugin({ __VUE_OPTIONS_API__: true, __VUE_PROD_DEVTOOLS__: true }), // to remove warn in browser console: runtime-core.esm-bundler.js:3607 Feature flags __VUE_OPTIONS_API__, __VUE_PROD_DEVTOOLS__ are not explicitly defined...
    new VueLoaderPlugin(),
    new CopyWebpackPlugin({
      patterns:
        [
          { from: path.resolve(__dirname, "src/assets/images"), to: path.resolve(__dirname, "dist/images") }
        ]
    }),
  ],

}