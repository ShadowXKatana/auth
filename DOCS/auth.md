การออกแบบระบบ Auth ด้วย Apollo (Go + Next.js)
หัวใจสำคัญยังคงเหมือนเดิมคือ HttpOnly Cookie แต่ต้องตั้งค่า Apollo ให้ส่ง Cookie ไปด้วย
3.1 ฝั่ง Backend (Go)
ต้องใช้ Library เช่น 99designs/gqlgen
สร้าง Schema สำหรับ Login, RefreshToken, User
Middleware ยังคงใช้ตรวจสอบ JWT จาก Cookie เหมือนเดิม
graphql
12345678910

# schema.graphql

type Mutation {
login(email: String!, password: String!): AuthPayload!
refreshToken: AuthPayload!
logout: Boolean!
}

type Query {
me: User
}
3.2 ฝั่ง Frontend (Next.js + Apollo)
ต้องตั้งค่า ApolloClient ให้ส่ง Cookie (credentials: 'include') และจัดการ Token Refresh ผ่าน onError Link
ติดตั้ง:
bash
1
npm install @apollo/client graphql @apollo/experimental-nextjs-app-support
ตั้งค่า Apollo Client (พร้อม Auth Logic):
typescript
12345678910111213141516171819202122232425262728293031323334353637
// lib/apolloClient.ts
import { ApolloClient, InMemoryCache, createHttpLink, from } from '@apollo/client';
import { onError } from '@apollo/client/link/error';
import { setContext } from '@apollo/client/link/context'; // ถ้าต้องใช้ Header เพิ่มเติม

// 1. HttpLink ตั้งค่าให้ส่ง Cookie
const httpLink = createHttpLink({
uri: process.env.NEXT_PUBLIC_GO_GRAPHQL_URL,
credentials: 'include', // สำคัญมาก! ไม่งั้น Cookie ไม่ส่งไป Go
});

ข้อควรระวังเรื่อง Token Refresh กับ Apollo:
การทำ Refresh Token แบบอัตโนมัติใน Apollo ซับซ้อนกว่า Axios Interceptor เพราะต้องจัดการกับ Observable Stream
ทางเลือกแนะนำ: ถ้า Token หมดอายุ ให้ Catch Error แล้วเด้งไปหน้า Login เลย (ง่ายสุด)
ทางเลือกขั้นสูง: ใช้ Library เช่น apollo-link-refresh ช่วยจัดการ Queue Request ระหว่างรอ Refresh Token
3.3 การใช้งานใน Next.js App Router (SSR)
การใช้ Apollo กับ Next.js App Router ต้องใช้ Wrapper พิเศษเพื่อให้ Cache Hydration ทำงานได้
typescript
12345678910111213
// app/ApolloWrapper.tsx
'use client';
import { NextAppDirEmulator } from '@apollo/experimental-nextjs-app-support';
import { ApolloProvider } from '@apollo/client';
import { client } from '@/lib/apolloClient';

export function ApolloWrapper({ children }: { children: React.ReactNode }) {
return (
<ApolloProvider client={client}>
{children}
