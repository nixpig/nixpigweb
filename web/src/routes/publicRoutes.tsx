import loadable from "@loadable/component";

const Home = loadable(() => import("../pages/Home"));
const Login = loadable(() => import("../pages/Login"));

export const publicRoutes = [
  {
    path: "/",
    element: <Home />,
  },
  {
    path: "/login",
    element: <Login />,
  },
];
