import { NextRequest, NextResponse } from 'next/server'
import { cookies } from 'next/headers'
import { mockDb, MOCK_SESSION_COOKIE } from '@/mocks/db'

async function getAuthenticatedUserId(): Promise<string | null> {
  const cookieStore = await cookies()
  const userId = cookieStore.get(MOCK_SESSION_COOKIE)?.value
  return userId ?? null
}

export async function GET(
  _request: NextRequest,
  { params }: { params: Promise<{ storageId: string }> },
) {
  const userId = await getAuthenticatedUserId()
  if (!userId) {
    return NextResponse.json({ error: 'unauthorized' }, { status: 401 })
  }

  const { storageId } = await params
  const ownsStorage = mockDb.getUserStorages(userId).some((s) => s.id === storageId)
  if (!ownsStorage) {
    return NextResponse.json({ error: 'not found' }, { status: 404 })
  }

  const items = mockDb.getStorageItems(storageId)
  return NextResponse.json(items)
}

export async function POST(
  request: NextRequest,
  { params }: { params: Promise<{ storageId: string }> },
) {
  const userId = await getAuthenticatedUserId()
  if (!userId) {
    return NextResponse.json({ error: 'unauthorized' }, { status: 401 })
  }

  const { storageId } = await params
  const ownsStorage = mockDb.getUserStorages(userId).some((s) => s.id === storageId)
  if (!ownsStorage) {
    return NextResponse.json({ error: 'not found' }, { status: 404 })
  }

  const body = (await request.json().catch(() => ({}))) as {
    name?: string
    sizeMb?: number
    tags?: string[]
  }
  const { name, sizeMb = 0, tags = [] } = body

  if (!name?.trim()) {
    return NextResponse.json({ error: 'name is required' }, { status: 400 })
  }

  const item = mockDb.createItem(storageId, name.trim(), sizeMb, tags)
  return NextResponse.json(item, { status: 201 })
}
