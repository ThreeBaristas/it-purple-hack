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

type CheckPriceProps = React.HTMLAttributes<HTMLDivElement>

export function CheckPriceCard({ className, ...props }: CheckPriceProps) {
  return (
    <Card className={cn('mx-auto max-w-md', className)} {...props}>
      <CardHeader>
        <CardTitle>Проверить цену</CardTitle>
        <CardDescription>
          Укажите локацию и категорию чтобы проверить цену для нее. Вы так же
          можете указать рекламный сегменент, либо оставить его пустым.
        </CardDescription>
      </CardHeader>
      <CardContent>
        <form className="space-y-4">
          <SelectCategory
            value={undefined}
            onChange={console.log}
            className="w-full"
          />
          <SelectLocation
            value={undefined}
            onChange={console.log}
            className="w-full"
          />
        </form>
      </CardContent>
      <CardFooter>
        <Button>Отправить</Button>
      </CardFooter>
    </Card>
  )
}
