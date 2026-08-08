# API and Admin Endpoints Reference

This document lists the available HTTP endpoints in the FilmyFly Go Fiber project, their purpose, expected inputs, and example usage.

## Base URL

If running locally:

```bash
http://localhost:3000
```

If deployed, replace the host with your server address.

---

## Authentication Notes

- Public routes do not require authentication.
- Admin routes require a valid session cookie created by logging in through the Firebase-based admin login flow.
- The login endpoint is:

```bash
POST /admin/login
```

It expects a JSON body like:

```json
{
  "idToken": "YOUR_FIREBASE_ID_TOKEN"
}
```

---

## Public Routes

### GET /
Renders the public homepage HTML view.

Usage:

```bash
curl http://localhost:3000/
```

What it returns:
- HTML page rendered from the homepage template
- trending movies
- recent movies
- categories

---

## Public API Endpoints

### GET /api/home
Returns homepage API data for the Astro frontend.

Query parameters:
- `page` (optional, default: `1`)

Example:

```bash
curl "http://localhost:3000/api/home?page=1"
```

Response contains:
- `trendingMovies`
- `recentMovies`
- `categories`
- `pagination`

---

### GET /api/movies
Returns a paginated list of movies.

Query parameters:
- `page` (optional, default: `1`)
- `limit` (optional, default: `20`)

Example:

```bash
curl "http://localhost:3000/api/movies?page=1&limit=20"
```

Response contains:
- `success`
- `data` (array of movies)
- `pagination`

---

### GET /api/movies/trending
Returns the movies marked as trending.

Example:

```bash
curl http://localhost:3000/api/movies/trending
```

Response contains:
- `success`
- `data` (array of movies)

---

### GET /api/movies/:slug
Returns details for a specific movie by slug.

Example:

```bash
curl http://localhost:3000/api/movies/avengers-endgame
```

Response contains:
- `movie`
- `relatedMovies`
- `downloadRedirectUrl`

---

### GET /api/categories
Returns all categories with movie counts.

Example:

```bash
curl http://localhost:3000/api/categories
```

Response contains:
- `success`
- `data` (array of categories with `_count.movies`)

---

### GET /api/categories/:slug
Returns a category and its movies.

Query parameters:
- `page` (optional, default: `1`)
- `limit` (optional, default: `20`)

Example:

```bash
curl "http://localhost:3000/api/categories/action?page=1&limit=20"
```

Response contains:
- `category`
- `movies`
- `pagination`

---

### GET /api/search
Searches movies by title.

Query parameters:
- `q` (required for search)
- `page` (optional, default: `1`)
- `limit` (optional, default: `20`)

Example:

```bash
curl "http://localhost:3000/api/search?q=avengers"
```

Response contains:
- `query`
- `movies`
- `pagination`

---

### GET /api/static-pages/:slug
Returns a published static page by slug.

Example:

```bash
curl http://localhost:3000/api/static-pages/about
```

Response contains:
- `success`
- `data` (static page object)

---

### GET /api/astro-settings
Returns Astro-specific settings as a key/value object.

Example:

```bash
curl http://localhost:3000/api/astro-settings
```

Response contains:
- `success`
- `data` (map of settings)

---

## Admin HTML Routes

These routes render HTML pages and require authentication.

### GET /admin/login
Shows the admin login page.

Example:

```bash
curl http://localhost:3000/admin/login
```

---

### POST /admin/login
Authenticates an admin using a Firebase ID token.

Example:

```bash
curl -X POST http://localhost:3000/admin/login \
  -H "Content-Type: application/json" \
  -d '{"idToken":"YOUR_FIREBASE_ID_TOKEN"}'
```

---

### POST /admin/logout
Logs the current admin out and destroys the session.

Example:

```bash
curl -X POST http://localhost:3000/admin/logout
```

---

### GET /admin
Shows the admin dashboard.

Example:

```bash
curl http://localhost:3000/admin
```

---

### GET /admin/system-check
Shows database and system status information.

Example:

```bash
curl http://localhost:3000/admin/system-check
```

---

## Admin Content Management Routes

### GET /admin/settings
Displays the site settings page.

Example:

```bash
curl http://localhost:3000/admin/settings
```

### POST /admin/settings
Updates site settings. It reads all form fields and saves them as setting records.

