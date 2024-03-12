import qs from 'qs'

import { axiosInstance } from './axios'

type Response = {
  price: number
  location_id: number
  category_id: number
  matrix_id: number
  user_segment_id: number
}

export type Request = {
  location_id: number
  category_id: number
  segment_id: number
}

export async function getPrice(request: Request): Promise<Response> {
  const query = qs.stringify(request)
  const data = await axiosInstance.get<Response>(`/price?${query}`)
  return data.data
}
