# Quick Fix Applied

## Issue
PostgreSQL column names are case-sensitive when quoted. The GORM models use camelCase (`createdAt`, `releaseYear`, `categoryId`) but queries weren't quoting them properly.

## Solution
Updated all API handlers to use quoted column names in SQL queries:

### Files Fixed:
- ✅ `internal/handlers/api/home.go` - Homepage data
- ✅ `internal/handlers/api/movies.go` - Movies endpoints  
- ✅ `internal/handlers/api/categories.go` - Categories endpoints
- ✅ `internal/handlers/api/search.go` - Search endpoint

### Changes Made:
- `createdAt` → `"createdAt"`
- `releaseYear` → `"releaseYear"`  
- `categoryId` → `"categoryId"`
- `order` → `"order"`

## Test Now
The server is running. Try these endpoints:

```bash
# Homepage with movies
curl http://localhost:3000/api/home

# Movies list
curl http://localhost:3000/api/movies

# Search
curl "http://localhost:3000/api/search?q=action"
```

All endpoints should now return proper movie data! 🎉
