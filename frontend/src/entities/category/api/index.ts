import { queryOptions } from '@tanstack/react-query'

import { getCategories, getCategoryById } from '@/shared/api'

export const getCategoriesQueryOptions = (query?: string) =>
  queryOptions({
    queryKey: ['categories', { query }],
    queryFn: () => getCategories(query)
  })

export const getCategoryByIdQueryOptions = (id: number) =>
  queryOptions({
    queryKey: ['categories', { id }],
    queryFn: () => getCategoryById(id)
  })
