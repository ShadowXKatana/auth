import { Card } from '@/components/common/card'
import { PageShell } from '@/components/common/page-shell'
import { useStorageController } from '@/app/storage/hook/useStorageController'

export const StorageContent = () => {
  const { items } = useStorageController()

  return (
    <PageShell title="Storage" subtitle="List of your stored items.">
      <Card>
        <h2 className="mb-3 text-lg font-medium">Items</h2>
        <ul className="divide-y divide-black/10 text-sm dark:divide-white/10">
          {items.map((item) => (
            <li key={item.id} className="flex items-center justify-between py-3">
              <div>
                <p className="font-medium">{item.name}</p>
                <p className="text-foreground/70">Updated: {item.updatedAt}</p>
              </div>
              <span>{item.sizeMb} MB</span>
            </li>
          ))}
        </ul>
      </Card>
    </PageShell>
  )
}
