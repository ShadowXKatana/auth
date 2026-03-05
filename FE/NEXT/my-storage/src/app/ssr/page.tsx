import { headers } from 'next/headers'
import { Card } from '@/components/common/card'
import { PageShell } from '@/components/common/page-shell'

type Todo = {
  id: number
  title: string
  completed: boolean
}

type SsrPayload = {
  generatedAt: string
  userAgent: string
  todo: Todo | null
}

export const dynamic = 'force-dynamic'

async function getSsrPayload(): Promise<SsrPayload> {
  const headerStore = await headers()
  const userAgent = headerStore.get('user-agent') ?? 'Unknown'

  let todo: Todo | null = null

  try {
    const response = await fetch('https://jsonplaceholder.typicode.com/todos/1', {
      cache: 'no-store',
    })

    if (response.ok) {
      todo = (await response.json()) as Todo
    }
  } catch {
    todo = null
  }

  return {
    generatedAt: new Date().toISOString(),
    userAgent,
    todo,
  }
}

export default async function SsrPage() {
  const data = await getSsrPayload()

  return (
    <PageShell title="SSR Demo" subtitle="This page is rendered on the server for every request.">
      <div className="space-y-4">
        <Card>
          <h2 className="mb-2 text-lg font-medium">Request Info</h2>
          <p className="text-sm">Rendered at: {data.generatedAt}</p>
          <p className="mt-1 text-sm break-all">User-Agent: {data.userAgent}</p>
        </Card>

        <Card>
          <h2 className="mb-2 text-lg font-medium">Server Fetch Result</h2>
          {data.todo ? (
            <div className="text-sm">
              <p>Todo #{data.todo.id}</p>
              <p className="mt-1">Title: {data.todo.title}</p>
              <p className="mt-1">Completed: {data.todo.completed ? 'Yes' : 'No'}</p>
            </div>
          ) : (
            <p className="text-sm">Unable to fetch remote data in current environment.</p>
          )}
        </Card>
      </div>
    </PageShell>
  )
}
