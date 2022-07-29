import AsyncStorage from '@react-native-async-storage/async-storage';
import axios, { AxiosResponse } from 'axios';

export const http = axios.create({
  baseURL: 'http://localhost:8080',
  withCredentials: true,
  transformResponse: (data) => {
    if (typeof data === 'string') {
      try {
        return JSON.parse(data, (k: string, val: unknown) => {
          if (typeof val === 'string') {
            const date = new Date(val);
            if (Number.isNaN(date.getDate())) {
              return val;
            }
            return date;
          }
          return val;
        });
      } catch (e) {
        return data;
      }
    }
    return data;
  },
});
http.interceptors.request.use(async (config) => {
  const jwtToken = await AsyncStorage.getItem('jwtToken');
  console.log(jwtToken);
  if (jwtToken && config.headers) {
    config.headers.Authorization = `Bearer ${jwtToken}`;
  }
  return config;
});

export async function postJson<
  T = unknown,
  R = AxiosResponse<T, unknown>,
  D = unknown
>(uri: string, data: D): Promise<R> {
  const resp = http.post<T, R, D>(uri, data, {
    headers: {
      'Content-Type': 'application/json',
    },
  });
  return resp;
}

export async function putJson<
  T = unknown,
  R = AxiosResponse<T, unknown>,
  D = unknown
>(uri: string, data: D): Promise<R> {
  const resp = http.put<T, R, D>(uri, data, {
    headers: {
      'Content-Type': 'application/json',
    },
  });
  return resp;
}

export function isErrorStatus(code: number): boolean {
  return code >= 400;
}
