import { z } from 'zod'

const envSchema = z.object({
  NEXT_PUBLIC_APP_NAME: z.string().min(1).default('my-storage'),
  NEXT_PUBLIC_GRAPHQL_ENDPOINT: z.string().url().default('http://localhost:4000/graphql'),
  NEXT_PUBLIC_API_URL: z.string().url().default('http://localhost:8080'),
})

export const env = envSchema.parse({
  NEXT_PUBLIC_APP_NAME: process.env.NEXT_PUBLIC_APP_NAME,
  NEXT_PUBLIC_GRAPHQL_ENDPOINT: process.env.NEXT_PUBLIC_GRAPHQL_ENDPOINT,
  NEXT_PUBLIC_API_URL: process.env.NEXT_PUBLIC_API_URL,
})

