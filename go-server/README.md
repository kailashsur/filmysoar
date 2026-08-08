<div align="center">

# 🎬 FilmyFly Go Fiber

### **Fast. Modern. Scalable.**

A production-focused Go backend for the FilmyFly movie discovery platform — built with **Go, Fiber, PostgreSQL, GORM, Firebase Admin SDK, and EJS**.

<p>
  <img src="https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go">
  <img src="https://img.shields.io/badge/Fiber-v2-00A98F?style=for-the-badge&logo=go&logoColor=white" alt="Fiber">
  <img src="https://img.shields.io/badge/PostgreSQL-Database-4169E1?style=for-the-badge&logo=postgresql&logoColor=white" alt="PostgreSQL">
  <img src="https://img.shields.io/badge/GORM-ORM-59666C?style=for-the-badge" alt="GORM">
  <img src="https://img.shields.io/badge/Firebase-Admin-FFCA28?style=for-the-badge&logo=firebase&logoColor=black" alt="Firebase">
  <img src="https://img.shields.io/badge/License-ISC-orange?style=for-the-badge" alt="License">
</p>

<p>
  <a href="#-features">Features</a> •
  <a href="#-architecture">Architecture</a> •
  <a href="#-quick-start">Quick Start</a> •
  <a href="#-api">API</a> •
  <a href="#-project-structure">Structure</a> •
  <a href="#-deployment">Deployment</a>
</p>

</div>

---

## ✨ Overview

**FilmyFly Go Fiber** is a high-performance backend designed to power a modern movie discovery and content-management experience.

The project was migrated from a Node.js/Express architecture to **Go + Fiber**, with a focus on:

- ⚡ Low-latency HTTP handling
- 🧩 Clean modular architecture
- 🗄️ PostgreSQL persistence through GORM
- 🔐 Secure administrator authentication
- 🚀 RESTful APIs for an Astro frontend
- 🎛️ Server-rendered EJS administration
- 📊 Logging, health checks, and operational tooling
- 📦 Simple production deployment

> **Note:** This repository provides backend and content-management infrastructure. Only publish, host, or distribute content for which you have the necessary rights and permissions.

---

## 🖼️ Project Preview

<div align="center">

<img src="docs/assets/filmyfly-go-fiber-banner.png" alt="FilmyFly Go Fiber — Project Preview" width="100%">

</div>

---

## 🧠 Why Go Fiber?

The backend was designed around a simple principle:

> **Keep the API fast, the architecture clean, and the operational surface predictable.**

### Core stack

| Layer | Technology |
|---|---|
| Language | **Go 1.21+** |
| HTTP Framework | **Fiber** |
| ORM | **GORM** |
| Database | **PostgreSQL** |
| Authentication | **Firebase Admin SDK** |
| Admin UI | **EJS** |
| Frontend API Consumer | **Astro** |
| Configuration | **Environment variables** |
| Logging | Application logging / structured logs |
| Deployment | Linux / VPS / Container-ready |

---

# 🚀 Features

### ⚡ Performance

- Fiber-powered HTTP server
- Lightweight request handling
- Efficient PostgreSQL access
- Pagination for large movie collections
- API-first architecture
- Production-oriented middleware

### 🎬 Content Management

- Movie CRUD operations
- Trending movie management
- Category management
- Search
- Static page management
- Bulk movie operations
- Slug-based movie URLs

### 🔐 Security

- Firebase Admin authentication
- Protected admin routes
- Session-based admin access
- Environment-based secrets
- Input validation
- Authentication middleware

### 🛠️ Administration

- Dedicated admin dashboard
- Movie management
- Static page management
- Site settings
- Astro frontend settings
- Application logs
- System/database health checks

### 🌐 API

- RESTful JSON responses
- Homepage aggregation endpoint
- Movie pagination
- Category filtering
- Search
- Movie detail endpoints
- Public configuration endpoints

---

# 🏗️ Architecture

