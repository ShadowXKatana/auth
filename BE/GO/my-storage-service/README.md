# my-storage-service

Core Backend Service หลักของระบบที่พัฒนาด้วย **Go + Gin** พร้อมใช้ **Clean Architecture** (ต่อยอดมาจากโครงสร้าง `init-go-gin`) 

ระบบนี้ให้บริการ 2 ส่วนหลักคือ:
1. **Authentication:** จัดการผู้ใช้, Login/Register และระบบ JSON Web Token (JWT)
2. **Storage Management:** จัดการคลังเก็บข้อมูลส่วนตัวของผู้ใช้ (Storage Bucket) และรายการไฟล์ (Items)

---

## 📁 โครงสร้างโปรเจกต์ (Project Structure)

```
my-storage-service/
├── cmd/api/                  # Application entrypoint
├── internal/
│   ├── app/                  # Dependency wiring (ต่อทุก layer เข้าด้วยกัน)
│   ├── gorm/                 # การเชื่อมต่อ Database และ Auto Migration
│   ├── domain/               # Entity Models + Repository interfaces (User, Storage, Item)
│   ├── usecase/              # Business logic ของระบบ Auth และข้อมูล
│   ├── repository/           # PostgreSQL repositories ผ่าน GORM
│   └── handler/http/         # HTTP handlers + routes + JWT Middleware
├── pkg/                      # Utilities ของระบบ (JWT Token Service, Password Hashing)
└── DB/migrations/            # (ถ้ามี) สำหรับ manage database migration
```

---

## 🛠 Features หลัก

- **Authentication Module**
  - Register, Login 
  - Token-based Authentication (Access Token + Refresh Token via HttpOnly Cookies)
  - Profile (Me) & Logout
  - Hash password อย่างปลอดภัย (bcrypt)
- **Storage Module**
  - สร้าง, ดูรายการ, ดูรายละเอียด, ลบ Storage
- **Item Module**
  - เพิ่ม Items ลงใน Storage, จัดการ Items และ อัปเดต Tags ได้
- **Security & Database**
  - Database: PostgreSQL (ผ่าน `gorm`) พร้อม auto-migration ตอนเริ่มแอป
  - ป้องกัน Route ด้วย JWT Auth Middleware

---

## 🚀 การเริ่มต้นระบบ (Setup & Run)

### Prerequisites
- Docker & Docker Compose (สำหรับรันฐานข้อมูล PostgreSQL)
- Go 1.22+

### 1. เริ่มต้น Database
ใช้ Docker compose ยกฐานข้อมูล PostgreSQL ขึ้นมาทำงานเบื้องหลัง:
```bash
docker compose up -d auth-db
```

### 2. ดาวน์โหลด Dependencies
```bash
go mod tidy
```

### 3. รัน Service
```bash
go run ./cmd/api
```
*(แอปจะทำงานที่พอร์ต `8080` เป็นค่าปริยาย)*

---

## ⚙️ Environment Variables (ตั้งค่า)

คุณสามารถตั้งค่า Environment หรือ config ระบบ ได้ตามตัวแปรนี้:

- `APP_PORT` (default: `8080`)
- `APP_ENV` (เช่น `production` จะบังคับใช้ Secure Cookie)
- `DB_DSN` (default: `host=localhost user=auth password=auth dbname=auth port=5432...`)
- `APP_CORS_ORIGIN` (default: `http://localhost:3000`)
- `JWT_SECRET` (default: `dev-secret-change-me`)
- `JWT_EXPIRES_IN_MINUTES` (default: `60`)
- `JWT_REFRESH_SECRET` (default: `dev-refresh-secret-change-me`)
- `JWT_REFRESH_EXPIRES_IN_MINUTES` (default: `10080` = 7 วัน)

**การ Seed ข้อมูลเบื้องต้น (ทิ้งไว้เป็น Optional):**
- `AUTH_SEED_EMAIL` เเละ `AUTH_SEED_PASSWORD` (ถ้าระบุระบบจะสร้าง user ให้อัตโนมัติในตอนเปิดแอป)

---

## 📡 API Endpoints

| Method | Path | Description | Access | 
|--------|------|-------------|---|
| `GET` | `/health` | Health check (System) | Public |
| `POST` | `/api/v1/auth/register` | สมัครสมาชิก | Public |
| `POST` | `/api/v1/auth/login` | เข้าสู่ระบบ | Public |
| `POST` | `/api/v1/auth/refresh` | ต่ออายุ Token Session | Public (ใช้ Cookie) |
| `POST` | `/api/v1/auth/logout` | ออกจากระบบ | Public (ลบ Cookie) |
| `GET` | `/api/v1/auth/me` | ดูข้อมูลโปรไฟล์ปัจจุบัน | **Protected** |
| `POST` | `/api/v1/storages` | สร้าง Storage ใหม่ | **Protected** |
| `GET` | `/api/v1/storages` | ดู Storage ทั้งหมดของตัวเอง | **Protected** |
| `GET` | `/api/v1/storages/:id` | ดูรายละเอียด Storage | **Protected** |
| `DELETE`| `/api/v1/storages/:id` | ลบ Storage | **Protected** |
| `POST` | `/api/v1/storages/:id/items` | เพิ่ม Item ใน Storage | **Protected** |
| `GET` | `/api/v1/storages/:id/items` | ลิสต์ Items ใน Storage | **Protected** |
| `PATCH` | `/api/v1/storages/:id/items/:itemId/tags`| แก้ไข Tags ของ Item | **Protected** |
| `DELETE`| `/api/v1/storages/:id/items/:itemId`| ลบ Item ทิ้ง | **Protected** |

> บริการนี้ติดต่อกับ Client ผ่าน **HttpOnly cookie** อัตโนมัติ (`access_token` เเละ `refresh_token`) ช่วยลดความเสี่ยงจากการถูกขโมย Token ฝั่ง Frontend
