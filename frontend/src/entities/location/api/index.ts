import { queryOptions } from '@tanstack/react-query'

import { getLocationById, getLocations } from '@/shared/api'

export const getLocationsQueryOptions = (query?: string) =>
  queryOptions({
    queryKey: ['locations', { query }],
    queryFn: () => getLocations(query)
  })

export const getLocationByIdQueryOptions = (id: number) =>
  queryOptions({
    queryKey: ['locations', { id }],
    queryFn: () => getLocationById(id)
  })
