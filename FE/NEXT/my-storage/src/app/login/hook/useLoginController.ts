'use client'

import { useState } from 'react'

type LoginFormState = {
  username: string
  password: string
}

export const useLoginController = () => {
  const [form, setForm] = useState<LoginFormState>({
    username: '',
    password: '',
  })

  const updateField = (field: keyof LoginFormState, value: string) => {
    setForm((previous) => ({ ...previous, [field]: value }))
  }

  const handleLogin = () => {
    console.log('login', form)
  }

  const handleGoogleLogin = () => {
    console.log('google login')
  }

  return {
    form,
    updateField,
    handleLogin,
    handleGoogleLogin,
  }
}
