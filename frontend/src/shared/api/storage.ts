import {
  queryOptions,
  useMutation,
  useQueryClient
} from '@tanstack/react-query'

import { axiosInstance } from './axios'

export type StorageDiscount = {
  matrix_id: number
  segment_id: number
}

export type Storage = {
  baseline_matrix_id: number
  discounts: Array<StorageDiscount>
}

export const getStorageQueryOptions = () =>
  queryOptions({
    queryKey: ['storage'],
    queryFn: async () => {
      const { data } = await axiosInstance.get<Storage>('/admin/storage')
      data.discounts.sort((a, b) => a.segment_id - b.segment_id)
      return data
    }
  })

export function useSaveStorageMutation() {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: async (storage: Storage) => {
      await axiosInstance.post<Storage>('/admin/storage', storage)
    },
    onMutate: async () => {
      await queryClient.invalidateQueries(getStorageQueryOptions())
    }
  })
}
