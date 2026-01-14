import { unwrapResponse } from "@/helpers/response";
import api from "@/lib/api";
import { ApiResponse } from "@/types/api.type";
import { MedicineResponse } from "@/types/medicine.type";

const BASE_PATH = "/medicines";

export async function searchMedicine(
  name: string
): Promise<MedicineResponse[]> {
  const response = await api.get<ApiResponse<MedicineResponse[]>>(
    `${BASE_PATH}/search?name=${name}`
  );
  return unwrapResponse(response.data);
}
