import { NextResponse } from 'next/server'
import { cookies } from 'next/headers'
import { mockDb, MOCK_SESSION_COOKIE } from '@/mocks/db'

export async function POST() {
  const cookieStore = await cookies()
  const userId = cookieStore.get(MOCK_SESSION_COOKIE)?.value
  const user = userId ? mockDb.findUserById(userId) : undefined

  if (!user) {
    return NextResponse.json({ error: 'session expired' }, { status: 401 })
  }

  return NextResponse.json({
    accessToken: 'mock-access-token',
    refreshToken: 'mock-refresh-token',
    user: { id: user.id, email: user.email },
  })
}
