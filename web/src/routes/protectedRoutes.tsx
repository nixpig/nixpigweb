import loadable from "@loadable/component";
import { ProtectedRoute } from "./ProtectedRoute";

const Admin = loadable(() => import("../pages/Admin"));
const Config = loadable(() => import("../pages/Config"));

export const protectedRoutes = [
  {
    path: "/admin",
    element: <ProtectedRoute />,
    children: [
      {
        path: "/admin",
        element: <Admin />,
      },
      {
        path: "/admin/config",
        element: <Config />,
      },
    ],
  },
];
