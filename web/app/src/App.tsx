import React, { useEffect } from "react";
import "./App.css";
import AuthProvider from "./provider/authProvider";
import Routes from "./routes";
import { http } from "./services";

function App() {
  useEffect(() => {
    http.get("/api/user").then((res) => {
      console.log(res.data);
    });
  });

  return (
    <AuthProvider>
      <Routes />
    </AuthProvider>
  );
}

export default App;
