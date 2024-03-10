import { axiosInstance } from './axios'

type LocationDTO = {
  id: number
  name: string
}

type Response = Array<LocationDTO>

export async function getLocations(query?: string): Promise<Response> {
  if (process.env.NODE_ENV === 'production') {
    let qs = ''
    if (query) {
      qs += `?search=${query}`
    }
    const response = await axiosInstance.get<Response>(`/locations${qs}`)
    return response.data
  }
  await new Promise((r) => setTimeout(r, 500))
  const data: Response = [
    {
      id: 1,
      name: 'Санкт-Петербург'
    },
    {
      id: 2,
      name: 'Москва'
    }
  ]
  return data.filter((it) =>
    it.name.toLowerCase().includes(query?.toLocaleLowerCase() ?? '')
  )
}
