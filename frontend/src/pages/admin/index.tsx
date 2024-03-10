import { createRoute } from '@tanstack/react-router'

import { CheckPriceCard } from '@/widgets/check-price'

import { rootRoute } from '../root'

export const indexRoute = createRoute({
  getParentRoute: () => rootRoute,
  component: AdminPageComponent,
  path: '/'
})

function AdminPageComponent() {
  return (
    <div>
      <CheckPriceCard />
    </div>
  )
}
