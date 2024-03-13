import { valibotResolver } from '@hookform/resolvers/valibot'
import { createRoute } from '@tanstack/react-router'
import { useForm } from 'react-hook-form'
import {
  coerce,
  number,
  object,
  type Output,
  parse,
  partial,
  string
} from 'valibot'

import { SelectCategory } from '@/entities/category'
import { SelectLocation } from '@/entities/location'
import {
  Button,
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle
} from '@/shared/ui'
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel
} from '@/shared/ui/form'
import { Input } from '@/shared/ui/input'

import { rootRoute } from '../root'
import { getPriceQueryOptions, useSavePriceMutation } from './api'

const locationSchema = object({
  id: number(),
  name: string()
})

export const priceFormSchema = object({
  location: locationSchema,
  category: locationSchema,
  matrix_id: coerce(number(), Number),
  price: coerce(number(), Number)
})

export const priceSearchSchema = partial(priceFormSchema)

export type PriceRequest = Output<typeof priceFormSchema>

export const priceRoute = createRoute({
  getParentRoute: () => rootRoute,
  component: PriceRouteComponent,
  path: '/price',
  validateSearch: (data) => parse(priceSearchSchema, data),
  loaderDeps: ({ search }) => ({ search }),
  loader: ({ context: { queryClient }, deps }) => {
    const categoryId = deps.search.category?.id
    const locationId = deps.search.location?.id
    const matrix = deps.search.matrix_id
    if (
      categoryId != undefined &&
      locationId != undefined &&
      matrix != undefined
    ) {
      return queryClient.ensureQueryData(
        getPriceQueryOptions({
          category_id: categoryId,
          location_id: locationId,
          matrix_id: matrix
        })
      )
    }
  }
})

export function PriceRouteComponent() {
  const price = priceRoute.useLoaderData()
  const search = priceRoute.useSearch()
  const form = useForm<Output<typeof priceFormSchema>>({
    resolver: valibotResolver(priceFormSchema),
    defaultValues: {
      ...search,
      price: price?.price
    }
  })

  const { mutate, isPending } = useSavePriceMutation()

  function handleSubmit(values: Output<typeof priceFormSchema>) {
    mutate({
      location_id: values.location.id,
      category_id: values.category.id,
      price: values.price,
      matrix_id: values.matrix_id
    })
  }

  return (
    <div>
      <Card className="mx-auto max-w-md">
        <CardHeader>
          <CardTitle>Изменение цены</CardTitle>
        </CardHeader>
        <CardContent className="space-y-4">
          <Form {...form}>
            <form
              id="price-route-form"
              onSubmit={form.handleSubmit(handleSubmit)}
              className="space-y-4"
            >
              <FormField
                control={form.control}
                name="category"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Категория</FormLabel>
                    <FormControl>
                      <SelectCategory {...field} className="w-full" />
                    </FormControl>
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name="location"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Локация</FormLabel>
                    <FormControl>
                      <SelectLocation {...field} className="w-full" />
                    </FormControl>
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name="matrix_id"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Матрица</FormLabel>
                    <FormControl>
                      <Input {...field} />
                    </FormControl>
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name="price"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Цена</FormLabel>
                    <FormControl>
                      <Input {...field} />
                    </FormControl>
                  </FormItem>
                )}
              />
            </form>
          </Form>
        </CardContent>
        <CardFooter>
          <Button type="submit" form="price-route-form" disabled={isPending}>
            Сохранить
          </Button>
        </CardFooter>
      </Card>
    </div>
  )
}
