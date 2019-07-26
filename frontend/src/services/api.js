import axios from "axios";
import {logout} from "./loginService";

const api = axios.create({
  baseURL: "http://localhost:3000",
  withCredentials: true
}, {
  headers: {    
    'Access-Control-Request-Headers': 'Content-Type',
    'Content-Type': 'application/json',
  },
});

api.interceptors.response.use((response) => {
  return response;
}, function (error) {
  // Do something with response error  
  if (error.response.status === 401) {
      console.log('unauthorized, logging out ...');
      logout();      
      window.location.reload();
  }
  return Promise.reject(error.response);
});

export default api;
