import { QueryClient } from '@tanstack/react-query'
import { createRootRouteWithContext } from '@tanstack/react-router'

import { Layout } from './layout'

export const rootRoute = createRootRouteWithContext<{
  queryClient: QueryClient
}>()({
  component: Layout
})