```text
                         ┌──────────────────────┐
                         │     Astro Frontend   │
                         │   Public Website UI  │
                         └──────────┬───────────┘
                                    │
                              REST / JSON
                                    │
                                    ▼
┌───────────────────────────────────────────────────────────┐
│                     Go + Fiber Backend                    │
│                                                           │
│  ┌────────────┐   ┌────────────┐   ┌─────────────────┐  │
│  │   Routes   │ → │ Middleware │ → │    Handlers     │  │
│  └────────────┘   └────────────┘   └────────┬────────┘  │
│                                             │             │
│                                      ┌──────▼──────┐      │
│                                      │    GORM     │      │
│                                      └──────┬──────┘      │
└─────────────────────────────────────────────┼─────────────┘
                                              │
                                              ▼
                                     ┌─────────────────┐
                                     │   PostgreSQL    │
                                     │    Database     │
                                     └─────────────────┘

                         Admin Browser
                              │
                              ▼
                     ┌──────────────────┐
                     │ Firebase Admin   │
                     │ Authentication   │
                     └────────┬─────────┘
                              │
                              ▼
                     ┌──────────────────┐
                     │ EJS Admin Panel  │
                     └──────────────────┘
```

---

# 📁 Project Structure

```text
filmyfly-go-fiber/
│
├── cmd/
│   └── server/
│       └── main.go                 # Application entry point
│
├── internal/
│   ├── config/                     # Environment/config loading
│   ├── database/                   # DB connection and models
│   ├── handlers/                   # HTTP handlers
│   ├── middleware/                 # Auth/security middleware
│   ├── routes/                     # Route registration
│   └── utils/                      # Shared utilities
│
├── views/                          # EJS admin templates
│
├── public/                         # Static assets
│
├── docs/
│   ├── ENDPOINTS.md                # API documentation
│   ├── SCHEMA.md                   # Database documentation
│   └── DEPLOYMENT.md               # Deployment documentation
│
├── logs/                           # Runtime logs
│
├── .env.example                    # Environment template
├── go.mod
├── go.sum
└── README.md
```

---

# ⚡ Quick Start

## 1. Clone

```bash
git clone https://github.com/YOUR_USERNAME/filmyfly-go-fiber.git
cd filmyfly-go-fiber
```

## 2. Configure environment

```bash
cp .env.example .env
```

Then configure your PostgreSQL and Firebase credentials.

Example:

```env
APP_ENV=development
PORT=3000

DATABASE_URL=postgres://username:password@localhost:5432/filmyfly

FIREBASE_PROJECT_ID=your-project-id
FIREBASE_CLIENT_EMAIL=your-client-email
FIREBASE_PRIVATE_KEY="your-private-key"
```

> Never commit `.env`, Firebase private keys, database passwords, or other secrets.

## 3. Install dependencies

```bash
go mod download
```

## 4. Run the application

```bash
go run ./cmd/server
```

The server will be available at:

```text
http://localhost:3000
```

---

# 🧪 Development

### Run tests

```bash
go test ./...
```

### Run with verbose output

```bash
go test -v ./...
```

### Build

```bash
go build -o filmyfly ./cmd/server
```

### Run the compiled binary

```bash
./filmyfly
```

### Optional: Air hot reload

Install Air and run:

```bash
air
```

---

# 🔌 API

The backend exposes a clean JSON API for the Astro frontend.

## Public endpoints

| Method | Endpoint | Description |
|---|---|---|
| `GET` | `/api/home` | Homepage data |
| `GET` | `/api/movies` | Paginated movies |
| `GET` | `/api/movies/trending` | Trending movies |
| `GET` | `/api/movies/:slug` | Movie details |
| `GET` | `/api/categories` | All categories |
| `GET` | `/api/categories/:slug` | Category + movies |
| `GET` | `/api/search?q=query` | Search movies |
| `GET` | `/api/static-pages/:slug` | Static page content |
| `GET` | `/api/astro-settings` | Public frontend settings |

## Admin routes

| Method | Endpoint | Access |
|---|---|---|
| `GET` | `/admin/login` | Public |
| `POST` | `/admin/login` | Public |
| `POST` | `/admin/logout` | Authenticated |
| `GET` | `/admin` | Protected |
| `GET` | `/admin/movies` | Protected |

For the complete endpoint reference:

```text
docs/ENDPOINTS.md
```

---

# 🔐 Authentication Flow

```text
Admin
  │
  ▼
/admin/login
  │
  ▼
Firebase Authentication
  │
  ├── Invalid → Access denied
  │
  └── Valid
       │
       ▼
  Session Created
       │
       ▼
 Protected Admin Routes
```

### Security principles

- Secrets belong in environment variables.
- Firebase service credentials must never be committed.
- Administrative endpoints must remain protected.
- Sessions should use secure cookie settings in production.
- Production deployments should use HTTPS.
- Keep dependencies updated.

---

# 🗄️ Database

The application uses **PostgreSQL** with **GORM**.

Typical entities include:

```text
Movies
Categories
Static Pages
Site Settings
Astro Settings
Admin / Session Data
```

