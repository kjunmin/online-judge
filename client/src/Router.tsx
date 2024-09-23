import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import { HomePage } from './pages/Home.page';
import { Problem } from './pages/Problem/Problem';
import { HealthCheck } from './pages/Healthcheck';
import { CreateProblem } from './pages/CreateProblem/CreateProblem';
import { ProblemsList } from './pages/ProblemsList/ProblemsList';

const router = createBrowserRouter([
  {
    path: '/',
    element: <HomePage />,
  },
  {
    path: "/problemslist",
    element: <ProblemsList />
  },
  {
    path: "/problem/create",
    element: <CreateProblem />
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
