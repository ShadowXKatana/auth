type StorageItem = {
  id: string
  name: string
  sizeMb: number
  updatedAt: string
}

export const useStorageController = () => {
  const items: StorageItem[] = [
    { id: '1', name: 'photo-2026-01.png', sizeMb: 2.4, updatedAt: '2026-02-10' },
    { id: '2', name: 'invoice-001.pdf', sizeMb: 0.8, updatedAt: '2026-02-17' },
    { id: '3', name: 'presentation.pptx', sizeMb: 12.1, updatedAt: '2026-02-21' },
  ]

  return {
    items,
  }
}
