'use client'

import Link from 'next/link'
import { Card } from '@/components/common/card'
import { PageShell } from '@/components/common/page-shell'
import { AuthGuard } from '@/components/common/auth-guard'
import { useAuthStore } from '@/store/useAuthStore'

const navLinks = [
  { name: 'Profile', href: '/profile', description: 'View your account details' },
  { name: 'Storage', href: '/storage', description: 'Manage your storages and items' },
]

export default function Home() {
  const { user } = useAuthStore()

  return (
    <AuthGuard>
      <PageShell
        title="My Storage"
        subtitle={user ? `Welcome, ${user.email}` : 'Select a page to navigate.'}
        showLogout
      >
        <div className="space-y-3">
          {navLinks.map((link) => (
            <Card key={link.href}>
              <Link href={link.href} className="block">
                <p className="text-sm font-medium hover:underline">{link.name}</p>
                <p className="mt-1 text-xs text-foreground/60">{link.description}</p>
              </Link>
            </Card>
          ))}
        </div>
      </PageShell>
    </AuthGuard>
  )
}

// test
