import Home from '@/app/page'
import { render } from '@testing-library/react'

jest.mock('@/components/common/auth-guard', () => ({
  AuthGuard: ({ children }: { children: React.ReactNode }) => <>{children}</>,
}))

jest.mock('@/store/useAuthStore', () => ({
  useAuthStore: () => ({
    user: { id: '1', email: 'snapshot@example.com' },
    clearUser: jest.fn(),
  }),
}))

jest.mock('next/navigation', () => ({
  useRouter: () => ({
    replace: jest.fn(),
  }),
}))

describe('Home page', () => {
  it('matches snapshot', () => {
    const { asFragment } = render(<Home />)
    expect(asFragment()).toMatchSnapshot()
  })
})
