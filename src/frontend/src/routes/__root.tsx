import { createRootRoute, Link, Outlet } from '@tanstack/react-router'

export const Route = createRootRoute({
  component: () => (
    <>
      <div>
        <Link to="/">
          Home
        </Link>{' '}
        <Link to="/connections">
          Connections
        </Link>
      </div>
      <hr />
      <Outlet />
    </>
  ),
})
