import { NextRequest, NextResponse } from 'next/server'
import { cookies } from 'next/headers'
import { mockDb, MOCK_SESSION_COOKIE } from '@/mocks/db'

export async function POST(request: NextRequest) {
  const body = (await request.json().catch(() => ({}))) as {
    email?: string
    password?: string
  }
  const { email, password } = body

  if (!email || !password) {
    return NextResponse.json({ error: 'email and password are required' }, { status: 400 })
  }

  let user = mockDb.findUserByEmail(email)
  if (!user) {
    // Auto-register new users so any email/password combo works in mock mode.
    user = mockDb.createUser(email, password)
  } else if (user.password !== password) {
    return NextResponse.json({ error: 'invalid credentials' }, { status: 401 })
  }

  const cookieStore = await cookies()
  cookieStore.set(MOCK_SESSION_COOKIE, user.id, {
    httpOnly: true,
    path: '/',
    maxAge: 60 * 60 * 24, // 24 hours
    sameSite: 'lax',
  })

  return NextResponse.json({
    accessToken: 'mock-access-token',
    refreshToken: 'mock-refresh-token',
    user: { id: user.id, email: user.email },
  })
}
