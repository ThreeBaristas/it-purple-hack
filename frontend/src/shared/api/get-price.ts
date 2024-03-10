import qs from 'qs'

import { axiosInstance } from './axios'

type Response = {
  price: number
  location_id: number
  microcategory_id: number
  matrix_id: number
  user_segment_id: number
}

type Request = {
  location_id?: number
  category_id?: number
} & ({ segment_id?: number } | { user_id?: number })

export async function getPrice(request: Request): Promise<Response> {
  if (process.env.NODE_ENV == 'production') {
    const query = qs.stringify(request)
    const data = await axiosInstance.get<Response>(`/price${query}`)
    return data.data
  }
  return {
    price: 1000,
    location_id: 1,
    microcategory_id: 1,
    matrix_id: 100,
    user_segment_id: 10
  }
}
