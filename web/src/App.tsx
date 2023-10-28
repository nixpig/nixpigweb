import "./App.css";
import AuthProvider from "./providers/AuthProvider";
import Router from "./routes/Router";

function App() {
  return (
    <AuthProvider>
      <Router />
    </AuthProvider>
  );
}

export default App;
