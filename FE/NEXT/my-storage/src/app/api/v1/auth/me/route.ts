import { NextResponse } from 'next/server'
import { cookies } from 'next/headers'
import { mockDb, MOCK_SESSION_COOKIE } from '@/mocks/db'

export async function GET() {
  const cookieStore = await cookies()
  const userId = cookieStore.get(MOCK_SESSION_COOKIE)?.value
  const user = userId ? mockDb.findUserById(userId) : undefined

  if (!user) {
    return NextResponse.json({ error: 'unauthorized' }, { status: 401 })
  }

  return NextResponse.json({ id: user.id, email: user.email })
}
