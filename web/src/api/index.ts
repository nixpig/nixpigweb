import axios from "axios";

const API_BASE_URL = "http://localhost";
const API_PORT = 3001;
const API_CONTEXT = "api";

const baseUrl = `${API_BASE_URL}${
  API_PORT ? `:${API_PORT}` : ""
}/${API_CONTEXT}`;

export const api = axios.create({
  baseURL: baseUrl,
  headers: {
    Accept: "application/json",
  },
});

export type ApiResponse = {
  data: { token: string } | null;
  error: boolean;
  message: string;
};
