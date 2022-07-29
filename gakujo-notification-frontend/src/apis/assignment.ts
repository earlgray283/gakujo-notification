import { http } from '../libs/axios';

export interface Assignment {
  kind: string;
  subjectName: string;
  semester: string;
  title: string;
  since: Date;
  until: Date;
  description: string;
  message: string;
  year: number;
  status: string;
}

export async function fetchAssignments(): Promise<Assignment[]> {
  const resp = await http.get('/assignments');
  return resp.data;
}
