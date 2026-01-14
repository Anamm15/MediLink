import api from "@/lib/api";
import { ApiResponse } from "@/types/api.type";
import { unwrapResponse } from "@/helpers/response";
import {
  DoctorProfileResponse,
  DoctorUpdateRequest,
} from "@/types/doctor.type";

import {
  DoctorScheduleResponse,
  CreateScheduleRequest,
  UpdateScheduleRequest,
  UpdateStatusScheduleRequest,
} from "@/types/schedule.type";

const BASE_PATH = "/doctors";
const SCHEDULE_BASE_PATH = `${BASE_PATH}/schedules`;

export async function searchDoctor(
  name: string,
  page: number
): Promise<DoctorProfileResponse[]> {
  const response = await api.get<ApiResponse<DoctorProfileResponse[]>>(
    `${BASE_PATH}/search?name=${name}&page=${page}`
  );
  return unwrapResponse(response.data);
}

export async function getDoctor(id: string): Promise<DoctorProfileResponse> {
  const response = await api.get<ApiResponse<DoctorProfileResponse>>(
    `${BASE_PATH}/${id}`
  );
  return unwrapResponse(response.data);
}

export async function doctorMe(): Promise<DoctorProfileResponse> {
  const response = await api.get<ApiResponse<DoctorProfileResponse>>(
    `${BASE_PATH}/me`
  );
  return unwrapResponse(response.data);
}

export async function updateDoctor(
  payload: DoctorUpdateRequest
): Promise<DoctorProfileResponse> {
  const response = await api.put<ApiResponse<DoctorProfileResponse>>(
    BASE_PATH,
    payload
  );
  return unwrapResponse(response.data);
}

export async function getSchedules(
  doctorid: string
): Promise<DoctorScheduleResponse[]> {
  const response = await api.get<ApiResponse<DoctorScheduleResponse[]>>(
    `${SCHEDULE_BASE_PATH}?doctor_id=${doctorid}`
  );
  return unwrapResponse(response.data);
}

export async function getScheduleByID(id: string) {
  const response = await api.get<ApiResponse<DoctorScheduleResponse>>(
    `${SCHEDULE_BASE_PATH}/${id}`
  );
  return unwrapResponse(response.data);
}

export async function getAvailableSchedules(
  doctor_id: string,
  date: string,
  day: string
): Promise<DoctorScheduleResponse[]> {
  const response = await api.get<ApiResponse<DoctorScheduleResponse[]>>(
    `${SCHEDULE_BASE_PATH}/availability?doctor_id=${doctor_id}&date=${date}&day=${day}`
  );
  return unwrapResponse(response.data);
}

export async function createSchedule(
  payload: CreateScheduleRequest
): Promise<DoctorScheduleResponse> {
  const response = await api.post<ApiResponse<DoctorScheduleResponse>>(
    SCHEDULE_BASE_PATH,
    payload
  );
  return unwrapResponse(response.data);
}

export async function updateSchedule(
  id: string,
  payload: UpdateScheduleRequest
): Promise<DoctorScheduleResponse> {
  const response = await api.put<ApiResponse<DoctorScheduleResponse>>(
    `${SCHEDULE_BASE_PATH}/${id}`,
    payload
  );
  return unwrapResponse(response.data);
}

export async function updateStatusSchedule(
  id: string,
  is_active: boolean
): Promise<DoctorScheduleResponse> {
  const response = await api.patch<ApiResponse<DoctorScheduleResponse>>(
    `${SCHEDULE_BASE_PATH}/${id}/status`,
    { is_active }
  );
  return unwrapResponse(response.data);
}

export async function deleteSchedule(id: string): Promise<null> {
  const response = await api.delete<ApiResponse<null>>(
    `${SCHEDULE_BASE_PATH}/${id}`
  );
  return unwrapResponse(response.data);
}
