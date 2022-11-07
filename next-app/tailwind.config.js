/** @type {import('tailwindcss').Config} */
const myColors = require("./themeOveride.js")
const myTailwindColors = Object.entries(myColors).reduce((object, [key, value]) => {
  let tmp = Object.entries(value).reduce((obj, [k, v]) => {
    return {
      ...obj,
      [key + "-" + k]: v
    }
  }, {})
  return {
    ...object,
    ...tmp,
  }
}, {})

console.log(JSON.stringify(myTailwindColors, 10, 10))

module.exports = {
  important: true,
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx}",
    "./src/@styled-components/**/*.{js,ts,jsx,tsx}",
    "./src/@components/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        ...myTailwindColors
      }
    },
  },
  plugins: [],
}
