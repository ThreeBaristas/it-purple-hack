import { axiosInstance } from './axios'

type CategoryDTO = {
  id: number
  name: string
}

type Response = Array<CategoryDTO>

export async function GetCategories(query?: string): Promise<Response> {
  let qs = ''
  if (query) {
    qs += `?search=${query}`
  }
  const response = await axiosInstance.get<Response>(`/categories${qs}`)
  return response.data
}
