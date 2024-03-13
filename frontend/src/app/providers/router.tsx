import {
  createRouter,
  RouterProvider as TanstackRouterProvider
} from '@tanstack/react-router'

import { indexRoute } from '@/pages/admin'
import { priceRoute } from '@/pages/price'
import { rootRoute } from '@/pages/root'
import { rulesRoute } from '@/pages/rules'
import { storageRoute } from '@/pages/storage'

import { client } from './query'

export const routeTree = rootRoute.addChildren([
  indexRoute,
  priceRoute,
  rulesRoute,
  storageRoute
])

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
