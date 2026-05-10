# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Run dev server
make run               # go run ./cmd/api

# Build binary
make build             # outputs to bin/api

# Tests
make test              # all packages
make test-v            # verbose
make test-pkg pkg=internal/service  # single package

# Database
make docker-up         # start Postgres in Docker
make migrate-up        # apply all migrations
make migrate-down      # roll back 1 migration

# Code generation
make sqlc              # regenerate repository code from db/queries/*.sql

# Tidy
make tidy              # go mod tidy
```

## Architecture

**Request flow:** `HTTP → Echo router → Middleware → Handler → Service → Repository (sqlc) → Postgres`

```
cmd/api/main.go          # Bootstrap: config, DB pool, Echo setup, graceful shutdown
internal/
  config/config.go       # Env-driven config struct (godotenv)
  db/postgres.go         # pgxpool connection with tuned pool settings
  handler/               # Thin HTTP handlers — bind, validate, call service, respond
  service/               # Business logic — one file per domain
  repository/            # sqlc-generated DB access (do not edit by hand)
  middleware/
    auth.go              # JWT extraction → sets user_id/email/role on context
pkg/
  jwt/jwt.go             # GenerateTokenPair, ValidateToken
  response/response.go   # Typed JSON helpers (OK, Created, BadRequest, …)
  validator/validator.go # Echo validator wired to go-playground/validator
db/
  migrations/            # golang-migrate SQL files (up/down pairs)
  queries/               # Raw SQL read by sqlc to generate internal/repository
sqlc.yaml                # sqlc codegen config
```

## Key conventions

- **Adding a new domain** (e.g., `courses`): create `internal/handler/course.go`, `internal/service/course.go`, write SQL in `db/queries/course.sql`, run `make sqlc`, register routes in `cmd/api/main.go:registerRoutes`.
- **Auth middleware** — `middleware.JWT(cfg)` validates the Bearer token and sets `user_id`, `email`, `role` on the Echo context. Use `middleware.RequireRole("admin")` for RBAC.
- **DB access** — always go through the sqlc-generated `Queries` in `internal/repository`. Never write raw SQL outside `db/queries/`.
- **Response shape** — use `pkg/response` helpers; all responses are `{"success": bool, "data": …}` or `{"success": false, "message": "…", "errors": …}`.
- **JWT** — access token: 15 min (HS256), refresh token: 7 days. Secrets in env vars `JWT_ACCESS_SECRET` / `JWT_REFRESH_SECRET`.

## Stack

| Concern | Choice |
|---|---|
| HTTP | Echo v4 |
| Postgres driver | pgx v5 (pgxpool) |
| SQL codegen | sqlc |
| JWT | golang-jwt/jwt v5 |
| Validation | go-playground/validator v10 |
| Env config | godotenv |
| Migrations | golang-migrate |
