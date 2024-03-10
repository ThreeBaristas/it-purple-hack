import { QueryClient, QueryClientProvider } from '@tanstack/react-query'

import { RouterProvider } from './router'

type Props = {
  queryClient: QueryClient
}

export function Providers({ queryClient }: Props) {
  return (
    <QueryClientProvider client={queryClient}>
      <RouterProvider />
    </QueryClientProvider>
  )
}
