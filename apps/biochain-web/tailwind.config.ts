/** @type {import('tailwindcss').Config} */
import type { Config } from 'tailwindcss'

export default <Partial<Config>>{
  content: [
    './app.vue',

    './formkit.theme.ts',
    './components/**/*.{js,vue,ts}',
    './layouts/**/*.vue',
    './**/*.vue',
    './plugins/**/*.{js,ts}',
    './primevue/presets/**/*.{js,vue,ts}',
    './error.vue',
    './lib/primevue/presets/wind/**/*.{js,vue,ts}',
  ],
  darkMode: 'class',

  theme: {
    extend: {
      colors: {
        surface: {
          950: 'hsl(0 0 21)',
          900: 'hsl(0 0 33)',
          800: 'hsl(0 0 40)',
          700: 'hsl(0 0 48)',
          600: 'hsl(0 0 53)',
          500: 'hsl(0 0 60)',
          400: 'hsl(0 0 65)',
          300: 'hsl(0 0 70)',
          200: 'hsl(0 0 87)',
          100: 'hsl(0 0 93)',
          50: 'hsl(0 0 97)',
          0: 'hsl(0 0 100)',
        },
        primary: {
          950: 'hsl(0 0 21)',
          900: 'hsl(0 0 33)',
          800: 'hsl(0 0 40)',
          700: 'hsl(0 0 48)',
          600: 'hsl(0 0 53)',
          500: 'hsl(0 0 60)',
          400: 'hsl(0 0 65)',
          300: 'hsl(0 0 70)',
          200: 'hsl(0 0 87)',
          100: 'hsl(0 0 93)',
          50: 'hsl(0 0 97)',
          0: 'hsl(0 0 100)',
        },
      },
    },
  },
  safelist: [
    'bg-surface-50',
    'bg-surface-100',
    'bg-surface-200',
    'bg-surface-300',
    'bg-surface-400',
    'bg-surface-500',
    'bg-surface-600',
    'bg-surface-700',
    'bg-surface-800',
    'bg-surface-900',
    'bg-surface-950',
    'bg-blue-50',
    'bg-blue-100',
    'bg-blue-200',
    'bg-blue-300',
    'bg-blue-400',
    'bg-blue-500',
    'bg-blue-600',
  ],
  plugins: [],
}
