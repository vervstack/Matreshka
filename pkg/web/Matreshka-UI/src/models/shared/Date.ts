export default function dateToString(d: Date): string{
    return `${d.getFullYear()}-${d.getMonth() + 1}-${d.getDate()}  ${d.getHours()}:${d.getMinutes()}`;
}
