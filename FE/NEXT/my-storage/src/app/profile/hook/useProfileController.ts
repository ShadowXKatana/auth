'use client'

import { useState, useEffect } from 'react'
import { authService, type AuthUser } from '@/lib/auth-service'

export type ProfileInfo = {
  id: string
  email: string
}

export const useProfileController = () => {
  const [profile, setProfile] = useState<ProfileInfo | null>(null)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')

  useEffect(() => {
    authService
      .getMe()
      .then((user: AuthUser) => {
        setProfile({
          id: user.id,
          email: user.email,
        })
      })
      .catch(() => {
        setError('Failed to load profile.')
      })
      .finally(() => {
        setLoading(false)
      })
  }, [])

  return {
    profile,
    loading,
    error,
  }
}

