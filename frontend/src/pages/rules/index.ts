import { createRoute } from '@tanstack/react-router'

import { rootRoute } from '../root'
import { RulesPageComponent } from './ui'

export const rulesRoute = createRoute({
  getParentRoute: () => rootRoute,
  component: RulesPageComponent,
  path: '/rules'
})
