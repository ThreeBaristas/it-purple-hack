import {
  createRouter,
  RouterProvider as TanstackRouterProvider
} from '@tanstack/react-router'

import { routeTree } from '@/pages/root'

import { client } from './query'

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
