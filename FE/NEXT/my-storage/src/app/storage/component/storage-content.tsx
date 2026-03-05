'use client'

import { useState } from 'react'
import Link from 'next/link'
import { Card } from '@/components/common/card'
import { PageShell } from '@/components/common/page-shell'
import { Button } from '@/components/common/button'
import { TextInput } from '@/components/common/text-input'
import { useStorageController } from '@/app/storage/hook/useStorageController'
import { AuthGuard } from '@/components/common/auth-guard'

export const StorageContent = () => {
  const { storages, loading, error, creating, createStorage, deleteStorage } =
    useStorageController()
  const [newName, setNewName] = useState('')

  const handleCreate = async () => {
    if (!newName.trim()) return
    await createStorage(newName)
    setNewName('')
  }

  return (
    <AuthGuard>
      <PageShell title="Storage" subtitle="Manage your storages." showBackLink showLogout>
        <div className="space-y-4">
          <Card>
            <h2 className="mb-3 text-lg font-medium">Create Storage</h2>
            <div className="flex gap-2">
              <div className="flex-1">
                <TextInput
                  label="Name"
                  value={newName}
                  onChange={(e) => setNewName(e.target.value)}
                  placeholder="Enter storage name"
                />
              </div>
              <div className="flex items-end">
                <Button type="button" onClick={handleCreate} disabled={creating}>
                  {creating ? 'Creating...' : 'Create'}
                </Button>
              </div>
            </div>
          </Card>

          <Card>
            <h2 className="mb-3 text-lg font-medium">Your Storages</h2>
            {loading ? (
              <p className="text-sm text-foreground/60">Loading...</p>
            ) : error ? (
              <p className="text-sm text-red-500">{error}</p>
            ) : storages.length === 0 ? (
              <p className="text-sm text-foreground/60">No storages yet. Create one above.</p>
            ) : (
              <ul className="divide-y divide-black/10 text-sm dark:divide-white/10">
                {storages.map((storage) => (
                  <li key={storage.id} className="flex items-center justify-between py-3">
                    <Link
                      href={`/storage/${storage.id}`}
                      className="font-medium hover:underline"
                    >
                      {storage.name}
                    </Link>
                    <button
                      type="button"
                      onClick={() => deleteStorage(storage.id)}
                      className="text-xs text-red-500 hover:underline"
                    >
                      Delete
                    </button>
                  </li>
                ))}
              </ul>
            )}
          </Card>
        </div>
      </PageShell>
    </AuthGuard>
  )
}

