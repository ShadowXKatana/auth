import Link from 'next/link'
import { Card } from '@/components/common/card'
import { PageShell } from '@/components/common/page-shell'
import { pages } from '@/app/constant'

export default function Home() {
  return (
    <PageShell title="My Storage" subtitle="Select a page to preview.">
      <div className="space-y-3">
        {pages.map((page) => (
          <Card key={page.href}>
            <Link href={page.href} className="text-sm font-medium hover:underline">
              {page.name}
            </Link>
          </Card>
        ))}
      </div>
    </PageShell>
  )
}
