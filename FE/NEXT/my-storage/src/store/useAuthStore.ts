import { create } from 'zustand'

type AuthUser = {
    id: string
    email: string
}

type AuthState = {
    user: AuthUser | null
    isLoggedIn: boolean
    setUser: (user: AuthUser) => void
    clearUser: () => void
}

export const useAuthStore = create<AuthState>((set) => ({
    user: null,
    isLoggedIn: false,
    setUser: (user) => set({ user, isLoggedIn: true }),
    clearUser: () => set({ user: null, isLoggedIn: false }),
}))
