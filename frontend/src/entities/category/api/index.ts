import { queryOptions } from '@tanstack/react-query'

import { getCategories } from '@/shared/api'

export const getCategoriesQueryOptions = (query?: string) =>
  queryOptions({
    queryKey: ['categories', { query }],
    queryFn: () => getCategories(query)
  })
