import axios from "axios";

const api = axios.create({
  baseURL: "http://localhost:3000"
}, {
  headers: {
    'Content-Type': 'application/json',    
  },
  credentials: "include", 
});

export default api;
