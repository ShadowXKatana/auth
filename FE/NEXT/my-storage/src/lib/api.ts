'use client'

import { env } from '@/lib/env'

const BASE_URL = env.NEXT_PUBLIC_API_URL

type RequestOptions = {
    body?: unknown
    headers?: Record<string, string>
}

async function request<T>(
    method: string,
    path: string,
    options: RequestOptions = {},
): Promise<T> {
    const url = `${BASE_URL}${path}`

    const res = await fetch(url, {
        method,
        credentials: 'include',
        headers: {
            'Content-Type': 'application/json',
            ...options.headers,
        },
        body: options.body ? JSON.stringify(options.body) : undefined,
    })

    if (res.status === 401 && !path.includes('/auth/')) {
        const refreshRes = await fetch(`${BASE_URL}/api/v1/auth/refresh`, {
            method: 'POST',
            credentials: 'include',
        })

        if (refreshRes.ok) {
            const retryRes = await fetch(url, {
                method,
                credentials: 'include',
                headers: {
                    'Content-Type': 'application/json',
                    ...options.headers,
                },
                body: options.body ? JSON.stringify(options.body) : undefined,
            })

            if (!retryRes.ok) {
                throw new ApiError(retryRes.status, await retryRes.text())
            }

            return retryRes.json() as Promise<T>
        }

        throw new ApiError(401, 'session expired')
    }

    if (!res.ok) {
        throw new ApiError(res.status, await res.text())
    }

    if (res.status === 204) {
        return {} as T
    }

    return res.json() as Promise<T>
}

export class ApiError extends Error {
    constructor(
        public status: number,
        message: string,
    ) {
        super(message)
        this.name = 'ApiError'
    }
}

export const api = {
    get: <T>(path: string) => request<T>('GET', path),
    post: <T>(path: string, body?: unknown) => request<T>('POST', path, { body }),
    patch: <T>(path: string, body?: unknown) => request<T>('PATCH', path, { body }),
    delete: <T>(path: string) => request<T>('DELETE', path),
}
