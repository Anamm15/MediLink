import { unwrapResponse } from "@/helpers/response";
import api from "@/lib/api";
import { ApiResponse } from "@/types/api.type";
import {
  PrescriptionCreateRequest,
  PrescriptionPaginateResponse,
  PrescriptionResponse,
  PrescriptionUpdateRequest,
} from "@/types/prescription.type";

const BASE_PATH = "/prescriptions";

export async function getPatientPrescriptions(
  id: string,
  page: number,
  limit: number
): Promise<PrescriptionPaginateResponse> {
  const response = await api.get<ApiResponse<PrescriptionPaginateResponse>>(
    `${BASE_PATH}/patient/${id}?page=${page}&limit=${limit}`
  );
  return unwrapResponse(response.data);
}

export async function getDoctorPrescriptions(
  id: string,
  page: number,
  limit: number
): Promise<PrescriptionPaginateResponse> {
  const response = await api.get<ApiResponse<PrescriptionPaginateResponse>>(
    `${BASE_PATH}/doctor/${id}?page=${page}&limit=${limit}`
  );
  return unwrapResponse(response.data);
}

export async function getDetailPrescription(
  id: string
): Promise<PrescriptionResponse> {
  const response = await api.get<ApiResponse<PrescriptionResponse>>(
    `${BASE_PATH}/${id}`
  );
  return unwrapResponse(response.data);
}

export async function createPrescription(
  payload: PrescriptionCreateRequest
): Promise<PrescriptionResponse> {
  const response = await api.post<ApiResponse<PrescriptionResponse>>(
    `${BASE_PATH}`,
    payload
  );
  return unwrapResponse(response.data);
}

export async function updatePrescription(
  id: string,
  payload: PrescriptionUpdateRequest
): Promise<PrescriptionResponse> {
  const response = await api.put<ApiResponse<PrescriptionResponse>>(
    `${BASE_PATH}/${id}`,
    payload
  );
  return unwrapResponse(response.data);
}
