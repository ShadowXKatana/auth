'use client'

import { useRouter } from 'next/navigation'

export const useAppRouter = () => {
  const router = useRouter()

  const navigateTo = (path: string) => {
    router.push(path)
  }

  const replaceTo = (path: string) => {
    router.replace(path)
  }

  const goBack = () => {
    router.back()
  }

  return {
    navigateTo,
    replaceTo,
    goBack,
  }
}
