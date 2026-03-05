# Auth Structure (my-storage-service)

เอกสารนี้อธิบายโครงสร้างระบบ Auth ใน BE/GO/my-storage-service ตามโค้ดปัจจุบัน

ดู sequence diagram เพิ่มเติมได้ที่: `docs/auth-sequence.md`

## 1) ภาพรวมสถาปัตยกรรม

ระบบ Auth ใช้แนวทางแยกชั้นแบบ Clean-ish Architecture:

- Delivery Layer: รับ/ส่ง HTTP request-response
- Usecase Layer: กติกาธุรกิจของ register/login/refresh/me
- Domain Layer: model และสัญญา repository
- Repository Layer: ที่เก็บข้อมูลผู้ใช้ (memory)
- Repository Layer: ที่เก็บข้อมูลผู้ใช้ผ่าน PostgreSQL (GORM)
- Security Layer: password hashing และ JWT
- App Wiring Layer: ประกอบ dependency ทั้งหมดเข้าด้วยกัน

## 2) โครงสร้างไฟล์ที่เกี่ยวกับ Auth

- internal/app/app.go
- internal/delivery/http/handler/auth/handler.go
- internal/delivery/http/middleware/auth.go
- internal/delivery/http/middleware/cors.go
- internal/delivery/http/router/router.go
- internal/usecase/user/usecase.go
- internal/domain/user/entity.go
- internal/domain/user/repository.go
- internal/repository/postgres/db.go
- internal/repository/postgres/user/repository.go
- internal/security/password.go
- internal/security/jwt.go

## 3) App Wiring

ไฟล์: internal/app/app.go

หน้าที่หลัก:

- สร้าง Gin engine และติดตั้ง CORS middleware
- สร้าง Access Token Service และ Refresh Token Service (JWT คนละ secret/ttl)
- สร้าง User repository (PostgreSQL/GORM) และ Password service (bcrypt)
- สร้าง User usecase และ Auth handler
- ติดตั้ง JWT middleware เพื่อป้องกัน endpoint ที่ต้องล็อกอิน
- seed user เริ่มต้นจาก ENV (ถ้าตั้งค่าไว้)

## 4) HTTP Endpoints (REST Auth)

ไฟล์: internal/delivery/http/router/router.go

กลุ่ม endpoint:

- POST /api/v1/auth/register
- POST /api/v1/auth/login
- POST /api/v1/auth/refresh
- POST /api/v1/auth/logout
- GET /api/v1/auth/me (ต้องผ่าน JWT middleware)

หมายเหตุ:

- /graphql มีไว้สำหรับ use case อื่น ไม่ใช้เป็น auth endpoint
- /api/v1/storages ใช้ JWT middleware เช่นเดียวกับ /api/v1/auth/me

## 5) Handler Layer

ไฟล์: internal/delivery/http/handler/auth/handler.go

ความรับผิดชอบ:

- แปลง request body เป็น input ของ usecase
- map error ของ usecase เป็น HTTP status ที่เหมาะสม
- set/clear HttpOnly cookies: access_token และ refresh_token
- อ่าน refresh_token จาก cookie เมื่อ refresh
- ใช้ข้อมูล auth_user_email จาก context สำหรับ me endpoint

Cookie policy ที่ใช้:

- HttpOnly: true
- SameSite: Lax
- Secure: true เมื่อ APP_ENV=production
- Path: /

## 6) Middleware Layer

### 6.1 JWT Middleware

ไฟล์: internal/delivery/http/middleware/auth.go

หน้าที่:

- อ่าน access_token จาก cookie
- parse และ validate JWT
- ใส่ claims ลง context: auth_user_id, auth_user_email
- ถ้า token ไม่ถูกต้อง/ไม่มี token จะตอบ 401

### 6.2 CORS Middleware

ไฟล์: internal/delivery/http/middleware/cors.go

หน้าที่:

