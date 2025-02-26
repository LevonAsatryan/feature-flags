import axios, { AxiosInstance } from 'axios';

const DEV_URL = 'http://localhost:8080';
const PROD_URL = 'https://example.com';
/**
 * Creating an axios instance and exporting ourselves to later on configure interceptors & error handlers
 */
let axiosInstance: AxiosInstance = axios.create({
  baseURL: import.meta.env.DEV ? DEV_URL : PROD_URL,
});

export default axiosInstance;
