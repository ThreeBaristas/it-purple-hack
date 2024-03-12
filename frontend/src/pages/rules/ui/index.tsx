import { useNavigate } from '@tanstack/react-router'
import { ChevronLeft, ChevronRight, ChevronsLeft, ChevronsRight } from 'lucide-react'

import { Button } from '@/shared/ui'
import { RulesTable } from '@/widgets/rules-table/ui'

import { rulesRoute } from '..'

export function RulesPageComponent() {
  const navigate = useNavigate({})
  const request = rulesRoute.useSearch()
  const response = rulesRoute.useLoaderData()

  const canGoPrevious = response.page > 0
  const canGoNext = response.page < response.totalPages - 1

  function goToPage(page: number) {
    navigate({
      search: (prev) => ({
        ...prev,
        page
      })
    })
  }

  return (
    <div className="space-y-4">
      <RulesTable request={request} />
      <nav className="flex items-center justify-end space-x-8 text-sm">
        <span className="font-medium">
          Страница {response.page + 1} из {response.totalPages}
        </span>
        <div className="flex flex-row items-center space-x-2">
          <Button
            className="size-8 p-0"
            variant="outline"
            disabled={!canGoPrevious}
            onClick={() => {
              goToPage(0)
            }}
          >
            <ChevronsLeft className="size-4" />
          </Button>
          <Button
            className="size-8 p-0"
            variant="outline"
            disabled={!canGoPrevious}
            onClick={() => {
              goToPage(request.page - 1)
            }}
          >
            <ChevronLeft className="size-4" />
          </Button>
          <Button
            className="size-8 p-0"
            variant="outline"
            disabled={!canGoNext}
            onClick={() => {
              goToPage(request.page + 1)
            }}
          >
            <ChevronRight className="size-4" />
          </Button>
          <Button
            className="size-8 p-0"
            variant="outline"
            disabled={!canGoNext}
            onClick={() => {
              goToPage(response.totalPages - 1)
            }}
          >
            <ChevronsRight className="size-4" />
          </Button>
        </div>
      </nav>
    </div>
  )
}
