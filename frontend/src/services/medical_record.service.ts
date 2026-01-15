import { unwrapResponse } from "@/helpers/response";
import api from "@/lib/api";
import { ApiResponse } from "@/types/api.type";
import {
  MedicalRecordCreateRequest,
  MedicalRecordPaginateResponse,
  MedicalRecordResponse,
  MedicalRecordUpdateRequest,
} from "@/types/medical_record.type";

const BASE_PATH = "/medical-records";

export async function getPatientMedicalRecords(
  id: string,
  page: number,
  limit: number
): Promise<MedicalRecordPaginateResponse> {
  const response = await api.get<ApiResponse<MedicalRecordPaginateResponse>>(
    `${BASE_PATH}/patient/${id}?page=${page}&limit=${limit}`
  );
  return unwrapResponse(response.data);
}

export async function getDoctorMedicalRecords(
  id: string,
  page: number,
  limit: number
): Promise<MedicalRecordPaginateResponse> {
  const response = await api.get<ApiResponse<MedicalRecordPaginateResponse>>(
    `${BASE_PATH}/doctor/${id}?page=${page}&limit=${limit}`
  );
  return unwrapResponse(response.data);
}

export async function getMedicalRecord(
  id: string
): Promise<MedicalRecordResponse> {
  const response = await api.get<ApiResponse<MedicalRecordResponse>>(
    `${BASE_PATH}/${id}`
  );
  return unwrapResponse(response.data);
}

export async function createMedicalRecord(
  payload: MedicalRecordCreateRequest
): Promise<MedicalRecordResponse> {
  const response = await api.post<ApiResponse<MedicalRecordResponse>>(
    BASE_PATH,
    payload
  );
  return unwrapResponse(response.data);
}

export async function updateMedicalRecord(
  id: string,
  payload: MedicalRecordUpdateRequest
): Promise<MedicalRecordResponse> {
  const response = await api.patch<ApiResponse<MedicalRecordResponse>>(
    `${BASE_PATH}/${id}`,
    payload
  );
  return unwrapResponse(response.data);
}

export async function deleteMedicalRecord(id: string): Promise<null> {
  const response = await api.delete<ApiResponse<null>>(`${BASE_PATH}/${id}`);
  return unwrapResponse(response.data);
}
