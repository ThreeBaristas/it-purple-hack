import { valibotResolver } from '@hookform/resolvers/valibot'
import { useMutation, useSuspenseQuery } from '@tanstack/react-query'
import { createRoute } from '@tanstack/react-router'
import { PlusCircle } from 'lucide-react'
import { useEffect, useMemo } from 'react'
import { useFieldArray, useForm } from 'react-hook-form'
import { array, coerce, number, object, Output } from 'valibot'

import {
  getStorageQueryOptions,
  StorageDiscount,
  useSaveStorageMutation
} from '@/shared/api/storage'
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
  FormField,
  FormItem,
  FormLabel
} from '@/shared/ui/form'
import { Input } from '@/shared/ui/input'

import { rootRoute } from '../root'

export const storageRoute = createRoute({
  getParentRoute: () => rootRoute,
  component: Component,
  path: '/storage',
  loader: ({ context: { queryClient } }) => {
    return queryClient.ensureQueryData(getStorageQueryOptions())
  }
})

function Component() {
  const { data, refetch } = useSuspenseQuery(getStorageQueryOptions())
  const schema = object({
    baseline_matrix_id: coerce(number(), Number),
    discounts: array(
      object({
        segment_id: coerce(number(), Number),
        matrix_id: coerce(number(), Number)
      })
    )
  })

  const form = useForm<Output<typeof schema>>({
    resolver: valibotResolver(schema),
    defaultValues: useMemo(() => {
      return data
    }, [data])
  })

  useEffect(() => {
    console.log('resetting form data to ', data)
    form.reset(data)
  }, [data, form])

  const { mutate, isPending } = useSaveStorageMutation()

  const { fields, append, remove } = useFieldArray({
    control: form.control,
    name: 'discounts'
  })

  function onSubmit(values: Output<typeof schema>) {
    mutate(values, {
      onSuccess: () => {
        refetch()
      }
    })
  }

  return (
    <Card className="mx-auto max-w-md">
      <CardHeader>
        <CardTitle>Конфигурация стораджа</CardTitle>
      </CardHeader>
      <CardContent>
        <Form {...form}>
          <form
            id="storage-form"
            onSubmit={form.handleSubmit(onSubmit)}
            className="space-y-8"
          >
            <FormField
              control={form.control}
              name="baseline_matrix_id"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Baseline матрица</FormLabel>
                  <FormControl>
                    <Input {...field} />
                  </FormControl>
                </FormItem>
              )}
            />
            <div className="space-y-4">
              <div className="grid grid-cols-2 gap-4 space-x-0 space-y-0">
                <FormLabel>Сегмент</FormLabel>
                <FormLabel>Матрица</FormLabel>
              </div>
              {fields.map((field, index) => (
                <FormItem
                  className="grid grid-cols-2 gap-4 space-x-0 space-y-0"
                  key={field.id}
                >
                  <Input {...form.register(`discounts.${index}.segment_id`)} />
                  <Input {...form.register(`discounts.${index}.matrix_id`)} />
                </FormItem>
              ))}
              <div className="grid grid-cols-2 gap-4 space-x-0 space-y-0">
                <Button
                  variant="outline"
                  onClick={() => append({ segment_id: 0, matrix_id: 0 })}
                >
                  <PlusCircle className="mr-2 size-4" />
                  Добавить
                </Button>
                <Button
                  variant="destructive"
                  disabled={fields.length == 0}
                  onClick={() => {
                    remove(fields.length - 1)
                  }}
                >
                  Удалить
                </Button>
              </div>
            </div>
          </form>
        </Form>
      </CardContent>
      <CardFooter className="space-x-4">
        <Button type="submit" form="storage-form" disabled={isPending}>
          Сохранить
        </Button>
      </CardFooter>
    </Card>
  )
}
