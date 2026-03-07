import { NextResponse } from 'next/server'
import { cookies } from 'next/headers'
import { MOCK_SESSION_COOKIE } from '@/mocks/db'

export async function POST() {
  const cookieStore = await cookies()
  cookieStore.delete(MOCK_SESSION_COOKIE)
  return NextResponse.json({ success: true })
}
