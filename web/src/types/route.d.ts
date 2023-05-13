import 'vue-router'

declare module 'vue-router' {
  interface RouteMeta {
    title: string
    layout?: string
    needAuth?: boolean
    role?: string[]
  }
}