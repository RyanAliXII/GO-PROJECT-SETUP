const Webpack = require('webpack')
const walk = require('walk');
const path = require('path')
var CopyWebpackPlugin = require('copy-webpack-plugin');
const { VueLoaderPlugin } = require('vue-loader')

const getEntries = (filesLocation) => {
  const entries = {}
  const relativePathFilesLoc = filesLocation.replace("./", "")
  const options = {
    listeners: {
      file: function (root, fileStats, next) {
        folder = root.replace(`${filesLocation}`, "")
        const entryKey = `${folder}/${fileStats.name}`
        entries[entryKey] = `./${relativePathFilesLoc}${entryKey}`
      },

    },
  };
  walker = walk.walkSync(filesLocation, options);
  return entries
}

module.exports = {
  mode: "production",
  watch: true,
  entry: getEntries("./assets/js"),
  resolve: {
    extensions: ['.vue', '.js'],
    alias: {
      '@components': path.resolve(__dirname, "templates/vue-components"),
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
      }
    ]
  },
  plugins: [
    new Webpack.DefinePlugin({ __VUE_OPTIONS_API__: true, __VUE_PROD_DEVTOOLS__: true }), // to remove warn in browser console: runtime-core.esm-bundler.js:3607 Feature flags __VUE_OPTIONS_API__, __VUE_PROD_DEVTOOLS__ are not explicitly defined...
    new VueLoaderPlugin(),
    new CopyWebpackPlugin({
      patterns:
        [
          { from: path.resolve(__dirname, "assets/images"), to: path.resolve(__dirname, "dist/images") }
        ]
    }),
  ],

}