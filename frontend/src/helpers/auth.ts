import axios from "axios";
import { BASE_URL } from "@/helpers/constant";

export async function forceRefreshToken() {
  const response = await axios.get<{ data: string }>(
    `${BASE_URL}/auth/refresh-token`,
    { withCredentials: true }
  );

  const newAccessToken = response.data.data;

  if (typeof window !== "undefined") {
    localStorage.setItem("access_token", newAccessToken);
  }

  return newAccessToken;
}
