import { type UserInfoProps } from '@/app/profile/interface'
import { Card } from '@/components/common/card'

export const UserInfo = ({ profile }: UserInfoProps) => {
  return (
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
  )
}

