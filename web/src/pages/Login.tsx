import { NavigateFunction, useNavigate } from "react-router-dom";
import { useAuth } from "../providers/AuthProvider";
import { AxiosResponse } from "axios";
import { useState } from "react";
import { api, ApiResponse } from "../api";

const BASE_URL = "http://localhost:3001";

interface LoginData {
  username: string;
  password: string;
}

async function login(
  e: React.MouseEvent,
  data: LoginData,
  navigate: NavigateFunction,
  setToken: React.Dispatch<React.SetStateAction<string | null>>
): Promise<any> {
  e.preventDefault();

  let res: AxiosResponse<ApiResponse>;

  try {
    res = await api.post(`${BASE_URL}/api/user/login`, {
      ...data,
    });

    setToken(res.data.data?.token ?? null);
    if (navigator) {
      navigate("/admin", { replace: true });
    } else {
      navigate("/admin", { replace: true });
      return res.data;
    }
  } catch (e: any) {
    return e.response.data;
  }
}

const Login = () => {
  const navigate = useNavigate();
  const { setToken } = useAuth();

  const [username, setUsername] = useState<string>("");
  const [password, setPassword] = useState<string>("");

  return (
    <div>
      <h2>Login</h2>
      <form>
        <input
          type="text"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
        />
        <input
          type="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />

        <button
          onClick={(e) => login(e, { username, password }, navigate, setToken)}
          type="submit"
        >
          Login
        </button>
      </form>
    </div>
  );
};

export default Login;
