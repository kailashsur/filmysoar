IMPORTANT: This message contains the COMPLETE CONTEXT for the task. Do not rely on previous conversation context.

PROJECT:
FilmySoar Astro frontend.

TARGET:
C:\Users\Kailash\Desktop\Projects\filmyfly

REFERENCE ONLY:
https://github.com/kailashsur/filmyfly-astro

The reference project is NOT necessarily correctly implementing ISR. Use it only to understand existing patterns. Do not blindly copy it.

==================================================
GOAL
==================================================

I want to convert the Astro frontend of FilmySoar into a production-quality Vercel ISR architecture.

The desired architecture is:

                    GO BACKEND
                        |
                        | API
                        v
                  ASTRO + VERCEL
                        |
          +-------------+-------------+
          |             |             |
         ISR           SSR         STATIC
          |             |             |
       movies        search       about/etc
       categories
       pagination

The most important requirement:

DO NOT generate every movie page during `astro build`.

If the backend contains:

10,000 movies
50,000 movies
100,000 movies

the Astro build must NOT generate all of those movie pages.

Movie pages must be generated ON DEMAND when a visitor requests them and then cached using Vercel ISR.

Example:

First request:

/movie/avatar-2026

Vercel
  -> no cached page
  -> Astro server renders page
  -> Astro requests Go backend
  -> HTML generated
  -> Vercel caches result
  -> response returned

Later requests:

/movie/avatar-2026

Vercel ISR/CDN
  -> cached HTML
  -> fast response

==================================================
RENDERING REQUIREMENTS
==================================================

These routes should use ISR:

/
 
/page/[page]

/movie/[slug]

/categories/[slug]

/categories/[slug]/page/[page]

These routes should remain dynamic SSR and MUST NOT use ISR:

/search

/search?q=anything

/search?q=anything&page=2

These informational pages should remain static/prerendered where appropriate:

/about
/contact
/dmca
/privacy-policy
/movie-request

==================================================
VERY IMPORTANT: PAGINATION
==================================================

The existing project currently uses query-string pagination such as:

/?page=2

/categories/action?page=2

I want pagination converted to pathname-based routes.

Homepage:

/
 
/page/2

/page/3

/page/4

Category:

/categories/action

/categories/action/page/2

/categories/action/page/3

/categories/action/page/4

Do NOT use:

/?page=2

/categories/action?page=2

for the new internal pagination URLs.

Search is the exception and must remain query based:

/search?q=avatar

/search?q=avatar&page=2

==================================================
MOVIE ROUTES
==================================================

The movie route should remain conceptually:

/movie/[slug]

For example:

/movie/avatar-2026

CRITICAL:

Do NOT use getStaticPaths() to fetch the complete movie catalog.

Do NOT fetch 1,000 / 10,000 / 50,000 movies during build just to generate movie pages.

Do NOT pre-generate the entire movie catalog.

The movie page must be dynamically rendered and cached through Vercel ISR.

The existing Go backend remains the source of truth.

Preserve the existing movie API contract.

==================================================
GO BACKEND
==================================================

The architecture must remain:

PostgreSQL
    |
    v
Go Fiber backend
    |
    v
Astro
    |
    v
Vercel ISR/CDN

Do not move the movie database into Astro.

Do not create a second database.

Do not unnecessarily modify the existing API.

Inspect the existing API helpers and preserve them.

==================================================
VERCEL
==================================================

The project should use:

@astrojs/vercel

with:

output: "server"

and the official ISR mechanism supported by the ACTUAL installed version of @astrojs/vercel.

Before implementing ISR:

1. Inspect package.json.
2. Determine the installed @astrojs/vercel version.
3. Inspect the package typings/documentation if necessary.
4. Use only configuration supported by that version.

Do NOT invent unsupported configuration properties.

==================================================
SEARCH EXCLUSION
==================================================

Search MUST NOT become ISR.

The search page reads:

Astro.url.searchParams

and sends queries to the backend.

Keep this behavior dynamic.

Do not generate search paths using getStaticPaths().

Do not create static search pages.

Do not cache arbitrary search queries as ISR pages.

If the installed Vercel adapter provides an official route exclusion mechanism, use it correctly.

If the version uses another official mechanism, use that instead.

==================================================
ISR INVALIDATION
==================================================

The Go backend is the source of truth.

I want on-demand invalidation when movie data changes.

Example:

Cached:

/movie/avatar

Admin updates Avatar in the Go backend.

The backend should be able to invalidate:

/movie/avatar

Then the next request should fetch fresh data from the Go API and create a new ISR cache.

Use the OFFICIAL Astro/Vercel ISR invalidation mechanism supported by the installed adapter version.

Do not invent a custom cache implementation.

If a secret/token is required:

