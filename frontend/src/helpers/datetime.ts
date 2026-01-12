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
