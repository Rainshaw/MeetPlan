/// <reference types="vite/client" />

interface ImportMetaEnv {
  readonly VITE_APP_TITLE: string
  readonly VITE_APP_TITLE_EN: string
  readonly VITE_BASE_URL: string
  readonly VITE_HOST_URL: string
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}