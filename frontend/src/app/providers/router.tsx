import { QueryClient } from '@tanstack/react-query'
import {
  createRootRouteWithContext,
  createRoute,
  createRouter,
  RouterProvider as TanstackRouterProvider
} from '@tanstack/react-router'

import { Layout } from '@/pages/layout'

import { AdminPageComponent } from '../../pages/admin'
import { client } from './query'

const rootRoute = createRootRouteWithContext<{ queryClient: QueryClient }>()({
  component: Layout
})

const indexRoute = createRoute({
  getParentRoute: () => rootRoute,
  component: AdminPageComponent,
  path: '/'
})

const routeTree = rootRoute.addChildren([indexRoute])

const router = createRouter({
  routeTree,
  context: {
    queryClient: client
  }
})

declare module '@tanstack/react-router' {
  interface Register {
    router: typeof router
  }
}

export function RouterProvider() {
  return <TanstackRouterProvider router={router} />
}
