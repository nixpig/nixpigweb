import loadable from "@loadable/component";
import { ProtectedRoute } from "./ProtectedRoute";

const Admin = loadable(() => import("../pages/admin/Admin"));
const Config = loadable(() => import("../pages/admin/Config"));
const User = loadable(() => import("../pages/admin/User"));
const Post = loadable(() => import("../pages/admin/Post"));

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
      {
        path: "/admin/user",
        element: <User />,
      },
      { path: "/admin/post", element: <Post /> },
    ],
  },
];
