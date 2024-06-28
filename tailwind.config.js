/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: 'selector',
  content: [ 
    "./views/**/*.{html,templ,go}",
    "./blog_posts/**/*.{go,md}",
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}

