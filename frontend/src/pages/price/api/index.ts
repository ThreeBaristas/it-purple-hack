import { queryOptions } from '@tanstack/react-query'

import { getPrice } from '@/shared/api'

import { PriceRequest } from '..'

export const getPriceQueryOptions = (req: PriceRequest) =>
  queryOptions({
    queryKey: ['price', req],
    queryFn: () => getPrice(req)
  })
