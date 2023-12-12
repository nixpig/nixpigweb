import {
  RouterProvider,
  createBrowserRouter,
  RouteObject,
} from "react-router-dom";
import { Content, Create, Edit, Dashboard, Login, Users } from "../pages";
import { Layout } from "../partials/Layout";
import { useAuth } from "../provider/authProvider";
import { ProtectedRoute } from "./ProtectedRoute";

const Routes = () => {
  const { token } = useAuth();

  const routesForPublic: RouteObject[] = [];

  const routesForAuthenticatedOnly = [
    {
      element: <Layout />,
      children: [
        {
          path: "/",
          element: <ProtectedRoute />,
          children: [
            {
              path: "/dashboard",
              element: <Dashboard />,
            },
            {
              path: "/content",
              element: <Content />,
            },
            {
              path: "/edit/:id",
              element: <Edit />,
            },
            {
              path: "/create",
              element: <Create />,
            },
            {
              path: "/users",
              element: <Users />,
            },
          ],
        },
      ],
    },
  ];

  const routesForNotAuthenticatedOnly = [
    {
      path: "/",
      element: <Login />,
    },
  ];

  const router = createBrowserRouter(
    [
      ...routesForPublic,
      ...(!token ? routesForNotAuthenticatedOnly : []),
      ...routesForAuthenticatedOnly,
      {
        path: "*",
        element: <div>404</div>,
      },
    ],
    {
      basename: "/admin",
    }
  );

  return <RouterProvider router={router} />;
};

export default Routes;
