# FilmyFly Website

FilmyFly is a mobile-first movie discovery website built with Astro. It displays movies, trending titles, categories, search results, and single movie details using the FilmyFly backend API.

## Tech stack

- Astro 5
- Node.js adapter
- Server-side rendering
- Responsive HTML and CSS
- REST API integration
- Optional GSAP animation support

## Requirements

- Node.js 18.17 or newer
- npm 9 or newer
- Running FilmyFly backend API

## Installation

```bash
npm install
```

Create a local environment file:

```bash
cp .env.example .env
```

On Windows PowerShell:

```powershell
Copy-Item .env.example .env
```

Set the backend API URL in `.env`:

```env
PUBLIC_API_BASE_URL=http://127.0.0.1:3000
```

For production, replace it with the public backend URL.

## Development

```bash
npm run dev
```

Open `http://localhost:4321` in a browser.

## Production build

```bash
npm run build
npm run preview
```

The production server can be started with:

```bash
node ./dist/server/entry.mjs
```

## Website routes

| Route | Purpose |
| --- | --- |
| `/` | Home page with trending and latest movies |
| `/search?q=movie` | Search results |
| `/categories` | All movie categories |
| `/categories/:slug` | Movies inside one category |
| `/movie/:slug` | Single movie details |

## Backend endpoints

The website expects these GET endpoints:

| Endpoint | Usage |
| --- | --- |
| `/api/home?page=1` | Homepage movies, trending data, categories, and pagination |
| `/api/search?q=query&page=1` | Search movies |
| `/api/categories` | All categories and movie counts |
| `/api/categories/:slug?page=1` | Movies for one category |
| `/api/movies/:slug` | Single movie and related movies |

The API response should use a `data` object. Movie lists should provide a `pagination` object when pagination is supported.

## Central website configuration

Update [onboard.json](./onboard.json) for shared website information:

- Website name
- Domain
- Logo mark and logo text
- Favicon and theme colors
- SEO title and description
- Keywords
- Social profiles
- Homepage copy
- Copyright information

The Astro pages access this configuration through [src/config/site.ts](./src/config/site.ts).

When changing the brand, update the values in `onboard.json` first. Do not add API keys, passwords, database credentials, or private tokens to this file.

## SEO and deployment checklist

Before deployment:

1. Set `PUBLIC_API_BASE_URL` to the production API URL.
2. Update `site.domain` in `onboard.json`.
3. Update the organization URL and default social image in `onboard.json`.
4. Confirm the API allows requests from the website domain.
5. Run `npm run build` successfully.
6. Test home, search, category, and movie routes.
7. Test the website on a narrow mobile viewport.
8. Confirm poster images load over HTTPS.

## Performance notes

- Movie posters use lazy loading except for the first visible items.
- API requests use a short timeout and safe empty-state fallbacks.
- The layout is mobile-first.
- Images should be served in WebP or AVIF where possible.
- Avoid adding large client-side libraries to content pages.

## Encoding

All source files must be saved as UTF-8. Keep this tag in every HTML document:

```html
<meta charset="UTF-8" />
```

When writing files from PowerShell, use UTF-8 encoding to prevent corrupted characters such as `â€œ` or `â†`.

## Project files

```text
src/
  config/site.ts       Shared onboard configuration
  pages/index.astro    Homepage
  pages/search.astro   Search page
  pages/categories.astro
  pages/categories/[slug].astro
  pages/movie/[slug].astro
onboard.json           Brand and SEO configuration
setup.json              Initial setup and deployment metadata
.env.example            Environment variable template
astro.config.mjs        Astro server configuration
```

## License

This project is private unless the owner specifies otherwise.
