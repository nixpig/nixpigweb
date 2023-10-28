import loadable from "@loadable/component";

const Home = loadable(() => import("../pages/Home"));
const Login = loadable(() => import("../pages/Login"));
const Register = loadable(() => import("../pages/Register"));
const About = loadable(() => import("../pages/About"));

export const publicRoutes = [
  {
    path: "/",
    element: <Home />,
  },
  {
    path: "/login",
    element: <Login />,
  },
  {
    path: "/register",
    element: <Register />,
  },
  {
    path: "/about",
    element: <About />,
  },
];
