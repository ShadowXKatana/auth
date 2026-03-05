'use client'

import { useState } from 'react'
import { Card } from '@/components/common/card'
import { PageShell } from '@/components/common/page-shell'
import { Button } from '@/components/common/button'
import { TextInput } from '@/components/common/text-input'
import { useStorageDetailController } from '@/app/storage/[id]/hook/useStorageDetailController'
import { AuthGuard } from '@/components/common/auth-guard'

export const StorageDetailContent = () => {
    const {
        storageId,
        items,
        loading,
        error,
        creating,
        createItem,
        deleteItem,
        updateItemTags,
    } = useStorageDetailController()

    const [newItemName, setNewItemName] = useState('')
    const [newItemSize, setNewItemSize] = useState('')
    const [newItemTags, setNewItemTags] = useState('')
    const [editingTagsId, setEditingTagsId] = useState<string | null>(null)
    const [editingTagsValue, setEditingTagsValue] = useState('')

    const handleCreateItem = async () => {
        if (!newItemName.trim()) return
        const sizeMb = parseFloat(newItemSize) || 0
        const tags = newItemTags
            .split(',')
            .map((t) => t.trim())
            .filter(Boolean)
        await createItem(newItemName, sizeMb, tags)
        setNewItemName('')
        setNewItemSize('')
        setNewItemTags('')
    }

    const handleSaveTags = async (itemId: string) => {
        const tags = editingTagsValue
            .split(',')
            .map((t) => t.trim())
            .filter(Boolean)
        await updateItemTags(itemId, tags)
        setEditingTagsId(null)
        setEditingTagsValue('')
    }

    return (
        <AuthGuard>
            <PageShell
                title={`Storage`}
                subtitle={`ID: ${storageId}`}
                showBackLink
                showLogout
            >
                <div className="space-y-4">
                    <Card>
                        <h2 className="mb-3 text-lg font-medium">Add Item</h2>
                        <div className="space-y-3">
                            <div className="flex gap-2">
                                <div className="flex-1">
                                    <TextInput
                                        label="Name"
                                        value={newItemName}
                                        onChange={(e) => setNewItemName(e.target.value)}
                                        placeholder="Item name"
                                    />
                                </div>
                                <div className="w-28">
                                    <TextInput
                                        label="Size (MB)"
                                        type="number"
                                        value={newItemSize}
                                        onChange={(e) => setNewItemSize(e.target.value)}
                                        placeholder="0"
                                    />
                                </div>
                            </div>
                            <TextInput
                                label="Tags (comma separated)"
                                value={newItemTags}
                                onChange={(e) => setNewItemTags(e.target.value)}
                                placeholder="photo, vacation, 2026"
                            />
                            <Button type="button" onClick={handleCreateItem} disabled={creating}>
                                {creating ? 'Adding...' : 'Add Item'}
                            </Button>
                        </div>
                    </Card>

                    <Card>
                        <h2 className="mb-3 text-lg font-medium">Items</h2>
                        {loading ? (
                            <p className="text-sm text-foreground/60">Loading...</p>
                        ) : error ? (
                            <p className="text-sm text-red-500">{error}</p>
                        ) : items.length === 0 ? (
                            <p className="text-sm text-foreground/60">No items yet. Add one above.</p>
                        ) : (
                            <ul className="divide-y divide-black/10 dark:divide-white/10">
                                {items.map((item) => (
                                    <li key={item.ID} className="py-3">
                                        <div className="flex items-start justify-between">
                                            <div className="min-w-0 flex-1">
                                                <p className="text-sm font-medium">{item.Name}</p>
                                                <p className="mt-0.5 text-xs text-foreground/60">
                                                    {item.SizeMb} MB
                                                </p>
                                                <div className="mt-2 flex flex-wrap gap-1">
                                                    {item.Tags.map((tag) => (
                                                        <span
                                                            key={tag}
                                                            className="rounded-full bg-black/5 px-2 py-0.5 text-xs dark:bg-white/10"
                                                        >
                                                            {tag}
                                                        </span>
                                                    ))}
                                                    {item.Tags.length === 0 && (
                                                        <span className="text-xs text-foreground/40">No tags</span>
                                                    )}
                                                </div>
                                            </div>
                                            <div className="ml-3 flex gap-2">
                                                <button
                                                    type="button"
                                                    onClick={() => {
                                                        setEditingTagsId(item.ID)
                                                        setEditingTagsValue(item.Tags.join(', '))
                                                    }}
                                                    className="text-xs text-foreground/60 hover:underline"
                                                >
                                                    Edit Tags
                                                </button>
                                                <button
                                                    type="button"
                                                    onClick={() => deleteItem(item.ID)}
                                                    className="text-xs text-red-500 hover:underline"
                                                >
                                                    Delete
                                                </button>
                                            </div>
                                        </div>

                                        {editingTagsId === item.ID && (
                                            <div className="mt-3 flex gap-2">
                                                <div className="flex-1">
                                                    <TextInput
                                                        label="Tags"
                                                        value={editingTagsValue}
                                                        onChange={(e) => setEditingTagsValue(e.target.value)}
                                                        placeholder="tag1, tag2, tag3"
                                                    />
                                                </div>
                                                <div className="flex items-end gap-1">
                                                    <Button
                                                        type="button"
                                                        onClick={() => handleSaveTags(item.ID)}
                                                    >
                                                        Save
                                                    </Button>
                                                    <Button
                                                        type="button"
                                                        variant="secondary"
                                                        onClick={() => setEditingTagsId(null)}
                                                    >
                                                        Cancel
                                                    </Button>
                                                </div>
                                            </div>
                                        )}
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
