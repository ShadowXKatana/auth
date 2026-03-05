'use client'

import Link from 'next/link'
import { useRouter } from 'next/navigation'
import { authService } from '@/lib/auth-service'
import { useAuthStore } from '@/store/useAuthStore'

type PageShellProps = {
  title: string
  subtitle?: string
  showBackLink?: boolean
  showLogout?: boolean
  children: React.ReactNode
}

export const PageShell = ({
  title,
  subtitle,
  showBackLink = false,
  showLogout = false,
  children,
}: PageShellProps) => {
  const router = useRouter()
  const { clearUser } = useAuthStore()

  const handleLogout = async () => {
    try {
      await authService.logout()
    } catch {
      // ignore
    }
    clearUser()
    router.replace('/login')
  }

  return (
    <main className="mx-auto w-full max-w-3xl p-6 md:p-10">
      <header className="mb-6">
        <div className="flex items-center justify-between">
          <div>
            {showBackLink ? (
              <Link
                href="/"
                className="mb-2 inline-block text-sm text-foreground/60 hover:text-foreground"
              >
                ← Back to Home
              </Link>
            ) : null}
            <h1 className="text-2xl font-semibold tracking-tight">{title}</h1>
            {subtitle ? (
              <p className="mt-1 text-sm text-foreground/80">{subtitle}</p>
            ) : null}
          </div>
          {showLogout ? (
            <button
              type="button"
              onClick={handleLogout}
              className="rounded-lg border border-black/15 px-3 py-1.5 text-sm font-medium transition hover:bg-black/5 dark:border-white/20 dark:hover:bg-white/10"
            >
              Logout
            </button>
          ) : null}
        </div>
      </header>
      {children}
    </main>
  )
}