- use an environment variable
- never hardcode it
- protect the invalidation mechanism
- do not expose an unauthenticated public cache purge endpoint

If backend changes are required, make the minimum production-quality changes necessary.

==================================================
SEO
==================================================

Do NOT break existing SEO.

Preserve:

- title
- meta description
- canonical URLs
- Open Graph
- Twitter metadata
- structured data
- movie metadata
- category metadata
- sitemap behavior
- robots behavior

Because pagination is changing, make sure canonical URLs use the new routes.

For example:

/page/2

NOT:

/?page=2

And:

/categories/action/page/2

NOT:

/categories/action?page=2

Avoid duplicate indexable URLs.

==================================================
UI
==================================================

This is NOT a redesign.

Do not change the existing visual design unless absolutely necessary.

Preserve:

- header
- footer
- navigation
- movie cards
- movie detail UI
- category UI
- search UI
- pagination appearance
- colors
- typography
- responsive design
- animations
- components

Only change routing/rendering/data-fetching logic required for this architecture.

==================================================
IMPORTANT: INSPECT THE EXISTING PROJECT
==================================================

Before modifying anything, inspect:

astro.config.*
package.json

and all relevant files under:

src/pages/
src/components/
src/layouts/
src/lib/

Also inspect:

- API helper functions
- movie page
- homepage
- category page
- search page
- pagination component
- SEO components
- sitemap configuration
- environment variable usage

Search the entire repository for:

getStaticPaths

?page=

Astro.url.searchParams

/movie/

categories/

pagination

Do not assume the exact current file structure.

Use the ACTUAL repository state.

==================================================
IMPLEMENTATION STRATEGY
==================================================

Implement this in phases, but DO NOT spend 15+ minutes only thinking/planning.

PHASE 1:
Inspect the project and implement route restructuring.

Create the required pathname-based pagination routes:

/page/[page]

/categories/[slug]/page/[page]

Update internal pagination links.

Keep search query parameters unchanged.

PHASE 2:
Configure Vercel ISR using the installed compatible @astrojs/vercel version.

Make:

/ 
/page/[page]
/movie/[slug]
/categories/[slug]
/categories/[slug]/page/[page]

ISR.

Make search dynamic SSR.

Keep informational pages static.

Remove any movie getStaticPaths implementation that generates the entire movie catalog.

PHASE 3:
Implement official Vercel/Astro on-demand invalidation for movie pages.

Document how the Go backend can invalidate:

/movie/<slug>

PHASE 4:
Verify everything.

==================================================
BUILD REQUIREMENT
==================================================

Run:

npm run check

npm run build

The build MUST NOT generate all movie pages.

A correct build should show static informational routes being prerendered, while movie and pagination routes remain server/ISR routes.

Do not consider the task complete merely because `npm run build` succeeds.

Verify the actual generated Vercel output/configuration where possible.

==================================================
EXPECTED FINAL ROUTE ARCHITECTURE
==================================================

Approximately:

src/pages/

    index.astro

    page/
        [page].astro

    movie/
        [slug].astro

    categories/
        [slug].astro
        [slug]/
            page/
                [page].astro

    search.astro

    about.astro
    contact.astro
    dmca.astro
    privacy-policy.astro
    movie-request.astro

The exact structure may differ if Astro's routing or the existing project makes another structure cleaner.

==================================================
FINAL VERIFICATION
==================================================

Test:

/
/page/2
/page/3

/movie/<real-existing-slug>

/categories/<real-existing-category>

/categories/<real-existing-category>/page/2

/search?q=<real-search-term>

/search?q=<real-search-term>&page=2

/about
/contact
/dmca
/privacy-policy
/movie-request

Verify:

1. Pages render.
2. API requests work.
3. Movie pages are NOT generated during build.
4. Pagination pages are NOT all generated during build.
5. Movie pages are on-demand ISR.
6. Category pages are ISR.
7. Pagination pages are ISR.
8. Search remains dynamic SSR.
9. Static pages remain prerendered.
10. SEO remains correct.
11. Internal links use the new pagination URLs.
12. No redirect loops.
13. No TypeScript/Astro errors.
14. Vercel build output is valid.

==================================================
FINAL REPORT
==================================================

When finished, report:

1. Files changed.
2. New route structure.
3. ISR routes.
4. SSR routes.
5. Static routes.
6. How movie pages are generated on demand.
7. How ISR invalidation works.
8. Required environment variables.
9. How the Go backend should trigger movie invalidation.
10. npm run check result.
11. npm run build result.
12. Any remaining issues.

IMPORTANT:

Actually modify the repository.

Do not merely tell me how to implement it.

Do not stop after only changing astro.config.

Do not revert successful work simply because route restructuring touches multiple files.

Do not use getStaticPaths() to generate the movie catalog.

Start by inspecting the current repository and then implement Phase 1.