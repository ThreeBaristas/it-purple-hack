import { Category } from '@/entities/category'
import { Location } from '@/entities/location'

export type Rule = {
  location: Location
  category: Category
  matrix_id: number
  price: number
}
