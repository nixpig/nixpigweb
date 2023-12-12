import React, { useEffect } from "react";
import "./App.css";
import AuthProvider from "./provider/authProvider";
import Routes from "./routes";
import axios from "axios";

function App() {
  useEffect(() => {
    axios.get("/api/user").then((res) => {
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
