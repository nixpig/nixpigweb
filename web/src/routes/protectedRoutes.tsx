import loadable from "@loadable/component";
import { ProtectedRoute } from "./ProtectedRoute";

const Admin = loadable(() => import("../pages/Admin"));

export const protectedRoutes = [
  {
    path: "/admin",
    element: <ProtectedRoute />,
    children: [
      {
        path: "/admin",
        element: <Admin />,
      },
    ],
  },
];
