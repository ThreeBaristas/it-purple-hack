import { useSuspenseQuery } from '@tanstack/react-query'
import { Eye, Trash } from 'lucide-react'
import React from 'react'

import { GetRulesRequest } from '@/shared/api'
import { Button } from '@/shared/ui'
import { Skeleton } from '@/shared/ui/skeleton'
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow
} from '@/shared/ui/table'

import { getRulesQueryOptions } from '../api'
import { Rule } from '../model'

function RulesTable({ request }: { request: GetRulesRequest }) {
  const { data } = useSuspenseQuery(getRulesQueryOptions(request))
  return (
    <Table>
      <TableHeader>
        <TableRow>
          <TableHead>Локация</TableHead>
          <TableHead>Категория</TableHead>
          <TableHead>Сегмент</TableHead>
          <TableHead>Цена</TableHead>
          <TableHead>Действия</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {data.data.map((it) => (
          <RuleRow
            key={`${it.location.id}${it.category.id}${it.segment}`}
            rule={it}
          />
        ))}
      </TableBody>
    </Table>
  )
}

function RulesTableFallback({ nRows }: { nRows: number }) {
  return (
    <Table>
      <TableHeader>
        <TableRow>
          <TableHead>Локация</TableHead>
          <TableHead>Категория</TableHead>
          <TableHead>Сегмент</TableHead>
          <TableHead>Цена</TableHead>
          <TableHead>Действия</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {new Array(nRows).fill(0).map((_, ind) => (
          <TableRow key={ind}>
            <TableCell>
              <Skeleton className="h-8 w-full" />
            </TableCell>
            <TableCell>
              <Skeleton className="h-8 w-full" />
            </TableCell>
            <TableCell>
              <Skeleton className="h-8 w-full" />
            </TableCell>
            <TableCell>
              <Skeleton className="h-8 w-full" />
            </TableCell>
            <TableCell>
              <Skeleton className="h-8 w-full" />
            </TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  )
}

const RulesTableContainer = ({ request }: { request: GetRulesRequest }) => {
  return (
    <React.Suspense fallback={<RulesTableFallback nRows={10} />}>
      <RulesTable request={request} />
    </React.Suspense>
  )
}

function RuleRow({ rule }: { rule: Rule }) {
  return (
    <TableRow>
      <TableCell>{rule.location.name}</TableCell>
      <TableCell>{rule.category.name}</TableCell>
      <TableCell>{rule.segment ? rule.segment : 'Baseline'}</TableCell>
      <TableCell>
        {new Intl.NumberFormat('ru', {
          style: 'currency',
          currency: 'RUB'
        }).format(rule.price)}
      </TableCell>
      <TableCell>
        <Button variant="outline" size="icon" className="size-8">
          <Eye className="size-4" />
        </Button>
        <Button variant="destructive" size="icon" className="ml-2 size-8">
          <Trash className="size-4" />
        </Button>
      </TableCell>
    </TableRow>
  )
}

export { RulesTableContainer as RulesTable }
