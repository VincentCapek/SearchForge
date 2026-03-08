import { defineConfig } from 'unocss'
import presetWind3 from '@unocss/preset-wind3'
import transformerVariantGroup from '@unocss/transformer-variant-group'
import transformerDirectives from '@unocss/transformer-directives'

export default defineConfig({
  presets: [
    presetWind3(),
  ],
  transformers: [
    transformerVariantGroup(),
    transformerDirectives(),
  ],
  theme: {
    fontFamily: {
      sans: [
        'Inter',
        'system-ui',
        '-apple-system',
        'BlinkMacSystemFont',
        '"Segoe UI"',
        'sans-serif',
      ],
    },
    colors: {
      canvas: {
        page: '#f5f6f8',
      },
      surface: {
        base: '#ffffff',
      },
      line: {
        soft: '#ebebeb',
        default: '#dfe1e5',
      },
      text: {
        primary: '#202124',
        secondary: '#4d5156',
        muted: '#5f6368',
        faint: '#70757a',
        placeholder: '#9aa0a6',
      },
      brand: {
        500: '#1a73e8',
        600: '#1765cc',
        700: '#1558d6',
        ring: '#4285f4',
      },
      success: {
        600: '#0d652d',
      },
      danger: {
        600: '#d93025',
      },
    },
    boxShadow: {
      soft: '0 0 0 1px rgba(16, 24, 40, 0.02)',
    },
  },
  shortcuts: {
    // Layout
    'page-shell':
      'min-h-screen bg-canvas-page font-sans text-text-primary antialiased',
    'page-container':
      'mx-auto w-full max-w-[740px] px-5 pt-10 pb-20',

    // Search
    'search-form':
      'grid grid-cols-1 gap-2 my-6 md:grid-cols-[1fr_auto]',
    'search-input':
        'flex-1 h-9 rounded-lg border border-solid border-line-default bg-surface-base px-3 text-sm text-text-primary appearance-none shadow-none outline-none transition-colors ' +
        'placeholder:text-text-placeholder focus:border-brand-ring focus:ring-2 focus:ring-brand-ring/15 disabled:cursor-not-allowed disabled:opacity-60',
    'btn-search':
      'h-9 shrink-0 rounded-lg border-0 bg-brand-500 px-4 text-sm font-medium text-white cursor-pointer transition-colors ' +
      'hover:bg-brand-600 active:bg-brand-700 disabled:cursor-not-allowed disabled:opacity-60',

    // Results
    'result-card':
      'group border-b border-line-soft py-5 last:border-0',
    'result-domain-row':
      'mb-1.5 flex items-center gap-1 text-xs leading-none text-text-muted',
    'result-title-link':
      'text-xl leading-snug font-normal text-brand-700 group-hover:underline',
    'result-url-text':
      'mt-0.5 block truncate text-sm text-success-600',
    'result-snippet':
      'mt-1.5 text-sm leading-relaxed text-text-secondary',

    // Feedback
    'feedback-panel':
      'flex items-center gap-2.5 py-6 text-sm text-text-muted',
    'error-panel':
      'py-4 text-sm text-danger-600',
    'count-line':
      'mb-1 border-b border-line-soft py-2 text-xs text-text-faint',
  },
})