import { createRoute } from '@tanstack/react-router'
import { fallback, number, object, parse } from 'valibot'

import { getRulesQueryOptions } from '@/widgets/rules-table/api'

import { rootRoute } from '../root'
import { RulesPageComponent } from './ui'

const schema = object({
  page: fallback(number(), 0),
  pageSize: fallback(number(), 2)
})

export const rulesRoute = createRoute({
  getParentRoute: () => rootRoute,
  component: RulesPageComponent,
  path: '/rules',
  validateSearch: (data) => parse(schema, data),
  loaderDeps: ({ search }) => ({ ...search }),
  loader: ({ context: { queryClient }, deps }) => {
    return queryClient.ensureQueryData(getRulesQueryOptions(deps))
  }
})
