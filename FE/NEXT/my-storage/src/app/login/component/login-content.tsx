'use client'

import { Button } from '@/components/common/button'
import { Card } from '@/components/common/card'
import { PageShell } from '@/components/common/page-shell'
import { TextInput } from '@/components/common/text-input'
import { useLoginController } from '@/app/login/hook/useLoginController'

export const LoginContent = () => {
  const { form, updateField, handleLogin, handleGoogleLogin } = useLoginController()

  return (
    <PageShell title="Login" subtitle="Sign in with username/password or Google.">
      <Card>
        <div className="space-y-4">
          <TextInput
            label="Username"
            value={form.username}
            onChange={(event) => updateField('username', event.target.value)}
            placeholder="Enter your username"
          />
          <TextInput
            label="Password"
            type="password"
            value={form.password}
            onChange={(event) => updateField('password', event.target.value)}
            placeholder="Enter your password"
          />
          <div className="flex flex-col gap-2 sm:flex-row">
            <Button type="button" onClick={handleLogin} className="w-full sm:w-auto">
              Login
            </Button>
            <Button
              type="button"
              variant="secondary"
              onClick={handleGoogleLogin}
              className="w-full sm:w-auto"
            >
              Continue with Google
            </Button>
          </div>
        </div>
      </Card>
    </PageShell>
  )
}
