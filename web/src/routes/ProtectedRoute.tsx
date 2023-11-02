import { Link, Navigate, Outlet } from "react-router-dom";
import { useAuth } from "../providers/AuthProvider";

export const ProtectedRoute = () => {
  const { token } = useAuth();

  if (!token) {
    return <Navigate to="/login" />;
  }

  return (
    <div>
      <h1>Admin</h1>
      <ul>
        <li>
          <Link to="/admin/config">Config</Link>
        </li>
        <li>
          <Link to="/admin/user">User</Link>
        </li>
      </ul>
      <Outlet />
    </div>
  );
};
