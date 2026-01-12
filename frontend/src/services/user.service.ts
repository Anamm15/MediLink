import api from "@/lib/api";
import { ApiResponse } from "@/types/api.type";
import { unwrapResponse } from "@/helpers/response";
import { OnBoardPatientRequest, PatientResponse } from "@/types/patient.type";
import {
  UserProfileResponse,
  UserResponse,
  UserUpdateProfileRequest,
  VerifyEmailRequest,
} from "@/types/user.type";

const BASE_PATH = "/users";

/* =======================
   QUERY
======================= */

export async function getUsers(): Promise<UserResponse[]> {
  const response = await api.get<ApiResponse<UserResponse[]>>(BASE_PATH);
  return unwrapResponse(response.data);
}

export async function me(): Promise<UserResponse> {
  const response = await api.get<ApiResponse<UserResponse>>(`${BASE_PATH}/me`);
  return unwrapResponse(response.data);
}

export async function getProfile(): Promise<UserProfileResponse> {
  const response = await api.get<ApiResponse<UserProfileResponse>>(
    `${BASE_PATH}/profile`
  );
  return unwrapResponse(response.data);
}

export async function sendEmailVerification(): Promise<null> {
  const response = await api.post<ApiResponse<null>>(
    `${BASE_PATH}/send-email-verification`
  );
  return unwrapResponse(response.data);
}

export async function verifyEmail(payload: VerifyEmailRequest): Promise<null> {
  const response = await api.post<ApiResponse<null>>(
    `${BASE_PATH}/verify-email`,
    payload
  );
  return unwrapResponse(response.data);
}

export async function onBoardPatient(
  payload: OnBoardPatientRequest
): Promise<PatientResponse> {
  const response = await api.post<ApiResponse<PatientResponse>>(
    `${BASE_PATH}/onboard-patient`,
    payload
  );
  return unwrapResponse(response.data);
}

export async function onBoardDoctor(): Promise<null> {
  const response = await api.post<ApiResponse<null>>(
    `${BASE_PATH}/onboard-doctor`
  );
  return unwrapResponse(response.data);
}

export async function updateProfile(
  payload: UserUpdateProfileRequest
): Promise<UserProfileResponse> {
  const response = await api.put<ApiResponse<UserProfileResponse>>(
    BASE_PATH,
    payload
  );
  return unwrapResponse(response.data);
}
