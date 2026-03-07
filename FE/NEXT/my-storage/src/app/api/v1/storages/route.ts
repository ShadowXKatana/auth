import { NextRequest, NextResponse } from 'next/server'
import { cookies } from 'next/headers'
import { mockDb, MOCK_SESSION_COOKIE } from '@/mocks/db'

async function getAuthenticatedUserId(): Promise<string | null> {
  const cookieStore = await cookies()
  const userId = cookieStore.get(MOCK_SESSION_COOKIE)?.value
  return userId ?? null
}

export async function GET() {
  const userId = await getAuthenticatedUserId()
  if (!userId) {
    return NextResponse.json({ error: 'unauthorized' }, { status: 401 })
  }

  const storages = mockDb.getUserStorages(userId)
  return NextResponse.json(storages)
}

export async function POST(request: NextRequest) {
  const userId = await getAuthenticatedUserId()
  if (!userId) {
    return NextResponse.json({ error: 'unauthorized' }, { status: 401 })
  }

  const body = (await request.json().catch(() => ({}))) as { name?: string }
  const { name } = body

  if (!name?.trim()) {
    return NextResponse.json({ error: 'name is required' }, { status: 400 })
  }

  const storage = mockDb.createStorage(userId, name.trim())
  return NextResponse.json(storage, { status: 201 })
}
