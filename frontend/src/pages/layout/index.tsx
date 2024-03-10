import { ReactQueryDevtools } from '@tanstack/react-query-devtools'
import { Outlet } from '@tanstack/react-router'

export function Layout() {
  return (
    <>
      <header className="bg-primary text-primary-foreground">
        <nav className="container mx-auto px-4 py-2">
          <h1 className="font-mono text-2xl font-bold">
            Three Baristas ft. ОЛПРОГА
          </h1>
        </nav>
      </header>
      <main className="container mx-auto mt-8">
        <Outlet />
      </main>
      <ReactQueryDevtools />
    </>
  )
}
