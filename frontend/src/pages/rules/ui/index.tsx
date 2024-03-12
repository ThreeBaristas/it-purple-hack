import { useSuspenseQuery } from '@tanstack/react-query'
import { useNavigate } from '@tanstack/react-router'
import {
  ChevronLeft,
  ChevronRight,
  ChevronsLeft,
  ChevronsRight
} from 'lucide-react'
import React from 'react'

import { GetRulesRequest } from '@/shared/api'
import { Button } from '@/shared/ui'
import { getRulesQueryOptions } from '@/widgets/rules-table/api'
import { RulesTable } from '@/widgets/rules-table/ui'

import { rulesRoute } from '..'

export function RulesPageComponent() {
  const request = rulesRoute.useSearch()
  return (
    <div className="space-y-4">
      <RulesTable request={request} />
      <React.Suspense fallback={<></>}>
        <Pagination request={request} />
      </React.Suspense>
    </div>
  )
}

function Pagination({ request }: { request: GetRulesRequest }) {
  const navigate = useNavigate({})
  const { data: response } = useSuspenseQuery(getRulesQueryOptions(request))

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
            goToPage(response.page - 1)
          }}
        >
          <ChevronLeft className="size-4" />
        </Button>
        <Button
          className="size-8 p-0"
          variant="outline"
          disabled={!canGoNext}
          onClick={() => {
            goToPage(response.page + 1)
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
  )
}
