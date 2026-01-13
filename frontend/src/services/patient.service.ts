import { unwrapResponse } from "@/helpers/response";
import api from "@/lib/api";
import { ApiResponse } from "@/types/api.type";
import { PatientResponse, PatientUpdateRequest } from "@/types/patient.type";

const BASE_PATH = "/patients";

export async function patientMe(): Promise<PatientResponse> {
  const response = await api.get<ApiResponse<PatientResponse>>(
    `${BASE_PATH}/me`
  );
  return unwrapResponse(response.data);
}

export async function updatePatient(payload: PatientUpdateRequest) {
  const response = await api.put<ApiResponse<PatientResponse>>(
    `${BASE_PATH}`,
    payload
  );
  return unwrapResponse(response.data);
}
