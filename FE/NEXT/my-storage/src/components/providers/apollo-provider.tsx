'use client'

import { ApolloNextAppProvider } from '@apollo/client-integration-nextjs'
import { makeApolloClient } from '@/lib/apollo-client'

type ApolloProviderProps = {
  children: React.ReactNode
}

export function ApolloProvider({ children }: ApolloProviderProps) {
  return <ApolloNextAppProvider makeClient={makeApolloClient}>{children}</ApolloNextAppProvider>
}