- อนุญาต origin ตาม APP_CORS_ORIGIN
- เปิดใช้งาน credentials เพื่อให้ browser ส่ง cookie ได้
- จัดการ preflight OPTIONS

## 7) Usecase Layer

ไฟล์: internal/usecase/user/usecase.go

Usecase methods:

- Register(ctx, input)
- Login(ctx, input)
- Refresh(ctx, refreshToken)
- Me(ctx, email)

ตรรกะสำคัญ:

- validate email/password เบื้องต้น
- ป้องกัน email ซ้ำตอนสมัคร
- hash password ตอนสมัคร และ compare ตอนล็อกอิน
- ออกทั้ง access token และ refresh token
- refresh โดย parse refresh token แล้วโหลด user จาก email ใน claims

## 8) Domain Layer

ไฟล์: internal/domain/user/entity.go, internal/domain/user/repository.go

- User entity: ID, Email, PasswordHash
- Repository contract:
  - Create(ctx, user)
  - GetByEmail(ctx, email)

## 9) Repository Layer (PostgreSQL + GORM)

ไฟล์: internal/repository/postgres/db.go, internal/repository/postgres/user/repository.go

- เชื่อมต่อ DB ด้วย `gorm.io/gorm` + `gorm.io/driver/postgres`
- map domain user กับ table `auth_users`
- normalize email เป็น lowercase+trim
- query ไม่พบข้อมูล map เป็น `ErrUserNotFound`

โครงสร้างตาราง auth ถูกจัดการโดย migration:

- DB/migrations/V2\_\_auth_users.sql
- DB/migrations/U2\_\_auth_users.sql

## 10) Security Layer

### 10.1 Password

ไฟล์: internal/security/password.go

- bcrypt hash
- bcrypt compare

### 10.2 JWT

ไฟล์: internal/security/jwt.go

- Issue(userID, email)
- Parse(token)
- ใช้ HMAC signing method
- claims มี user_id, email และ registered claims เช่น expiresAt/issuedAt

## 11) Auth Flow

### 11.1 Register

1. Client ส่ง email/password ไป register
2. Handler เรียก usecase.Register
3. Usecase ตรวจ input + กัน email ซ้ำ + hash password + create user
4. Usecase ออก access/refresh token
5. Handler set cookies และตอบข้อมูล auth

### 11.2 Login

1. Client ส่ง email/password ไป login
2. Usecase ดึง user จาก repository
3. Compare password hash
4. ออก access/refresh token
5. Handler set cookies

### 11.3 Refresh

1. Client เรียก refresh endpoint พร้อม refresh_token cookie
2. Usecase parse refresh token
3. Usecase โหลด user แล้วออก token ชุดใหม่
4. Handler set cookies ใหม่

### 11.4 Me / Protected API

1. Middleware อ่าน access_token cookie
2. Validate token แล้ว set auth_user_id/auth_user_email ใน context
3. Handler ใช้ข้อมูลจาก context เพื่อดึง profile (me) หรือให้ผ่านไปยัง protected endpoint

## 12) Environment Variables ที่เกี่ยวข้อง

- APP_PORT
- APP_ENV
- APP_CORS_ORIGIN
- JWT_SECRET
- JWT_EXPIRES_IN_MINUTES
- JWT_REFRESH_SECRET
- JWT_REFRESH_EXPIRES_IN_MINUTES
- AUTH_SEED_EMAIL
- AUTH_SEED_PASSWORD
- DB_DSN

## 13) หมายเหตุด้านการใช้งานจริง

- Production ควรตั้งค่า APP_ENV=production เพื่อเปิด Secure cookie
- ควรใช้ HTTPS เสมอใน production
- ควรพิจารณาเก็บ refresh token แบบมี revocation (DB/redis) หากต้องการ logout ทุกอุปกรณ์หรือ revoke เฉพาะ session
- ควรย้าย user repository จาก memory ไป persistent storage ก่อนใช้งานจริง
