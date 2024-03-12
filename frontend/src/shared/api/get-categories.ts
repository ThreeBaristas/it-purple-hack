import { axiosInstance } from './axios'

type CategoryDTO = {
  id: number
  name: string
}

type Response = Array<CategoryDTO>

export async function getCategories(query?: string): Promise<Response> {
  let qs = ''
  if (query) {
    qs += `?search=${query}`
  }
  const response = await axiosInstance.get<Response>(`/categories${qs}`)
  return response.data
}

export async function getCategoryById(id: number): Promise<CategoryDTO> {
  const response = await axiosInstance.get<CategoryDTO>(`/categories/${id}`)
  return response.data
}
