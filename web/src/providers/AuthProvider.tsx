import { api } from "../api";
import { createContext, useContext, useEffect, useMemo, useState } from "react";

interface TokenContextValue {
  token: string | null;
  setToken: React.Dispatch<React.SetStateAction<string | null>>;
}

const AuthContext = createContext<TokenContextValue>({
  token: "",
  setToken: () => null,
});

const AuthProvider = ({ children }: any) => {
  const [token, setToken] = useState(localStorage.getItem("token"));
  // const [token, setToken] = useState<string | null>("");

  useEffect(() => {
    if (token) {
      api.defaults.headers.common.Authorization = `Bearer ${token}`;
      localStorage.setItem("token", token);
    } else {
      delete api.defaults.headers.common.Authorization;
      localStorage.removeItem("token");
    }
  }, [token]);

  const contextValue = useMemo(() => ({ token, setToken }), [token]);

  return (
    <AuthContext.Provider value={contextValue}>{children}</AuthContext.Provider>
  );
};

export const useAuth = () => useContext(AuthContext);

export default AuthProvider;
