import axios from 'axios';

const baseUrl = process.env.ENVIRONMENT === 'development' ? process.env.URL_API_DEV : process.env.URL_API_PROD

const API = axios.create({
  baseURL: baseUrl,
  headers: {
    'Content-Type': 'application/json',
    'Accept': 'application/json'
  }
});

export default API
