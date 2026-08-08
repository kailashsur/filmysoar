# Smoke tests

Run the service first, then execute from the repository root:

```powershell
python scripts/smoke_test.py
```

The script uses only Python’s standard library. It runs `gofmt`, `go vet`,
`go test`, and `go build`, then checks health, public API endpoints,
pagination edge cases, authentication protection, and (when supplied) CSRF.

For authenticated admin checks, provide the browser session cookie:

```powershell
python scripts/smoke_test.py --admin-cookie "session_id=YOUR_VALUE"
```

Use `--skip-build` when Go checks were already run, or `--base-url` to test a
deployed instance.
