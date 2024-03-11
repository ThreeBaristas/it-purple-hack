import QueryString from 'qs'

import { axiosInstance } from './axios'

export type Request = {
  page?: number
  pageSize?: number
}

type IdAndName = {
  name: string
  id: number
}

type RuleDTO = {
  location: IdAndName
  category: IdAndName
  segment: number
  price: number
}

type Response = {
  data: Array<RuleDTO>
  page: number
  pageSize: number
  totalPages: number
}

export async function getRules(req: Request): Promise<Response> {
  const qs = QueryString.stringify(req)
  const data = await axiosInstance.get<Response>('/admin/rules?' + qs)
  return data.data
}
