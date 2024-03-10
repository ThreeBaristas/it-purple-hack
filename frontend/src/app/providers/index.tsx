import { QueryClient, QueryClientProvider } from '@tanstack/react-query'

type Props = {
  queryClient: QueryClient
}

export function Providers({ queryClient }: Props) {
  return (
    <QueryClientProvider client={queryClient}>
      <h1 className="text-purple-600">Hello</h1>
    </QueryClientProvider>
  )
}
