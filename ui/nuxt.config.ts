export default defineNuxtConfig({
  compatibilityDate: '2024-11-01',
  ssr: true,
  devtools: { enabled: true },
  modules: ['@nuxtjs/i18n', '@sidebase/nuxt-auth'],
  runtimeConfig: {
    public: {
      baseUrl: process.env.NUXT_PUBLIC_BASE_URL,
      messengerBaseUrl:
        process.env.NUXT_PUBLIC_MESSENGER_BASE_URL || 'http://localhost:8083',
    },
  },
  nitro: {
    routeRules: {
      '/api/**': {
        proxy: process.env.NUXT_PUBLIC_API_BASE_URL + '/api/**',
      },
      '/upload': {
        proxy: process.env.NUXT_PUBLIC_FILE_BASE_URL + '/upload',
      },
      '/messenger/**': {
        proxy: process.env.NUXT_PUBLIC_MESSENGER_BASE_URL + '/api/**',
      },
    },
    devProxy: {
      '/messenger/ws': {
        target:
          process.env.NUXT_PUBLIC_MESSENGER_BASE_URL?.replace('http', 'ws') ||
          'ws://localhost:8083',
        ws: true,
      },
    },
    publicAssets: [{ dir: 'public', maxAge: 60 }],
  },
  auth: {
    isEnabled: true,
    baseURL: process.env.NUXT_PUBLIC_BASE_URL + '/api/auth',
    provider: {
      type: 'local',
      endpoints: {
        signIn: { path: '/login', method: 'post' },
        signUp: { path: '/register', method: 'post' },
        signOut: { path: '/logout', method: 'get' },
        getSession: { path: '/session', method: 'get' },
      },
      session: {
        dataType: {
          user_id: 'string',
          first_name: 'string',
          last_name: 'string',
          role: 'string',
          is_admin: 'boolean',
        },
      },
      token: {
        type: 'Bearer',
        cookieName: 'auth.token',
        cookieDomain: process.env.NUXT_PUBLIC_BASE_URL,
        headerName: 'Authorization',
        signInResponseTokenPointer: '/access_token',
        maxAgeInSeconds: 52200,
        secureCookieAttribute: true, // баг библы, secure кука удаляется при перезагрузке страницы
        httpOnlyCookieAttribute: true,
        sameSiteAttribute: 'strict',
      },
      refresh: {
        isEnabled: true,
        endpoint: { path: '/refresh-token', method: 'post' },
        refreshOnlyToken: false,
        token: {
          signInResponseRefreshTokenPointer: '/refresh_token',
          refreshRequestTokenPointer: '/refresh_token',
          refreshResponseTokenPointer: '/access_token',
          maxAgeInSeconds: 604800,
          secureCookieAttribute: true,
          httpOnlyCookieAttribute: true,
          sameSiteAttribute: 'strict',
          cookieName: 'auth.refresh_token',
          cookieDomain: process.env.NUXT_PUBLIC_BASE_URL,
        },
      },
    },
    sessionRefresh: {
      enablePeriodically: 3399999,
      enableOnWindowFocus: false,
    },
  },
  i18n: {
    locales: [
      { code: 'en', language: 'en-US' },
      { code: 'ru', language: 'ru-RU' },
    ],
    defaultLocale: 'ru',
  },
  css: ['~/assets/css/transitions.css'],
  app: {
    head: {
      link: [
        { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
        {
          rel: 'apple-touch-icon',
          sizes: '180x180',
          href: '/apple-touch-icon.png',
        },
        {
          rel: 'icon',
          type: 'image/png',
          sizes: '32x32',
          href: '/favicon-32x32.png',
        },
        {
          rel: 'icon',
          type: 'image/png',
          sizes: '16x16',
          href: '/favicon-16x16.png',
        },
        { rel: 'manifest', href: '/manifest.json' },
      ],
    },
    pageTransition: { name: 'page', mode: 'out-in' },
  },
})
