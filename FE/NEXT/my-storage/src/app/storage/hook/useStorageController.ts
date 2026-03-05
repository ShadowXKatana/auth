'use client'

import { useState, useEffect, useCallback } from 'react'
import { api } from '@/lib/api'

type StorageItem = {
  id: string
  userId: string
  name: string
  createdAt: string
  updatedAt: string
}

export const useStorageController = () => {
  const [storages, setStorages] = useState<StorageItem[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')
  const [creating, setCreating] = useState(false)

  const fetchStorages = useCallback(async () => {
    try {
      const data = await api.get<StorageItem[]>('/api/v1/storages')
      setStorages(data ?? [])
    } catch {
      setError('Failed to load storages.')
    } finally {
      setLoading(false)
    }
  }, [])

  useEffect(() => {
    fetchStorages()
  }, [fetchStorages])

  const createStorage = async (name: string) => {
    if (!name.trim()) return
    setCreating(true)
    try {
      await api.post('/api/v1/storages', { name })
      await fetchStorages()
    } catch {
      setError('Failed to create storage.')
    } finally {
      setCreating(false)
    }
  }

  const deleteStorage = async (id: string) => {
    try {
      await api.delete(`/api/v1/storages/${id}`)
      setStorages((prev) => prev.filter((s) => s.id !== id))
    } catch {
      setError('Failed to delete storage.')
    }
  }

  return {
    storages,
    loading,
    error,
    creating,
    createStorage,
    deleteStorage,
  }
}

