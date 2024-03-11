import { QueryClient } from '@tanstack/react-query'
import { createRootRouteWithContext } from '@tanstack/react-router'

import { indexRoute } from './admin'
import { Layout } from './layout'
import { priceRoute } from './price'
import { rulesRoute } from './rules'

export const rootRoute = createRootRouteWithContext<{
  queryClient: QueryClient
}>()({
  component: Layout
})

export const routeTree = rootRoute.addChildren([
  indexRoute,
  priceRoute,
  rulesRoute
])