Example:

```bash
curl -X POST http://localhost:3000/admin/settings \
  -d "siteName=FilmyFly" \
  -d "downloadRedirectUrl=https://example.com"
```

---

### GET /admin/astro-settings
Displays the Astro settings page.

Example:

```bash
curl http://localhost:3000/admin/astro-settings
```

### POST /admin/astro-settings
Updates Astro settings. It reads all submitted form fields and stores them in the Astro settings table.

Example:

```bash
curl -X POST http://localhost:3000/admin/astro-settings \
  -d "siteTitle=FilmyFly" \
  -d "theme=dark"
```

---

### GET /admin/movies
Lists movies in the admin panel.

Query parameters:
- `page` (optional)
- `search` (optional)
- `category` (optional)

Example:

```bash
curl "http://localhost:3000/admin/movies?page=1&search=avengers"
```

### GET /admin/movies/add
Shows the add-movie form.

### POST /admin/movies/add
Creates a new movie.

Example body (form fields):

```bash
curl -X POST http://localhost:3000/admin/movies/add \
  -d "title=New Movie" \
  -d "slug=new-movie" \
  -d "description=Great movie"
```

### GET /admin/movies/bulk-add
Shows the bulk import page.

### POST /admin/movies/bulk-add
Imports multiple movies from JSON.

Example:

```bash
curl -X POST http://localhost:3000/admin/movies/bulk-add \
  -d "moviesJson=[{\"title\":\"Movie 1\",\"slug\":\"movie-1\"}]"
```

### GET /admin/movies/edit/:id
Shows the edit form for a movie.

### POST /admin/movies/edit/:id
Updates an existing movie.

### POST /admin/movies/delete/:id
Deletes a movie.

### POST /admin/movies/trending/add/:id
Adds a movie to the trending list.

### POST /admin/movies/trending/remove/:id
Removes a movie from the trending list.

---

## Static Page Administration

### GET /admin/static-pages
Lists static pages.

### GET /admin/static-pages/add
Shows the add-static-page form.

### POST /admin/static-pages/add
Creates a new static page.

Example:

```bash
curl -X POST http://localhost:3000/admin/static-pages/add \
  -d "title=About" \
  -d "slug=about" \
  -d "content=This is the about page" \
  -d "isPublished=on"
```

### GET /admin/static-pages/edit/:id
Shows the edit form for a static page.

### POST /admin/static-pages/edit/:id
Updates a static page.

### POST /admin/static-pages/delete/:id
Deletes a static page.

---

## Logs Routes

### GET /admin/logs
Shows the logs viewing page.

### GET /admin/logs/data
Returns log data in JSON form.

Example:

```bash
curl http://localhost:3000/admin/logs/data
```

### POST /admin/logs/clear
Clears log files.

Example:

```bash
curl -X POST http://localhost:3000/admin/logs/clear \
  -H "Content-Type: application/json" \
  -d '{"type":"all"}'
```

Valid values for `type`:
- `all`
- `app`
- `error`

### GET /admin/logs/download
Downloads a log file.

Example:

```bash
curl "http://localhost:3000/admin/logs/download?type=app" -o app.log
```

---

## 404 Handling

Any unknown route returns a JSON response like:

```json
{
  "success": false,
  "error": "Route not found"
}
```

---

## Quick Reference Summary

| Route | Method | Purpose |
| --- | --- | --- |
| / | GET | Public homepage |
| /api/home | GET | Homepage API data |
| /api/movies | GET | Movie list |
| /api/movies/trending | GET | Trending movies |
| /api/movies/:slug | GET | Single movie |
| /api/categories | GET | Category list |
| /api/categories/:slug | GET | Category details |
| /api/search | GET | Movie search |
| /api/static-pages/:slug | GET | Static page |
| /api/astro-settings | GET | Astro settings |
| /admin/login | GET/POST | Admin login page and auth |
| /admin/logout | POST | Logout |
| /admin | GET | Dashboard |
| /admin/settings | GET/POST | Settings management |
| /admin/astro-settings | GET/POST | Astro settings management |
| /admin/movies | GET/POST | Movie management |
| /admin/static-pages | GET/POST | Static page management |
| /admin/logs | GET/POST | Logs management |
