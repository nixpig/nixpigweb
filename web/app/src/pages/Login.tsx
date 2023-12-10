import { Navigate } from "react-router-dom";
import { useAuth } from "../provider/authProvider";
import axios from "axios";
import { useState } from "react";

const login = async (
  e: any,
  username: string,
  password: string,
  setToken: any
) => {
  e.preventDefault();

  try {
    let res = await axios.post("https://nixpig.dev/api/auth/login", {
      username,
      password,
    });

    setToken(res.data.data.token);
  } catch (e) {
    alert("Login failed");
  }
};

export const Login = () => {
  const { token, setToken } = useAuth();

  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  if (token) {
    return <Navigate to="/dashboard" />;
  }

  return (
    <div>
      <form>
        <div>
          <label htmlFor="username">Username: </label>
          <input id="username" onChange={(e) => setUsername(e.target.value)} />
        </div>

        <div>
          <label htmlFor="password">Password: </label>
          <input
            type="password"
            id="password"
            onChange={(e) => setPassword(e.target.value)}
          />
        </div>

        <div>
          <button
            type="submit"
            id="login-btn"
            onClick={(e) => login(e, username, password, setToken)}
          >
            Login
          </button>
        </div>
      </form>
    </div>
  );
};