A database connection is initialized during application startup.

For schema information, see:

```text
docs/SCHEMA.md
```

---

# 📦 Production Build

Build a Linux binary:

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
go build -ldflags="-s -w" -o filmyfly ./cmd/server
```

Then run:

```bash
./filmyfly
```

For a production environment, place the application behind a reverse proxy such as Nginx and enable HTTPS.

Recommended production flow:

```text
Internet
   │
   ▼
Cloudflare / CDN
   │
   ▼
Nginx
   │
   ▼
Go Fiber
   │
   ▼
PostgreSQL
```

---

# 🩺 Health & Operations

A production backend should be observable.

Recommended operational checks:

```text
Application
├── HTTP availability
├── Database connectivity
├── Error logs
├── Authentication status
├── Request latency
└── Resource usage
```

The built-in **System Check** and **Logs** sections provide an administration surface for operational visibility.

---

# 🎨 Admin Panel

The admin interface provides centralized control over the platform.

```text
Admin Dashboard
│
├── 🎬 Movies
│   ├── Add Movie
│   ├── Edit Movie
│   ├── Search
│   ├── Categories
│   ├── Trending
│   └── Bulk Add
│
├── 📄 Static Pages
│
├── ⚙️ Settings
│
├── 🚀 Astro Settings
│
├── 📊 Logs
│
└── 🔍 System Check
```

---

# 🧩 Design Principles

### Separation of concerns

```text
Routes
  ↓
Middleware
  ↓
Handlers
  ↓
Business Logic
  ↓
Database
```

### API-first

The backend does not depend on the frontend rendering architecture.

That makes it possible to connect:

- Astro
- Mobile applications
- Internal tools
- Other web clients
- Future frontend implementations

without rewriting the core backend.

---

# 📈 Performance Goals

FilmyFly Go Fiber is designed around:

- Minimal request overhead
- Efficient database queries
- Pagination
- Reusable middleware
- Stateless public APIs
- Production-ready compilation
- Clean separation between public and administrative workloads

Always benchmark real workloads before making production capacity claims.

---

# 🛡️ Security Checklist

Before production:

```text
[ ] HTTPS enabled
[ ] Strong database password
[ ] Firebase credentials secured
[ ] .env excluded from Git
[ ] Admin sessions configured securely
[ ] CORS configured intentionally
[ ] Rate limiting considered
[ ] Request validation enabled
[ ] Production logging configured
[ ] Database backups configured
[ ] Dependencies updated
[ ] Error responses do not expose secrets
```

---

# 🤝 Contributing

Contributions are welcome.

### Development workflow

```bash
git checkout -b feature/my-feature

go test ./...

git add .
git commit -m "feat: add my feature"

git push origin feature/my-feature
```

Then open a pull request.

### Commit convention

Recommended:

```text
feat: add movie search
fix: resolve category pagination
perf: optimize movie query
refactor: simplify auth middleware
docs: update API documentation
chore: update dependencies
```

---

# 🗺️ Roadmap

- [x] Go + Fiber backend
- [x] PostgreSQL integration
- [x] GORM integration
- [x] Firebase Admin authentication
- [x] Admin dashboard
- [x] Movie management
- [x] REST API
- [x] Astro integration
- [ ] Automated API documentation
- [ ] Advanced caching
- [ ] Rate limiting
- [ ] Metrics dashboard
- [ ] Docker deployment
- [ ] CI/CD pipeline
- [ ] Automated database backups

---

# 📚 Documentation

| Document | Description |
|---|---|
| `docs/ENDPOINTS.md` | Complete API reference |
| `docs/SCHEMA.md` | Database schema |
| `docs/DEPLOYMENT.md` | Production deployment |
| `.env.example` | Configuration reference |

---

# ⚖️ Content & Copyright

FilmyFly Go Fiber is backend software for managing and presenting movie-related metadata and content.

You are responsible for ensuring that any movies, images, metadata, streams, downloads, or other media distributed through a deployment of this software comply with applicable copyright laws, licenses, platform policies, and other regulations.

**Do not use this project to distribute content you do not have permission to distribute.**

---

# 📜 License

This project is licensed under the **ISC License**.

See the `LICENSE` file for details.

---

<div align="center">

### Built for speed. Designed for scale. Engineered for the next generation of movie platforms.

<br>

**FilmyFly Go Fiber**

<sub>Go • Fiber • PostgreSQL • GORM • Firebase • EJS</sub>

<br><br>

⭐ **If this project helps you, consider giving it a star.**

</div>
