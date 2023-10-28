import loadable from "@loadable/component";
import "./App.css";
import { Routes, Route, Outlet, Link } from "react-router-dom";

const Home = loadable(() => import("./pages/Home"));
const Login = loadable(() => import("./pages/Login"));
const Register = loadable(() => import("./pages/Register"));
const Admin = loadable(() => import("./pages/Admin"));

function App() {
  return (
    <Routes>
      <Route path="/" element={<Wrapper />}>
        <Route index element={<Home />} />
        <Route path="login" element={<Login />} />
        <Route path="register" element={<Register />} />
        <Route path="admin" element={<Admin />} />
      </Route>
    </Routes>
  );
}

function Wrapper() {
  return (
    <div>
      <Link to="/">Home</Link>
      <Link to="/login">Login</Link>
      <Link to="/register">Register</Link>
      <Link to="/admin">Admin</Link>
      <Outlet />
    </div>
  );
}

export default App;
