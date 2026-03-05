'use client'

import { useState } from 'react'
import { useRouter } from 'next/navigation'
import { authService } from '@/lib/auth-service'
import { useAuthStore } from '@/store/useAuthStore'

type LoginFormState = {
  email: string
  password: string
}

export const useLoginController = () => {
  const router = useRouter()
  const { setUser } = useAuthStore()
  const [form, setForm] = useState<LoginFormState>({
    email: '',
    password: '',
  })
  const [error, setError] = useState('')
  const [loading, setLoading] = useState(false)

  const updateField = (field: keyof LoginFormState, value: string) => {
    setForm((previous) => ({ ...previous, [field]: value }))
    setError('')
  }

  const handleLogin = async () => {
    if (!form.email || !form.password) {
      setError('Please enter both email and password.')
      return
    }

    setLoading(true)
    setError('')

    try {
      const result = await authService.login(form.email, form.password)
      setUser(result.user)
      router.replace('/')
    } catch {
      setError('Invalid email or password.')
    } finally {
      setLoading(false)
    }
  }

  return {
    form,
    error,
    loading,
    updateField,
    handleLogin,
  }
}

