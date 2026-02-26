'use client'

import { ApolloProvider } from '@/components/providers/apollo-provider'

type ProvidersProps = {
  children: React.ReactNode
}

export const Providers = ({ children }: ProvidersProps) => {
  return <ApolloProvider>{children}</ApolloProvider>
}
