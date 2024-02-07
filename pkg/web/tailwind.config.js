/** @type {import('tailwindcss').Config} */
export default {
  content: [
    './app.vue',

    './formkit.theme.ts', // <-- add your theme file
    './components/**/*.{js,vue,ts}',
    './layouts/**/*.vue',
    './pages/**/*.vue',
    './plugins/**/*.{js,ts}',
    './error.vue',
  ],
  darkMode: 'class',

  theme: {
    extend: {},
  },
  plugins: [],
}
