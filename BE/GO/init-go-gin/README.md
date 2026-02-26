# init-go-gin

Go + Gin starter with a lightweight clean architecture structure.

## Structure

- `cmd/server`: application entrypoint
- `internal/app`: dependency wiring
- `internal/domain`: entities and interfaces
- `internal/usecase`: business logic
- `internal/repository`: data implementations
- `internal/delivery/http`: handlers and routers

## Run

```bash
go mod tidy
go run ./cmd/server
```

Default port is `8080`. You can override with `APP_PORT`.

## Endpoints

- `GET /health`
- `POST /api/v1/users`
- `GET /api/v1/users`

### Create user payload

```json
{
  "name": "John Doe",
  "email": "john@example.com"
}
```
