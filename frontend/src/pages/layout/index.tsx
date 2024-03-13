import { ReactQueryDevtools } from '@tanstack/react-query-devtools'
import { Link, Outlet } from '@tanstack/react-router'

export function Layout() {
  return (
    <>
      <header className="bg-primary text-primary-foreground">
        <nav className="container mx-auto flex items-center justify-between px-4 py-2">
          <h1 className="font-mono text-2xl font-bold">
            <Link search={{ page: 0, pageSize: 10 }} to="/rules">
              Three Baristas ft. ОЛПРОГА
            </Link>
          </h1>
          <ul className="flex flex-row space-x-8">
            <li>
              <Link to="/storage">Сторадж</Link>
            </li>
            <li>
              <Link to="/rules" search={{ page: 0, pageSize: 10 }}>
                Правила цен
              </Link>
            </li>
          </ul>
        </nav>
      </header>
      <main className="container mx-auto mt-8">
        <Outlet />
      </main>
      <ReactQueryDevtools />
    </>
  )
}
