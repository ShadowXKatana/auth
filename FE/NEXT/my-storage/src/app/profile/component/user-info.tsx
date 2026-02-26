import { UserInfoProps } from '@/app/profile/interface'
import { Card } from '@/components/common/card'

export const UserInfo = ({ profile }: UserInfoProps) => {
  return (
    <Card>
      <h2 className="mb-3 text-lg font-medium">User Info</h2>
      <dl className="grid grid-cols-1 gap-2 text-sm sm:grid-cols-2">
        <div>
          <dt className="text-foreground/70">Name</dt>
          <dd>{profile.name}</dd>
        </div>
        <div>
          <dt className="text-foreground/70">Email</dt>
          <dd>{profile.email}</dd>
        </div>
        <div>
          <dt className="text-foreground/70">Plan</dt>
          <dd>{profile.plan}</dd>
        </div>
      </dl>
    </Card>
  )
}
