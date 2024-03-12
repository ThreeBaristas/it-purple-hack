import { valibotResolver } from '@hookform/resolvers/valibot'
import { createRoute } from '@tanstack/react-router'
import { useForm } from 'react-hook-form'
import { number, object, type Output, parse, partial } from 'valibot'

import { getLocationsQueryOptions } from '@/entities/location/api'
import { Card, CardContent, CardHeader, CardTitle } from '@/shared/ui'

import { rootRoute } from '../root'
import { getPriceQueryOptions } from './api'

const locationSchema = object({
  id: number(),
  name: number()
})

export const priceFormSchema = object({
  location: locationSchema,
  category: locationSchema,
  segment_id: number()
})

export const priceSearchSchema = partial(priceFormSchema)

export type PriceRequest = Output<typeof priceFormSchema>

export const priceRoute = createRoute({
  getParentRoute: () => rootRoute,
  component: PriceRouteComponent,
  path: '/price',
  validateSearch: (data) => parse(priceSearchSchema, data),
  loaderDeps: ({ search }) => ({ search }),
  loader: ({ deps, context: { queryClient } }) => {
    return queryClient.ensureQueryData(getPriceQueryOptions(deps.search))
  }
})

export function PriceRouteComponent() {
  const form = useForm<Output<typeof priceFormSchema>>({
    resolver: valibotResolver(priceFormSchema)
  })

  return (
    <div>
      <Card className="mx-auto max-w-md">
        <CardHeader>
          <CardTitle>Изменение цены</CardTitle>
        </CardHeader>
        <CardContent className="space-y-4"></CardContent>
      </Card>
    </div>
  )
}
