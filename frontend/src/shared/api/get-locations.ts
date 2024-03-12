import { axiosInstance } from './axios'

type LocationDTO = {
  id: number
  name: string
}

type Response = Array<LocationDTO>

export async function getLocations(query?: string): Promise<Response> {
  let qs = ''
  if (query) {
    qs += `?search=${query}`
  }
  const response = await axiosInstance.get<Response>(`/locations${qs}`)
  return response.data
}

export async function getLocation(id: number): Promise<LocationDTO> {
  const response = await axiosInstance.get<LocationDTO>(`/locations/${id}`)
  return response.data
}
