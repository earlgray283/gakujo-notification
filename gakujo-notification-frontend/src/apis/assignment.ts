import { http } from 'libs/axios';

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
}

export async function fetchAssignments(
  jwtToken: string
): Promise<Assignment[]> {
  const resp = await http.get('/assignments', {
    headers: {
      Authorization: `Bearer ${jwtToken}`,
    },
  });
  return resp.data;
}
