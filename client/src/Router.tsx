import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import { HomePage } from './pages/Home.page';
import { Problem } from './pages/Problem/Problem';
import { HealthCheck } from './pages/Healthcheck';

const router = createBrowserRouter([
  {
    path: '/',
    element: <HomePage />,
  },
  {
    path: "/problem/:problemId",
    element: <Problem />
  },
  {
    path: "/healthcheck",
    element: <HealthCheck />
  }
]);

export function Router() {
  return <RouterProvider router={router} />;
}
