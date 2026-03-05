'use client'

import { api } from '@/lib/api'

export type AuthUser = {
    id: string
    email: string
}

type AuthResult = {
    accessToken: string
    refreshToken: string
    user: AuthUser
}

export const authService = {
    login: (email: string, password: string) =>
        api.post<AuthResult>('/api/v1/auth/login', { email, password }),

    logout: () => api.post<{ success: boolean }>('/api/v1/auth/logout'),

    refresh: () => api.post<AuthResult>('/api/v1/auth/refresh'),

    getMe: () => api.get<AuthUser>('/api/v1/auth/me'),
}
