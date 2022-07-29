import axios from 'axios';
import { http } from '../libs/axios';

export async function signin(
  username: string,
  password: string
): Promise<string> {
  const data = new FormData();
  data.append('username', username);
  data.append('password', password);
  const resp = await http.post('/auth/signin', data);
  return resp.data;
}

export async function signup(
  username: string,
  password: string,
  gakujoId: string,
  gakujoPassword: string
): Promise<string> {
  const data = new FormData();
  data.append('username', username);
  data.append('password', password);
  data.append('gakujoId', gakujoId);
  data.append('gakujoPassword', gakujoPassword);
  const resp = await http.post('/auth/signup', data, {
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  });
  return resp.data;
}
