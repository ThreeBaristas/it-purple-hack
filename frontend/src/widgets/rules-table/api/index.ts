import {
  queryOptions,
  useMutation,
  useQueryClient
} from '@tanstack/react-query'

import { getPriceQueryOptions } from '@/pages/price/api'
import { getRules, GetRulesRequest } from '@/shared/api'
import { DeletePriceRequest, deletePriceRule } from '@/shared/api/delete-price'

export const getRulesQueryOptions = (req: GetRulesRequest) =>
  queryOptions({
    queryKey: ['rules', req],
    queryFn: () => getRules(req)
  })

export function useDeletePriceMutation(req: DeletePriceRequest) {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: () => deletePriceRule(req),
    onSuccess: () => {
      queryClient.invalidateQueries(getPriceQueryOptions({ ...req }))
      queryClient.invalidateQueries({ queryKey: ['rules'] })
    }
  })
}
