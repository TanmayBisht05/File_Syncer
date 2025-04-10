import axios from 'axios';

const axiosInstance = axios.create({
  baseURL: 'http://localhost:5000', 
  withCredentials: true, // important if you’re using cookies (like for logout)
});

export default axiosInstance;
