export function toSimpleDateString(date: Date): string {
  return `${date.getMonth() + 1}/${date.getDay()} ${date
    .getUTCHours()
    .toString()
    .padStart(2, '0')}:${date.getUTCMinutes().toString().padStart(2, '0')}`;
}
