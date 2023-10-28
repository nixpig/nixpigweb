import { RouterProvider, createBrowserRouter } from "react-router-dom";
import { protectedRoutes } from "./protectedRoutes";
import { publicRoutes } from "./publicRoutes";

const Router = () => {
  const router = createBrowserRouter([...publicRoutes, ...protectedRoutes]);

  return <RouterProvider router={router} />;
};

export default Router;
