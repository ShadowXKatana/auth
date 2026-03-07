import { randomUUID } from 'crypto'

export type MockUser = {
  id: string
  email: string
  password: string
}

export type MockStorage = {
  id: string
  userId: string
  name: string
  createdAt: string
  updatedAt: string
}

export type MockItem = {
  ID: string
  StorageID: string
  Name: string
  SizeMb: number
  Tags: string[]
  CreatedAt: string
  UpdatedAt: string
}

export const MOCK_SESSION_COOKIE = 'mock_uid'

const isoNow = () => new Date().toISOString()

const seedUsers: MockUser[] = [
  { id: 'user-mock-1', email: 'demo@example.com', password: 'demo1234' },
]

const seedStorages: MockStorage[] = [
  {
    id: 'storage-mock-1',
    userId: 'user-mock-1',
    name: 'Photos',
    createdAt: '2025-01-01T00:00:00Z',
    updatedAt: '2025-01-01T00:00:00Z',
  },
  {
    id: 'storage-mock-2',
    userId: 'user-mock-1',
    name: 'Documents',
    createdAt: '2025-01-02T00:00:00Z',
    updatedAt: '2025-01-02T00:00:00Z',
  },
  {
    id: 'storage-mock-3',
    userId: 'user-mock-1',
    name: 'Music',
    createdAt: '2025-01-03T00:00:00Z',
    updatedAt: '2025-01-03T00:00:00Z',
  },
]

const seedItems: MockItem[] = [
  {
    ID: 'item-mock-1',
    StorageID: 'storage-mock-1',
    Name: 'vacation.jpg',
    SizeMb: 2.5,
    Tags: ['vacation', '2025'],
    CreatedAt: '2025-01-01T00:00:00Z',
    UpdatedAt: '2025-01-01T00:00:00Z',
  },
  {
    ID: 'item-mock-2',
    StorageID: 'storage-mock-1',
    Name: 'family.jpg',
    SizeMb: 1.8,
    Tags: ['family', 'photo'],
    CreatedAt: '2025-01-01T00:00:00Z',
    UpdatedAt: '2025-01-01T00:00:00Z',
  },
  {
    ID: 'item-mock-3',
    StorageID: 'storage-mock-2',
    Name: 'report.pdf',
    SizeMb: 0.5,
    Tags: ['work', '2025'],
    CreatedAt: '2025-01-02T00:00:00Z',
    UpdatedAt: '2025-01-02T00:00:00Z',
  },
  {
    ID: 'item-mock-4',
    StorageID: 'storage-mock-2',
    Name: 'resume.docx',
    SizeMb: 0.2,
    Tags: ['personal', 'work'],
    CreatedAt: '2025-01-02T00:00:00Z',
    UpdatedAt: '2025-01-02T00:00:00Z',
  },
  {
    ID: 'item-mock-5',
    StorageID: 'storage-mock-3',
    Name: 'favorite-song.mp3',
    SizeMb: 5.1,
    Tags: ['music', 'favorite'],
    CreatedAt: '2025-01-03T00:00:00Z',
    UpdatedAt: '2025-01-03T00:00:00Z',
  },
  {
    ID: 'item-mock-6',
    StorageID: 'storage-mock-3',
    Name: 'podcast-ep1.mp3',
    SizeMb: 48.3,
    Tags: ['podcast', 'tech'],
    CreatedAt: '2025-01-03T00:00:00Z',
    UpdatedAt: '2025-01-03T00:00:00Z',
  },
]

class MockDb {
  users: MockUser[] = [...seedUsers]
  storages: MockStorage[] = [...seedStorages]
  items: MockItem[] = [...seedItems]

  findUserByEmail(email: string): MockUser | undefined {
    return this.users.find((u) => u.email === email)
  }

  findUserById(id: string): MockUser | undefined {
    return this.users.find((u) => u.id === id)
  }

  createUser(email: string, password: string): MockUser {
    const user: MockUser = { id: randomUUID(), email, password }
    this.users.push(user)
    return user
  }

  getUserStorages(userId: string): MockStorage[] {
    return this.storages.filter((s) => s.userId === userId)
  }

  createStorage(userId: string, name: string): MockStorage {
    const storage: MockStorage = {
      id: randomUUID(),
      userId,
      name,
      createdAt: isoNow(),
      updatedAt: isoNow(),
    }
    this.storages.push(storage)
    return storage
  }

  deleteStorage(id: string): void {
    this.storages = this.storages.filter((s) => s.id !== id)
    this.items = this.items.filter((i) => i.StorageID !== id)
  }

  getStorageItems(storageId: string): MockItem[] {
    return this.items.filter((i) => i.StorageID === storageId)
  }

  createItem(storageId: string, name: string, sizeMb: number, tags: string[]): MockItem {
    const item: MockItem = {
      ID: randomUUID(),
      StorageID: storageId,
      Name: name,
      SizeMb: sizeMb,
      Tags: tags,
      CreatedAt: isoNow(),
      UpdatedAt: isoNow(),
    }
    this.items.push(item)
    return item
  }

  deleteItem(id: string): void {
    this.items = this.items.filter((i) => i.ID !== id)
  }

  updateItemTags(id: string, tags: string[]): MockItem | undefined {
    const item = this.items.find((i) => i.ID === id)
    if (!item) return undefined
    item.Tags = tags
    item.UpdatedAt = isoNow()
    return item
  }
}

// Use a global singleton so the in-memory state survives HMR reloads in development.
const globalForMockDb = globalThis as unknown as { __mockDb?: MockDb }

export const mockDb = globalForMockDb.__mockDb ?? (globalForMockDb.__mockDb = new MockDb())
