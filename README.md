# Auth & Storage Platform Monorepo

Repository นี้เป็น Monorepo ที่รวมทั้งส่วน Backend และ Frontend ของโปรเจกต์ระบบ Authentication และระบบ File Storage 

## 🚀 โครงสร้างโปรเจกต์ (Projects Overview)

### Backend (Go)

- 🔹 [**init-go-gin**](./BE/GO/init-go-gin/) — Starter template สำหรับ Go + Gin ที่ใช้โครงสร้างแบบ Clean Architecture เป็นตัวตั้งต้นของ service อื่นๆ 
- 🔹 [**my-storage-service**](./BE/GO/my-storage-service/) — Service หลักของระบบ (Core backend) ทำหน้าที่จัดการ Authentication (JWT + bcrypt), การสร้าง Storage containers, และ Item management โดยใช้ PostgreSQL เป็นฐานข้อมูล

### Frontend (React & Next.js)

- 🔸 [**init (React + Vite)**](./FE/REACT/init/) — Starter template ขา Frontend ที่ใช้ React + Vite ถูก setup พร้อมใช้ด้วย TailwindCSS v4, Zustand, React Router, ESLint, Prettier, และ Husky
- 🔸 [**my-storage (Next.js)**](./FE/NEXT/my-storage/) — หน้าเว็บแอปพลิเคชันหลักของโปรเจกต์ สร้างด้วย Next.js (App Router) เชื่อมต่อกับ backend เพื่อแสดงผลหน้า Login, จัดการไฟล์ และ Storages ต่างๆ

---

## 🗄️ Database Architecture

ไฟล์ Database Diagram (DBML) ถูกเก็บรักษาไว้ที่ `./DB/dbdiagram.dbml` 

🔗 **ดูแผนผัง DB ออนไลน์:** [dbdiagram.io/d/69a11340a3f0aa31e141f961](https://dbdiagram.io/d/69a11340a3f0aa31e141f961)

---

## 💻 การเริ่มต้นใช้งาน (Getting Started)

สำหรับวิธีการ Setup และ Run แนะนำให้เข้าไปดูที่ `README.md` ของแต่ละโปรเจกต์:

1. ดูวิธีรัน Backend (Database + Go API) ได้ที่ 👉 [`my-storage-service/README.md`](./BE/GO/my-storage-service/README.md)
2. ดูวิธีรัน Frontend (Next.js) ได้ที่ 👉 [`FE/NEXT/my-storage/README.md`](./FE/NEXT/my-storage/README.md)

### 🔑 Environment Variables
โปรเจกต์นี้ใช้ไฟล์ `.env` ในการจัดการ Configuration:
1. Copy ไฟล์ตัวอย่าง: `cp .env.example .env`
2. แก้ไขค่าในไฟล์ `.env` ตามต้องการ (เช่น Database password, JWT secret)
3. Docker Compose จะดึงค่าจากไฟล์ `.env` มาใช้โดยอัตโนมัติคครับ

---

## TODO

- [ ] unit test
- [ ] e2e test
