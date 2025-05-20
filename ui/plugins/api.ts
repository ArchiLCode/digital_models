import { ofetch } from 'ofetch'

export default defineNuxtPlugin(() => {
  const runtimeConfig = useRuntimeConfig()
  const baseURL = runtimeConfig.public.baseUrl + '/api'

  const api = {
    get: async (url: string, queryParams?: URLSearchParams) => {
      const { token } = useAuth()

      const fullUrl = queryParams ? `${url}?${queryParams.toString()}` : url

      return ofetch(fullUrl, {
        baseURL,
        method: 'GET',
        headers: {
          ...(token.value ? { Authorization: `${token.value}` } : {}),
        },
      })
    },
    post: async (url: string, body?: Record<string, any>) => {
      const { token } = useAuth()

      return ofetch(url, {
        baseURL,
        method: 'POST',
        headers: {
          ...(token.value ? { Authorization: `${token.value}` } : {}),
        },
        body,
      })
    },
    update: async (url: string, body?: Record<string, any>) => {
      const { token } = useAuth()

      return ofetch(url, {
        baseURL,
        method: 'UPDATE',
        headers: {
          ...(token.value ? { Authorization: `${token.value}` } : {}),
        },
        body,
      })
    },
    put: async (url: string, body?: Record<string, any>) => {
      const { token } = useAuth()

      return ofetch(url, {
        baseURL,
        method: 'PUT',
        headers: {
          ...(token.value ? { Authorization: `${token.value}` } : {}),
        },
        body,
      })
    },
  }

  return {
    provide: {
      api,
    },
  }
})
