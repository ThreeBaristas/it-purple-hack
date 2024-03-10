import { valibotResolver } from '@hookform/resolvers/valibot'
import { useForm } from 'react-hook-form'
import { number, object, Output, string } from 'valibot'

import { SelectCategory } from '@/entities/category'
import { SelectLocation } from '@/entities/location'
import { cn } from '@/shared/lib'
import {
  Button,
  Card,
  CardContent,
  CardDescription,
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
} from '@/shared/ui/ui/form'
import { Separator } from '@/shared/ui/ui/separator'

type CheckPriceProps = React.HTMLAttributes<HTMLDivElement>

export function CheckPriceCard({ className, ...props }: CheckPriceProps) {
  const locationSchema = object({
    id: number(),
    name: string()
  })
  const schema = object({
    location: locationSchema,
    category: locationSchema
  })
  const form = useForm<Output<typeof schema>>({
    resolver: valibotResolver(schema),
    defaultValues: {}
  })

  function handleSubmit(values: Output<typeof schema>) {
    console.log(values)
  }

  return (
    <Card className={cn('mx-auto max-w-md', className)} {...props}>
      <CardHeader>
        <CardTitle>Проверить цену</CardTitle>
        <CardDescription>
          Укажите локацию и категорию чтобы проверить цену для нее. Вы так же
          можете указать рекламный сегменент, либо оставить его пустым. Перед
          тем, как изменять цену, нужно ее сначала проверить.
        </CardDescription>
      </CardHeader>
      <CardContent>
        <Form {...form}>
          <form
            className="space-y-4"
            id="check-price-form"
            onSubmit={form.handleSubmit(handleSubmit)}
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
          </form>
        </Form>
      </CardContent>
      <CardFooter>
        <Button type="submit" form="check-price-form">
          Отправить
        </Button>
      </CardFooter>
    </Card>
  )
}
