import { unwrapResponse } from "@/helpers/response";
import api from "@/lib/api";
import { ApiResponse } from "@/types/api.type";
import {
  AppointmentDetailResponse,
  BookingResponse,
  CreateBookingRequest,
} from "@/types/appointment.type";

const BASE_PATH = "/appointments";

export async function getAppointments(): Promise<AppointmentDetailResponse> {
  const response = await api.get<ApiResponse<AppointmentDetailResponse>>(
    BASE_PATH
  );
  return unwrapResponse(response.data);
}

export async function getAppointment(
  id: string
): Promise<AppointmentDetailResponse> {
  const response = await api.get<ApiResponse<AppointmentDetailResponse>>(
    `${BASE_PATH}/${id}`
  );
  return unwrapResponse(response.data);
}

export async function createAppointment(
  payload: CreateBookingRequest
): Promise<BookingResponse> {
  const response = await api.post<ApiResponse<BookingResponse>>(
    BASE_PATH,
    payload
  );
  return unwrapResponse(response.data);
}

export async function getDoctorAppointments(): Promise<AppointmentDetailResponse> {
  const response = await api.get<ApiResponse<AppointmentDetailResponse>>(
    `${BASE_PATH}/doctor`
  );
  return unwrapResponse(response.data);
}

export async function getPatientAppointments(): Promise<AppointmentDetailResponse> {
  const response = await api.get<ApiResponse<AppointmentDetailResponse>>(
    `${BASE_PATH}/patient`
  );
  return unwrapResponse(response.data);
}

export async function completeAppointment(id: string): Promise<null> {
  const response = await api.post<ApiResponse<null>>(
    `${BASE_PATH}/${id}/complete`
  );
  return unwrapResponse(response.data);
}

export async function cancelAppointment(id: string): Promise<null> {
  const response = await api.post<ApiResponse<null>>(
    `${BASE_PATH}/${id}/cancel`
  );
  return unwrapResponse(response.data);
}

export async function deleteAppointment(id: string): Promise<null> {
  const response = await api.delete<ApiResponse<null>>(`${BASE_PATH}/${id}`);
  return unwrapResponse(response.data);
}
