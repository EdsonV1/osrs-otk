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
        // Refined shadow system for depth and elevation
        'xs': '0 1px 2px 0 rgb(0 0 0 / 0.05)',
        'sm': '0 1px 3px 0 rgb(0 0 0 / 0.1), 0 1px 2px -1px rgb(0 0 0 / 0.1)',
        'md': '0 4px 6px -1px rgb(0 0 0 / 0.1), 0 2px 4px -2px rgb(0 0 0 / 0.1)',
        'lg': '0 10px 15px -3px rgb(0 0 0 / 0.1), 0 4px 6px -4px rgb(0 0 0 / 0.1)',
        'xl': '0 20px 25px -5px rgb(0 0 0 / 0.1), 0 8px 10px -6px rgb(0 0 0 / 0.1)',
        '2xl': '0 25px 50px -12px rgb(0 0 0 / 0.25)',
        
        // Specialized shadows for dark theme
        'dark-sm': '0 2px 4px 0 rgb(0 0 0 / 0.3)',
        'dark-md': '0 4px 8px 0 rgb(0 0 0 / 0.3), 0 2px 4px 0 rgb(0 0 0 / 0.2)',
        'dark-lg': '0 8px 16px 0 rgb(0 0 0 / 0.3), 0 4px 8px 0 rgb(0 0 0 / 0.2)',
        'dark-xl': '0 16px 32px 0 rgb(0 0 0 / 0.3), 0 8px 16px 0 rgb(0 0 0 / 0.2)',
        
        // Component-specific shadows
        'card': '0 4px 6px -1px rgb(0 0 0 / 0.15), 0 2px 4px -2px rgb(0 0 0 / 0.1)',
        'card-hover': '0 8px 12px -2px rgb(0 0 0 / 0.2), 0 4px 8px -4px rgb(0 0 0 / 0.1)',
        'button': '0 1px 2px 0 rgb(0 0 0 / 0.15)',
        'button-hover': '0 2px 4px 0 rgb(0 0 0 / 0.2)',
        'navbar': '0 4px 6px -1px rgb(0 0 0 / 0.1), 0 2px 4px -2px rgb(0 0 0 / 0.05)',
        'modal': '0 20px 25px -5px rgb(0 0 0 / 0.3), 0 8px 10px -6px rgb(0 0 0 / 0.2)',
        
        // Inner shadow effects
        'inner': 'inset 0 2px 4px 0 rgb(0 0 0 / 0.1)',
        'inner-border': 'inset 0px 0px 0px 1px hsl(213 93% 67% / 0.2), 0 1px 2px 0 rgb(0 0 0 / 0.1)',
        'glow': '0 0 20px hsl(213 93% 67% / 0.3)',
      },
      colors: {
        // Base gray scale - refined for better contrast
        'gray-950': 'hsl(220 25% 3%)',
        'gray-900': 'hsl(220 24% 6%)',
        'gray-800': 'hsl(220 23% 9%)',
        'gray-700': 'hsl(220 22% 13%)',
        'gray-600': 'hsl(220 20% 18%)',
        'gray-500': 'hsl(220 18% 25%)',
        'gray-400': 'hsl(220 16% 35%)',
        'gray-300': 'hsl(220 14% 50%)',
        'gray-200': 'hsl(220 12% 70%)',
        'gray-100': 'hsl(220 10% 85%)',
        
        // Consistent theme system
        'theme': {
          'bg': {
            'primary': 'hsl(220 24% 6%)',    // Main background - darker for depth
            'secondary': 'hsl(220 23% 9%)',  // Card/section backgrounds
            'tertiary': 'hsl(220 22% 13%)',  // Elevated elements
            'elevated': 'hsl(220 20% 16%)',  // Modals, dropdowns
            'inverse': 'hsl(220 10% 95%)',   // Light mode alternative
          },
          'text': {
            'primary': 'hsl(220 8% 92%)',    // Main content - higher contrast
            'secondary': 'hsl(220 10% 75%)', // Secondary content  
            'tertiary': 'hsl(220 12% 60%)',  // Muted text
            'quaternary': 'hsl(220 14% 45%)', // Very subtle text
            'inverse': 'hsl(220 24% 15%)',   // Dark text for light backgrounds
          },
          'border': {
            'primary': 'hsl(220 20% 18%)',   // Standard borders
            'secondary': 'hsl(220 18% 25%)', // Input borders
            'accent': 'hsl(213 93% 67%)',    // Interactive borders
            'subtle': 'hsl(220 22% 12%)',    // Very subtle dividers
          },
          'accent': {
            'primary': 'hsl(213 93% 67%)',   // Main brand color - OSRS blue
            'hover': 'hsl(213 93% 60%)',     // Hover state
            'active': 'hsl(213 93% 55%)',    // Active state
            'subtle': 'hsl(213 93% 67% / 0.1)', // Background tint
            'muted': 'hsl(213 93% 67% / 0.6)',   // Secondary accent
          },
          'success': {
            'primary': 'hsl(142 71% 45%)',
            'subtle': 'hsl(142 71% 45% / 0.1)',
          },
          'warning': {
            'primary': 'hsl(38 92% 50%)',
            'subtle': 'hsl(38 92% 50% / 0.1)',
          },
          'error': {
            'primary': 'hsl(0 84% 60%)',
            'subtle': 'hsl(0 84% 60% / 0.1)',
          }
        },
        
        // Legacy support - will be phased out
        'discord-blue': 'hsl(213 93% 67%)',
        'discord-blue-hover': 'hsl(213 93% 60%)',
        'patreon-orange': '#F96855',
        'patreon-orange-hover': '#c95040',
        
        // Backwards compatibility (deprecated - use theme.* instead)
        'theme-bg-primary': 'hsl(220 24% 6%)',
        'theme-bg-secondary': 'hsl(220 23% 9%)',
        'theme-bg-tertiary': 'hsl(220 22% 13%)',
        'theme-card-bg': 'hsl(220 23% 9%)',
        'theme-text-primary': 'hsl(220 8% 92%)',
        'theme-text-secondary': 'hsl(220 10% 75%)',
        'theme-text-tertiary': 'hsl(220 12% 60%)',
        'theme-border': 'hsl(220 20% 18%)',
        'theme-border-input': 'hsl(220 18% 25%)',
        'theme-accent': 'hsl(213 93% 67%)',
        'theme-accent-hover': 'hsl(213 93% 60%)',
        'theme-focus-ring': 'hsl(213 93% 67% / 0.3)',
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
      spacing: {
        '4.5': '1.125rem',
        '5.5': '1.375rem',
        '13': '3.25rem',
        '15': '3.75rem',
        '17': '4.25rem',
        '18': '4.5rem',
        '19': '4.75rem',
        '21': '5.25rem',
        '22': '5.5rem',
        '88': '22rem',
        '100': '25rem',
        '112': '28rem',
        '128': '32rem',
      },
      borderRadius: {
        'card': '0.75rem',   // 12px - standard card radius
        'button': '0.5rem',  // 8px - button radius
        'input': '0.5rem',   // 8px - input radius
        'modal': '1rem',     // 16px - modal radius
        'pill': '9999px',    // Full pill shape
      },
      backdropBlur: {
        xs: '2px',
        sm: '4px',
        md: '8px',
        lg: '12px',
        xl: '16px',
        '2xl': '24px',
        '3xl': '32px',
      },
      backgroundImage: {
        // Refined gradient system
        'hero-gradient': 'radial-gradient(ellipse 100% 100% at bottom, hsl(213 93% 67% / 0.1), transparent 50%)',
        'feature-gradient': 'radial-gradient(ellipse 100% 100% at 50% 75%, hsl(213 93% 67% / 0.07), transparent 60%)',
        'howto-gradient': 'radial-gradient(ellipse 100% 100% at center, hsl(213 93% 67% / 0.1), transparent 50%)',
        'accent-gradient': 'linear-gradient(135deg, hsl(213 93% 67%), hsl(213 93% 55%))',
        'card-gradient': 'linear-gradient(135deg, hsl(220 23% 9% / 0.8), hsl(220 22% 13% / 0.4))',
        'glass': 'linear-gradient(135deg, hsl(220 23% 9% / 0.7), hsl(220 22% 13% / 0.3))',
      },
      animation: {
        'fade-in': 'fadeIn 0.2s ease-out',
        'slide-up': 'slideUp 0.3s ease-out',
        'slide-down': 'slideDown 0.3s ease-out',
        'scale-in': 'scaleIn 0.2s ease-out',
        'glow': 'glow 2s ease-in-out infinite alternate',
      },
      keyframes: {
        fadeIn: {
          '0%': { opacity: '0' },
          '100%': { opacity: '1' }
        },
        slideUp: {
          '0%': { transform: 'translateY(10px)', opacity: '0' },
          '100%': { transform: 'translateY(0)', opacity: '1' }
        },
        slideDown: {
          '0%': { transform: 'translateY(-10px)', opacity: '0' },
          '100%': { transform: 'translateY(0)', opacity: '1' }
        },
        scaleIn: {
          '0%': { transform: 'scale(0.95)', opacity: '0' },
          '100%': { transform: 'scale(1)', opacity: '1' }
        },
        glow: {
          '0%': { boxShadow: '0 0 20px hsl(213 93% 67% / 0.3)' },
          '100%': { boxShadow: '0 0 30px hsl(213 93% 67% / 0.5)' }
        }
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