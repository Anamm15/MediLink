import api from "@/lib/api";
import { OnBoardPatientRequest, PatientResponse } from "@/types/patient.type";
import {
  UserProfileResponse,
  UserResponse,
  UserUpdateProfileRequest,
  VerifyEmailRequest,
} from "@/types/user.type";

const BASE_PATH = "/users";

export async function getUsers(): Promise<UserResponse[]> {
  const response = await api.get(`${BASE_PATH}`);
  return response.data.data;
}

export async function me(): Promise<UserResponse> {
  const response = await api.get(`${BASE_PATH}/me`);
  return response.data.data;
}

export async function getProfile(): Promise<UserProfileResponse> {
  const response = await api.get(`${BASE_PATH}/profile`);
  return response.data.data;
}

export async function sendEmailVerification() {
  const response = await api.post(`${BASE_PATH}/send-email-verification`);
  return response.data;
}

export async function verifyEmail(payload: VerifyEmailRequest) {
  const response = await api.post(`${BASE_PATH}/verify-email`, payload);
  return response.data;
}

export async function onBoardPatient(
  payload: OnBoardPatientRequest
): Promise<PatientResponse> {
  const response = await api.post(`${BASE_PATH}/onboard-patient`, payload);
  return response.data.data;
}

export async function onBoardDoctor() {
  const response = await api.post(`${BASE_PATH}/onboard-doctor`);
  return response.data.data;
}

export async function updateProfile(payload: UserUpdateProfileRequest) {
  const response = await api.put(`${BASE_PATH}`, payload);
  return response.data.data;
}
