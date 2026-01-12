import api from "@/lib/api";
import { ApiResponse } from "@/types/api.type";
import { unwrapResponse } from "@/helpers/response";
import {
  LoginRequest,
  RegistrationRequest,
  RegistrationResponse,
  ResetPasswordRequest,
} from "@/types/auth.type";

const BASE_PATH = "/auth";

export async function login(payload: LoginRequest): Promise<string> {
  const response = await api.post<ApiResponse<string>>(
    `${BASE_PATH}/login`,
    payload
  );

  return unwrapResponse(response.data);
}

export async function register(
  payload: RegistrationRequest
): Promise<RegistrationResponse> {
  const response = await api.post<ApiResponse<RegistrationResponse>>(
    `${BASE_PATH}/register`,
    payload
  );

  return unwrapResponse(response.data);
}

export async function logout(): Promise<null> {
  const response = await api.post<ApiResponse<null>>(`${BASE_PATH}/logout`);

  return unwrapResponse(response.data);
}

export async function changePassword(
  oldPassword: string,
  newPassword: string
): Promise<null> {
  const response = await api.post<ApiResponse<null>>(
    `${BASE_PATH}/change-password`,
    {
      oldPassword,
      newPassword,
    }
  );

  return unwrapResponse(response.data);
}

export async function requestResetPassword(email: string): Promise<null> {
  const response = await api.post<ApiResponse<null>>(
    `${BASE_PATH}/request-reset-password`,
    {
      email,
    }
  );

  return unwrapResponse(response.data);
}

export async function resetPassword(
  payload: ResetPasswordRequest
): Promise<null> {
  const response = await api.post<ApiResponse<null>>(
    `${BASE_PATH}/reset-password`,
    payload
  );

  return unwrapResponse(response.data);
}
