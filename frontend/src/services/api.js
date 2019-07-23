import axios from "axios";

const api = axios.create({
  baseURL: "http://localhost:3000",    
}, {
  headers: {    
    "Access-Control-Allow-Origin": "localhost:3001",
    "Access-Control-Allow-Credentials": "true",
    "Access-Control-Allow-Methods": "GET,HEAD,OPTIONS,POST,PUT",
    "Accept": "application/json",
    'Content-Type': 'application/json',    
  },
  crossDomain: true,
  withCredentials: true,    
});

export default api;
