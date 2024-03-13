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
  matrix_id: number
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
  const { data } = await axiosInstance.get<
    Omit<Response, 'data'> & { data: Array<RuleDTO> | undefined }
  >('/admin/rules?' + qs)
  const newData: Response = {
    data: data.data ?? [],
    page: data.page,
    pageSize: data.pageSize,
    totalPages: data.totalPages
  }
  return newData
}
