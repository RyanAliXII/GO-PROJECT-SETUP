/** @type {import('tailwindcss').Config} */
module.exports = {
  mode:"jit",
  content: [
    "./src/templates/*.{html,js,vue}",
    "./src/templates/**/*.{html,js,vue}",
    "./src/vue/**/*.{html,js,vue}"
  ],
  theme: {
    extend: {},
  },
  plugins:[],
}
