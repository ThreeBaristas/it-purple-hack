import {
  queryOptions,
  useMutation,
  useQueryClient
} from '@tanstack/react-query'
import QueryString from 'qs'

import { getPrice, type GetPriceRequest } from '@/shared/api'
import { axiosInstance } from '@/shared/api/axios'

export const getPriceQueryOptions = (req: GetPriceRequest) =>
  queryOptions({
    queryKey: ['price', req],
    queryFn: () => getPrice(req)
  })

type SaveReq = {
  location_id: number
  category_id: number
  matrix_id: number
  price: number
}

export function useSavePriceMutation() {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: (req: SaveReq) => {
      const qs = QueryString.stringify(req)
      return axiosInstance.put('/admin/price?' + qs)
    },
    onMutate: () => {
      queryClient.invalidateQueries({ queryKey: ['rules'] })
    }
  })
}
