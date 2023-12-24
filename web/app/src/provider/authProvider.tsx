import {
  createContext,
  ReactNode,
  useContext,
  useEffect,
  useMemo,
  useState,
} from "react";

import { http } from "../services";

const AuthContext = createContext<{
  token: string | null;
  setToken: (data: string) => void;
}>({ token: "", setToken: () => {} });

const AuthProvider = ({ children }: { children: ReactNode }) => {
  const [token, setToken_] = useState(localStorage.getItem("token"));

  useEffect(() => {
    if (token) {
      http.defaults.headers.common["Authorization"] = `Bearer ${token}`;
      localStorage.setItem("token", token);
    } else {
      delete http.defaults.headers.common["Authorization"];
      localStorage.removeItem("token");
    }
  }, [token]);

  const setToken = (data: string) => {
    setToken_(data);
  };

  const contextValue = useMemo(
    () => ({
      token,
      setToken,
    }),
    [token]
  );

  return (
    <AuthContext.Provider value={contextValue}>{children}</AuthContext.Provider>
  );
};

export const useAuth = () => {
  return useContext(AuthContext);
};

export default AuthProvider;
