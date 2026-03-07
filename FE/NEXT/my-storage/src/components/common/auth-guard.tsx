'use client'

import { useEffect, useState } from 'react'
import { useRouter } from 'next/navigation'
import { authService } from '@/lib/auth-service'
import { useAuthStore } from '@/store/useAuthStore'

type AuthGuardProps = {
    children: React.ReactNode
}

export const AuthGuard = ({ children }: AuthGuardProps) => {
    const router = useRouter()
    const { isLoggedIn, setUser } = useAuthStore()
    const [checking, setChecking] = useState(!isLoggedIn)

    useEffect(() => {
        if (isLoggedIn) {
            return
        }

        authService
            .getMe()
            .then((user) => {
                setUser(user)
                setChecking(false)
            })
            .catch(() => {
                router.replace('/login')
            })
    }, [isLoggedIn, setUser, router])

    if (checking) {
        return (
            <div className="flex min-h-screen items-center justify-center">
                <p className="text-sm text-foreground/60">Loading...</p>
            </div>
        )
    }

    return <>{children}</>
}
