/** @type {import('tailwindcss').Config} */

import tailwindcssAnimate from 'tailwindcss-animate';
import tailwindcssContainerQueries from '@tailwindcss/container-queries';
import tailwindcssTypography from '@tailwindcss/typography';
import tailwindcssForms from '@tailwindcss/forms';

export default {
  content: [
    './src/**/*.{html,js,svelte,ts}',
  ],
  theme: {
    extend: {
      fontFamily: {
        sans: ['Inter', 'ui-sans-serif', 'system-ui', '-apple-system', 'BlinkMacSystemFont', '"Segoe UI"', 'Roboto', '"Helvetica Neue"', 'Arial', '"Noto Sans"', 'sans-serif', '"Apple Color Emoji"', '"Segoe UI Emoji"', '"Segoe UI Symbol"', '"Noto Color Emoji"'],
      },
      boxShadow: {
        button: '0 1px 2px 0 rgb(0 0 0 / 0.15)',
        'inner-border': 'inset 0px 0px 0px 1px rgba(255, 255, 255, 0.1), 0 1px 2px 0 rgb(0 0 0 / 0.15)',
        'card': '0 4px 6px -1px rgb(0 0 0 / 0.1), 0 2px 4px -2px rgb(0 0 0 / 0.1)',
      },
      colors: {
        'gray-950': 'hsl(220 23% 4%)',
        'gray-900': 'hsl(220 23% 8%)',
        'gray-800': 'hsl(220 23% 11%)',
        'gray-700': 'hsl(220 23% 15%)',
        'gray-600': 'hsl(220 23% 20%)',
        'gray-500': 'hsl(220 23% 25%)',
        'gray-400': 'hsl(220 23% 31%)',
        'gray-300': 'hsl(220 23% 43%)',
        'gray-200': 'hsl(220 20% 64%)',
        'gray-100': 'hsl(220 18% 83%)',
        'discord-blue': '#5865F2',
        'discord-blue-hover': '#4853cf',
        'patreon-orange': '#F96855',
        'patreon-orange-hover': '#c95040',
        'theme-bg': 'hsl(220 23% 8%)', 
        'theme-bg-secondary': 'hsl(220 23% 11%)', 
        'theme-card-bg': 'hsl(220 23% 11%)', 
        'theme-text-primary': 'hsl(220 18% 83%)',
        'theme-text-secondary': 'hsl(220 20% 64%)',
        'theme-text-tertiary': 'hsl(220 23% 43%)',
        'theme-border': 'hsl(220 23% 20%)',
        'theme-border-input': 'hsl(220 23% 25%)',
        'theme-accent': '#5865F2',
        'theme-accent-hover': '#4853cf',
        'theme-focus-ring': 'hsl(220 86% 70%)',
      },
      fontSize: {
        h1: ['2.25rem', { lineHeight: '2.5rem', fontWeight: '700' }],
        h2: ['1.875rem', { lineHeight: '2.25rem', fontWeight: '700' }],
        h3: ['1.5rem', { lineHeight: '2rem', fontWeight: '600' }], 
        body: ['0.875rem', { lineHeight: '1.5rem' }],
        sm: ['0.875rem', { lineHeight: '1.25rem' }],
        base: ['1rem', { lineHeight: '1.5rem' }],
        lg: ['1.125rem', { lineHeight: '1.75rem' }],
        xl: ['1.25rem', { lineHeight: '1.75rem' }],
      },
      backgroundImage: {
        'hero-gradient': 'radial-gradient(ellipse 100% 100% at bottom, rgba(37, 99, 235, .1), rgba(37, 99, 235, 0) 50%);',
        'feature-gradient': 'radial-gradient(ellipse 100% 100% at 50% 75%, hsla(227, 89%, 64%, 0.07), hsla(227, 89%, 64%, 0) 60%)',
        'howto-gradient': 'radial-gradient(ellipse 100% 100% at center, hsla(227, 89%, 64%, 0.1), hsla(227, 89%, 64%, 0) 50%)',
      },
      screens: {
        xs: '450px',
      },
    },
  },
  plugins: [
    tailwindcssAnimate,
    tailwindcssContainerQueries,
    tailwindcssTypography,
    tailwindcssForms,
  ],
};