import {
  Button,
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle
} from '@/shared/ui'

export function AdminPageComponent() {
  return (
    <div>
      <Card className="mx-auto max-w-md">
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
    </div>
  )
}
