import { queryOptions } from '@tanstack/react-query'

import { getPrice, type GetPriceRequest } from '@/shared/api'

export const getPriceQueryOptions = (req: GetPriceRequest) =>
  queryOptions({
    queryKey: ['price', req],
    queryFn: () => getPrice(req)
  })
