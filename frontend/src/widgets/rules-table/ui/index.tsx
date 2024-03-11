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
      <TableCell>{rule.segment}</TableCell>
      <TableCell>{rule.price}</TableCell>
      <TableCell></TableCell>
    </TableRow>
  )
}
