'use client'

import { Card } from '@/components/common/card'
import { PageShell } from '@/components/common/page-shell'
import { useProfileController } from '@/app/profile/hook/useProfileController'
import { useAppRouter } from '@/hooks/useAppRouter'
import { AuthGuard } from '@/components/common/auth-guard'

export const ProfileContent = () => {
  const { profile, loading, error } = useProfileController()
  const { navigateTo } = useAppRouter()

  return (
    <AuthGuard>
      <PageShell title="Profile" subtitle="Your account information." showBackLink showLogout>
        <div className="space-y-4">
          {loading ? (
            <Card>
              <p className="text-sm text-foreground/60">Loading profile...</p>
            </Card>
          ) : error ? (
            <Card>
              <p className="text-sm text-red-500">{error}</p>
            </Card>
          ) : profile ? (
            <Card>
              <h2 className="mb-3 text-lg font-medium">User Info</h2>
              <dl className="grid grid-cols-1 gap-2 text-sm sm:grid-cols-2">
                <div>
                  <dt className="text-foreground/70">ID</dt>
                  <dd className="font-mono text-xs">{profile.id}</dd>
                </div>
                <div>
                  <dt className="text-foreground/70">Email</dt>
                  <dd>{profile.email}</dd>
                </div>
              </dl>
            </Card>
          ) : null}

          <Card>
            <button
              type="button"
              onClick={() => navigateTo('/storage')}
              className="block w-full text-left"
            >
              <h2 className="mb-1 text-lg font-medium">Storage</h2>
              <span className="text-sm font-medium hover:underline">Go to storage →</span>
            </button>
          </Card>
        </div>
      </PageShell>
    </AuthGuard>
  )
}

