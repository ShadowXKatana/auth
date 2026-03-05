# init-go-gin

Go + Gin starter template ที่ใช้ Clean Architecture สำหรับเริ่มต้น project ใหม่

---

## 📁 Project Structure

```
init-go-gin/
├── cmd/api/                  # Application entrypoint
│   └── main.go               # สร้าง app แล้ว start server
├── internal/
│   ├── app/                  # Dependency wiring (ต่อทุก layer เข้าด้วยกัน)
│   │   └── app.go
│   ├── domain/               # Entity + Repository interface
│   │   └── user.go
│   ├── usecase/              # Business logic
│   │   └── user_usecase.go
│   ├── repository/           # Data layer implementation
│   │   └── user.go
│   └── handler/http/         # HTTP handlers + router
│       ├── user_handler.go
│       └── router.go
├── go.mod
└── go.sum
```

### แต่ละ layer ทำอะไร

| Layer | Directory | หน้าที่ |
|-------|-----------|---------|
| **Domain** | `internal/domain/` | กำหนด struct (entity) และ interface ของ repository — เป็น core ที่ไม่ depend อะไรเลย |
| **Usecase** | `internal/usecase/` | Business logic ทั้งหมด เช่น validate input, เรียก repo — depend แค่ domain |
| **Repository** | `internal/repository/` | Implement interface จาก domain เพื่อเก็บ/ดึงข้อมูล (ตอนนี้เป็น in-memory) |
| **Handler** | `internal/handler/http/` | รับ HTTP request, เรียก usecase, ตอบ response — มี router สำหรับ register routes |
| **App** | `internal/app/` | ต่อทุก layer เข้าด้วยกัน (wire dependencies) แล้ว return `*gin.Engine` พร้อมใช้ |
| **Entrypoint** | `cmd/api/` | จุดเริ่มต้น — เรียก `app.New()` แล้ว `engine.Run()` |

### Data flow

```
Request → Router → Handler → Usecase → Repository → Data store
                                ↑
                             Domain (entity + interface)
```

---

## 🔗 ความเกี่ยวข้องกับ my-storage-service

`init-go-gin` เป็น **template โครงสร้าง** ที่ถูก scale ขึ้นใน `my-storage-service` ซึ่งเป็นบริการจริงที่มี:

### Features ของ my-storage-service

| Feature | รายละเอียด |
|---------|-----------|
| **Auth** | Register, Login, Refresh token, Logout, Me (JWT + bcrypt) |
| **Storage** | CRUD สำหรับ storage — สร้าง, ดูรายการ, ดูตัวเดียว, ลบ |
| **Item** | CRUD สำหรับ item ภายใน storage — สร้าง, ดูรายการ, ลบ, แก้ tags |
| **Database** | PostgreSQL ผ่าน GORM + auto-migrate |
| **Security** | JWT middleware, CORS, HttpOnly cookie |

### Stack เปรียบเทียบ

| | init-go-gin | my-storage-service |
|---|---|---|
| **Framework** | Gin | Gin |
| **Database** | In-memory (map) | PostgreSQL + GORM |
| **Auth** | ❌ | ✅ JWT + bcrypt |
| **Entities** | User | User, Storage, Item |
| **Package** | `pkg` ❌ | `pkg` (JWT, Password) |

---

## 🚀 Setup & Run

### Prerequisites

- **Go** 1.22+

### 1. Install dependencies

```bash
cd BE/GO/init-go-gin
go mod tidy
```

### 2. Run

```bash
go run ./cmd/api
```

Server จะเริ่มที่ port `8080` — เปลี่ยนได้ด้วย environment variable:

```bash
APP_PORT=3000 go run ./cmd/api
```

---

## 📡 API Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/health` | Health check |
| `POST` | `/api/v1/users` | สร้าง user |
| `GET` | `/api/v1/users` | ดูรายการ user ทั้งหมด |

### Create User

```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "email": "john@example.com"}'
```

Response:
```json
{
  "id": "1",
  "name": "John Doe",
  "email": "john@example.com"
}
```

### List Users

```bash
curl http://localhost:8080/api/v1/users
```
