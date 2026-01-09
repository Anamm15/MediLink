import api from "@/lib/api";
import {
  LoginRequest,
  RegistrationRequest,
  RegistrationResponse,
  ResetPasswordRequest,
} from "@/types/auth.type";

const BASE_PATH = "/auth";

export async function login(payload: LoginRequest): Promise<string> {
  const response = await api.post(`${BASE_PATH}/login`, payload);
  return response.data.data;
}

export async function register(
  payload: RegistrationRequest
): Promise<RegistrationResponse> {
  const response = await api.post(`${BASE_PATH}/register`, payload);
  return response.data.data;
}

export async function logout() {
  const response = await api.post(`${BASE_PATH}/logout`);
  return response.data;
}

export async function changePassword(oldPassword: string, newPassword: string) {
  const response = await api.post(`${BASE_PATH}/change-password`, {
    oldPassword,
    newPassword,
  });
  return response.data;
}

export async function requestResetPassword(email: string) {
  const response = await api.post(`${BASE_PATH}/request-reset-password`, {
    email,
  });
  return response.data;
}

export async function resetPassword(payload: ResetPasswordRequest) {
  const response = await api.post(`${BASE_PATH}/reset-password`, payload);
  return response.data.data;
}
