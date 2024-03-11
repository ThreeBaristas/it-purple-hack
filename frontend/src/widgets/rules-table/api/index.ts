import { queryOptions } from '@tanstack/react-query'

import { getRules, GetRulesRequest } from '@/shared/api'

export const getRulesQueryOptions = (req: GetRulesRequest) =>
  queryOptions({
    queryKey: ['rules', req],
    queryFn: () => getRules(req)
  })
