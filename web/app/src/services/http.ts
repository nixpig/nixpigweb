import axios from "axios";

export const http = axios.create({
  baseURL:
    process.env.NODE_ENV === "development" ? "http://localhost:8081" : "/",
});
