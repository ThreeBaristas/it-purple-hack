import { axiosInstance } from './axios'

type CategoryDTO = {
  id: number
  name: string
}

type Response = Array<CategoryDTO>

export async function getCategories(query?: string): Promise<Response> {
  if (process.env.NODE_ENV === 'production') {
    let qs = ''
    if (query) {
      qs += `?search=${query}`
    }
    const response = await axiosInstance.get<Response>(`/categories${qs}`)
    return response.data
  }
  await new Promise((r) => setTimeout(r, 500))
  const data: Response = [
    {
      id: 1,
      name: 'Личные товары'
    },
    {
      id: 2,
      name: 'Транспорт'
    }
  ]
  return data.filter((it) =>
    it.name.toLowerCase().includes(query?.toLocaleLowerCase() ?? '')
  )
}
