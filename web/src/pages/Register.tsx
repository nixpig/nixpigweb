import { useState } from "react";
import { AxiosResponse } from "axios";
import { api, ApiResponse } from "../api";

interface RegisterData {
  username: string;
  email: string;
  password: string;
}

async function register(data: RegisterData): Promise<ApiResponse> {
  let res: AxiosResponse<ApiResponse>;

  try {
    res = await api.post("http://localhost:3001/api/user/register", {
      ...data,
    });
  } catch (e: any) {
    res = e.response;
  }

  return res.data;
}

const Register = () => {
  const [username, setUsername] = useState<string>("");
  const [email, setEmail] = useState<string>("");
  const [password, setPassword] = useState<string>("");

  return (
    <div id="register" className="register">
      <h2 className="register__heading">Register</h2>
      <form name="register" className="register__form">
        <div className="register__form-group">
          <label htmlFor="username-input" className="register__input-label">
            Username
          </label>
          <input
            id="username-input"
            className="register__input-field"
            type="text"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
        </div>

        <div className="register__form-group">
          <label htmlFor="email-input" className="register__input-label">
            Email
          </label>
          <input
            id="email-input"
            className="register__input-field"
            type="text"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
        </div>

        <div className="register__form-group">
          <label htmlFor="password-input" className="register__input-label">
            Password
          </label>
          <input
            id="password-input"
            className="register__input-field"
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
        </div>

        <div className="register__form-group">
          <button
            className="register__input-button"
            type="submit"
            onClick={() => register({ username, email, password })}
          >
            Register
          </button>
        </div>
      </form>
    </div>
  );
};
export default Register;
