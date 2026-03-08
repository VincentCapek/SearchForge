# CLAUDE.md

## Project overview

This project is a local search app with:

- Go backend API
- Vue 3 + TypeScript frontend
- Vite for frontend tooling
- UnoCSS for styling

The frontend lives in `assets/`.
The Go API lives at the project root with folders like `server/`, `search/`, and `utils/`.

## Package manager

- Use `pnpm` only
- Do not use npm, yarn, or bun
- Keep scripts and dependencies consistent with pnpm

## Frontend stack

- Vue 3
- TypeScript
- Vite
- UnoCSS

## Backend stack

- Go
- Standard library HTTP server
- JSON API endpoints under `/api/*`

## Styling rules

- Prefer UnoCSS utility classes over large custom CSS files
- Keep custom CSS minimal and only for:
    - complex animations
    - third-party overrides
    - very specific global resets
- Prefer utility-first composition directly in Vue templates
- Reuse repeated utility groups through UnoCSS `shortcuts`
- Do not introduce Tailwind
- Do not introduce CSS frameworks on top of UnoCSS unless explicitly requested

## UnoCSS rules

- Use a dedicated `uno.config.ts` at the frontend root
- Use the Vite integration for UnoCSS
- Import `virtual:uno.css` in the frontend entry file
- Prefer static class names that UnoCSS can detect at build time
- Avoid building classes dynamically like:
    - `` `text-${color}-500` ``
    - `` `p-${size}` ``
- If dynamic styling is needed:
    - use a static mapping object
    - or update UnoCSS `safelist`
    - or configure content extraction explicitly
- If utilities are stored in `.ts` files, remember UnoCSS does not scan `.ts` by default unless configured

## Class authoring conventions

- Prefer readable class groups
- Keep class order roughly like:
    - layout
    - spacing
    - sizing
    - border
    - background
    - typography
    - effects
    - states
- Example:
    - `flex items-center gap-3 rounded-xl border border-slate-800 bg-slate-900/80 px-4 py-3 text-sm text-slate-100 shadow-lg`

## Design direction

- Modern search UI
- Clean, dense, readable
- Dark-first aesthetic unless asked otherwise
- Soft borders, subtle glass effects, restrained shadows
- Good spacing and hierarchy
- Avoid overly flashy gradients
- Prioritize usability over visual gimmicks

## Component conventions

- Keep Vue components small and focused
- Prefer splitting UI into:
    - `components/`
    - `composables/`
    - `lib/`
    - `types/`
- Keep `App.vue` thin
- Move stateful search logic to composables
- Extract presentational blocks into dedicated components

## API integration conventions

- Frontend should call relative endpoints like `/api/search?q=...`
- Assume Vite proxy is used in development
- Handle loading, error, empty, and success states explicitly
- Always type API responses in TypeScript

## Code style expectations

- Write production-ready code
- Keep naming clear and explicit
- Avoid unnecessary abstractions
- Prefer simple functions over complex patterns
- Do not add dependencies unless clearly useful
- Preserve existing project structure unless there is a strong reason to improve it

## When editing frontend code

Claude should:

1. Prefer UnoCSS utilities in templates
2. Reuse existing component patterns
3. Avoid adding large scoped `<style>` blocks unless necessary
4. Suggest `shortcuts` in `uno.config.ts` when utility groups repeat a lot
5. Keep templates readable
6. Keep TypeScript types explicit
7. Maintain accessible markup where possible

## When editing UnoCSS config

Claude may update:

- `uno.config.ts`
- `vite.config.ts`
- `src/main.ts`

Claude should prefer:

- `presetUno()`
- `presetAttributify()` if useful and consistent with the project
- `presetIcons()` only if icons are actually used
- `transformerDirectives()`
- `transformerVariantGroup()`

Claude should not add presets or transformers that are unused.

## Recommended UnoCSS patterns

Use shortcuts for repeated UI patterns like:

- buttons
- cards
- inputs
- badges
- panels
- result items

Example categories of shortcuts:

- `btn-primary`
- `search-input`
- `result-card`
- `panel-glass`

## Search page expectations

For the search interface, always account for:

- query input
- submit button
- loading state
- error state
- empty state
- results list
- responsive layout

## Performance and maintainability

- Prefer static utility strings
- Minimize custom CSS
- Avoid huge monolithic components
- Avoid premature state libraries unless needed
- Keep the frontend easy to refactor into multiple pages later

## Commands Claude can assume

Frontend:

```bash
cd assets
pnpm install
pnpm dev
pnpm build
```

Backend:

```bash
go run .
```

## If Claude proposes styling changes

Claude should:

- explain the component structure briefly
- prefer incremental refactors
- avoid rewriting the whole app unnecessarily
- keep the UI consistent across components

## If Claude adds dynamic classes

Claude must ensure UnoCSS can detect them, either by:

- replacing them with static mappings
- adding a safelist
- or updating extraction config