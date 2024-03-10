import './index.css'

import React from 'react'
import { createRoot } from 'react-dom/client'

import { Providers } from './providers'
import { client } from './providers/query'

const container = document.getElementById('root') as HTMLDivElement
const root = createRoot(container)

root.render(
  <React.StrictMode>
    <Providers queryClient={client} />
  </React.StrictMode>
)
