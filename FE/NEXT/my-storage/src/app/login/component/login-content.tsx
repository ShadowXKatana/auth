'use client'

import { Button } from '@/components/common/button'
import { Card } from '@/components/common/card'
import { PageShell } from '@/components/common/page-shell'
import { TextInput } from '@/components/common/text-input'
import { useLoginController } from '@/app/login/hook/useLoginController'

export const LoginContent = () => {
  const { form, error, loading, updateField, handleLogin } = useLoginController()

  return (
    <PageShell title="Login" subtitle="Sign in to your account.">
      <Card>
        <div className="space-y-4">
          <TextInput
            label="Email"
            type="email"
            value={form.email}
            onChange={(event) => updateField('email', event.target.value)}
            placeholder="Enter your email"
          />
          <TextInput
            label="Password"
            type="password"
            value={form.password}
            onChange={(event) => updateField('password', event.target.value)}
            placeholder="Enter your password"
          />
          {error ? (
            <p className="text-sm text-red-500">{error}</p>
          ) : null}
          <Button type="button" onClick={handleLogin} disabled={loading} className="w-full">
            {loading ? 'Signing in...' : 'Login'}
          </Button>
        </div>
      </Card>
    </PageShell>
  )
}

