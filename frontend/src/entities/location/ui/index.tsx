import { ChevronsUpDown } from 'lucide-react'
import React from 'react'

import { cn } from '@/shared/lib'
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
  PopoverContent,
  PopoverTrigger
} from '@/shared/ui'
import { Button, Popover } from '@/shared/ui'

import { Location } from '../model'

type Props = Omit<
  React.ComponentPropsWithoutRef<typeof Button>,
  'value' | 'onChange'
> & {
  value: number | undefined
  onChange: (value: number | undefined) => void
}

export function SelectLocation({
  className,
  value,
  onChange,
  ...props
}: Props) {
  const [open, setOpen] = React.useState(false)
  const categories: Array<Location> = [
    {
      id: 1,
      name: 'test location'
    }
  ]

  return (
    <Popover open={open} onOpenChange={setOpen}>
      <PopoverTrigger asChild>
        <Button
          variant="outline"
          role="combobox"
          aria-expanded={open}
          className={cn('w-[200px] justify-between', className)}
          {...props}
        >
          {value
            ? categories.find((category) => category.id === value)?.name
            : 'Выберите локацию'}
          <ChevronsUpDown className="ml-2 size-4 shrink-0 opacity-50" />
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-[200px] p-0">
        <Command>
          <CommandInput placeholder="Выберите локацию" />
          <CommandEmpty>Не найдено</CommandEmpty>
          <CommandList>
            <CommandGroup>
              {categories.map((category) => (
                <CommandItem
                  value={category.name}
                  key={category.id}
                  onSelect={() => onChange(category.id)}
                />
              ))}
            </CommandGroup>
          </CommandList>
        </Command>
      </PopoverContent>
    </Popover>
  )
}
