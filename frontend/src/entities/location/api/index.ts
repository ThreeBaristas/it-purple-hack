import { queryOptions } from '@tanstack/react-query'

import { getLocations } from '@/shared/api'

export const getLocationsQueryOptions = (query?: string) =>
  queryOptions({
    queryKey: ['locations', { query }],
    queryFn: () => getLocations(query)
  })
