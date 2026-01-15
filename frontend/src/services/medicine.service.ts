import { unwrapResponse } from "@/helpers/response";
import api from "@/lib/api";
import { ApiResponse } from "@/types/api.type";
import {
  MedicinePaginationResponse,
  MedicineResponse,
} from "@/types/medicine.type";

const BASE_PATH = "/medicines";

export async function searchMedicine(
  name: string,
  page: number,
  limit: number
): Promise<MedicinePaginationResponse> {
  const response = await api.get<ApiResponse<MedicinePaginationResponse>>(
    `${BASE_PATH}/search?name=${name}&page=${page}&limit=${limit}`
  );
  return unwrapResponse(response.data);
}
