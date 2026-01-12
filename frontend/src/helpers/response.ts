import { ApiResponse } from "@/types/api.type";

export function unwrapResponse<T>(res: ApiResponse<T>): T {
  if (!res.status) {
    throw new Error(res.message);
  }
  return res.data;
}
