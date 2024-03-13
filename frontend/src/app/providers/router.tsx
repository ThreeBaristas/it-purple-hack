import {
  createRouter,
  RouterProvider as TanstackRouterProvider
} from '@tanstack/react-router'

import { priceRoute } from '@/pages/price'
import { rootRoute } from '@/pages/root'
import { rulesRoute } from '@/pages/rules'
import { storageRoute } from '@/pages/storage'

import { client } from './query'

export const routeTree = rootRoute.addChildren([
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
