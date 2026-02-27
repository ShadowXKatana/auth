# my-storage-service

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
docker compose up -d auth-db
go mod tidy
go run ./cmd/server
```

Default port is `8080`. You can override with `APP_PORT`.

## Auth Config

- `JWT_SECRET` (default: `dev-secret-change-me`)
- `JWT_EXPIRES_IN_MINUTES` (default: `60`)
- `JWT_REFRESH_SECRET` (default: `dev-refresh-secret-change-me`)
- `JWT_REFRESH_EXPIRES_IN_MINUTES` (default: `10080`)
- `APP_CORS_ORIGIN` (default: `http://localhost:3000`)
- `DB_DSN` (default: `host=localhost user=auth password=auth dbname=auth port=5432 sslmode=disable TimeZone=UTC`)

Optional seed user for login:

- `AUTH_SEED_EMAIL`
- `AUTH_SEED_PASSWORD`

## Endpoints

- `GET /health`
- `GET /graphql`
- `POST /graphql`
- `POST /api/v1/auth/register`
- `POST /api/v1/auth/login`
- `POST /api/v1/auth/refresh`
- `POST /api/v1/auth/logout`
- `GET /api/v1/auth/me`
- `POST /api/v1/storages` (requires `access_token` cookie)
- `GET /api/v1/storages` (requires `access_token` cookie)

Auth ใช้ REST API เท่านั้น ส่วน GraphQL (`/graphql`) แยกไว้สำหรับ use case อื่น.
Auth persistence ใช้ PostgreSQL ผ่าน GORM (`auth_users` table)

## Database Schema (Redesigned)

Migration อยู่ที่ `DB/migrations` โดยโครงสร้างหลักแบ่งเป็น:

- `auth_users`: ผู้ใช้งานระบบ auth
- `auth_user_profiles`: โปรไฟล์ผู้ใช้แบบ 1:1 กับ `auth_users`
- `auth_refresh_tokens`: เก็บ refresh token แบบ hash เพื่อรองรับ revocation/session management
- `storages`: ข้อมูล storage ของผู้ใช้ (1 ผู้ใช้มีหลาย storage ได้)
- `storage_items`: รายการ item ภายใน storage
- `storage_tags`: master tag
- `storage_item_tags`: ตารางเชื่อม many-to-many ระหว่าง item และ tag

หมายเหตุ:

- ใช้ `UUID` เป็น primary key ทุกตารางหลัก
- ใช้ `TIMESTAMPTZ` สำหรับ field เวลาเพื่อให้ timezone-safe
- ใช้ `ON DELETE CASCADE` ในความสัมพันธ์หลักเพื่อลด orphan rows

## Register / Login Payload

```json
{
  "email": "user@example.com",
  "password": "secret123"
}
```

`register` และ `login` จะคืนค่า `accessToken`, `refreshToken`, `user` และ set `HttpOnly` cookies.

## Refresh Token

`POST /api/v1/auth/refresh` จะอ่าน `refresh_token` cookie และออก token ใหม่

## Me Endpoint

`GET /api/v1/auth/me` ต้องมี `access_token` cookie ที่ยัง valid.

## Create Storage Payload

```json
{
  "name": "Main Bucket",
  "path": "/data/main"
}
```

ใช้ cookie ที่ได้จาก REST auth เรียก storage endpoints ต่อได้ทันที.
