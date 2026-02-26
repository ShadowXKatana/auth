'use client'

import { Card } from '@/components/common/card'
import { PageShell } from '@/components/common/page-shell'
import { useProfileController } from '@/app/profile/hook/useProfileController'
import { useAppRouter } from '@/hooks/useAppRouter'
import { UserInfo } from '@/app/profile/component/user-info'

export const ProfileContent = () => {
  const { profile, storage } = useProfileController()
  const { navigateTo } = useAppRouter()

  return (
    <PageShell title="Profile" subtitle="User information and storage summary.">
      <div className="space-y-4">
        <UserInfo profile={profile} />

        <Card>
          <button
            type="button"
            onClick={() => navigateTo('/storage')}
            className="block w-full text-left"
          >
            <h2 className="mb-3 text-lg font-medium">Storage</h2>
            <p className="text-sm">
              Used {storage.usedGb} GB / {storage.totalGb} GB
            </p>
            <span className="mt-3 inline-block text-sm font-medium hover:underline">
              Go to storage
            </span>
          </button>
        </Card>
      </div>
    </PageShell>
  )
}
