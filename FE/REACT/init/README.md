# react-init (React + Vite Starter)

Template เริ่มต้นสำหรับฝั่ง Frontend ที่ Setup แบบเบสิค ครบจบสำหรับการพร้อมพัฒนาด้วยเทคโนโลยีที่รวดเร็วและเป็นมาตรฐานล่าสุด

โปรเจกต์นี้ขับเคลื่อนด้วย **[React](https://react.dev)** + **[Vite](https://vite.dev)** แบบ TypeScript เต็มรูปแบบ 

## 📦 Stack ปัจจุบัน

- React Router v7 สำหรับจัดการ Navigation แบบมาตรฐาน
- Zustand แบบเบาบางสำหรับการแชร์ State ทั่วทั้งแอป (Global State Management)
- TailwindCSS v4 แบบไร้ config อันยุ่งยาก สำหรับงานฝั่ง Styling (Utility-First)
- Lucide React ชุดไอคอนสวยๆ ที่เข้าถึงง่าย
- TypeScript สำหรับความปลอดภัยของข้อมูล
- ESLint + Prettier + Husky สำหรับคุมคุณภาพโค้ดก่อนการ Commit ไปที่ Repository

---

## 🚀 การเริ่มต้นระบบ (Getting Started)

### 1. ติดตั้ง Dependencies
```bash
npm install
```

### 2. ตั้งค่า Environment Variables
ทำการคัดลอกไฟล์ `.env.example` ไปเป็น `.env.local` เพื่อใช้งานต่างๆ 
```bash
cp .env.example .env.local
```

### 3. รัน Development Server
```bash
npm run dev
```
แล้วเปิด [http://localhost:5173](http://localhost:5173) บน Browser เพื่อเริ่มพัฒนาได้ทันที

---

## 🛠 คำสั่ง (Scripts) ที่ใช้งานได้

| Command | รายละเอียด |
|---|---|
| `npm run dev` | รัน Vite development server (เร็วปานจรวด!) |
| `npm run build` | บิลด์โค้ดเตรียมเสิร์ฟเป็น Static Asset บน Server |
| `npm run preview` | เปิด Server ชั่วคราวลองเทสไฟล์ที่เพิ่ง Build ไปหมาดๆ |
| `npm run lint` | ตรวจจับข้อผิดพลาดและกลิ่นโค้ดแปลกๆ ผ่าน ESLint |
| `npm run typecheck` | ให้ TypeScript เช็คว่าเราเขียน Type ตรงตามโมเดลไหม |
| `npm run format` | สั่งให้ Prettier จัดเรียงหน้ากระดาษและวรรคตอนให้เรียบร้อย |
| `npm test` | รัน Unit tests ด้วย Vitest และ Testing Library |

---

## 📚 แหล่งอ้างอิงและศูนย์การเรียนรู้

- [Vite Documentation](https://vite.dev/guide/) - ทำความรู้จักพลังเบื้องหลัง
- [React Documentation](https://react.dev) - เจาะลึก Hooks เเละ Components ของ React
- [TypeScript](https://www.typescriptlang.org/docs/) - เรียนรู้จัก Type และ Interface
- [TailwindCSS v4 Docs](https://tailwindcss.com/docs) - สูตรโกง Utilities เพื่อสร้าง UI สวยๆ
