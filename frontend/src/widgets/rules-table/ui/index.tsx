import { Eye, Trash } from 'lucide-react'

import { Button } from '@/shared/ui'
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow
} from '@/shared/ui/table'

import { Rule } from '../model'

export function RulesTable({ rows }: { rows: Rule[] }) {
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
        {rows.map((it) => (
          <RuleRow
            key={`${it.location.id}${it.category.id}${it.segment}`}
            rule={it}
          />
        ))}
      </TableBody>
    </Table>
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
