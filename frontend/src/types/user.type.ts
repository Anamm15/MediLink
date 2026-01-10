import { PatientResponse } from "./patient.type";

export type UserResponse = {
  id: string;
  email: string;
  role: string;
  name: string;
  phone_number: string;
  status: string;
  avatar_url?: string;
  is_verified: boolean;
  created_at: string;
  updated_at: string;
};

export type UserProfileResponse = {
  user: UserResponse;
  patient: PatientResponse | null;
};

export type UserUpdateProfileRequest = Pick<
  UserResponse,
  "name" | "email" | "phone_number"
>;

export type VerifyEmailRequest = {
  otp: string;
};
