export type ProfileInfo = {
  name: string
  email: string
  plan: string
}

type StorageSummary = {
  usedGb: number
  totalGb: number
}

export const useProfileController = () => {
  const profile: ProfileInfo = {
    name: 'John Doe',
    email: 'john@example.com',
    plan: 'Pro',
  }

  const storage: StorageSummary = {
    usedGb: 128,
    totalGb: 256,
  }

  return {
    profile,
    storage,
  }
}
