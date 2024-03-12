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

import { getLocationsQueryOptions } from '../api'
import { Location } from '../model'

type Props = Omit<
  React.ComponentPropsWithoutRef<typeof Button>,
  'value' | 'onChange'
> & {
  value: Location | undefined
  onChange: (value: Location) => void
}

const SelectLocation = React.forwardRef<
  React.ComponentRef<typeof Button>,
  Props
>(({ className, value, onChange, ...props }, ref) => {
  const [open, setOpen] = React.useState(false)
  const [search, setSearch] = React.useState('')
  const [query] = useDebounce(search, 500)
  const { data: locations, isLoading } = useQuery(
    getLocationsQueryOptions(query)
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
          ref={ref}
        >
          {value ? value.name : 'Выберите локацию'}
          <ChevronsUpDown className="ml-2 size-4 shrink-0 opacity-50" />
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-[200px] p-0">
        <Command shouldFilter={false}>
          <CommandInput
            placeholder="Выберите локацию"
            value={search}
            onValueChange={setSearch}
          />
          <CommandGroup>
            {isLoading && <CommandLoading>Загрузка</CommandLoading>}
            {!isLoading && <CommandEmpty>Не найдено</CommandEmpty>}
            {locations &&
              locations.map((location) => (
                <CommandItem
                  key={location.id}
                  value={String(location.id)}
                  onSelect={() => {
                    onChange(location)
                    setOpen(false)
                  }}
                >
                  {location.name}
                  <CheckIcon
                    className={cn(
                      'ml-auto h-4 w-4',
                      value?.id === location.id ? 'opacity-100' : 'opacity-0'
                    )}
                  />
                </CommandItem>
              ))}
          </CommandGroup>
        </Command>
      </PopoverContent>
    </Popover>
  )
})
SelectLocation.displayName = 'SelectLocation'

export { SelectLocation }
