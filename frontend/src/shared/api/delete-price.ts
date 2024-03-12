import QueryString from 'qs'

import { axiosInstance } from './axios'

export type DeletePriceRequest = {
  location_id: number
  category_id: number
  segment_id: number
}

export async function deletePriceRule(req: DeletePriceRequest) {
  const qs = QueryString.stringify(req)
  await axiosInstance.delete('/admin/price?' + qs)
  return
}
