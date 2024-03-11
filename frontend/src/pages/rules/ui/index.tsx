import { RulesTable } from '@/widgets/rules-table/ui'

import { rulesRoute } from '..'

export function RulesPageComponent() {
  const rows = rulesRoute.useLoaderData()
  return (
    <div>
      <RulesTable rows={rows.data} />
    </div>
  )
}
