# Chirpy

Chirpy is a small REST API written in Go for posting and reading short messages called chirps. It includes user registration and login, JWT authentication, refresh tokens, Chirpy Red membership upgrades, and a PostgreSQL-backed data store.

## Why Chirpy?

Chirpy demonstrates the building blocks of a practical backend service: HTTP routing, database access, password hashing, authentication, and RESTful CRUD operations. It is useful for learning Go web development or as a starting point for a microblogging API.

## Prerequisites

- Go 1.27 or newer
- PostgreSQL
- `goose` for database migrations (optional if you apply the SQL files another way)

## Install

Clone the repository and download the Go dependencies:

```bash
git clone <repository-url>
cd chirpy
go mod download
```

Create a PostgreSQL database:

```bash
createdb chirpy
```

Apply the migrations from `sql/schema`:

```bash
goose -dir sql/schema postgres "postgres://postgres:postgres@localhost:5432/chirpy?sslmode=disable" up
```

Adjust the connection string if your PostgreSQL username, password, host, or database name differs.

## Configuration

Create `.env` in the project root:

```dotenv
DB_URL="postgres://postgres:postgres@localhost:5432/chirpy?sslmode=disable"
PLATFORM="dev"
JWT_SECRET="replace-with-a-long-random-secret"
POLKA_KEY="replace-with-your-webhook-api-key"
```

- `DB_URL`: PostgreSQL connection string.
- `PLATFORM`: Set to `dev` to enable the development database reset endpoint.
- `JWT_SECRET`: Secret used to sign access tokens.
- `POLKA_KEY`: API key used by the Polka webhook endpoint.

## Run

```bash
go run .
```

The server listens on `http://localhost:8080`. Verify it with:

```bash
curl http://localhost:8080/api/healthz
```

To build and run a binary:

```bash
go build -o out
./out
```

The static app page is available at `http://localhost:8080/app/`, and metrics are available at `http://localhost:8080/admin/metrics`.

## Test

```bash
go test ./...
```
