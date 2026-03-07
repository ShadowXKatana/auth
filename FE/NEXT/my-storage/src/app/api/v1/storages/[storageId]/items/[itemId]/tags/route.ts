import { NextRequest, NextResponse } from 'next/server'
import { cookies } from 'next/headers'
import { mockDb, MOCK_SESSION_COOKIE } from '@/mocks/db'

async function getAuthenticatedUserId(): Promise<string | null> {
  const cookieStore = await cookies()
  const userId = cookieStore.get(MOCK_SESSION_COOKIE)?.value
  return userId ?? null
}

export async function PATCH(
  request: NextRequest,
  { params }: { params: Promise<{ storageId: string; itemId: string }> },
) {
  const userId = await getAuthenticatedUserId()
  if (!userId) {
    return NextResponse.json({ error: 'unauthorized' }, { status: 401 })
  }

  const { storageId, itemId } = await params
  const ownsStorage = mockDb.getUserStorages(userId).some((s) => s.id === storageId)
  if (!ownsStorage) {
    return NextResponse.json({ error: 'not found' }, { status: 404 })
  }

  const body = (await request.json().catch(() => ({}))) as { tags?: string[] }
  const tags = Array.isArray(body.tags) ? body.tags : []

  const updated = mockDb.updateItemTags(itemId, tags)
  if (!updated) {
    return NextResponse.json({ error: 'not found' }, { status: 404 })
  }

  return NextResponse.json(updated)
}
