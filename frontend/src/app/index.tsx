import './index.css'

import React from 'react'
import { createRoot } from 'react-dom/client'

import { Providers } from './providers'
import { QueryClient } from '@tanstack/react-query'

const client = new QueryClient()

const container = document.getElementById('root') as HTMLDivElement
const root = createRoot(container)

root.render(
  <React.StrictMode>
    <Providers queryClient={client} />
  </React.StrictMode>
)
