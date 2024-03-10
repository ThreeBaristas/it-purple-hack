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
          <p>Укажите локацию и категорию чтобы проверить цену для нее.</p>
          <p>
            Вы так же Вы так же можете указать рекламный сегменент, либо
            оставить его пустым.
          </p>
        </CardDescription>
      </CardHeader>
      <CardContent>
        <form></form>
      </CardContent>
      <CardFooter>
        <Button>Отправить</Button>
      </CardFooter>
    </Card>
  )
}
