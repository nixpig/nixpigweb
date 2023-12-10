import { Navigate } from "react-router-dom";
import { useAuth } from "../provider/authProvider";

export const Login = () => {
  const { token } = useAuth();

  if (token) {
    return <Navigate to="/dashboard" />;
  }

  return <div>Login screen</div>;
};
