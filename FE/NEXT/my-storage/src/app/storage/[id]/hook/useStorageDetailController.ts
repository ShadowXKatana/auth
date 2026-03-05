'use client'

import { useState, useEffect, useCallback } from 'react'
import { useParams } from 'next/navigation'
import { api } from '@/lib/api'

type ItemData = {
    ID: string
    StorageID: string
    Name: string
    SizeMb: number
    Tags: string[]
    CreatedAt: string
    UpdatedAt: string
}

export const useStorageDetailController = () => {
    const params = useParams()
    const storageId = params.id as string

    const [items, setItems] = useState<ItemData[]>([])
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState('')
    const [creating, setCreating] = useState(false)

    const fetchItems = useCallback(async () => {
        try {
            const data = await api.get<ItemData[]>(`/api/v1/storages/${storageId}/items`)
            setItems(data ?? [])
        } catch {
            setError('Failed to load items.')
        } finally {
            setLoading(false)
        }
    }, [storageId])

    useEffect(() => {
        fetchItems()
    }, [fetchItems])

    const createItem = async (name: string, sizeMb: number, tags: string[]) => {
        if (!name.trim()) return
        setCreating(true)
        try {
            await api.post(`/api/v1/storages/${storageId}/items`, { name, sizeMb, tags })
            await fetchItems()
        } catch {
            setError('Failed to create item.')
        } finally {
            setCreating(false)
        }
    }

    const deleteItem = async (itemId: string) => {
        try {
            await api.delete(`/api/v1/storages/${storageId}/items/${itemId}`)
            setItems((prev) => prev.filter((i) => i.ID !== itemId))
        } catch {
            setError('Failed to delete item.')
        }
    }

    const updateItemTags = async (itemId: string, tags: string[]) => {
        try {
            const updated = await api.patch<ItemData>(
                `/api/v1/storages/${storageId}/items/${itemId}/tags`,
                { tags },
            )
            setItems((prev) => prev.map((i) => (i.ID === itemId ? updated : i)))
        } catch {
            setError('Failed to update tags.')
        }
    }

    return {
        storageId,
        items,
        loading,
        error,
        creating,
        createItem,
        deleteItem,
        updateItemTags,
    }
}
