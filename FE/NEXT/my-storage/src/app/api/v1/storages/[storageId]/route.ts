import { NextRequest, NextResponse } from 'next/server'
import { cookies } from 'next/headers'
import { mockDb, MOCK_SESSION_COOKIE } from '@/mocks/db'

async function getAuthenticatedUserId(): Promise<string | null> {
  const cookieStore = await cookies()
  const userId = cookieStore.get(MOCK_SESSION_COOKIE)?.value
  return userId ?? null
}

export async function DELETE(
  _request: NextRequest,
  { params }: { params: Promise<{ storageId: string }> },
) {
  const userId = await getAuthenticatedUserId()
  if (!userId) {
    return NextResponse.json({ error: 'unauthorized' }, { status: 401 })
  }

  const { storageId } = await params
  const storages = mockDb.getUserStorages(userId)
  const exists = storages.some((s) => s.id === storageId)

  if (!exists) {
    return NextResponse.json({ error: 'not found' }, { status: 404 })
  }

  mockDb.deleteStorage(storageId)
  return new NextResponse(null, { status: 204 })
}
