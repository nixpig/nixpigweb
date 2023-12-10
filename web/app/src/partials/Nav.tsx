import { Link } from "react-router-dom";
import { useAuth } from "../provider/authProvider";

export const Nav = () => {
  const { token, setToken } = useAuth();

  return (
    <div>
      <ul>
        <li>
          <Link to="/dashboard">Dashboard</Link>
        </li>
        <li>
          <Link to="/content">Content</Link>
        </li>
        <li>
          <Link to="/create">Create</Link>
        </li>
        <li>
          <Link to="/users">Users</Link>
        </li>
        <li>
          <a href="/" onClick={() => setToken("")}>
            Logout
          </a>
        </li>
      </ul>
    </div>
  );
};
