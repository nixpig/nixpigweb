import axios, { AxiosResponse } from "axios";
import { useState } from "react";

const BASE_URL = "http://localhost:3001";

type LoginResponse = {
  data: { token: string } | null;
  error: boolean;
  message: string;
};

async function login(
  username: string,
  password: string
): Promise<LoginResponse> {
  let res: AxiosResponse<LoginResponse>;

  try {
    res = await axios.post(`${BASE_URL}/api/user/login`, {
      username,
      password,
    });

    return res.data;
  } catch (e: any) {
    return e.response.data;
  }
}

const Login = () => {
  const [username, setUsername] = useState<string>("");
  const [password, setPassword] = useState<string>("");

  return (
    <div>
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

      <button onClick={() => login(username, password)} type="submit">
        Login
      </button>
    </div>
  );
};

export default Login;
