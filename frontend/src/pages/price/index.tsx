import { createRoute } from '@tanstack/react-router'
import { number, object, optional, type Output, parse } from 'valibot'

import { Card, CardContent, CardHeader, CardTitle } from '@/shared/ui'
import { Label } from '@/shared/ui/label'

import { rootRoute } from '../root'
import { getPriceQueryOptions } from './api'

export const priceRouteSchema = object({
  location_id: optional(number()),
  category_id: optional(number()),
  segment_id: optional(number()),
  user_id: optional(number())
})

export type PriceRequest = Output<typeof priceRouteSchema>

export const priceRoute = createRoute({
  getParentRoute: () => rootRoute,
  component: PriceRouteComponent,
  path: '/price',
  validateSearch: (data) => parse(priceRouteSchema, data),
  loaderDeps: ({ search }) => ({ search }),
  loader: ({ deps, context: { queryClient } }) => {
    return queryClient.ensureQueryData(getPriceQueryOptions(deps.search))
  }
})

export function PriceRouteComponent() {
  const {
    location_id,
    category_id,
    user_segment_id: segment_id,
    matrix_id,
    price
  } = priceRoute.useLoaderData()
  const { user_id } = priceRoute.useSearch()

  return (
    <div>
      <Card className="mx-auto max-w-md">
        <CardHeader>
          <CardTitle>Изменение цены</CardTitle>
        </CardHeader>
        <CardContent className="space-y-4">
          <div>
            <Label>Категория</Label>
            <p>Категория #{category_id != undefined ? category_id : 'Н/А'}</p>
          </div>
          <div>
            <Label>Локация</Label>
            <p>Локация #{location_id != undefined ? location_id : 'Н/А'}</p>
          </div>
          {user_id != undefined ? (
            <div>
              <Label>Пользователь</Label>
              <p>ID {user_id}</p>
            </div>
          ) : null}
          {segment_id && (
            <div>
              <Label>Загружено из сегмента</Label>
              <p>Сегмент #{segment_id}</p>
            </div>
          )}
          <div>
            <Label>Загружено из матрицы</Label>
            <p>Матрица #{matrix_id}</p>
          </div>
          <div>
            <Label>Цена</Label>
            <p>
              {new Intl.NumberFormat('ru-RU', {
                currency: 'RUB',
                style: 'currency'
              }).format(price)}
            </p>
          </div>
        </CardContent>
      </Card>
    </div>
  )
}
