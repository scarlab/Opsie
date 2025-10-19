// api/ApiManager.ts
import axios, { type AxiosInstance } from 'axios';

const ApiManager: AxiosInstance = axios.create({
    baseURL: '/api/v1',
    withCredentials: true,
    headers: {
        'X-Session-Key': '92348793h42348',
    },
});

export default ApiManager;
