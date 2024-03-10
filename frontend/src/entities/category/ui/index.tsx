import { useQuery } from '@tanstack/react-query'
import { CheckIcon, ChevronsUpDown } from 'lucide-react'
import React from 'react'
import { useDebounce } from 'use-debounce'

import { cn } from '@/shared/lib'
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandLoading,
  PopoverContent,
  PopoverTrigger
} from '@/shared/ui'
import { Button, Popover } from '@/shared/ui'

import { getCategoriesQueryOptions } from '../api'
import { Category } from '../model'

type Props = Omit<
  React.ComponentPropsWithoutRef<typeof Button>,
  'value' | 'onChange'
> & {
  value: Category | undefined
  onChange: (value: Category | undefined) => void
}

export function SelectCategory({
  className,
  value,
  onChange,
  ...props
}: Props) {
  const [open, setOpen] = React.useState(false)
  const [searchValue, setSearchValue] = React.useState('')
  const [debouncedSearchValue] = useDebounce(searchValue, 500)
  const { data: categories, isLoading } = useQuery(
    getCategoriesQueryOptions(debouncedSearchValue)
  )
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
          {value ? value.name : 'Выберите категорию'}
          <ChevronsUpDown className="ml-2 size-4 shrink-0 opacity-50" />
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-[200px] p-0">
        <Command shouldFilter={false}>
          <CommandInput
            placeholder="Выберите категорию"
            value={searchValue}
            onValueChange={setSearchValue}
          />
          <CommandGroup>
            {isLoading && <CommandLoading>Загрузка</CommandLoading>}
            {!isLoading && <CommandEmpty>Не найдено</CommandEmpty>}
            {categories &&
              categories.map((category) => (
                <CommandItem
                  key={category.id}
                  value={String(category.id)}
                  onSelect={() => onChange(category)}
                >
                  {category.name}
                  <CheckIcon
                    className={cn(
                      'ml-auto h-4 w-4',
                      value?.id === category.id ? 'opacity-100' : 'opacity-0'
                    )}
                  />
                </CommandItem>
              ))}
          </CommandGroup>
        </Command>
      </PopoverContent>
    </Popover>
  )
}
