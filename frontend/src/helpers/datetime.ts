import { DAYS } from "./constant";

export const getCurrentTime = () => {
  const now = new Date();
  const hours = String(now.getHours()).padStart(2, "0");
  const minutes = String(now.getMinutes()).padStart(2, "0");
  return `${hours}:${minutes}`;
};

export const getCurrentDate = () => {
  const now = new Date();
  const formattedDate = [
    now.getFullYear(),
    String(now.getMonth() + 1).padStart(2, "0"),
    String(now.getDate()).padStart(2, "0"),
  ].join("-");
  return formattedDate;
};

export const getCurrentDay = () => {
  return DAYS[new Date().getDay()];
};

export const getToday = () => {
  const now = new Date();
  const date = [
    now.getFullYear(),
    String(now.getMonth() + 1).padStart(2, "0"),
    String(now.getDate()).padStart(2, "0"),
  ].join("-");

  const day = DAYS[new Date().getDay()];
  return { date, day };
};

export const getDayOfDate = (date: string) => {
  return DAYS[new Date(date).getDay()];
};

export const formatIDDate = (dateString: string) => {
  return new Date(dateString).toLocaleDateString("id-ID", {
    day: "numeric",
    month: "short",
    year: "numeric",
  });
};

export const formatIDDatetime = (dateString: string) => {
  return new Date(dateString).toLocaleDateString("id-ID", {
    day: "numeric",
    month: "short",
    year: "numeric",
    hour: "numeric",
    minute: "numeric",
  });
};

export const isSameDate = (d1: Date, d2: Date) => {
  return (
    d1.getDate() === d2.getDate() &&
    d1.getMonth() === d2.getMonth() &&
    d1.getFullYear() === d2.getFullYear()
  );
};

export const isToday = (d: Date) => isSameDate(d, new Date());

export function getDaysInMonth(date: Date) {
  const year = date.getFullYear();
  const month = date.getMonth();
  const daysInMonth = new Date(year, month + 1, 0).getDate();
  const firstDayObj = new Date(year, month, 1);
  const startDay = firstDayObj.getDay();

  const res = [];
  for (let i = 0; i < startDay; i++) res.push(null);
  for (let i = 1; i <= daysInMonth; i++) res.push(new Date(year, month, i));
  return res;
}
