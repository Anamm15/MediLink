import api from "@/lib/api";
import { PatientUpdateRequest } from "@/types/patient.type";

const BASE_PATH = "/patients";

export async function updatePatient(payload: PatientUpdateRequest) {
  const response = await api.put(`${BASE_PATH}`, payload);
  return response.data.data;
}
