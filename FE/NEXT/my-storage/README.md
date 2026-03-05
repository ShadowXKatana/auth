# my-storage (Next.js Application)

โปรเจกต์หน้าเว็บแอปพลิเคชัน (Frontend) หลักของระบบ Auth & File Storage ที่ถูกสร้างขึ้นด้วย **[Next.js](https://nextjs.org)** (App Router)

แอปพลิเคชันตัวนี้ทำงานประสานกับ Backend `my-storage-service` เพื่อใช้จัดการบัญชี (Register, Login) และจัดการคลังไฟล์ (Storage Buckets & Items) ผ่าน HTTP HttpOnly Cookies

---

## 🚀 การเริ่มต้นระบบ (Getting Started)

ระบบนี้ใช้ `pnpm` เป็น Package Manager หลัก

### 1. ติดตั้ง Dependencies
```bash
pnpm install
```

### 2. ตั้งค่า Environment Variables
ทำการคัดลอกไฟล์ `.env.example` ไปเป็น `.env.local` เพื่อใช้งาน:
```bash
cp .env.example .env.local
```

### 3. รัน Development Server
```bash
pnpm dev
```
แล้วเปิด [http://localhost:3000](http://localhost:3000) บน Browser เพื่อเริ่มประเมินผล! 🎉

หน้าเว็บจะอัปเดตอัตโนมัติ (Fast Refresh) ทุกครั้งที่คุณแก้ไขไฟล์

---

## 🛠 คำสั่ง (Scripts) ที่ใช้งานได้

| Command | รายละเอียด |
|---|---|
| `pnpm dev` | รัน Next.js development server |
| `pnpm build` | Build โปรเจกต์เตรียมสำหรับ Production |
| `pnpm start` | รัน Web Server หลังจากที่ Build แล้ว |
| `pnpm lint` | ตรวจสอบ Code Style ด้วย ESLint |
| `pnpm format` | จัด Format ไฟล์ผ่าน Prettier |
| `pnpm typecheck` | ตรวจสอบ TypeScript Types ภายในโปรเจกต์ |
| `pnpm test` | รัน Unit Tests (ถ้ามี) |

---

## 📚 แหล่งเรียนรู้เพิ่มเติมสำหรับ Next.js

ลองทำความเข้าใจเพิ่มเติมเกี่ยวกับ Next.js ได้จากพื้นที่เหล่านี้:

- [Next.js Documentation](https://nextjs.org/docs) - เรียนรู้ฟีเจอร์ใหม่ๆ 
- [Learn Next.js](https://nextjs.org/learn) - Tutorial ปฏิบัติจริง
- [The Next.js GitHub repo](https://github.com/vercel/next.js)
